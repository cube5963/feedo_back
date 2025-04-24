package models

import (
	"time"
)

type User struct {
	UserID    uint      `json:"user_id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	Delete    bool      `json:"delete" gorm:"default:false"`
}
