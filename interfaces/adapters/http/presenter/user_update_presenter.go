package presenter

import (
    "net/http"

    "github.com/ktaki8ra/cleanarch-go/interfaces/adapters/http/http_json"
    "github.com/ktaki8ra/cleanarch-go/usecases/usecase"
)

func UserUpdateSuccessConvert(userUpdateOutputData usecase.UserUpdateOutputData) http_json.UserUpdateResponseJson {
    return http_json.UserUpdateResponseJson {
        StatusCode: http.StatusOK,
        Message: "User Updated.",
        UserId: userUpdateOutputData.UserId.Value,
    }
}
