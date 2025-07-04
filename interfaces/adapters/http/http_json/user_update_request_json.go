package http_json

type UserUpdateRequestJson struct {
    UserId string `json:"user_id"`
    NewUserId string `json:"new_user_id"`
    PlainTextPassword string `json:"password"`
}
