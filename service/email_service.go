package service

type EmailService interface {
	VerifyEmail(userId int64, token string) error
}
