package presenter

import (
    "net/http"

    "github.com/ktaki8ra/cleanarch-go/interfaces/adapters/http/http_json"
    "github.com/ktaki8ra/cleanarch-go/usecases/usecase"
)

func UserDeleteSuccessConvert(userDeleteOutputData usecase.UserDeleteOutputData) http_json.UserDeleteResponseJson {
    return http_json.UserDeleteResponseJson {
        StatusCode: http.StatusOK,
        Message: "User Deleted.",
        UserId: userDeleteOutputData.UserId.Value,
    }
}
