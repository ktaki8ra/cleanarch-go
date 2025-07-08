package usecase_test

import (
    "fmt"
    "net/http"
    "errors"
    "testing"
    "github.com/ktaki8ra/cleanarch-go/mock"

    "github.com/stretchr/testify/assert"
    "github.com/ktaki8ra/cleanarch-go/domain/domain_model"
    "github.com/ktaki8ra/cleanarch-go/usecases/usecase"
)

func TestUserDeleteUseCase(t *testing.T) {
    input := usecase.UserDeleteInputData{
        UserId:            domain_model.UserId{Value: "user01"},
        PlainTextPassword: domain_model.PlainTextPassword{Value: "???"},
    }
    t.Run("Success", func(t *testing.T) {
        mockRepo := &mock.MockUserRepository{
            FindUserByIdFunc: func(userId domain_model.UserId) (domain_model.User, error) {
                return domain_model.User{
                    UserId: domain_model.UserId{Value: "user01"},
                    EncryptedPassword: domain_model.EncryptedPassword{Value: "???"},
                }, nil
            },
            DeleteFunc: func(user domain_model.User) error {
                return nil
            },
        }
        mockCrypto := &mock.MockCryptoService{
            DecryptFunc: func(encryptedPassword domain_model.EncryptedPassword) (domain_model.PlainTextPassword, error) {
                return domain_model.PlainTextPassword{Value: "???"}, nil
            },
            MatchesFunc: func(value0 string, value1 string) bool {
                return true
            },
        }
        uc := usecase.NewUserDeleteUseCase(mockCrypto, mockRepo)
        output, err := uc.Execute(input)
        assert.Empty(t, err.Err)
        assert.Equal(t, input.UserId, output.UserId)
    })

    t.Run("Failed: User(UserId) Not Found", func(t *testing.T) {
        mockRepo := &mock.MockUserRepository{
            // Fail
            FindUserByIdFunc: func(userId domain_model.UserId) (domain_model.User, error) {
                return domain_model.User{}, errors.New("Input UserId Not found")
            },
            // Not Executed
            DeleteFunc: func(user domain_model.User) error {
                return nil
            },
        }
        mockCrypto := &mock.MockCryptoService{
            // Not Executed
            DecryptFunc: func(encryptedPassword domain_model.EncryptedPassword) (domain_model.PlainTextPassword, error) {
                return domain_model.PlainTextPassword{Value: "???"}, nil
            },
            // Not Executed
            MatchesFunc: func(value0 string, value1 string) bool {
                return true
            },
        }
        uc := usecase.NewUserDeleteUseCase(mockCrypto, mockRepo)
        _, err := uc.Execute(input)
        assert.Equal(t, http.StatusNotFound, err.StatusCode)
        assert.Equal(t, fmt.Sprintf("UserId %s Not Found", input.UserId.Value), err.Msg)
        assert.Error(t, err.Err)
    })

    t.Run("Failed: Password Decrypt Error", func(t *testing.T) {
        mockRepo := &mock.MockUserRepository{
            // Pass
            FindUserByIdFunc: func(userId domain_model.UserId) (domain_model.User, error) {
                return domain_model.User{
                    UserId: domain_model.UserId{Value: "user01"},
                    EncryptedPassword: domain_model.EncryptedPassword{Value: "???"},
                }, nil
            },
            // Not Executed
            DeleteFunc: func(user domain_model.User) error {
                return nil
            },
        }
        mockCrypto := &mock.MockCryptoService{
            // Fail
            DecryptFunc: func(encryptedPassword domain_model.EncryptedPassword) (domain_model.PlainTextPassword, error) {
                return domain_model.PlainTextPassword{}, errors.New("Password Decrypt Error")
            },
            // Not Executed
            MatchesFunc: func(value0 string, value1 string) bool {
                return true
            },
        }
        uc := usecase.NewUserDeleteUseCase(mockCrypto, mockRepo)
        _, err := uc.Execute(input)
        assert.Equal(t, http.StatusInternalServerError, err.StatusCode)
        assert.Equal(t, "Decrypt Password Error", err.Msg)
        assert.Error(t, err.Err)
    })

    t.Run("Failed: Password Matches Error", func(t *testing.T) {
        mockRepo := &mock.MockUserRepository{
            // Pass
            FindUserByIdFunc: func(userId domain_model.UserId) (domain_model.User, error) {
                return domain_model.User{
                    UserId: domain_model.UserId{Value: "user01"},
                    EncryptedPassword: domain_model.EncryptedPassword{Value: "???"},
                }, nil
            },
            // Not Executed
            DeleteFunc: func(user domain_model.User) error {
                return nil
            },
        }
        mockCrypto := &mock.MockCryptoService{
            // Pass
            DecryptFunc: func(encryptedPassword domain_model.EncryptedPassword) (domain_model.PlainTextPassword, error) {
                return domain_model.PlainTextPassword{Value: "???"}, nil
            },
            // Fail
            MatchesFunc: func(value0 string, value1 string) bool {
                return false
            },
        }
        uc := usecase.NewUserDeleteUseCase(mockCrypto, mockRepo)
        _, err := uc.Execute(input)
        assert.Equal(t, http.StatusInternalServerError, err.StatusCode)
        assert.Equal(t, "Password Matches Error", err.Msg)
        assert.Error(t, err.Err)
    })

    t.Run("Failed: Delete User failure", func(t *testing.T) {
        mockRepo := &mock.MockUserRepository{
            // Pass
            FindUserByIdFunc: func(userId domain_model.UserId) (domain_model.User, error) {
                return domain_model.User{
                    UserId: domain_model.UserId{Value: "user01"},
                    EncryptedPassword: domain_model.EncryptedPassword{Value: "???"},
                }, nil
            },
            // Fail
            DeleteFunc: func(user domain_model.User) error {
                return errors.New("Delete User failure")
            },
        }
        mockCrypto := &mock.MockCryptoService{
            // Pass
            DecryptFunc: func(encryptedPassword domain_model.EncryptedPassword) (domain_model.PlainTextPassword, error) {
                return domain_model.PlainTextPassword{Value: "???"}, nil
            },
            // Pass
            MatchesFunc: func(value0 string, value1 string) bool {
                return true
            },
        }
        uc := usecase.NewUserDeleteUseCase(mockCrypto, mockRepo)
        _, err := uc.Execute(input)
        assert.Equal(t, http.StatusInternalServerError, err.StatusCode)
        assert.Equal(t, "User Delete Error", err.Msg)
        assert.Error(t, err.Err)
    })
}

