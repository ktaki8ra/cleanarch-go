package mock

// Defined Mock for UserRepository

import (
    "github.com/ktaki8ra/cleanarch-go/domain/domain_model"
)

type MockUserRepository struct {
    FindUserByIdFunc    func(domain_model.UserId) (domain_model.User, error)
    FindUserByEmailFunc func(domain_model.Email) (domain_model.User, error)
    SaveFunc            func(domain_model.User) error
    DeleteFunc          func(domain_model.User) error
    UpdateFunc          func(domain_model.UserId, domain_model.User) error
}

func (m *MockUserRepository) FindUserById(id domain_model.UserId) (domain_model.User, error) {
    return m.FindUserByIdFunc(id)
}
func (m *MockUserRepository) FindUserByEmail(email domain_model.Email) (domain_model.User, error) {
    return m.FindUserByEmailFunc(email)
}
func (m *MockUserRepository) Save(user domain_model.User) error {
    return m.SaveFunc(user)
}
func (m *MockUserRepository) Delete(user domain_model.User) error {
    return m.DeleteFunc(user)
}
func (m *MockUserRepository) Update(currentUserId domain_model.UserId, newUser domain_model.User) error {
    return m.UpdateFunc(currentUserId, newUser)
}
