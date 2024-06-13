package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/withoutsecondd/ToDo/handlers"
)

func SetupAllRoutes(a *fiber.App) {
	setupLoginRoute(a)
	setupUserRoutes(a)
	setupListRoutes(a)
	setupTaskRoutes(a)
}

func setupLoginRoute(a *fiber.App) {
	a.Post("/login", handlers.LoginHandler)
}

func setupUserRoutes(a *fiber.App) {
	a.Get("/users/", handlers.GetAllUsersHandler)
	a.Get("/users/:id", handlers.GetUserByIdHandler)
}

func setupListRoutes(a *fiber.App) {
	a.Get("/lists", handlers.GetListsByUserIdHandler) // Requires user_id as a query variable
}

func setupTaskRoutes(a *fiber.App) {
	a.Get("/tasks", handlers.GetTasksByIdHandler) // Requires list_id or user_id as a query variable
}
