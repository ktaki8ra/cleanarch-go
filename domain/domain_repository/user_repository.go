package domain_repository

import (
    "github.com/ktaki8ra/cleanarch-go/domain/domain_model"
)

type UserRepository interface {
    FindUserById(userId domain_model.UserId) (domain_model.User, error)
    Save(user domain_model.User) error
}
