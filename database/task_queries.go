package database

import (
	"github.com/withoutsecondd/ToDo/models"
)

func GetTasksByUserId(userId int64) ([]models.Task, error) {
	tasks := make([]models.Task, 0)

	query := `
		SELECT l.id, l.title, t.id, t.list_id, t.title FROM
			(SELECT * FROM withoutsecondd.list WHERE user_id = ?) AS l
		INNER JOIN withoutsecondd.task AS t ON t.list_id = l.id
		ORDER BY t.id;
	`

	err := DB.Select(&tasks, query, userId)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func GetTasksByListId(listId int64) ([]models.Task, error) {
	tasks := make([]models.Task, 0)

	err := DB.Select(&tasks, "SELECT * FROM withoutsecondd.task WHERE list_id = ?", listId)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
