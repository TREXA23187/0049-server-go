package models

type UserModel struct {
	MODEL
	NickName string      `gorm:"size:36" json:"nick_name"`
	UserName string      `gorm:"size:36" json:"user_name"`
	Password string      `gorm:"size:128" json:"-"`
	Avatar   string      `gorm:"size:256" json:"avatar_id"`
	Email    string      `gorm:"size:128" json:"email"`
	Tel      string      `gorm:"size:18" json:"tel"`
	Address  string      `gorm:"size:64" json:"address"`
	IP       string      `gorm:"size:20" json:"ip"`
	Roles    []RoleModel `gorm:"many2many:user_role_models;joinForeignKey:UserID;JoinReferences:RoleID" json:"roles"`
}
