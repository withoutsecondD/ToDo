package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/withoutsecondd/ToDo/database"
	"github.com/withoutsecondd/ToDo/internal/utils"
)

func GetListsByCurrentUserHandler(c *fiber.Ctx) error {
	token, err := utils.ValidateJwtToken(c)
	if err != nil {
		return utils.FormatErrorResponse(c, fiber.StatusUnauthorized, err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return utils.FormatErrorResponse(c, fiber.StatusInternalServerError, errors.New("error getting token claims"))
	}

	userId := int64(claims["id"].(float64))

	lists, err := database.GetListsByUserId(userId)
	if err != nil {
		return utils.FormatErrorResponse(c, fiber.StatusBadRequest, err)
	}

	return utils.FormatSuccessResponse(c, lists)
}
