package middlewares

import (
	"Blog/utils"
	"github.com/gofiber/fiber/v2"
)

func AuthRequired(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	userId, err := utils.ValidateJWT(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}
	c.Locals("userId", userId)
	return c.Next()
}
