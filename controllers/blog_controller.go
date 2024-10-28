package controllers

import (
	"Blog/models"
	"Blog/services"
	"Blog/utils"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func CreatePost(c *fiber.Ctx) error {
	// Retrieve userID from locals and ensure it's uint
	userID, ok := c.Locals("userID").(uint)
	if !ok {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, "User ID not found or invalid")
	}

	var post models.CreatePost
	if err := c.BodyParser(&post); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid input")
	}

	baseModel := models.Post{
		Title:   post.Title,
		Content: post.Content,
		UserID:  userID,
	}

	if err := services.CreatePost(&baseModel); err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, "Post created successfully")
}

func GetPost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	post, err := services.GetPost(uint(id))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, "Post not found")
	}

	return c.JSON(post)
}
