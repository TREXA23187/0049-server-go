package user_service

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"strconv"
)

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
