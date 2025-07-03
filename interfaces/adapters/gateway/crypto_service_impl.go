package gateway

import (
    "github.com/ktaki8ra/cleanarch-go/domain/domain_model"
)

type CryptoServiceImpl struct {}

func (cs *CryptoServiceImpl) Encrypt(
    plainTextPassword domain_model.PlainTextPassword,
) (domain_model.EncryptedPassword, error) {
    return domain_model.EncryptedPassword{ Value: plainTextPassword.Value + "ENC" }, nil
}
