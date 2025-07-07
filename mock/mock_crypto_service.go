package mock

// Defined Mock for CryptoService

import (
    "github.com/ktaki8ra/cleanarch-go/domain/domain_model"
)

type MockCryptoService struct {
    EncryptFunc func(domain_model.PlainTextPassword) (domain_model.EncryptedPassword, error)
    DecryptFunc func(domain_model.EncryptedPassword) (domain_model.PlainTextPassword, error)
    MatchesFunc func(string, string) bool
}
func (m *MockCryptoService) Encrypt(plainTextPassword domain_model.PlainTextPassword) (domain_model.EncryptedPassword, error) {
    return m.EncryptFunc(plainTextPassword)
}
func (m *MockCryptoService) Decrypt(encryptedPassword domain_model.EncryptedPassword) (domain_model.PlainTextPassword, error) {
    return m.DecryptFunc(encryptedPassword)
}
func (m *MockCryptoService) Matches(value0 string, value1 string) bool {
    return m.MatchesFunc(value0, value1)
}
