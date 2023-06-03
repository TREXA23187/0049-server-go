package flag

import (
	"0049-server-go/global"
	"0049-server-go/services/user_service"
)

func CreateUser() {
	roleModel, err := user_service.UserService{}.InitRoleAndPermission()
	if err != nil {
		global.Log.Error(err)
	}

	var (
		userName = "admin"
		nickName = "admin"
		password = "admin"
		email    = "admin@admin.com"
		ip       = "127.0.0.1"
	)

	err = user_service.UserService{}.CreateUser(userName, nickName, password, email, ip, roleModel.ID)
	if err != nil {
		global.Log.Error(err)
	}

	global.Log.Infof("User %s was created successfully！", userName)
}
