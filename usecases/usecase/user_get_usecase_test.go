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

func TestUserGetUseCase(t *testing.T) {
    input := usecase.UserGetInputData{
        UserId:            domain_model.UserId{Value: "user01"},
    }
    t.Run("Success", func(t *testing.T) {
        mockRepo := &mock.MockUserRepository{
            FindUserByIdFunc: func(userId domain_model.UserId) (domain_model.User, error) {
                return domain_model.User{
                    UserId: domain_model.UserId{Value: "user01"},
                    EncryptedPassword: domain_model.EncryptedPassword{Value: "???"},
                }, nil
            },
        }
        uc := usecase.NewUserGetUseCase(mockRepo)
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
        }
        uc := usecase.NewUserGetUseCase(mockRepo)
        _, err := uc.Execute(input)
        assert.Equal(t, http.StatusNotFound, err.StatusCode)
        assert.Equal(t, fmt.Sprintf("UserId %s Not Found", input.UserId.Value), err.Msg)
        assert.Error(t, err.Err)
    })
}

