package service

import (
	"errors"
	"strings"

	"github.com/withoutsecondd/ToDo/database"
	"github.com/withoutsecondd/ToDo/internal/utils"
	"github.com/withoutsecondd/ToDo/models"
	"golang.org/x/crypto/bcrypt"
)

type DefaultEntityService struct {
	db database.Database
	v  utils.Validator
}

func NewDefaultEntityService(d database.Database, v utils.Validator) *DefaultEntityService {
	return &DefaultEntityService{db: d, v: v}
}

func (s *DefaultEntityService) GetUserById(id int64) (*models.UserResponse, error) {
	return s.db.GetUserById(id)
}

func (s *DefaultEntityService) GetUserByEmail(email string) (*models.UserResponse, error) {
	return s.db.GetUserByEmail(email)
}

func (s *DefaultEntityService) GetUserPasswordByEmail(email string) ([]byte, error) {
	return s.db.GetUserPasswordByEmail(email)
}

func (s *DefaultEntityService) CreateUser(user *models.User) (*models.UserResponse, error) {
	err := s.v.ValidateStruct(user)
	if err != nil {
		return nil, errors.New("refused to create a user: some fields are invalid")
	}

	existingUser, _ := s.db.GetUserByEmail(user.Email)
	if existingUser != nil {
		return nil, errors.New("user with such email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.ID = 0
	user.Password = string(hashedPassword)
	user.Email = strings.ToLower(user.Email)

	return s.db.CreateUser(user)
}

func (s *DefaultEntityService) GetListById(listId int64, userId int64) (*models.List, error) {
	list, err := s.db.GetListById(listId)
	if err != nil {
		return nil, utils.NewDBError("list not found")
	}

	if list.UserID != userId {
		return nil, utils.NewForbiddenError("this list doesn't belong to current user")
	}

	return list, nil
}

func (s *DefaultEntityService) GetListsByUserId(id int64) ([]models.List, error) {
	// Check if the user exists first
	_, err := s.db.GetUserById(id)
	if err != nil {
		return nil, err
	}

	return s.db.GetListsByUserId(id)
}

func (s *DefaultEntityService) CreateList(listDto *models.ListCreateDto, userId int64) (*models.List, error) {
	err := s.v.ValidateStruct(listDto)
	if err != nil {
		return nil, errors.New("refused to create a list: some fields are invalid")
	}

	list := &models.List{
		ID:     0,
		UserID: userId,
		Title:  listDto.Title,
	}

	return s.db.CreateList(list)
}

func (s *DefaultEntityService) GetTasksByUserId(userId int64) ([]models.Task, error) {
	// Check if the user exists first
	_, err := s.db.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	return s.db.GetTasksByUserId(userId)
}

func (s *DefaultEntityService) GetTasksByListId(listId int64, userId int64) ([]models.Task, error) {
	list, err := s.db.GetListById(listId)
	if err != nil {
		return nil, utils.NewDBError("list not found")
	}

	if list.UserID != userId {
		return nil, utils.NewForbiddenError("this list doesn't belong to current user")
	}

	tasks, err := s.db.GetTasksByListId(list.ID)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *DefaultEntityService) GetTaskById(taskId int64, userId int64) (*models.Task, error) {
	task, err := s.db.GetTaskById(taskId)
	if err != nil {
		return nil, utils.NewDBError("task not found")
	}

	list, err := s.db.GetListById(task.ListID)
	if err != nil {
		return nil, utils.NewDBError("task is invalid: list this task belongs to is not found")
	}

	if list.UserID != userId {
		return nil, utils.NewForbiddenError("this task doesn't belong to current user")
	}

	return task, nil
}

func (s *DefaultEntityService) CreateTask(task *models.Task) (*models.Task, error) {
	return s.db.CreateTask(task)
}

func (s *DefaultEntityService) GetTagsByUserId(userId int64) ([]models.Tag, error) {
	// Check if the user exists first
	_, err := s.db.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	return s.db.GetTagsByUserId(userId)
}

func (s *DefaultEntityService) GetTagById(tagId int64, userId int64) (*models.Tag, error) {
	tag, err := s.db.GetTagById(tagId)
	if err != nil {
		return nil, err
	}

	if tag.UserID != userId {
		return nil, utils.NewForbiddenError("this tag doesn't belong to current user")
	}

	return tag, nil
}

func (s *DefaultEntityService) GetTagsByTaskId(taskId int64, userId int64) ([]models.Tag, error) {
	task, err := s.db.GetTaskById(taskId)
	if err != nil {
		return nil, err
	}

	list, err := s.db.GetListById(task.ListID)
	if err != nil {
		return nil, err
	}

	if list.UserID != userId {
		return nil, utils.NewForbiddenError("this task belongs to other user's list")
	}

	tags, err := s.db.GetTagsByTaskId(taskId)
	if err != nil {
		return nil, err
	}

	return tags, nil
}

func (s *DefaultEntityService) CreateTag(tagDto *models.TagCreateDto, userId int64) (*models.Tag, error) {
	err := s.v.ValidateStruct(tagDto)
	if err != nil {
		return nil, errors.New("refused to create a tag: some fields are invalid")
	}

	tag := &models.Tag{
		ID:     0,
		UserID: userId,
		Title:  tagDto.Title,
		Color:  tagDto.Color,
	}

	return s.db.CreateTag(tag)
}
