package usecase

import (
    "github.com/ktaki8ra/cleanarch-go/domain/domain_model"
    "github.com/ktaki8ra/cleanarch-go/domain/domain_repository"
    "github.com/ktaki8ra/cleanarch-go/usecases/service"

    "net/http"
)

type UserCreateInputData struct {
    UserId domain_model.UserId
    Email domain_model.Email
    PlainTextPassword domain_model.PlainTextPassword
}
type UserCreateOutputData struct {
    UserId domain_model.UserId
}

type UserCreateUseCase struct {
    Cs service.CryptoService
    Ur domain_repository.UserRepository
}
func NewUserCreateUseCase(
    cs service.CryptoService,
    ur domain_repository.UserRepository,
) *UserCreateUseCase {
    return &UserCreateUseCase {
        Cs: cs,
        Ur: ur,
    }
}
func (uc *UserCreateUseCase) Execute(userCreateInputData UserCreateInputData) (UserCreateOutputData, UseCaseError) {

    _, findUserError := uc.Ur.FindUserById(userCreateInputData.UserId)
    if findUserError != nil {
        findUserErr := UseCaseError {
            StatusCode: http.StatusInternalServerError,
            Msg: "User Already Exists",
            Err: findUserError,
        }
        return UserCreateOutputData{}, findUserErr
    }

    encryptedPassword, encryptPasswordError := uc.Cs.Encrypt(userCreateInputData.PlainTextPassword)
    if encryptPasswordError != nil {
        encryptPasswordErr := UseCaseError {
            StatusCode: http.StatusInternalServerError,
            Msg: "Encrypt Password Error",
            Err: encryptPasswordError,
        }
        return UserCreateOutputData{}, encryptPasswordErr
    }

    user := domain_model.GenerateUser(userCreateInputData.UserId, userCreateInputData.Email, encryptedPassword)

    userCreateError := uc.Ur.Save(user)
    if userCreateError != nil {
        userCreateErr := UseCaseError {
            StatusCode: http.StatusInternalServerError,
            Msg: "User Create Error",
            Err: userCreateError,
        }
        return UserCreateOutputData{}, userCreateErr
    }

    userCreateOutputData := UserCreateOutputData{UserId: userCreateInputData.UserId}
    return userCreateOutputData, UseCaseError{}
}

