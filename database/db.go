package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"os"
)

var DB *sqlx.DB

func InitConnection() error {
	godotenv.Load()

	var err error
	DB, err = sqlx.Connect(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_ADDRESS"),
			os.Getenv("DB_SCHEMA"),
		),
	)
	if err != nil {
		return err
	}

	if err := DB.Ping(); err != nil {
		return err
	}

	return nil
}
