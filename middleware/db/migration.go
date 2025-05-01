package db

import (
	"gorm.io/gorm"

	"feedo_back/middleware/models"
)

func Migrate(database *gorm.DB) {
	database.AutoMigrate(
		&models.UserDB{},
	)
}