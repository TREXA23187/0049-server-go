package user_service

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/ctype"
	"errors"
)

func CreatePermission(url string, method ctype.PermissionMethod, description string) error {
	permissionModel := models.PermissionModel{
		URL:         url,
		Method:      method,
		Description: description,
	}
	err := global.DB.Take(&permissionModel, "url = ? and method = ?", url, method).Error
	if err == nil {
		return errors.New("permission already exists")
	}
	err = global.DB.Create(&permissionModel).Error
	if err != nil {
		global.Log.Error(err)
		return err
	}

	return nil
}
