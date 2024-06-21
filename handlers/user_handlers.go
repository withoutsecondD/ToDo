package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/withoutsecondd/ToDo/database"
	"github.com/withoutsecondd/ToDo/internal/utils"
)

func GetCurrentUserHandler(c *fiber.Ctx) error {
	token, err := utils.ValidateJwtToken(c)
	if err != nil {
		return utils.FormatErrorResponse(c, fiber.StatusUnauthorized, err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return utils.FormatErrorResponse(c, fiber.StatusInternalServerError, errors.New("error getting token claims"))
	}

	userId := int64(claims["id"].(float64))

	user, err := database.GetUserById(userId)
	if err != nil {
		return utils.FormatErrorResponse(c, fiber.StatusNotFound, err)
	}

	return utils.FormatSuccessResponse(c, user)
}
