package domain_model

import (
    "strings"
    "errors"
)

type Email struct {
    Value string
}

func ValidateEmail(value string) (Email, error) {
    if !strings.Contains(value, "@") {
        return Email{}, errors.New("value must contain '@' symbol")
    }
    return Email{Value: value}, nil
}
