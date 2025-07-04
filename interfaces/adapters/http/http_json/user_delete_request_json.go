package http_json

type UserDeleteRequestJson struct {
    UserId string `json:"user_id"`
    PlainTextPassword string `json:"password"`
}
