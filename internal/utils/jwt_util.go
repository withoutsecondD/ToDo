package utils

import (
	"crypto/rand"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

var JwtKey []byte

type Claims struct {
	ID int64 `json:"id"`
	jwt.RegisteredClaims
}

func GenerateJwtToken(id int64) (string, error) {
	claims := &Claims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(JwtKey)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func ValidateJwtToken(c *fiber.Ctx) (*jwt.Token, error) {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return nil, errors.New("no authorization token provided")
	}

	tokenSlice := strings.Split(authHeader, "Bearer ")
	if len(tokenSlice) != 2 {
		return nil, errors.New("invalid token format")
	}
	tokenStr := tokenSlice[1]

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func GenerateJwtKey() error {
	JwtKey = make([]byte, 32)
	if _, err := rand.Read(JwtKey); err != nil {
		return err
	}

	return nil
}
