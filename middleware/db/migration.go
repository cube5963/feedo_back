package db

import (
	"gorm.io/gorm"

	"feedo_back/middleware/db/models"
)

func Migrate(database *gorm.DB) {
	database.AutoMigrate(
		&models.User{},
	)
}