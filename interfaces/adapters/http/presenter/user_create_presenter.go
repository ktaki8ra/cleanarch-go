package presenter

import (
    "net/http"

    "github.com/ktaki8ra/cleanarch-go/interfaces/adapters/http/http_json"
    "github.com/ktaki8ra/cleanarch-go/usecases/usecase"
)

func UserCreateSuccessConvert(userCreateOutputData usecase.UserCreateOutputData) http_json.UserCreateResponseJson {
    return http_json.UserCreateResponseJson {
        StatusCode: http.StatusOK,
        Message: "User Created.",
        UserId: userCreateOutputData.UserId.Value,
    }
}
