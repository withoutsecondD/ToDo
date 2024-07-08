package service

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthService interface {
	Authenticate(l *LoginRequest) (string, error)
	AuthorizeWithToken(tokenStr string) (int64, error)
}
