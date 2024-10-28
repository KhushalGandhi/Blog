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

// GetAllPosts controller to fetch all blog posts.
func GetAllPosts(c *fiber.Ctx) error {
	posts, err := services.GetAllPosts()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponse(c, posts)
}

func UpdatePost(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	postID, err := c.ParamsInt("id")
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid post ID")
	}

	var updatedPost models.Post
	if err := c.BodyParser(&updatedPost); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid input")
	}

	if err := services.UpdatePost(userID, postID, updatedPost); err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, "Post updated successfully")
}

// DeletePost controller to handle post deletion by owner.
func DeletePost(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	postID, err := c.ParamsInt("id")
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid post ID")
	}

	if err := services.DeletePost(userID, postID); err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, "Post deleted successfully")
}
