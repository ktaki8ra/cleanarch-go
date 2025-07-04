package http_json

type UserGetResponseJson struct {
    StatusCode int `json:"status_code"`
    UserId string `json:"user_id"`
    Email string `json:"email"`
}
