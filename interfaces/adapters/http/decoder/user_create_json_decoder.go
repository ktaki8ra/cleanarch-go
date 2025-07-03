package decoder

import (
    "github.com/ktaki8ra/cleanarch-go/interfaces/adapters/http/http_json"
    "github.com/ktaki8ra/cleanarch-go/interfaces/adapters/adapter_error"
    "github.com/labstack/echo/v4"

    "net/http"
    "fmt"
)

func DecodeUserCreateRequestJson(c echo.Context) (http_json.UserCreateRequestJson, adapter_error.DecodeJsonError) {
    var userCreateRequestJson http_json.UserCreateRequestJson
    err := c.Bind(&userCreateRequestJson)
    if err != nil {
        decodeJsonError := adapter_error.DecodeJsonError {
            StatusCode: http.StatusBadRequest,
            Msg: "Invalid Request Payload",
            Err: fmt.Errorf("Invalid Request Payload"),
        }
        return http_json.UserCreateRequestJson{}, decodeJsonError
    }
    return userCreateRequestJson, adapter_error.DecodeJsonError{}
}
