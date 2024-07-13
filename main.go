package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/withoutsecondd/ToDo/database"
	"github.com/withoutsecondd/ToDo/internal/utils"
	"github.com/withoutsecondd/ToDo/service"
	"github.com/withoutsecondd/ToDo/todo_handler"
	"log"
)

func main() {
	app := fiber.New()
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

	mainEntityService := service.NewDefaultEntityService(db, v)
	jwtAuthService, err := service.NewJwtAuthService(db)
	if err != nil {
		log.Fatal(err)
	}

	handler := todo_handler.NewHandler(mainEntityService, jwtAuthService)

	handler.SetupRoutes(app)

	err = app.Listen(":8080")
	if err != nil {
		log.Fatal(err)
		return
	}
}
