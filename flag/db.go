package flag

import (
	"0049-server-go/global"
	"0049-server-go/models"
)

func MakeMigrations() {
	var err error
	//global.DB.SetupJoinTable(&models.UserModel{}, "CollectsModels", &models.UserCollectModel{})
	global.DB.SetupJoinTable(&models.UserModel{}, "Role", &models.UserRoleModel{})
	global.DB.SetupJoinTable(&models.RoleModel{}, "Permissions", &models.RolePermissionModel{})

	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.UserModel{},
			&models.RoleModel{},
			&models.PermissionModel{},
			&models.UserRoleModel{},
			&models.RolePermissionModel{},
			//&log_stash.LogStashModel{},
		)
	if err != nil {
		global.Log.Error("[ error ] Failed to generate database table structure!")
		return
	}
	global.Log.Info("[ success ] Successfully generated database table structure！")
}
