package utils

import (
	"github.com/gofiber/fiber/v2"
)

func SuccessResponse(c *fiber.Ctx, message interface{}) error {
	return c.JSON(fiber.Map{
		"success": true,
		"message": message,
	})
}

func ErrorResponse(c *fiber.Ctx, code int, message string) error {
	return c.Status(code).JSON(fiber.Map{
		"success": false,
		"message": message,
	})
}
