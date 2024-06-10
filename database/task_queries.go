package database

import (
	"github.com/withoutsecondd/ToDo/models"
)

func GetTasksByListId(listId int64) ([]models.Task, error) {
	tasks := make([]models.Task, 0)

	err := DB.Select(&tasks, "SELECT * FROM withoutsecondd.task WHERE list_id = ?", listId)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
