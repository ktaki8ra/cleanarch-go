package validator

import (
    "net/http"

    "github.com/ktaki8ra/cleanarch-go/interfaces/adapters/http/http_json"
    "github.com/ktaki8ra/cleanarch-go/interfaces/adapters/adapter_error"
    "github.com/ktaki8ra/cleanarch-go/usecases/usecase"
    "github.com/ktaki8ra/cleanarch-go/domain/domain_model"
)

func ValidateUserGetData(
    userGetRequestJson http_json.UserGetRequestJson,
) (usecase.UserGetInputData, adapter_error.ValidationError) {

    userId, validateUserIdErr := domain_model.ValidateUserId(userGetRequestJson.UserId)
    if validateUserIdErr != nil {
        validationUserIdError := adapter_error.ValidationError {
            StatusCode: http.StatusUnprocessableEntity, // 422
            Msg: "Validation UserId Error",
            Err: validateUserIdErr,
        }
        return usecase.UserGetInputData{}, validationUserIdError
    }

    userGetInputData := usecase.UserGetInputData{
        UserId: userId,
    }

    return userGetInputData, adapter_error.ValidationError{}

}
