package gateway

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "fmt"
    "io"

    "github.com/ktaki8ra/cleanarch-go/domain/domain_model"
)

type CryptoServiceImpl struct {
    SecretKey string
}

// AES-256
func (cs *CryptoServiceImpl) Encrypt(
    plainTextPassword domain_model.PlainTextPassword,
) (domain_model.EncryptedPassword, error) {
    key := []byte(cs.SecretKey)
    if len(key) != 32 {
        return domain_model.EncryptedPassword{}, fmt.Errorf("SecretKey must be 32 bytes for AES-256")
    }
    block, err := aes.NewCipher(key)
    if err != nil {
        return domain_model.EncryptedPassword{}, fmt.Errorf("cipher init error: %w", err)
    }
    aesGCM, err := cipher.NewGCM(block)
    if err != nil {
        return domain_model.EncryptedPassword{}, fmt.Errorf("GCM init error: %w", err)
    }
    nonce := make([]byte, aesGCM.NonceSize())
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        return domain_model.EncryptedPassword{}, fmt.Errorf("nonce gen error: %w", err)
    }
    ciphertext := aesGCM.Seal(nonce, nonce, []byte(plainTextPassword.Value), nil)
    encoded := base64.StdEncoding.EncodeToString(ciphertext)
    return domain_model.EncryptedPassword{Value: encoded}, nil
}

func (cs *CryptoServiceImpl) Decrypt(
    encryptedPassword domain_model.EncryptedPassword,
) (domain_model.PlainTextPassword, error) {
    key := []byte(cs.SecretKey)
    if len(key) != 32 {
        return domain_model.PlainTextPassword{}, fmt.Errorf("SecretKey must be 32 bytes for AES-256")
    }
    data, err := base64.StdEncoding.DecodeString(encryptedPassword.Value)
    if err != nil {
        return domain_model.PlainTextPassword{}, fmt.Errorf("base64 decode error: %w", err)
    }
    block, err := aes.NewCipher(key)
    if err != nil {
        return domain_model.PlainTextPassword{}, fmt.Errorf("cipher init error: %w", err)
    }
    aesGCM, err := cipher.NewGCM(block)
    if err != nil {
        return domain_model.PlainTextPassword{}, fmt.Errorf("GCM init error: %w", err)
    }
    nonceSize := aesGCM.NonceSize()
    if len(data) < nonceSize {
        return domain_model.PlainTextPassword{}, fmt.Errorf("ciphertext too short")
    }
    nonce, ciphertext := data[:nonceSize], data[nonceSize:]
    plain, err := aesGCM.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        return domain_model.PlainTextPassword{}, fmt.Errorf("decrypt error: %w", err)
    }
    return domain_model.PlainTextPassword{Value: string(plain)}, nil
}

