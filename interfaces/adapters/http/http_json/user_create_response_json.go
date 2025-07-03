package http_json

type UserCreateResponseJson struct {
    StatusCode int `json:"status_code"`
    Message string `json:"message"`
    UserId string `json:"user_id"`
}
