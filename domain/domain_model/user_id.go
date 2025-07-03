package domain_model

import (
    "errors"
    "regexp"
)

type UserId struct {
    Value string
}

func ValidateUserId(value string) (UserId, error) {
    if len(value) < 6 || len(value) > 16 {
        return UserId{}, errors.New("UserId must be between 6 and 16 characters")
    }
    validPattern := regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
    if !validPattern.MatchString(value) {
        return UserId{}, errors.New("UserId contains invalid characters")
    }
    return UserId{Value: value}, nil
}
