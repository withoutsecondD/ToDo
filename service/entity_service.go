package service

import "github.com/withoutsecondd/ToDo/models"

type EntityService interface {
	GetUserById(id int64) (*models.UserResponse, error)
	GetUserByEmail(email string) (*models.UserResponse, error)
	GetUserPasswordByEmail(email string) ([]byte, error)
	CreateUser(user *models.User) (*models.UserResponse, error)

	GetListById(listId int64, userId int64) (*models.List, error)
	GetListsByUserId(id int64) ([]models.List, error)

	GetTasksByUserId(userId int64) ([]models.Task, error)
	GetTasksByListId(listId int64, userId int64) ([]models.Task, error)
	GetTaskById(taskId int64, userId int64) (*models.Task, error)
}
