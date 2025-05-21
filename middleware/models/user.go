package models

import (
	"time"
)

type UserDB struct {
	UserID    uint      `json:"user_id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	Delete    bool      `json:"delete" gorm:"default:false"`
}

type UserRequest struct {
	Name     string `json:"name" example:"田所浩二"`
	Email    string `json:"email" example:"example@example.com"`
	Password string `json:"password" example:"password"`
}

type UserLoginRequest struct {
	Email    string `json:"email" example:"example@example.com"`
	Password string `json:"password" example:"password"`
}

type UserVerifyRequest struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJpc3N1ZXIiLCJhbGciOiJIUzI1NiIsInN1YiI6InVzZXIiLCJhdWQiOlsiYXV0aWVuY2UiXSwiaWF0IjoxNzQ1NDY3MjgzLCJleHBpcmF0aW9uIjoxNzQ1NDY3MjgzLCJqdGkiOiJpZCJ9.dxgM6uH2F8ZglV_xcPhjCRnOSJBYq9oeS1TDLkLg_eg"`	
}
