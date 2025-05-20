package models

import (
	"time"
)

type FormType string

const (
	FormTypeRadio     FormType = "radio"
	FormTypeCheckbox  FormType = "checkbox"
	FormTypeText      FormType = "text"
	FormTypeStar      FormType = "star"
	FormTypeTwoChoice FormType = "two_choice"
	FormTypeSlider    FormType = "slider"
)

type FormDB struct {
	FormID    uint      `json:"form_id" gorm:"primaryKey"`
	FormName  string    `json:"form_name" gorm:"not null"`
	ImgID     string    `json:"img_id"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Delete    bool      `json:"delete" gorm:"default:false"`
}

type SectionDB struct {
	SectionID    uint      `json:"section_id" gorm:"primaryKey"`
	FormID       uint      `json:"form_id" gorm:"not null"`
	SectionName  string    `json:"section_name" gorm:"not null"`
	SectionOrder int       `json:"section_order" gorm:"not null"`
	SectionType  FormType  `json:"section_type" gorm:"not null"`
	SectionDesc  string    `json:"section_desc" gorm:"not null"` //encoding base64 from json
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Delete       bool      `json:"delete" gorm:"default:false"`
}

type AnswerDB struct {
	AnswerID  uint      `json:"answer_id" gorm:"primaryKey"`
	FormID    uint      `json:"form_id" gorm:"not null"`
	SectionID uint      `json:"section_id" gorm:"not null"`
	Answer    string    `json:"answer" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	Delete    bool      `json:"delete" gorm:"default:false"`
}

type CreateFormRequest struct {
	FormName string `json:"form_name" example:"アンケートフォーム"`
	ImgID    string `json:"img_id" example:"1234567890"`
	Sections []CreateSectionRequest `json:"sections"`
}

type CreateSectionRequest struct {
	SectionName  string   `json:"section_name" example:"質問1"`
	SectionType  FormType `json:"section_type" example:"radio"`
	SectionDesc  string   `json:"section_desc" example:"質問の中身"` // encoding base64 from json
}
