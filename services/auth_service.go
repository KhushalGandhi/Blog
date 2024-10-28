package services

import (
	"Blog/database"
	"Blog/models"
	"Blog/utils"
	"errors"
)

func RegisterUser(user *models.User) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	return database.DB.Create(user).Error
}

func LoginUser(email, password string) (string, error) {
	var user models.User

	// Check if the user exists
	err := database.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return "", errors.New("user not found")
	}

	// Check if the password is correct
	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("incorrect password")
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}

func UpdateUser(userID uint, updatedData models.User) error {
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return errors.New("user not found")
	}

	updatedData.Password = user.Password // Keep original password
	return database.DB.Model(&user).Updates(updatedData).Error
}
