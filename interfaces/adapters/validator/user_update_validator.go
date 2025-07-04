package validator

import (
    "net/http"

    "github.com/ktaki8ra/cleanarch-go/interfaces/adapters/http/http_json"
    "github.com/ktaki8ra/cleanarch-go/interfaces/adapters/adapter_error"
    "github.com/ktaki8ra/cleanarch-go/usecases/usecase"
    "github.com/ktaki8ra/cleanarch-go/domain/domain_model"
)

func ValidateUserUpdateData(
    userUpdateRequestJson http_json.UserUpdateRequestJson,
) (usecase.UserUpdateInputData, adapter_error.ValidationError) {

    userId, validateUserIdErr := domain_model.ValidateUserId(userUpdateRequestJson.UserId)
    if validateUserIdErr != nil {
        validationUserIdError := adapter_error.ValidationError {
            StatusCode: http.StatusUnprocessableEntity, // 422
            Msg: "Validation UserId Error",
            Err: validateUserIdErr,
        }
        return usecase.UserUpdateInputData{}, validationUserIdError
    }

    newUserId, validateNewUserIdErr := domain_model.ValidateUserId(userUpdateRequestJson.NewUserId)
    if validateNewUserIdErr != nil {
        validationNewUserIdError := adapter_error.ValidationError {
            StatusCode: http.StatusUnprocessableEntity, // 422
            Msg: "Validation NewUserId Error",
            Err: validateNewUserIdErr,
        }
        return usecase.UserUpdateInputData{}, validationNewUserIdError
    }

    plainTextPassword, validatePlainTextPassErr := domain_model.ValidatePlainTextPassword(userUpdateRequestJson.PlainTextPassword)
    if validatePlainTextPassErr != nil {
        validationPlainTextPassError := adapter_error.ValidationError {
            StatusCode: http.StatusUnprocessableEntity, // 422
            Msg: "Validation Password Error",
            Err: validatePlainTextPassErr,
        }
        return usecase.UserUpdateInputData{}, validationPlainTextPassError
    }

    userUpdateInputData := usecase.UserUpdateInputData{
        UserId: userId,
        NewUserId: newUserId,
        PlainTextPassword: plainTextPassword,
    }

    return userUpdateInputData, adapter_error.ValidationError{}
}
