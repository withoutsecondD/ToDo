package utils

import (
	"crypto/rand"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtKey []byte

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJwtToken(email string) (string, error) {
	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func GenerateJwtKey() error {
	jwtKey = make([]byte, 32)
	if _, err := rand.Read(jwtKey); err != nil {
		return err
	}

	return nil
}
