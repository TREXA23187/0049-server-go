package console_service

import (
	"0049-server-go/global"
	"0049-server-go/models"
)

func GetModelById(modelId string) (*models.ModelModel, error) {
	var modelModel models.ModelModel
	err := global.DB.Take(&modelModel, "id = ?", modelId).Error
	if err != nil {
		global.Log.Error(err)
		return nil, err
	}

	return &modelModel, nil
}
