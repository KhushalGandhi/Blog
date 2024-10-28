package services

import (
	"Blog/database"
	"Blog/models"
)

func CreatePost(post *models.Post) error {
	return database.DB.Create(post).Error
}

func GetPost(id uint) (*models.Post, error) {
	var post models.Post
	err := database.DB.First(&post, id).Error
	return &post, err
}
