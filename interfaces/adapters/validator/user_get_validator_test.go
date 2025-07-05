package validator_test

import (
    "net/http"
    "testing"
    "github.com/stretchr/testify/assert"

    "github.com/ktaki8ra/cleanarch-go/interfaces/adapters/http/http_json"
    "github.com/ktaki8ra/cleanarch-go/interfaces/adapters/validator"
)

func TestValidateUserGetData(t *testing.T) {
    t.Run("Success: Valid InputData", func(t *testing.T) {
        input := http_json.UserGetRequestJson{
            UserId:            "user01",
        }

        result, err := validator.ValidateUserGetData(input)

        assert.Equal(t, "", err.Msg)
        assert.Equal(t, input.UserId, result.UserId.Value)
    })

    t.Run("Failed: Invalid UserId", func(t *testing.T) {
        input := http_json.UserGetRequestJson{
            UserId:            "user 01",
        }
    
        _, err := validator.ValidateUserGetData(input)
    
        assert.Equal(t, http.StatusUnprocessableEntity, err.StatusCode)
        assert.Equal(t, "Validation UserId Error", err.Msg)
        assert.Error(t, err.Err)
    })
}

