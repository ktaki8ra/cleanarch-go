package decoder

import (
    "github.com/ktaki8ra/cleanarch-go/interfaces/adapters/http/http_json"
    "github.com/ktaki8ra/cleanarch-go/interfaces/adapters/adapter_error"
    "github.com/labstack/echo/v4"

    "net/http"
    "fmt"
)

func DecodeUserGetRequestJson(c echo.Context) (http_json.UserGetRequestJson, adapter_error.DecodeJsonError) {
    var userGetRequestJson http_json.UserGetRequestJson
    err := c.Bind(&userGetRequestJson)
    if err != nil {
        decodeJsonError := adapter_error.DecodeJsonError {
            StatusCode: http.StatusBadRequest,
            Msg: "Invalid Request Payload",
            Err: fmt.Errorf("Invalid Request Payload"),
        }
        return http_json.UserGetRequestJson{}, decodeJsonError
    }
    return userGetRequestJson, adapter_error.DecodeJsonError{}
}
