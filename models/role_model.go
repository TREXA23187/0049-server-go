package models

type RoleModel struct {
	MODEL
	RoleName    string          `gorm:"size:36" json:"role_name"`
	Description string          `gorm:"size:36" json:"description"`
	Permission  PermissionModel `gorm:"many2many:role_permission_models;joinForeignKey:RoleID;JoinReferences:PermissionID" json:"permission"`
}
