package utils

import (
	"gopkg.in/gomail.v2"
)

type DefaultDialer struct {
	from     string
	host     string
	port     int
	password string
}

func NewDefaultDialer(from string, host string, port int, password string) *DefaultDialer {
	return &DefaultDialer{
		from:     from,
		host:     host,
		port:     port,
		password: password,
	}
}

func (d *DefaultDialer) SendEmail(to string, subject string, body string) error {
	sender := gomail.NewDialer(d.host, d.port, d.from, d.password)

	msg := gomail.NewMessage()
	msg.SetHeaders(map[string][]string{
		"From":    {d.from},
		"To":      {to},
		"Subject": {subject},
	})
	msg.SetBody("text/html", body)

	err := sender.DialAndSend(msg)
	if err != nil {
		return err
	}

	return nil
}
