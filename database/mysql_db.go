package database

import (
	"errors"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/withoutsecondd/ToDo/internal/utils"
	"github.com/withoutsecondd/ToDo/models"
)

type MySqlDB struct {
	DB *sqlx.DB
}

func NewMySqlDB(db *sqlx.DB) *MySqlDB {
	return &MySqlDB{DB: db}
}

func (db *MySqlDB) GetUserById(id int64) (*models.UserResponse, error) {
	user := &models.UserResponse{}

	query := "SELECT id, age, first_name, last_name, city, email, email_verified FROM withoutsecondd.user WHERE id = ?"
	if err := db.DB.Get(user, query, id); err != nil {
		return nil, err
	}

	return user, nil
}

func (db *MySqlDB) GetUserByEmail(email string) (*models.UserResponse, error) {
	user := &models.UserResponse{}

	query := "SELECT id, age, first_name, last_name, city, email, email_verified FROM withoutsecondd.user WHERE email = ?"
	if err := db.DB.Get(user, query, email); err != nil {
		return nil, err
	}

	return user, nil
}

func (db *MySqlDB) GetUserPasswordByEmail(email string) ([]byte, error) {
	var password []byte

	query := "SELECT password FROM withoutsecondd.user WHERE email = ?"
	if err := db.DB.Get(&password, query, email); err != nil {
		return nil, err
	}

	return password, nil
}

func (db *MySqlDB) CreateUser(user *models.User) (*models.UserResponse, error) {
	query := `
		INSERT INTO withoutsecondd.user(age, first_name, last_name, city, email, password) 
		VALUES(?, ?, ?, ?, ?, ?)
	`

	result, err := db.DB.Exec(query, user.Age, user.FirstName, user.LastName, user.City, user.Email, user.Password)
	if err != nil {
		return nil, err
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	createdUser, err := db.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (db *MySqlDB) UpdateUserInformation(userDto *models.UserUpdateInformationDto) (*models.UserResponse, error) {
	query := `
		UPDATE withoutsecondd.user
		SET age = ?, first_name = ?, last_name = ?, city = ?
		WHERE id = ?
	`

	_, err := db.DB.Exec(query, userDto.Age, userDto.FirstName, userDto.LastName, userDto.City, userDto.ID)
	if err != nil {
		return nil, err
	}

	updatedUser, err := db.GetUserById(userDto.ID)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (db *MySqlDB) UpdateUserEmailStatus(userDto *models.UserEmailStatusDto) (*models.UserResponse, error) {
	query := `
		UPDATE withoutsecondd.user
		SET email_verified = ?
		WHERE id = ?
	`

	_, err := db.DB.Exec(query, userDto.EmailVerified, userDto.ID)
	if err != nil {
		return nil, err
	}

	updatedUser, err := db.GetUserById(userDto.ID)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (db *MySqlDB) GetListById(id int64) (*models.List, error) {
	list := &models.List{}

	query := "SELECT * FROM withoutsecondd.list WHERE id = ?"
	if err := db.DB.Get(list, query, id); err != nil {
		return nil, err
	}

	return list, nil
}

func (db *MySqlDB) GetListsByUserId(userId int64) ([]models.List, error) {
	lists := make([]models.List, 0)

	query := "SELECT * FROM withoutsecondd.list WHERE user_id = ?"
	if err := db.DB.Select(&lists, query, userId); err != nil {
		return nil, err
	}

	return lists, nil
}

func (db *MySqlDB) CreateList(list *models.List) (*models.List, error) {
	query := `
		INSERT INTO withoutsecondd.list(user_id, title) 
		VALUES(?, ?);
	`

	result, err := db.DB.Exec(query, list.UserID, list.Title)
	if err != nil {
		return nil, err
	}

	insertedListId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	createdList, err := db.GetListById(insertedListId)
	if err != nil {
		return nil, err
	}

	return createdList, nil
}

func (db *MySqlDB) GetTaskById(id int64) (*models.Task, error) {
	task := &models.Task{}

	query := "SELECT * FROM withoutsecondd.task WHERE id = ?"
	if err := db.DB.Get(task, query, id); err != nil {
		return nil, err
	}

	return task, nil
}

func (db *MySqlDB) GetTasksByUserId(userId int64) ([]models.Task, error) {
	tasks := make([]models.Task, 0)

	query := `
		SELECT t.id, t.list_id, t.title, t.description, t.status, t.deadline FROM
    		(SELECT id FROM withoutsecondd.list WHERE user_id = ?) AS l
        INNER JOIN withoutsecondd.task AS t ON t.list_id = l.id
		ORDER BY t.list_id;
	`
	if err := db.DB.Select(&tasks, query, userId); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (db *MySqlDB) GetTasksByListId(listId int64) ([]models.Task, error) {
	tasks := make([]models.Task, 0)

	query := "SELECT * FROM withoutsecondd.task WHERE list_id = ?"
	if err := db.DB.Select(&tasks, query, listId); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (db *MySqlDB) CreateTask(task *models.Task) (*models.Task, error) {
	query := `
		INSERT INTO withoutsecondd.task(id, list_id, title, description, status, deadline) 
		VALUES(?, ?, ?, ?, ?, ?)
	`
	_, err := db.DB.Queryx(query, task.ID, task.ListID, task.Title, task.Description, task.Status, task.Deadline)
	if err != nil {
		var sqlErr *mysql.MySQLError
		if errors.As(err, &sqlErr) && sqlErr.Number == 1062 {
			return nil, utils.NewDBError("task with such id already exists")
		} else {
			return nil, err
		}
	}

	createdTask, err := db.GetTaskById(task.ID)
	if err != nil {
		return nil, err
	}

	return createdTask, nil
}

func (db *MySqlDB) GetTagsByUserId(userId int64) ([]models.Tag, error) {
	tags := make([]models.Tag, 0)

	query := `
		SELECT * FROM withoutsecondd.tag
		WHERE user_id = ?
	`
	if err := db.DB.Select(&tags, query, userId); err != nil {
		return nil, err
	}

	return tags, nil
}

func (db *MySqlDB) GetTagById(id int64) (*models.Tag, error) {
	tag := &models.Tag{}

	query := "SELECT * FROM withoutsecondd.tag WHERE id = ?"
	if err := db.DB.Get(tag, query, id); err != nil {
		return nil, err
	}

	return tag, nil
}

func (db *MySqlDB) GetTagsByTaskId(taskId int64) ([]models.Tag, error) {
	tags := make([]models.Tag, 0)

	query := `
		SELECT tag.id, tag.title, tag.color, tag.user_id FROM withoutsecondd.task_tag
		INNER JOIN withoutsecondd.tag tag on task_tag.tag_id = tag.id
		WHERE task_id = ?
	`
	if err := db.DB.Select(&tags, query, taskId); err != nil {
		return nil, err
	}

	return tags, nil
}

func (db *MySqlDB) CreateTag(tag *models.Tag) (*models.Tag, error) {
	query := `
		INSERT INTO withoutsecondd.tag(user_id, title, color)
		VALUES(?, ?, ?)
	`
	result, err := db.DB.Exec(query, tag.UserID, tag.Title, tag.Color)
	if err != nil {
		return nil, err
	}

	createdTagId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	createdTag, err := db.GetTagById(createdTagId)
	if err != nil {
		return nil, err
	}

	return createdTag, nil
}

func InitMySqlConnection() (*sqlx.DB, error) {
	if err := godotenv.Load("C:\\Users\\Diyar Z\\go\\src\\github.com\\withoutsecondd\\ToDo\\.env"); err != nil {
		return nil, err
	}

	db, err := sqlx.Connect(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
			os.Getenv("MYSQL_DB_USER"),
			os.Getenv("MYSQL_DB_PASSWORD"),
			os.Getenv("MYSQL_DB_ADDRESS"),
			os.Getenv("MYSQL_DB_SCHEMA"),
		),
	)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
