package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/withoutsecondd/ToDo/database"
	"github.com/withoutsecondd/ToDo/internal/fmtResponse"
	"strconv"
)

func GetAllUsersHandler(c *fiber.Ctx) error {
	users, err := database.GetAllUsers()
	if err != nil {
		return fmtResponse.FormatErrorResponse(c, fiber.StatusInternalServerError, err)
	}

	return fmtResponse.FormatSuccessResponse(c, users)
}

func GetUserByIdHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fmtResponse.FormatErrorResponse(c, fiber.StatusBadRequest, err)
	}

	user, err := database.GetUserById(int64(id))
	if err != nil {
		return fmtResponse.FormatErrorResponse(c, fiber.StatusNotFound, err)
	}

	return fmtResponse.FormatSuccessResponse(c, user)
}
