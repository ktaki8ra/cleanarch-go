package domain_repository

import (
    "github.com/ktaki8ra/cleanarch-go/domain/domain_model"
)

type UserRepository interface {
    FindUserById(userId domain_model.UserId) (domain_model.User, error)
    FindUserByEmail(email domain_model.Email) (domain_model.User, error)
    Save(user domain_model.User) error
    Delete(user domain_model.User) error
    Update(currentUserId domain_model.UserId, newUser domain_model.User) error
}
