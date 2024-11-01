package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"user_name;unique;not null" json:"user_name"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}

type UserResponse struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}

type UserRequest struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateProfile struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}

type UserProfile struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}
