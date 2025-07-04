package usecase

import (
    "github.com/ktaki8ra/cleanarch-go/domain/domain_model"
    "github.com/ktaki8ra/cleanarch-go/domain/domain_repository"

    "net/http"
    "fmt"
)

type UserGetInputData struct {
    UserId domain_model.UserId
}
type UserGetOutputData struct {
    UserId domain_model.UserId
    Email domain_model.Email
}

type UserGetUseCase struct {
    Ur domain_repository.UserRepository
}
func NewUserGetUseCase(
    ur domain_repository.UserRepository,
) *UserGetUseCase {
    return &UserGetUseCase {
        Ur: ur,
    }
}
func (uc *UserGetUseCase) Execute(userGetInputData UserGetInputData) (UserGetOutputData, UseCaseError) {

    user, findUserError := uc.Ur.FindUserById(userGetInputData.UserId)
    if findUserError != nil {
        findUserErr := UseCaseError {
            StatusCode: http.StatusInternalServerError,
            Msg: fmt.Sprintf("UserId %s Not Exists", userGetInputData.UserId.Value),
            Err: fmt.Errorf("UserId %s Not Exists", userGetInputData.UserId.Value),
        }
        return UserGetOutputData{}, findUserErr
    }

    userGetOutputData := UserGetOutputData{UserId: user.UserId, Email: user.Email}
    return userGetOutputData, UseCaseError{}
}

