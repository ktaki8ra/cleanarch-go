package gateway_test

import (
    "testing"
    "github.com/stretchr/testify/assert"

    "github.com/ktaki8ra/cleanarch-go/interfaces/adapters/gateway"
    "github.com/ktaki8ra/cleanarch-go/domain/domain_model"
)

func TestCryptoServiceImpl(t *testing.T) {
    t.Run("Success", func(t *testing.T) {
        secretKey := "RTw9EHjFBcdJfrNehvL3eTEDYkFCZXFg"
        plainTextPassword := domain_model.PlainTextPassword{Value: "p4ssw0rd"}

        cs := &gateway.CryptoServiceImpl{SecretKey: secretKey}

        encryptedPassword, err := cs.Encrypt(plainTextPassword)
        assert.Nil(t, err)

        newPlainTextPassword, err := cs.Decrypt(encryptedPassword)
        assert.Nil(t, err)

        assert.Equal(t, plainTextPassword.Value, newPlainTextPassword.Value)
    })
}
