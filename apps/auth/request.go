package auth

type RegisterRequestPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
	Role     string `json:"role"`
}

type LoginRequestPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
