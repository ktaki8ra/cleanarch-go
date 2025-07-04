package decoder

import (
    "github.com/ktaki8ra/cleanarch-go/interfaces/adapters/http/http_json"
    "github.com/ktaki8ra/cleanarch-go/interfaces/adapters/adapter_error"
    "github.com/labstack/echo/v4"

    "net/http"
    "fmt"
)

func DecodeUserDeleteRequestJson(c echo.Context) (http_json.UserDeleteRequestJson, adapter_error.DecodeJsonError) {
    var userDeleteRequestJson http_json.UserDeleteRequestJson
    err := c.Bind(&userDeleteRequestJson)
    if err != nil {
        decodeJsonError := adapter_error.DecodeJsonError {
            StatusCode: http.StatusBadRequest,
            Msg: "Invalid Request Payload",
            Err: fmt.Errorf("Invalid Request Payload"),
        }
        return http_json.UserDeleteRequestJson{}, decodeJsonError
    }
    return userDeleteRequestJson, adapter_error.DecodeJsonError{}
}
