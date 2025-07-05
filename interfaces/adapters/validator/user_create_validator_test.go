package validator_test

import (
    "net/http"
    "testing"
    "github.com/stretchr/testify/assert"

    "github.com/ktaki8ra/cleanarch-go/interfaces/adapters/http/http_json"
    "github.com/ktaki8ra/cleanarch-go/interfaces/adapters/validator"
)

func TestValidateUserCreateData(t *testing.T) {
    t.Run("Success: Valid InputData", func(t *testing.T) {
        input := http_json.UserCreateRequestJson{
            UserId:            "user01",
            Email:             "user01@example.com",
            PlainTextPassword: "password",
        }

        result, err := validator.ValidateUserCreateData(input)

        assert.Equal(t, "", err.Msg)
        assert.Equal(t, input.UserId, result.UserId.Value)
        assert.Equal(t, input.Email, result.Email.Value)
        assert.Equal(t, input.PlainTextPassword, result.PlainTextPassword.Value)
    })

    t.Run("Failed: Invalid UserId", func(t *testing.T) {
        input := http_json.UserCreateRequestJson{
            UserId:            "user 01",
            Email:             "user01@example.com",
            PlainTextPassword: "password",
        }
    
        _, err := validator.ValidateUserCreateData(input)
    
        assert.Equal(t, http.StatusUnprocessableEntity, err.StatusCode)
        assert.Equal(t, "Validation UserId Error", err.Msg)
        assert.Error(t, err.Err)
    })

    t.Run("Failed: Invalid Email", func(t *testing.T) {
        input := http_json.UserCreateRequestJson{
            UserId:            "user01",
            Email:             "user01example.com",
            PlainTextPassword: "password",
        }

        _, err := validator.ValidateUserCreateData(input)

        assert.Equal(t, http.StatusUnprocessableEntity, err.StatusCode)
        assert.Equal(t, "Validation Email Error", err.Msg)
        assert.Error(t, err.Err)
    })

    t.Run("Failed: Invalid Password", func(t *testing.T) {
        input := http_json.UserCreateRequestJson{
            UserId:            "user01",
            Email:             "user01d@example.com",
            PlainTextPassword: "pass",
        }
    
        _, err := validator.ValidateUserCreateData(input)
    
        assert.Equal(t, http.StatusUnprocessableEntity, err.StatusCode)
        assert.Equal(t, "Validation Password Error", err.Msg)
        assert.Error(t, err.Err)
    })
}

