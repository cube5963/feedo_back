package models

import (
	"encoding/json"
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
	SectionID    uint            `json:"section_id" gorm:"primaryKey"`
	FormID       uint            `json:"form_id"`
	SectionName  string          `json:"section_name" gorm:"not null"`
	SectionOrder int             `json:"section_order" gorm:"not null"`
	SectionType  FormType        `json:"section_type" gorm:"not null"`
	SectionDesc  json.RawMessage `json:"section_desc"`
	CreatedAt    time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time       `json:"updated_at" gorm:"autoUpdateTime"`
	Delete       bool            `json:"delete" gorm:"default:false"`
}

type AnswerDB struct {
	AnswerID  uint      `json:"answer_id" gorm:"primaryKey"`
	FormID    uint      `json:"form_id"`
	SectionID uint      `json:"section_id"`
	Answer    string    `json:"answer"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	Delete    bool      `json:"delete" gorm:"default:false"`
}
