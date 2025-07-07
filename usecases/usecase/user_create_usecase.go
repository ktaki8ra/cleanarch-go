package usecase

import (
    "github.com/ktaki8ra/cleanarch-go/domain/domain_model"
    "github.com/ktaki8ra/cleanarch-go/domain/domain_repository"
    "github.com/ktaki8ra/cleanarch-go/usecases/service"

    "net/http"
    "fmt"
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

    _, findUserByIdError := uc.Ur.FindUserById(userCreateInputData.UserId)
    if findUserByIdError == nil {
        findUserByIdErr := UseCaseError {
            StatusCode: http.StatusConflict,
            Msg: "Input UserId Already Exists",
            Err: fmt.Errorf("Input UserId Already Exists"),
        }
        return UserCreateOutputData{}, findUserByIdErr
    }

    _, findUserByEmailError := uc.Ur.FindUserByEmail(userCreateInputData.Email)
    if findUserByEmailError == nil {
        findUserByEmailErr := UseCaseError {
            StatusCode: http.StatusConflict,
            Msg: "Input Email Already Exists",
            Err: fmt.Errorf("Input Email Already Exists"),
        }
        return UserCreateOutputData{}, findUserByEmailErr
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

    newUser := domain_model.GenerateUser(userCreateInputData.UserId, userCreateInputData.Email, encryptedPassword)

    userCreateError := uc.Ur.Save(newUser)
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

