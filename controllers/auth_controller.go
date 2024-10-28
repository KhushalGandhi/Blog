package controllers

import (
	"Blog/models"
	"Blog/services"
	"Blog/utils"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid input")
	}

	if err := services.RegisterUser(&user); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SuccessResponse(c, "User registered successfully")
}

func Login(c *fiber.Ctx) error {
	var loginData map[string]string
	if err := c.BodyParser(&loginData); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid input")
	}

	token, err := services.LoginUser(loginData["email"], loginData["password"])
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid credentials")
	}

	return c.JSON(fiber.Map{"token": token})
}

func UpdateProfile(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	var updatedData models.User
	if err := c.BodyParser(&updatedData); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid input")
	}

	if err := services.UpdateUser(userID, updatedData); err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, "Profile updated successfully")
}
