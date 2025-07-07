package usecase_test

import (
    "errors"
    "testing"
    "github.com/ktaki8ra/cleanarch-go/mock"

    "github.com/stretchr/testify/assert"
    "github.com/ktaki8ra/cleanarch-go/domain/domain_model"
    "github.com/ktaki8ra/cleanarch-go/usecases/usecase"
)

func TestUserUpdateUseCase(t *testing.T) {
    t.Run("Success", func(t *testing.T) {
        input := usecase.UserUpdateInputData{
            UserId:            domain_model.UserId{Value: "user01"},
            NewUserId:         domain_model.UserId{Value: "test01"},
            PlainTextPassword: domain_model.PlainTextPassword{Value: "???"},
        }
        mockRepo := &mock.MockUserRepository{
            FindUserByIdFunc: func(userId domain_model.UserId) (domain_model.User, error) {
                if userId.Value == "user01" {
                    return domain_model.User{
                        UserId: domain_model.UserId{Value: "user01"},
                        Email: domain_model.Email{Value: "user01@example.com"},
                        EncryptedPassword: domain_model.EncryptedPassword{Value: "???"},
                    }, nil
                } else {
                    return domain_model.User{}, errors.New("Input NewUserId already Exists Error")
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

}

