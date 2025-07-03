package http_json

type ErrorResponseJson struct {
    StatusCode int `json:"statusCode"`
    Msg string `json:"msg"`
}
