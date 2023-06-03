package user_service

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/utils/pwd"
	"errors"
)

const Avatar = "/uploads/avatar/default.png"

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
