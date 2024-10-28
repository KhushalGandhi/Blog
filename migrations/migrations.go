package migrations

import (
	"Blog/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Post{})
}
