package validator

import (
    "net/http"

    "github.com/ktaki8ra/cleanarch-go/interfaces/adapters/http/http_json"
    "github.com/ktaki8ra/cleanarch-go/interfaces/adapters/adapter_error"
    "github.com/ktaki8ra/cleanarch-go/usecases/usecase"
    "github.com/ktaki8ra/cleanarch-go/domain/domain_model"
)

func ValidateUserCreateData(
    userCreateRequestJson http_json.UserCreateRequestJson,
) (usecase.UserCreateInputData, adapter_error.ValidationError) {

    userId, validateUserIdErr := domain_model.ValidateUserId(userCreateRequestJson.UserId)
    if validateUserIdErr != nil {
        validationUserIdError := adapter_error.ValidationError {
            StatusCode: http.StatusUnprocessableEntity, // 422
            Msg: "Validation UserId Error",
            Err: validateUserIdErr,
        }
        return usecase.UserCreateInputData{}, validationUserIdError
    }

    email, validateEmailErr := domain_model.ValidateEmail(userCreateRequestJson.Email)
    if validateEmailErr != nil {
        validationEmailError := adapter_error.ValidationError {
            StatusCode: http.StatusUnprocessableEntity, // 422
            Msg: "Validation Email Error",
            Err: validateEmailErr,
        }
        return usecase.UserCreateInputData{}, validationEmailError
    }

    plainTextPassword, validatePlainTextPassErr := domain_model.ValidatePlainTextPassword(userCreateRequestJson.PlainTextPassword)
    if validatePlainTextPassErr != nil {
        validationPlainTextPassError := adapter_error.ValidationError {
            StatusCode: http.StatusUnprocessableEntity, // 422
            Msg: "Validation Password Error",
            Err: validatePlainTextPassErr,
        }
        return usecase.UserCreateInputData{}, validationPlainTextPassError
    }

    userCreateInputData := usecase.UserCreateInputData{
        UserId: userId,
        Email: email,
        PlainTextPassword: plainTextPassword,
    }

    return userCreateInputData, adapter_error.ValidationError{}

}
