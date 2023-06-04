package user_service

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/ctype"
	"errors"
	"strconv"
)

func CreateRole(role ctype.Role, description string) error {
	roleModel := models.RoleModel{
		RoleType:    role,
		Description: description,
	}
	err := global.DB.Take(&roleModel, "role_type = ?", role).Error
	if err == nil {
		return errors.New("role already exists")
	}
	err = global.DB.Create(&roleModel).Error
	if err != nil {
		global.Log.Error(err)
		return err
	}

	return nil
}

func GetRoleByUserId(userId string) (*models.RoleModel, error) {
	var userRoleModel models.UserRoleModel
	err := global.DB.Take(&userRoleModel, "user_id = ?", userId).Error
	if err != nil {
		global.Log.Error(err)
		return nil, err
	}

	return GetRoleById(strconv.Itoa(int(userRoleModel.RoleID)))
}

func GetRoleById(roleId string) (*models.RoleModel, error) {
	var roleModel models.RoleModel
	err := global.DB.Take(&roleModel, "id = ?", roleId).Error
	if err != nil {
		global.Log.Error(err)
		return nil, err
	}

	return &roleModel, nil
}

func GetRoleByRoleType(role ctype.Role) (*models.RoleModel, error) {
	var roleModel models.RoleModel
	err := global.DB.Take(&roleModel, "role_type = ?", role).Error
	if err != nil {
		global.Log.Error(err)
		return nil, err
	}

	return &roleModel, nil
}

func GrantPermissionToRole(role ctype.Role, url string, method ctype.PermissionMethod) error {
	roleModel, err := GetRoleByRoleType(role)
	if err != nil {
		global.Log.Error(err)
		return err
	}

	permissionModel, err := GetPermissionByURLAndMethod(url, method)
	if err != nil {
		global.Log.Error(err)
		return err
	}

	rolePermissionModel := models.RolePermissionModel{
		RoleID:       roleModel.ID,
		PermissionID: permissionModel.ID,
	}

	err = global.DB.Take(&rolePermissionModel, "role_id = ? and permission_id = ?", roleModel.ID, permissionModel.ID).Error
	if err == nil {
		global.Log.Error("the role already has the permission")
		return errors.New("the role already has the permission")
	}

	err = global.DB.Create(&rolePermissionModel).Error
	if err != nil {
		global.Log.Error(err)
		return err
	}

	return nil
}
