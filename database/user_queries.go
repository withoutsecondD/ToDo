package database

import "github.com/withoutsecondd/ToDo/models"

func GetAllUsers() ([]models.User, error) {
	users := make([]models.User, 0)

	err := DB.Select(&users, "SELECT * FROM withoutsecondd.user")
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserById(id int64) (models.User, error) {
	user := models.User{}

	err := DB.Get(&user, "SELECT * FROM withoutsecondd.user WHERE id = ?", id)
	if err != nil {
		return user, err
	}

	return user, nil
}
