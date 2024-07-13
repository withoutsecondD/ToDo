package service

import (
	"errors"
	"github.com/withoutsecondd/ToDo/database"
	"github.com/withoutsecondd/ToDo/internal/utils"
	"github.com/withoutsecondd/ToDo/models"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type DefaultEntityService struct {
	db database.Database
	v  utils.Validator
}

func NewDefaultEntityService(d database.Database, v utils.Validator) *DefaultEntityService {
	return &DefaultEntityService{db: d, v: v}
}

func (m *DefaultEntityService) GetUserById(id int64) (*models.UserResponse, error) {
	return m.db.GetUserById(id)
}

func (m *DefaultEntityService) GetUserByEmail(email string) (*models.UserResponse, error) {
	return m.db.GetUserByEmail(email)
}

func (m *DefaultEntityService) GetUserPasswordByEmail(email string) ([]byte, error) {
	return m.db.GetUserPasswordByEmail(email)
}

func (m *DefaultEntityService) CreateUser(user *models.User) (*models.UserResponse, error) {
	err := m.v.ValidateStruct(user)
	if err != nil {
		return nil, errors.New("refused to create a user: some fields are invalid")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)
	user.Email = strings.ToLower(user.Email)

	return m.db.CreateUser(user)
}

func (m *DefaultEntityService) GetListById(listId int64, userId int64) (*models.List, error) {
	list, err := m.db.GetListById(listId)
	if err != nil {
		return nil, utils.NewDBError("list not found")
	}

	if list.UserID != userId {
		return nil, utils.NewForbiddenError("this list doesn't belong to current user")
	}

	return list, nil
}

func (m *DefaultEntityService) GetListsByUserId(id int64) ([]models.List, error) {
	return m.db.GetListsByUserId(id)
}

func (m *DefaultEntityService) GetTasksByUserId(userId int64) ([]models.Task, error) {
	return m.db.GetTasksByUserId(userId)
}

func (m *DefaultEntityService) GetTasksByListId(listId int64, userId int64) ([]models.Task, error) {
	list, err := m.db.GetListById(listId)
	if err != nil {
		return nil, err
	}

	if list.UserID != userId {
		return nil, errors.New("this list doesn't belong to current user")
	}

	tasks, err := m.db.GetTasksByListId(list.ID)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (m *DefaultEntityService) GetTaskById(taskId int64, userId int64) (*models.Task, error) {
	task, err := m.db.GetTaskById(taskId)
	if err != nil {
		return nil, utils.NewDBError("task not found")
	}

	list, err := m.db.GetListById(task.ListID)
	if err != nil {
		return nil, utils.NewDBError("task is invalid: list this task belongs to is not found")
	}

	if list.UserID != userId {
		return nil, utils.NewForbiddenError("this task doesn't belong to current user")
	}

	return task, nil
}
