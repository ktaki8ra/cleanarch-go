package http_json

type UserCreateRequestJson struct {
    UserId string `json:"user_id"`
    Email string `json:"email"`
    PlainTextPassword string `json:"password"`
}
