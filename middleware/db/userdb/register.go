package userdb

import (
    "gorm.io/gorm"

    "feedo_back/middleware/models"
)

func RegisterUser(db *gorm.DB, user *models.UserRequest) bool {
    database := db.Model(&models.UserDB{})

    if err := database.Create(&models.UserDB{
        Name:     user.Name,
        Email:    user.Email,
        Password: user.Password,
    }).Error; err != nil {
        return false
    }
    return true
}