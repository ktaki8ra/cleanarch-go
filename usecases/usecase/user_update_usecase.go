package usecase

import (
    "github.com/ktaki8ra/cleanarch-go/domain/domain_model"
    "github.com/ktaki8ra/cleanarch-go/domain/domain_repository"
    "github.com/ktaki8ra/cleanarch-go/usecases/service"

    "net/http"
    "fmt"
)

type UserUpdateInputData struct {
    UserId domain_model.UserId
    NewUserId domain_model.UserId
    PlainTextPassword domain_model.PlainTextPassword
}
type UserUpdateOutputData struct {
    UserId domain_model.UserId
}

type UserUpdateUseCase struct {
    Cs service.CryptoService
    Ur domain_repository.UserRepository
}
func NewUserUpdateUseCase(
    cs service.CryptoService,
    ur domain_repository.UserRepository,
) *UserUpdateUseCase {
    return &UserUpdateUseCase {
        Cs: cs,
        Ur: ur,
    }
}
func (uc *UserUpdateUseCase) Execute(userUpdateInputData UserUpdateInputData) (UserUpdateOutputData, UseCaseError) {

    currentUser, findUserError := uc.Ur.FindUserById(userUpdateInputData.UserId)
    if findUserError != nil {
        findUserErr := UseCaseError {
            StatusCode: http.StatusInternalServerError,
            Msg: fmt.Sprintf("UserId %s Not Exists", userUpdateInputData.UserId.Value),
            Err: fmt.Errorf("UserId %s Not Exists", userUpdateInputData.UserId.Value),
        }
        return UserUpdateOutputData{}, findUserErr
    }

    _, findNewUserError := uc.Ur.FindUserById(userUpdateInputData.NewUserId)
    if findNewUserError == nil {
        newUserIdAlreadyExists := UseCaseError {
            StatusCode: http.StatusConflict,
            Msg: "Input NewUserId Already Exists",
            Err: fmt.Errorf("Input NewUserId Already Exists"),
        }
        return UserUpdateOutputData{}, newUserIdAlreadyExists
    }

    plainTextPassword, decryptPasswordError := uc.Cs.Decrypt(currentUser.EncryptedPassword)
    if decryptPasswordError != nil {
        decryptPasswordErr := UseCaseError {
            StatusCode: http.StatusInternalServerError,
            Msg: "Decrypt Password Error",
            Err: decryptPasswordError,
        }
        return UserUpdateOutputData{}, decryptPasswordErr
    }

    if !uc.Cs.Matches(userUpdateInputData.PlainTextPassword.Value, plainTextPassword.Value) {
        passwordMatchesErr := UseCaseError {
            StatusCode: http.StatusInternalServerError,
            Msg: "Password Matches Error",
            Err: fmt.Errorf("Password Matches Error"),
        }
        return UserUpdateOutputData{}, passwordMatchesErr
    }

    newUser := domain_model.GenerateUser(userUpdateInputData.NewUserId, currentUser.Email, currentUser.EncryptedPassword)

    userUpdateError := uc.Ur.Update(currentUser.UserId, newUser)
    if userUpdateError != nil {
        userUpdateErr := UseCaseError {
            StatusCode: http.StatusInternalServerError,
            Msg: "User Update Error",
            Err: userUpdateError,
        }
        return UserUpdateOutputData{}, userUpdateErr
    }

    userUpdateOutputData := UserUpdateOutputData{UserId: newUser.UserId}
    return userUpdateOutputData, UseCaseError{}
}

