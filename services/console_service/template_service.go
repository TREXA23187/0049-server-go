package console_service

import (
	"0049-server-go/global"
	"0049-server-go/models"
)

func GetTemplateById(templateId string) (*models.TemplateModel, error) {
	var templateModel models.TemplateModel
	err := global.DB.Take(&templateModel, "id = ?", templateId).Error
	if err != nil {
		global.Log.Error(err)
		return nil, err
	}

	return &templateModel, nil
}
