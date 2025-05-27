package formdb

import (
	"feedo_back/middleware/models"

	"gorm.io/gorm"
)

func CreateAnswer(db *gorm.DB, answer *models.CreateAnswerRequest) bool {
	answerdb := db.Model(&models.AnswerDB{})

	answerRecord := &models.AnswerDB{
		FormID:     answer.FormID,
		SectionID:  answer.SectionID,
		Answer:     answer.Answer,
	}
	
	if err := answerdb.Create(answerRecord).Error; err != nil {
		return false
	}
	return true
}