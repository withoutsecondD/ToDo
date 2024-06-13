package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/withoutsecondd/ToDo/database"
	"github.com/withoutsecondd/ToDo/internal/utils"
	"strconv"
)

func GetListsByUserIdHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Queries()["user_id"])
	if err != nil {
		return utils.FormatErrorResponse(c, fiber.StatusBadRequest, err)
	}

	lists, err := database.GetListsByUserId(int64(id))
	if err != nil {
		return utils.FormatErrorResponse(c, fiber.StatusBadRequest, err)
	}

	return utils.FormatSuccessResponse(c, lists)
}
