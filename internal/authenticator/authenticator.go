package authenticator

import (
	"errors"
	"github.com/withoutsecondd/ToDo/database"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Authenticate returns nil if user with such credentials is present, error otherwise
func Authenticate(l *LoginRequest) error {
	hashedPassword, err := database.GetUserPasswordByEmail(l.Email)
	if err != nil {
		return errors.New("no user is found with such email")
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(l.Password))
	if err != nil {
		return errors.New("incorrect email or password")
	}

	return nil
}
