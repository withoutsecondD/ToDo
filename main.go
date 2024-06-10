package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/withoutsecondd/ToDo/database"
	"github.com/withoutsecondd/ToDo/routes"
	"log"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000", // URL фронтенда
		AllowHeaders:     "Authorization, Content-Type",
		AllowCredentials: true,
	}))

	err := database.InitConnection()
	if err != nil {
		log.Fatal(err)
	}

	routes.SetupAllRoutes(app)

	app.Listen(":8080")
}
