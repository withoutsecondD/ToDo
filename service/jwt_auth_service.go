package service

import (
	"crypto/rand"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/withoutsecondd/ToDo/database"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type JwtAuthService struct {
	db     database.Database
	jwtKey []byte
}

func generateJwtKey() ([]byte, error) {
	jwtKey := make([]byte, 32)
	if _, err := rand.Read(jwtKey); err != nil {
		return nil, err
	}

	//return jwtKey, nil
	return []byte{22, 112, 222, 0, 209, 146, 167, 212, 158, 39, 193, 131, 191, 67, 190, 52, 15, 170, 254, 43, 6, 5, 3, 175, 134, 227, 118, 82, 6, 243, 98, 111}, nil
}

func NewJwtAuthService(db database.Database) (*JwtAuthService, error) {
	jwtKey, err := generateJwtKey()
	if err != nil {
		return nil, err
	}
	return &JwtAuthService{db: db, jwtKey: jwtKey}, nil
}

type Claims struct {
	ID int64 `json:"id"`
	jwt.RegisteredClaims
}

func (ja *JwtAuthService) checkCredentials(l *LoginRequest) error {
	hashedPassword, err := ja.db.GetUserPasswordByEmail(l.Email)
	if err != nil {
		return errors.New("no user is found with such email")
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(l.Password))
	if err != nil {
		return errors.New("incorrect email or password")
	}

	return nil
}

// Authenticate returns nil if user with such credentials is present, error otherwise
func (ja *JwtAuthService) Authenticate(l *LoginRequest) (string, error) {
	if err := ja.checkCredentials(l); err != nil {
		return "", err
	}

	user, err := ja.db.GetUserByEmail(l.Email)
	if err != nil {
		return "", err
	}

	claims := &Claims{
		ID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(ja.jwtKey)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func (ja *JwtAuthService) AuthorizeWithToken(tokenStr string) (int64, error) {
	token, err := ja.ValidateJwtToken(tokenStr)
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errors.New("error getting token claims")
	}

	return int64(claims["id"].(float64)), nil
}

func (ja *JwtAuthService) ValidateJwtToken(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return ja.jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}
