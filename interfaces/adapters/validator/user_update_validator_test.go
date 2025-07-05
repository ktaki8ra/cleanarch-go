package validator_test

import (
    "net/http"
    "testing"
    "github.com/stretchr/testify/assert"

    "github.com/ktaki8ra/cleanarch-go/interfaces/adapters/http/http_json"
    "github.com/ktaki8ra/cleanarch-go/interfaces/adapters/validator"
)

func TestValidateUserUpdateData(t *testing.T) {
    t.Run("Success: Valid InputData", func(t *testing.T) {
        input := http_json.UserUpdateRequestJson{
            UserId:            "user01",
            NewUserId:         "test01",
            PlainTextPassword: "password",
        }

        result, err := validator.ValidateUserUpdateData(input)

        assert.Equal(t, "", err.Msg)
        assert.Equal(t, input.UserId, result.UserId.Value)
        assert.Equal(t, input.NewUserId, result.NewUserId.Value)
        assert.Equal(t, input.PlainTextPassword, result.PlainTextPassword.Value)
    })

    t.Run("Failed: Invalid UserId", func(t *testing.T) {
        input := http_json.UserUpdateRequestJson{
            UserId:            "user 01",
            NewUserId:         "test01",
            PlainTextPassword: "password",
        }
    
        _, err := validator.ValidateUserUpdateData(input)
    
        assert.Equal(t, http.StatusUnprocessableEntity, err.StatusCode)
        assert.Equal(t, "Validation UserId Error", err.Msg)
        assert.Error(t, err.Err)
    })

    t.Run("Failed: Invalid New UserId", func(t *testing.T) {
        input := http_json.UserUpdateRequestJson{
            UserId:            "user01",
            NewUserId:         "test 01",
            PlainTextPassword: "password",
        }

        _, err := validator.ValidateUserUpdateData(input)

        assert.Equal(t, http.StatusUnprocessableEntity, err.StatusCode)
        assert.Equal(t, "Validation NewUserId Error", err.Msg)
        assert.Error(t, err.Err)
    })

    t.Run("Failed: Invalid Password", func(t *testing.T) {
        input := http_json.UserUpdateRequestJson{
            UserId:            "user01",
            NewUserId:         "test01",
            PlainTextPassword: "pass",
        }
    
        _, err := validator.ValidateUserUpdateData(input)
    
        assert.Equal(t, http.StatusUnprocessableEntity, err.StatusCode)
        assert.Equal(t, "Validation Password Error", err.Msg)
        assert.Error(t, err.Err)
    })
}

