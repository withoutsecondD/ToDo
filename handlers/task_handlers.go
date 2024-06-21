package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/withoutsecondd/ToDo/database"
	"github.com/withoutsecondd/ToDo/internal/utils"
	"strconv"
)

// GetTasksByIdHandler returns response with tasks specified by list id they belong to.
// If no list id is provided, returns tasks by user
func GetTasksByIdHandler(c *fiber.Ctx) error {
	token, err := utils.ValidateJwtToken(c)
	if err != nil {
		return utils.FormatErrorResponse(c, fiber.StatusUnauthorized, err)
	}

	listIdStr := c.Queries()["list_id"]

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return utils.FormatErrorResponse(c, fiber.StatusInternalServerError, errors.New("error getting token claims"))
	}

	userId := int64(claims["id"].(float64))

	if listIdStr == "" {
		tasks, err := database.GetTasksByUserId(userId)
		if err != nil {
			return utils.FormatErrorResponse(c, fiber.StatusInternalServerError, err)
		}

		return utils.FormatSuccessResponse(c, tasks)
	} else {
		listId, err := strconv.Atoi(listIdStr)
		if err != nil {
			return utils.FormatErrorResponse(c, fiber.StatusBadRequest, err)
		}

		list, err := database.GetListById(int64(listId))

		if list.UserID != userId {
			return utils.FormatErrorResponse(c, fiber.StatusForbidden, errors.New("this list doesn't belong to current user"))
		}

		tasks, err := database.GetTasksByListId(list.ID)
		if err != nil {
			return utils.FormatErrorResponse(c, fiber.StatusInternalServerError, err)
		}

		return utils.FormatSuccessResponse(c, tasks)
	}
}
