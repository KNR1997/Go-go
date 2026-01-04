package auth

type RegisterParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
