package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/withoutsecondd/ToDo/database"
	"github.com/withoutsecondd/ToDo/internal/fmtResponse"
	"strconv"
)

func GetListsByUserIdHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Queries()["user_id"])
	if err != nil {
		return fmtResponse.FormatErrorResponse(c, fiber.StatusBadRequest, err)
	}

	lists, err := database.GetListsByUserId(int64(id))
	if err != nil {
		return fmtResponse.FormatErrorResponse(c, fiber.StatusBadRequest, err)
	}

	return fmtResponse.FormatSuccessResponse(c, lists)
}
