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

func TestUserUpdateUseCase(t *testing.T) {
    input := usecase.UserUpdateInputData{
        UserId:            domain_model.UserId{Value: "user01"},
        NewUserId:         domain_model.UserId{Value: "test01"},
        PlainTextPassword: domain_model.PlainTextPassword{Value: "???"},
    }
    t.Run("Success", func(t *testing.T) {
        mockRepo := &mock.MockUserRepository{
            FindUserByIdFunc: func(userId domain_model.UserId) (domain_model.User, error) {
                if userId.Value == "user01" {
                    return domain_model.User{
                        UserId: domain_model.UserId{Value: "user01"},
                        Email: domain_model.Email{Value: "user01@example.com"},
                        EncryptedPassword: domain_model.EncryptedPassword{Value: "???"},
                    }, nil
                } else {
                    return domain_model.User{}, errors.New("Input NewUserId Not Found")
                }
            },
            UpdateFunc: func(currentUserId domain_model.UserId, newUser domain_model.User) error {
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
        uc := usecase.NewUserUpdateUseCase(mockCrypto, mockRepo)
        output, err := uc.Execute(input)
        assert.Empty(t, err.Err)
        assert.Equal(t, input.NewUserId, output.UserId)
    })

    t.Run("Failed: User(UserId) Not Found.", func(t *testing.T) {
        mockRepo := &mock.MockUserRepository{
            FindUserByIdFunc: func(userId domain_model.UserId) (domain_model.User, error) {
                if userId.Value == "user01" { // Fail
                    return domain_model.User{}, errors.New("Input UserId Not found")
                } else {
                    return domain_model.User{ // Not Executed
                        UserId: domain_model.UserId{Value: "test01"},
                        Email: domain_model.Email{Value: "test01@example.com"},
                        EncryptedPassword: domain_model.EncryptedPassword{Value: "???"},
                    }, nil
                }
            },
            // Not Executed
            UpdateFunc: func(currentUserId domain_model.UserId, newUser domain_model.User) error {
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
        uc := usecase.NewUserUpdateUseCase(mockCrypto, mockRepo)
        _, err := uc.Execute(input)
        assert.Equal(t, http.StatusNotFound, err.StatusCode)
        assert.Equal(t, fmt.Sprintf("UserId %s Not Found", input.UserId.Value), err.Msg)
        assert.Error(t, err.Err)
    })

    t.Run("Failed: User(NewUserId) Already Exists.", func(t *testing.T) {
        mockRepo := &mock.MockUserRepository{
            FindUserByIdFunc: func(userId domain_model.UserId) (domain_model.User, error) {
                if userId.Value == "user01" { // Pass
                    return domain_model.User{
                        UserId: domain_model.UserId{Value: "user01"},
                        Email: domain_model.Email{Value: "user01@example.com"},
                        EncryptedPassword: domain_model.EncryptedPassword{Value: "???"},
                    }, nil
                } else {
                    return domain_model.User{ // Fail
                        UserId: domain_model.UserId{Value: "test01"},
                        Email: domain_model.Email{Value: "test01@example.com"},
                        EncryptedPassword: domain_model.EncryptedPassword{Value: "???"},
                    }, nil
                }
            },
            // Not Executed
            UpdateFunc: func(currentUserId domain_model.UserId, newUser domain_model.User) error {
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
        uc := usecase.NewUserUpdateUseCase(mockCrypto, mockRepo)
        _, err := uc.Execute(input)
        assert.Equal(t, http.StatusConflict, err.StatusCode)
        assert.Equal(t, "Input NewUserId Already Exists", err.Msg)
        assert.Error(t, err.Err)
    })

    t.Run("Failed: Password Decrypt Error", func(t *testing.T) {
        mockRepo := &mock.MockUserRepository{
            FindUserByIdFunc: func(userId domain_model.UserId) (domain_model.User, error) {
                if userId.Value == "user01" { // Pass
                    return domain_model.User{
                        UserId: domain_model.UserId{Value: "user01"},
                        Email: domain_model.Email{Value: "user01@example.com"},
                        EncryptedPassword: domain_model.EncryptedPassword{Value: "???"},
                    }, nil
                } else { // Pass
                    return domain_model.User{}, errors.New("Input NewUserId Not Found")
                }
            },
            // Not Executed
            UpdateFunc: func(currentUserId domain_model.UserId, newUser domain_model.User) error {
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
        uc := usecase.NewUserUpdateUseCase(mockCrypto, mockRepo)
        _, err := uc.Execute(input)
        assert.Equal(t, http.StatusInternalServerError, err.StatusCode)
        assert.Equal(t, "Decrypt Password Error", err.Msg)
        assert.Error(t, err.Err)
    })

    t.Run("Failed: Password Matches Error", func(t *testing.T) {
        mockRepo := &mock.MockUserRepository{
            FindUserByIdFunc: func(userId domain_model.UserId) (domain_model.User, error) {
                if userId.Value == "user01" { // Pass
                    return domain_model.User{
                        UserId: domain_model.UserId{Value: "user01"},
                        Email: domain_model.Email{Value: "user01@example.com"},
                        EncryptedPassword: domain_model.EncryptedPassword{Value: "???"},
                    }, nil
                } else { // Pass
                    return domain_model.User{}, errors.New("Input NewUserId Not Found")
                }
            },
            // Not Executed
            UpdateFunc: func(currentUserId domain_model.UserId, newUser domain_model.User) error {
                return nil
            },
        }
        mockCrypto := &mock.MockCryptoService{
            // pass
            DecryptFunc: func(encryptedPassword domain_model.EncryptedPassword) (domain_model.PlainTextPassword, error) {
                return domain_model.PlainTextPassword{Value: "???"}, nil
            },
            // fail
            MatchesFunc: func(value0 string, value1 string) bool {
                return false
            },
        }
        uc := usecase.NewUserUpdateUseCase(mockCrypto, mockRepo)
        _, err := uc.Execute(input)
        assert.Equal(t, http.StatusInternalServerError, err.StatusCode)
        assert.Equal(t, "Password Matches Error", err.Msg)
        assert.Error(t, err.Err)
    })

    t.Run("Failed: Update User failure", func(t *testing.T) {
        mockRepo := &mock.MockUserRepository{
            FindUserByIdFunc: func(userId domain_model.UserId) (domain_model.User, error) {
                if userId.Value == "user01" { // Pass
                    return domain_model.User{
                        UserId: domain_model.UserId{Value: "user01"},
                        Email: domain_model.Email{Value: "user01@example.com"},
                        EncryptedPassword: domain_model.EncryptedPassword{Value: "???"},
                    }, nil
                } else { // Pass
                    return domain_model.User{}, errors.New("Input NewUserId Not Found")
                }
            },
            // fail
            UpdateFunc: func(currentUserId domain_model.UserId, newUser domain_model.User) error {
                return errors.New("Update User failure")
            },
        }
        mockCrypto := &mock.MockCryptoService{
            // pass
            DecryptFunc: func(encryptedPassword domain_model.EncryptedPassword) (domain_model.PlainTextPassword, error) {
                return domain_model.PlainTextPassword{Value: "???"}, nil
            },
            // pass
            MatchesFunc: func(value0 string, value1 string) bool {
                return true
            },
        }
        uc := usecase.NewUserUpdateUseCase(mockCrypto, mockRepo)
        _, err := uc.Execute(input)
        assert.Equal(t, http.StatusInternalServerError, err.StatusCode)
        assert.Equal(t, "User Update Error", err.Msg)
        assert.Error(t, err.Err)
    })
}

