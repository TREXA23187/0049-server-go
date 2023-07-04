package console_service

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/ctype"
	"errors"
)

func (ConsoleService) CreateInstance(instanceId, instanceName, title, description, templateId, modelId, url, dataFile string, status ctype.Status) (*models.InstanceModel, error) {

	// Check if the user exists
	var instanceModel models.InstanceModel
	err := global.DB.Take(&instanceModel, "title = ?", title).Error
	if err == nil {
		return nil, errors.New("title already exists")
	}

	instanceModel = models.InstanceModel{
		InstanceID:   instanceId,
		InstanceName: instanceName,
		Title:        title,
		Description:  description,
		TemplateID:   templateId,
		ModelID:      modelId,
		URL:          url,
		Status:       status,
		DataFile:     dataFile,
	}

	// save to database
	err = global.DB.Create(&instanceModel).Error
	if err != nil {
		global.Log.Error(err)
		return nil, err
	}

	//err = global.DB.Create(&models.UserRoleModel{
	//	UserID: userModel.ID,
	//	RoleID: roleID,
	//}).Error

	if err != nil {
		global.Log.Error(err)
		return nil, err
	}

	return &instanceModel, nil
}
