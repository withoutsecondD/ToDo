package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/withoutsecondd/ToDo/database"
	"github.com/withoutsecondd/ToDo/internal/utils"
	"strconv"
)

func GetAllUsersHandler(c *fiber.Ctx) error {
	users, err := database.GetAllUsers()
	if err != nil {
		return utils.FormatErrorResponse(c, fiber.StatusInternalServerError, err)
	}

	return utils.FormatSuccessResponse(c, users)
}

func GetUserByIdHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.FormatErrorResponse(c, fiber.StatusBadRequest, err)
	}

	user, err := database.GetUserById(int64(id))
	if err != nil {
		return utils.FormatErrorResponse(c, fiber.StatusNotFound, err)
	}

	return utils.FormatSuccessResponse(c, user)
}
