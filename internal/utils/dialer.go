package utils

type Dialer interface {
	SendEmail(to string, subject string, body string) error
}
