package main

import (
	"crypto/rand"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/withoutsecondd/ToDo/database"
	"github.com/withoutsecondd/ToDo/internal/utils"
	"github.com/withoutsecondd/ToDo/service"
	"github.com/withoutsecondd/ToDo/todo_handler"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:     "ToDo REST API application",
		JSONDecoder: utils.CustomJSONDecoder,
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000", // URL фронтенда
		AllowHeaders:     "Authorization, Content-Type",
		AllowCredentials: true,
	}))

	connection, err := database.InitMySqlConnection()
	if err != nil {
		log.Fatal(err)
	}

	db := database.NewMySqlDB(connection)
	v, err := utils.NewDefaultValidator()
	if err != nil {
		log.Fatal(err)
	}
	dialer := utils.NewDefaultDialer("marbiru15@gmail.com", "smtp.gmail.com", 587, "pwbt unpm cfgu laru")

	defaultEntityService := service.NewDefaultEntityService(db, v)

	jwtKey, err := generateJwtKey()
	if err != nil {
		log.Fatal(err)
	}
	jwtAuthService := service.NewJwtAuthService(db, jwtKey)

	jwtKey, err = generateJwtKey()
	if err != nil {
		log.Fatal(err)
	}
	defaultEmailService := service.NewDefaultEmailService(db, dialer, jwtKey)

	handler := todo_handler.NewHandler(defaultEntityService, jwtAuthService, defaultEmailService)

	handler.SetupRoutes(app)

	err = app.Listen("localhost:8080")
	if err != nil {
		log.Fatal(err)
		return
	}
}

func generateJwtKey() ([]byte, error) {
	jwtKey := make([]byte, 32)
	if _, err := rand.Read(jwtKey); err != nil {
		return nil, err
	}

	//return jwtKey, nil
	return []byte{22, 112, 222, 0, 209, 146, 167, 212, 158, 39, 193, 131, 191, 67, 190, 52, 15, 170, 254, 43, 6, 5, 3, 175, 134, 227, 118, 82, 6, 243, 98, 111}, nil
}
