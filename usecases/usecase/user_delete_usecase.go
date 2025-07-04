package usecase

import (
    "github.com/ktaki8ra/cleanarch-go/domain/domain_model"
    "github.com/ktaki8ra/cleanarch-go/domain/domain_repository"
    "github.com/ktaki8ra/cleanarch-go/usecases/service"

    "net/http"
    "fmt"
)

type UserDeleteInputData struct {
    UserId domain_model.UserId
    PlainTextPassword domain_model.PlainTextPassword
}
type UserDeleteOutputData struct {
    UserId domain_model.UserId
}

type UserDeleteUseCase struct {
    Cs service.CryptoService
    Ur domain_repository.UserRepository
}
func NewUserDeleteUseCase(
    cs service.CryptoService,
    ur domain_repository.UserRepository,
) *UserDeleteUseCase {
    return &UserDeleteUseCase {
        Cs: cs,
        Ur: ur,
    }
}
func (uc *UserDeleteUseCase) Execute(userDeleteInputData UserDeleteInputData) (UserDeleteOutputData, UseCaseError) {

    user, findUserError := uc.Ur.FindUserById(userDeleteInputData.UserId)
    if findUserError != nil {
        findUserErr := UseCaseError {
            StatusCode: http.StatusInternalServerError,
            Msg: fmt.Sprintf("UserId %s Not Exists", userDeleteInputData.UserId.Value),
            Err: fmt.Errorf("UserId %s Not Exists", userDeleteInputData.UserId.Value),
        }
        return UserDeleteOutputData{}, findUserErr
    }

    plainTextPassword, decryptPasswordError := uc.Cs.Decrypt(user.EncryptedPassword)
    if decryptPasswordError != nil {
        decryptPasswordErr := UseCaseError {
            StatusCode: http.StatusInternalServerError,
            Msg: "Decrypt Password Error",
            Err: decryptPasswordError,
        }
        return UserDeleteOutputData{}, decryptPasswordErr
    }

    if !uc.Cs.Matches(userDeleteInputData.PlainTextPassword.Value, plainTextPassword.Value) {
        passwordMatchesErr := UseCaseError {
            StatusCode: http.StatusInternalServerError,
            Msg: "Password Matches Error",
            Err: fmt.Errorf("Password Matches Error"),
        }
        return UserDeleteOutputData{}, passwordMatchesErr
    }

    userDeleteError := uc.Ur.Delete(user)
    if userDeleteError != nil {
        userDeleteErr := UseCaseError {
            StatusCode: http.StatusInternalServerError,
            Msg: "User Delete Error",
            Err: userDeleteError,
        }
        return UserDeleteOutputData{}, userDeleteErr
    }

    userDeleteOutputData := UserDeleteOutputData{UserId: userDeleteInputData.UserId}
    return userDeleteOutputData, UseCaseError{}
}

