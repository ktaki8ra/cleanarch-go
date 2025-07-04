package service

import (
    "github.com/ktaki8ra/cleanarch-go/domain/domain_model"
)

type CryptoService interface {
    Encrypt(plainTextPassword domain_model.PlainTextPassword) (domain_model.EncryptedPassword, error)
    Decrypt(encryptedPassword domain_model.EncryptedPassword) (domain_model.PlainTextPassword, error)
}
