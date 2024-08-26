package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/withoutsecondd/ToDo/database"
	"github.com/withoutsecondd/ToDo/internal/utils"
	"github.com/withoutsecondd/ToDo/models"
	"time"
)

type DefaultEmailService struct {
	db     database.Database
	d      utils.Dialer
	jwtKey []byte
}

func NewDefaultEmailService(db database.Database, d utils.Dialer, jK []byte) *DefaultEmailService {
	return &DefaultEmailService{db: db, d: d, jwtKey: jK}
}

func (s *DefaultEmailService) VerifyEmail(userId int64, token string) error {
	user, err := s.db.GetUserById(userId)
	if err != nil {
		return err
	}

	if token == "" {
		claims := &Claims{
			ID: 0,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
		}

		generatedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenStr, err := generatedToken.SignedString(s.jwtKey)
		if err != nil {
			return err
		}

		body := fmt.Sprintf(`
			<p>You're trying to create an account in our ToDo application,
			please click the link below to confirm your email:</p>
			<a href="http://localhost:8080/api/emails/verify?t=%s">Click!</a>
			<p>Please don't reply to this message</p>
		`, tokenStr)

		err = s.d.SendEmail(user.Email, "Email verification", body)
		if err != nil {
			return err
		}
	} else {
		_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return s.jwtKey, nil
		})
		if err != nil {
			return err
		}

		_, err = s.db.UpdateUserEmailStatus(&models.UserEmailStatusDto{
			ID:            userId,
			EmailVerified: true,
		})
		if err != nil {
			return err
		}

		return nil
	}

	return nil
}
