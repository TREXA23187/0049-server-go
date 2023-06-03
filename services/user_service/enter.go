package user_service

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/ctype"
	"errors"
)

type UserService struct {
}

func (UserService) InitRoleAndPermission() (*models.RoleModel, error) {
	permissionModel := models.PermissionModel{
		URL:         "*",
		Method:      ctype.ALL,
		Description: "super admin permission",
	}
	err := global.DB.Take(&permissionModel, "url = ? and method = ?", permissionModel.URL, permissionModel.Method).Error
	if err == nil {
		return nil, errors.New("permission already exists")
	}
	err = global.DB.Create(&permissionModel).Error
	if err != nil {
		global.Log.Error(err)
		return nil, err
	}

	roleModel := models.RoleModel{
		RoleType:    ctype.RoleSuperAdmin,
		Description: "super admin",
	}
	err = global.DB.Take(&roleModel, "role_type = ?", roleModel.RoleType).Error
	if err == nil {
		return nil, errors.New("role already exists")
	}
	err = global.DB.Create(&roleModel).Error
	if err != nil {
		global.Log.Error(err)
		return nil, err
	}

	err = global.DB.Create(&models.RolePermissionModel{
		RoleID:       roleModel.ID,
		PermissionID: permissionModel.ID,
	}).Error
	if err != nil {
		global.Log.Error(err)
		return nil, err
	}

	return &roleModel, nil
}
