package formdb

import (
	"feedo_back/middleware/models"

	"gorm.io/gorm"
)

func CreateForm(db *gorm.DB, form *models.CreateFormRequest) bool {
	formdb := db.Model(&models.FormDB{})
	sectiondb := db.Model(&models.SectionDB{})

	formRecord := &models.FormDB{
		FormName: form.FormName,
		ImgID:    form.ImgID,
	}

	if err := formdb.Create(formRecord).Error; err != nil {
		return false
	}

	for i := 0; i < len(form.Sections); i++ {
		section := form.Sections[i]

		if err := sectiondb.Create(&models.SectionDB{
			FormID:       formRecord.FormID,
			SectionName:  section.SectionName,
			SectionType:  section.SectionType,
			SectionDesc:  section.SectionDesc,
			SectionOrder: i,
		}).Error; err != nil {
			return false
		}
	}

	return true
}