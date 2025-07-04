package decoder

import (
    "github.com/ktaki8ra/cleanarch-go/interfaces/adapters/http/http_json"
    "github.com/ktaki8ra/cleanarch-go/interfaces/adapters/adapter_error"
    "github.com/labstack/echo/v4"

    "net/http"
    "fmt"
)

func DecodeUserUpdateRequestJson(c echo.Context) (http_json.UserUpdateRequestJson, adapter_error.DecodeJsonError) {
    var userUpdateRequestJson http_json.UserUpdateRequestJson
    err := c.Bind(&userUpdateRequestJson)
    if err != nil {
        decodeJsonError := adapter_error.DecodeJsonError {
            StatusCode: http.StatusBadRequest,
            Msg: "Invalid Request Payload",
            Err: fmt.Errorf("Invalid Request Payload"),
        }
        return http_json.UserUpdateRequestJson{}, decodeJsonError
    }
    return userUpdateRequestJson, adapter_error.DecodeJsonError{}
}
