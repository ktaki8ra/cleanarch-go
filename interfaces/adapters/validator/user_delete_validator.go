package validator

import (
    "net/http"

    "github.com/ktaki8ra/cleanarch-go/interfaces/adapters/http/http_json"
    "github.com/ktaki8ra/cleanarch-go/interfaces/adapters/adapter_error"
    "github.com/ktaki8ra/cleanarch-go/usecases/usecase"
    "github.com/ktaki8ra/cleanarch-go/domain/domain_model"
)

func ValidateUserDeleteData(
    userDeleteRequestJson http_json.UserDeleteRequestJson,
) (usecase.UserDeleteInputData, adapter_error.ValidationError) {

    userId, validateUserIdErr := domain_model.ValidateUserId(userDeleteRequestJson.UserId)
    if validateUserIdErr != nil {
        validationUserIdError := adapter_error.ValidationError {
            StatusCode: http.StatusUnprocessableEntity, // 422
            Msg: "Validation UserId Error",
            Err: validateUserIdErr,
        }
        return usecase.UserDeleteInputData{}, validationUserIdError
    }

    plainTextPassword, validatePlainTextPassErr := domain_model.ValidatePlainTextPassword(userDeleteRequestJson.PlainTextPassword)
    if validatePlainTextPassErr != nil {
        validationPlainTextPassError := adapter_error.ValidationError {
            StatusCode: http.StatusUnprocessableEntity, // 422
            Msg: "Validation Password Error",
            Err: validatePlainTextPassErr,
        }
        return usecase.UserDeleteInputData{}, validationPlainTextPassError
    }

    userDeleteInputData := usecase.UserDeleteInputData{
        UserId: userId,
        PlainTextPassword: plainTextPassword,
    }

    return userDeleteInputData, adapter_error.ValidationError{}

}
