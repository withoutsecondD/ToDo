package database

import (
	"github.com/withoutsecondd/ToDo/models"
)

func GetUserById(id int64) (models.UserResponse, error) {
	user := models.UserResponse{}

	err := DB.Get(&user, "SELECT id, age, first_name, last_name, city, email FROM withoutsecondd.user WHERE id = ?", id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func GetUserByEmail(email string) (models.UserResponse, error) {
	user := models.UserResponse{}

	err := DB.Get(&user, "SELECT id, age, first_name, last_name, city, email FROM withoutsecondd.user WHERE email = ?", email)
	if err != nil {
		return user, err
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
