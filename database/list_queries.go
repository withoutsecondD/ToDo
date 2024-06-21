package database

import (
	"github.com/withoutsecondd/ToDo/models"
)

func GetListById(id int64) (models.List, error) {
	var list models.List

	err := DB.Get(&list, "SELECT * FROM withoutsecondd.list WHERE id = ?", id)
	if err != nil {
		return models.List{}, err
	}

	return list, nil
}

func GetListsByUserId(id int64) ([]models.List, error) {
	lists := make([]models.List, 0)

	err := DB.Select(&lists, "SELECT * FROM withoutsecondd.list WHERE user_id = ?", id)
	if err != nil {
		return nil, err
	}

	return lists, nil
}
