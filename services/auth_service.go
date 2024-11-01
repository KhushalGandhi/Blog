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

	return database.DB.Create(user).Error // we can take it to db section too like service repo db section
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

func UpdateUser(userID uint, updatedData models.UpdateProfile) error {
	var user models.User

	if err := database.DB.First(&user, userID).Error; err != nil {
		return errors.New("user not found")
	}

	updateData := map[string]interface{}{
		"username": updatedData.UserName,
		"email":    updatedData.Email,
	}

	return database.DB.Model(&user).Updates(updateData).Error
}

func ViewProfile(userID uint) (*models.UserProfile, error) {
	var user models.User
	var profile models.UserProfile

	// Find the user by ID
	if err := database.DB.First(&user, userID).Error; err != nil {
		return nil, errors.New("user not found")
	}

	// Map user data to profile struct
	profile.UserName = user.Username
	profile.Email = user.Email

	return &profile, nil
}
