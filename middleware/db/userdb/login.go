package userdb

import (
	"gorm.io/gorm"

	"feedo_back/middleware/models"
)

func LoginUser(db *gorm.DB, email string) (bool, string) {
	var user struct {
		Password string
	}

	if err := db.Model(&models.UserDB{}).Select("password").Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, ""
		}
		return false, ""
	}

	return true, user.Password
}