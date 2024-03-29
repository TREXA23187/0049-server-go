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

func GetPermissionByURLAndMethod(url string, method ctype.PermissionMethod) (*models.PermissionModel, error) {
	var permissionModel models.PermissionModel
	err := global.DB.Take(&permissionModel, "url = ? and method = ?", url, method).Error
	if err != nil {
		global.Log.Error(err)
		return nil, err
	}

	return &permissionModel, nil
}

func GetPermissionById(id uint) (*models.PermissionModel, error) {
	var permissionModel models.PermissionModel
	err := global.DB.Take(&permissionModel, "id = ?", id).Error
	if err != nil {
		global.Log.Error(err)
		return nil, err
	}

	return &permissionModel, nil
}

func GetPermissionByRole(role ctype.Role) ([]models.RolePermissionModel, error) {
	var rolePermissionModelList []models.RolePermissionModel
	roleModel, err := GetRoleByRoleType(role)
	if err != nil {
		global.Log.Error(err)
		return nil, err
	}

	err = global.DB.Where("role_id = ?", roleModel.ID).Find(&rolePermissionModelList).Error
	if err != nil {
		global.Log.Error(err)
		return nil, err
	}

	return rolePermissionModelList, nil
}
