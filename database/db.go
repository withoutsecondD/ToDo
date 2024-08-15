package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/withoutsecondd/ToDo/models"
)

type Database interface {
	GetUserById(id int64) (*models.UserResponse, error)
	GetUserByEmail(email string) (*models.UserResponse, error)
	GetUserPasswordByEmail(email string) ([]byte, error)
	CreateUser(user *models.User) (*models.UserResponse, error)

	GetListById(id int64) (*models.List, error)
	GetListsByUserId(userId int64) ([]models.List, error)
	CreateList(list *models.List) (*models.List, error)

	GetTaskById(id int64) (*models.Task, error)
	GetTasksByUserId(userId int64) ([]models.Task, error)
	GetTasksByListId(listId int64) ([]models.Task, error)
	CreateTask(task *models.Task) (*models.Task, error)

	GetTagById(id int64) (*models.Tag, error)
	GetTagsByUserId(userId int64) ([]models.Tag, error)
	GetTagsByTaskId(taskId int64) ([]models.Tag, error)
	CreateTag(tag *models.Tag) (*models.Tag, error)
}
