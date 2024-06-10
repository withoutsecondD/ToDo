package database

import (
	"github.com/withoutsecondd/ToDo/models"
)

func GetListsByUserId(id int64) ([]models.List, error) {
	lists := make([]models.List, 0)

	err := DB.Select(&lists, "SELECT * FROM withoutsecondd.list WHERE user_id = ?", id)
	if err != nil {
		return nil, err
	}

	return lists, nil
}
