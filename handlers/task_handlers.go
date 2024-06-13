package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/withoutsecondd/ToDo/database"
	"github.com/withoutsecondd/ToDo/internal/utils"
	"strconv"
)

// GetTasksByIdHandler returns response with tasks specified by user id or by list id they belong to
func GetTasksByIdHandler(c *fiber.Ctx) error {
	userId := c.Queries()["user_id"]

	if userId != "" {
		c.Append("id", userId)
		return getTasksByUserIdHandler(c)
	}

	listId := c.Queries()["list_id"]

	if listId != "" {
		c.Append("id", listId)
		return getTasksByListIdHandler(c)
	}

	return utils.FormatErrorResponse(c, fiber.StatusBadRequest, errors.New("no user_id or list_id parameter is specified"))
}

func getTasksByUserIdHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.GetRespHeader("id"))
	if err != nil {
		return utils.FormatErrorResponse(c, fiber.StatusBadRequest, err)
	}

	tasks, err := database.GetTasksByUserId(int64(id))
	if err != nil {
		return utils.FormatErrorResponse(c, fiber.StatusBadRequest, err)
	}

	return utils.FormatSuccessResponse(c, tasks)
}

func getTasksByListIdHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.GetRespHeader("id"))
	if err != nil {
		return utils.FormatErrorResponse(c, fiber.StatusBadRequest, err)
	}

	tasks, err := database.GetTasksByListId(int64(id))
	if err != nil {
		return utils.FormatErrorResponse(c, fiber.StatusBadRequest, err)
	}

	return utils.FormatSuccessResponse(c, tasks)
}
