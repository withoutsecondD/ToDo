package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/withoutsecondd/ToDo/internal/authenticator"
	"github.com/withoutsecondd/ToDo/internal/utils"
)

func LoginHandler(c *fiber.Ctx) error {
	var loginRequest authenticator.LoginRequest
	if err := c.BodyParser(&loginRequest); err != nil {
		return utils.FormatErrorResponse(c, fiber.StatusBadRequest, errors.New("incorrect email or password"))
	}

	if err := authenticator.Authenticate(&loginRequest); err != nil {
		return utils.FormatErrorResponse(c, fiber.StatusUnauthorized, err)
	}

	token, err := utils.GenerateJwtToken(loginRequest.Email)
	if err != nil {
		return utils.FormatErrorResponse(c, fiber.StatusInternalServerError, err)
	}

	return utils.FormatSuccessResponse(c, fiber.Map{"token": token})
}
