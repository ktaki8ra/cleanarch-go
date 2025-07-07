package usecase_test

import (
    "net/http"
    "errors"
    "testing"
    "github.com/ktaki8ra/cleanarch-go/mock"

    "github.com/stretchr/testify/assert"
    "github.com/ktaki8ra/cleanarch-go/domain/domain_model"
    "github.com/ktaki8ra/cleanarch-go/usecases/usecase"
)

func TestUserCreateUseCase(t *testing.T) {
    input := usecase.UserCreateInputData{
        UserId:            domain_model.UserId{Value: "user01"},
        Email:             domain_model.Email{Value: "user01@example.com"},
        PlainTextPassword: domain_model.PlainTextPassword{Value: "???"},
    }
    t.Run("Success", func(t *testing.T) {
        mockRepo := &mock.MockUserRepository{
            FindUserByIdFunc: func(userId domain_model.UserId) (domain_model.User, error) {
                return domain_model.User{}, errors.New("Input UserId Not Found")
            },
            FindUserByEmailFunc: func(email domain_model.Email) (domain_model.User, error) {
                return domain_model.User{}, errors.New("Input Email not found")
            },
            SaveFunc: func(user domain_model.User) error {
                return nil
            },
        }
        mockCrypto := &mock.MockCryptoService{
            EncryptFunc: func(plainTextPassword domain_model.PlainTextPassword) (domain_model.EncryptedPassword, error) {
                return domain_model.EncryptedPassword{Value: "???"}, nil
            },
        }
        uc := usecase.NewUserCreateUseCase(mockCrypto, mockRepo)
        output, err := uc.Execute(input)
        assert.Empty(t, err.Err)
        assert.Equal(t, input.UserId, output.UserId)
    })

    t.Run("Failed: Already Exists UserId", func(t *testing.T) {
        mockRepo := &mock.MockUserRepository{
            FindUserByIdFunc: func(userId domain_model.UserId) (domain_model.User, error) {
                return domain_model.User{
                    UserId:            domain_model.UserId{Value: "user01"},
                    Email:             domain_model.Email{Value: "test01@example.com"},
                    EncryptedPassword: domain_model.EncryptedPassword{Value: "???"},
                }, nil
            },
            FindUserByEmailFunc: func(email domain_model.Email) (domain_model.User, error) {
                return domain_model.User{}, errors.New("Input Email not found")
            },
            SaveFunc: func(user domain_model.User) error {
                return nil
            },
        }
        mockCrypto := &mock.MockCryptoService{
            EncryptFunc: func(plainTextPassword domain_model.PlainTextPassword) (domain_model.EncryptedPassword, error) {
                return domain_model.EncryptedPassword{Value: "???"}, nil
            },
        }
        uc := usecase.NewUserCreateUseCase(mockCrypto, mockRepo)
        _, err := uc.Execute(input)
        assert.Equal(t, http.StatusConflict, err.StatusCode)
        assert.Equal(t, "Input UserId Already Exists", err.Msg)
        assert.Error(t, err.Err)
    })

    t.Run("Failed: Already Exists Email", func(t *testing.T) {
        mockRepo := &mock.MockUserRepository{
            FindUserByIdFunc: func(userId domain_model.UserId) (domain_model.User, error) {
                return domain_model.User{}, errors.New("Input UserId not found")
            },
            FindUserByEmailFunc: func(email domain_model.Email) (domain_model.User, error) {
                return domain_model.User{
                    UserId:            domain_model.UserId{Value: "test01"},
                    Email:             domain_model.Email{Value: "user01@example.com"},
                    EncryptedPassword: domain_model.EncryptedPassword{Value: "???"},
                }, nil
            },
            SaveFunc: func(user domain_model.User) error {
                return nil
            },
        }
        mockCrypto := &mock.MockCryptoService{
            EncryptFunc: func(plainTextPassword domain_model.PlainTextPassword) (domain_model.EncryptedPassword, error) {
                return domain_model.EncryptedPassword{Value: "???"}, nil
            },
        }
        uc := usecase.NewUserCreateUseCase(mockCrypto, mockRepo)
        _, err := uc.Execute(input)
        assert.Equal(t, http.StatusConflict, err.StatusCode)
        assert.Equal(t, "Input Email Already Exists", err.Msg)
        assert.Error(t, err.Err)
    })

    t.Run("Failed: Encrypt Password Error", func(t *testing.T) {
        mockRepo := &mock.MockUserRepository{
            FindUserByIdFunc: func(userId domain_model.UserId) (domain_model.User, error) {
                return domain_model.User{}, errors.New("Input UserId not found")
            },
            FindUserByEmailFunc: func(email domain_model.Email) (domain_model.User, error) {
                return domain_model.User{}, errors.New("Input Email not found")
            },
            SaveFunc: func(user domain_model.User) error {
                return nil
            },
        }
        mockCrypto := &mock.MockCryptoService{
            EncryptFunc: func(plainTextPassword domain_model.PlainTextPassword) (domain_model.EncryptedPassword, error) {
                return domain_model.EncryptedPassword{}, errors.New("Password Encryption Failed.")
            },
        }
        uc := usecase.NewUserCreateUseCase(mockCrypto, mockRepo)
        _, err := uc.Execute(input)
        assert.Equal(t, http.StatusInternalServerError, err.StatusCode)
        assert.Equal(t, "Encrypt Password Error", err.Msg)
        assert.Error(t, err.Err)
    })

    t.Run("Failed: Create User failure", func(t *testing.T) {
        mockRepo := &mock.MockUserRepository{
            FindUserByIdFunc: func(userId domain_model.UserId) (domain_model.User, error) {
                return domain_model.User{}, errors.New("Input UserId not found")
            },
            FindUserByEmailFunc: func(email domain_model.Email) (domain_model.User, error) {
                return domain_model.User{}, errors.New("Input Email not found")
            },
            SaveFunc: func(user domain_model.User) error {
                return errors.New("User Create Failed")
            },
        }
        mockCrypto := &mock.MockCryptoService{
            EncryptFunc: func(plainTextPassword domain_model.PlainTextPassword) (domain_model.EncryptedPassword, error) {
                return domain_model.EncryptedPassword{Value: "???"}, nil
            },
        }
        uc := usecase.NewUserCreateUseCase(mockCrypto, mockRepo)
        _, err := uc.Execute(input)
        assert.Equal(t, http.StatusInternalServerError, err.StatusCode)
        assert.Equal(t, "User Create Error", err.Msg)
        assert.Error(t, err.Err)
    })
}

