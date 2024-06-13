package utils

import "github.com/gofiber/fiber/v2"

// FormatErrorResponse Forms a JSON error response from fiber context, status code and error
func FormatErrorResponse(c *fiber.Ctx, status int, err error) error {
	return c.Status(status).JSON(fiber.Map{
		"success": false,
		"msg":     err.Error(),
	})
}

// FormatSuccessResponse Forms a JSON successful response from fiber context and data
func FormatSuccessResponse(c *fiber.Ctx, data interface{}) error {
	return c.JSON(fiber.Map{
		"success": true,
		"msg":     nil,
		"data":    data,
	})
}
