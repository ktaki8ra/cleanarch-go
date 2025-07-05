package validator_test

import (
    "net/http"
    "testing"
    "github.com/stretchr/testify/assert"

    "github.com/ktaki8ra/cleanarch-go/interfaces/adapters/http/http_json"
    "github.com/ktaki8ra/cleanarch-go/interfaces/adapters/validator"
)

func TestValidateUserDeleteData(t *testing.T) {
    t.Run("Success: Valid InputData", func(t *testing.T) {
        input := http_json.UserDeleteRequestJson{
            UserId:            "user01",
            PlainTextPassword: "password",
        }

        result, err := validator.ValidateUserDeleteData(input)

        assert.Equal(t, "", err.Msg)
        assert.Equal(t, input.UserId, result.UserId.Value)
        assert.Equal(t, input.PlainTextPassword, result.PlainTextPassword.Value)
    })

    t.Run("Failed: Invalid UserId", func(t *testing.T) {
        input := http_json.UserDeleteRequestJson{
            UserId:            "user 01",
            PlainTextPassword: "password",
        }
    
        _, err := validator.ValidateUserDeleteData(input)
    
        assert.Equal(t, http.StatusUnprocessableEntity, err.StatusCode)
        assert.Equal(t, "Validation UserId Error", err.Msg)
        assert.Error(t, err.Err)
    })

    t.Run("Failed: Invalid Password", func(t *testing.T) {
        input := http_json.UserDeleteRequestJson{
            UserId:            "user01",
            PlainTextPassword: "pass",
        }
    
        _, err := validator.ValidateUserDeleteData(input)
    
        assert.Equal(t, http.StatusUnprocessableEntity, err.StatusCode)
        assert.Equal(t, "Validation Password Error", err.Msg)
        assert.Error(t, err.Err)
    })
}

