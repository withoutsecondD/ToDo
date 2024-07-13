package database

import (
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/withoutsecondd/ToDo/models"
	"os"
	"strings"
)

type MySqlDB struct {
	DB *sqlx.DB
}

func NewMySqlDB(db *sqlx.DB) *MySqlDB {
	return &MySqlDB{DB: db}
}

func (db *MySqlDB) GetUserById(id int64) (*models.UserResponse, error) {
	user := &models.UserResponse{}

	err := db.DB.Get(user, "SELECT id, age, first_name, last_name, city, email FROM withoutsecondd.user WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (db *MySqlDB) GetUserByEmail(email string) (*models.UserResponse, error) {
	user := &models.UserResponse{}

	err := db.DB.Get(user, "SELECT id, age, first_name, last_name, city, email FROM withoutsecondd.user WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (db *MySqlDB) GetUserPasswordByEmail(email string) ([]byte, error) {
	var password []byte

	err := db.DB.Get(&password, "SELECT password FROM withoutsecondd.user WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	return password, nil
}

func (db *MySqlDB) CreateUser(user *models.User) (*models.UserResponse, error) {
	query := `
		INSERT INTO withoutsecondd.user(age, first_name, last_name, city, email, password) 
		VALUES(?, ?, ?, ?, ?, ?)
	`

	_, err := db.DB.Queryx(query, user.Age, user.FirstName, user.LastName, user.City, user.Email, user.Password)
	if err != nil {
		var sqlErr *mysql.MySQLError
		if errors.As(err, &sqlErr) && sqlErr.Number == 1062 {
			return nil, errors.New("user with such email already exists")
		}
	}

	createdUser, err := db.GetUserByEmail(strings.ToLower(user.Email))
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (db *MySqlDB) GetListById(id int64) (*models.List, error) {
	list := &models.List{}

	err := db.DB.Get(list, "SELECT * FROM withoutsecondd.list WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (db *MySqlDB) GetListsByUserId(id int64) ([]models.List, error) {
	lists := make([]models.List, 0)

	err := db.DB.Select(&lists, "SELECT * FROM withoutsecondd.list WHERE user_id = ?", id)
	if err != nil {
		return nil, err
	}

	return lists, nil
}

func (db *MySqlDB) GetTasksByUserId(userId int64) ([]models.Task, error) {
	tasks := make([]models.Task, 0)

	query := `
		SELECT l.id, l.title, t.id, t.list_id, t.title FROM
			(SELECT * FROM withoutsecondd.list WHERE user_id = ?) AS l
		INNER JOIN withoutsecondd.task AS t ON t.list_id = l.id
		ORDER BY t.id;
	`

	err := db.DB.Select(&tasks, query, userId)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (db *MySqlDB) GetTasksByListId(listId int64) ([]models.Task, error) {
	tasks := make([]models.Task, 0)

	err := db.DB.Select(&tasks, "SELECT * FROM withoutsecondd.task WHERE list_id = ?", listId)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (db *MySqlDB) GetTaskById(taskId int64) (*models.Task, error) {
	task := &models.Task{}

	err := db.DB.Get(task, "SELECT * FROM withoutsecondd.task WHERE id = ?", taskId)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func InitMySqlConnection() (*sqlx.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	db, err := sqlx.Connect(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_ADDRESS"),
			os.Getenv("DB_SCHEMA"),
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
