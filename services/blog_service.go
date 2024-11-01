package services

import (
	"Blog/database"
	"Blog/models"
	"errors"
)

func CreatePost(post *models.Post) error {
	return database.DB.Create(post).Error
}

func GetPost(id uint) (*models.Post, error) {
	var post models.Post
	err := database.DB.First(&post, id).Error
	return &post, err
}

// GetAllPosts retrieves all blog posts.
func GetAllPosts() ([]models.Post, error) {
	var posts []models.Post
	if err := database.DB.Find(&posts).Error; err != nil {
		return nil, errors.New("failed to retrieve posts")
	}
	return posts, nil
}

func UpdatePost(userID uint, postID int, updatedPost models.CreatePost) error {
	var post models.Post
	if err := database.DB.First(&post, postID).Error; err != nil {
		return errors.New("post not found")
	}

	if post.UserID != userID {
		return errors.New("you can only update your own posts")
	}

	post.Title = updatedPost.Title
	post.Content = updatedPost.Content
	return database.DB.Save(&post).Error
}

// DeletePost allows the owner to delete their post.
func DeletePost(userID uint, postID int) error {
	var post models.Post
	if err := database.DB.First(&post, postID).Error; err != nil {
		return errors.New("post not found")
	}

	if post.UserID != userID {
		return errors.New("you can only delete your own posts")
	}

	return database.DB.Delete(&post).Error
}
