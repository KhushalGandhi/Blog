package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"username;unique;not null" json:"username"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}

type UserResponse struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
}

type UserRequest struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateProfile struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
}

type UserProfile struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
}
