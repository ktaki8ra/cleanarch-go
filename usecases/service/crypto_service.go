package service

import (
    "github.com/ktaki8ra/cleanarch-go/domain/domain_model"
)

type CryptoService interface {
    Encrypt(plainTextPassword domain_model.PlainTextPassword) (domain_model.EncryptedPassword, error)
}
