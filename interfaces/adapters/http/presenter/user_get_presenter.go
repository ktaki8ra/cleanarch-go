package presenter

import (
    "net/http"

    "github.com/ktaki8ra/cleanarch-go/interfaces/adapters/http/http_json"
    "github.com/ktaki8ra/cleanarch-go/usecases/usecase"
)

func UserGetSuccessConvert(userGetOutputData usecase.UserGetOutputData) http_json.UserGetResponseJson {
    return http_json.UserGetResponseJson {
        StatusCode: http.StatusOK,
        UserId: userGetOutputData.UserId.Value,
        Email: userGetOutputData.Email.Value,
    }
}
