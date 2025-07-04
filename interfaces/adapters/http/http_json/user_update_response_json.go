package http_json

type UserUpdateResponseJson struct {
    StatusCode int `json:"status_code"`
    Message string `json:"message"`
    UserId string `json:"user_id"`
}
