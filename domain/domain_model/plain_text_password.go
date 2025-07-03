package domain_model

import (
    "errors"
)

type PlainTextPassword struct {
    Value string
}

func ValidatePlainTextPassword(value string) (PlainTextPassword, error) {
    if len(value) < 8 || len(value) > 16 {
        return PlainTextPassword{}, errors.New("Password must be between 8 and 16 characters")
    }
    return PlainTextPassword{Value: value}, nil
}
