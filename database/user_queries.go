package database

import (
	"errors"
	"github.com/go-sql-driver/mysql"
	"github.com/withoutsecondd/ToDo/internal/utils"
	"github.com/withoutsecondd/ToDo/models"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

func GetUserById(id int64) (*models.UserResponse, error) {
	user := &models.UserResponse{}

	err := DB.Get(user, "SELECT id, age, first_name, last_name, city, email FROM withoutsecondd.user WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByEmail(email string) (*models.UserResponse, error) {
	user := &models.UserResponse{}

	err := DB.Get(user, "SELECT id, age, first_name, last_name, city, email FROM withoutsecondd.user WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserPasswordByEmail(email string) ([]byte, error) {
	var password []byte

	err := DB.Get(&password, "SELECT password FROM withoutsecondd.user WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	return password, nil
}

func CreateUser(user *models.User) (*models.UserResponse, error) {
	err := utils.ValidateStruct(user)
	if err != nil {
		return nil, errors.New("refused to create a user: some fields are invalid")
	}

	query := `
		INSERT INTO withoutsecondd.user(age, first_name, last_name, city, email, password) 
		VALUES(?, ?, ?, ?, ?, ?)
	`

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	_, err = DB.Queryx(query, user.Age, user.FirstName, user.LastName, user.City, strings.ToLower(user.Email), hashedPassword)
	if err != nil {
		var sqlErr *mysql.MySQLError
		if errors.As(err, &sqlErr) && sqlErr.Number == 1062 {
			return nil, errors.New("user with such email already exists")
		}
	}

	createdUser, err := GetUserByEmail(strings.ToLower(user.Email))
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
