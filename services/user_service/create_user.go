package user_service

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/ctype"
	"0049-server-go/utils/pwd"
	"errors"
)

const Avatar = "/uploads/avatar/default.png"

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

func (UserService) CreateUser(userName, nickName, password string, email string, ip string, roleID uint) error {
	// Check if the user exists
	var userModel models.UserModel
	err := global.DB.Take(&userModel, "user_name = ?", userName).Error
	if err == nil {
		return errors.New("username already exists")
	}

	// Hash password
	hashPwd := pwd.HashPwd(password)

	userModel = models.UserModel{
		UserName: userName,
		NickName: nickName,
		Password: hashPwd,
		Email:    email,
		Avatar:   Avatar,
		IP:       ip,
		Address:  "Intranet",
	}
	// save to database
	err = global.DB.Create(&userModel).Error
	if err != nil {
		global.Log.Error(err)
		return err
	}

	err = global.DB.Create(&models.UserRoleModel{
		UserID: userModel.ID,
		RoleID: roleID,
	}).Error
	if err != nil {
		global.Log.Error(err)
		return err
	}

	return nil
}
