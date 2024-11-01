package controllers

import (
	"Blog/models"
	"Blog/services"
	"Blog/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Register(c *fiber.Ctx) error {
	var user models.UserRequest
	if err := c.BodyParser(&user); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid input")
	}

	if !utils.ValidateEmail(user.Email) {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid Email")
	}

	baseModel := models.User{
		Model:    gorm.Model{},
		Username: user.UserName,
		Email:    user.Email,
		Password: user.Password,
	}

	if err := services.RegisterUser(&baseModel); err != nil {
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

	var updatedData models.UpdateProfile
	if err := c.BodyParser(&updatedData); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid input")
	}

	if err := services.UpdateUser(userID, updatedData); err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, "Profile updated successfully")
}

func ViewProfile(c *fiber.Ctx) error {
	// Retrieve userID from the context
	userID, ok := c.Locals("userID").(uint)
	if !ok {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid user ID")
	}

	// Fetch the user profile
	userProfile, err := services.ViewProfile(userID)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	// Return a successful response with the user profile data
	return utils.SuccessResponse(c, userProfile)
}
