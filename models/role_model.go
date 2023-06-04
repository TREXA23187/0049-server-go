package models

import "0049-server-go/models/ctype"

type RoleModel struct {
	MODEL
	RoleType    ctype.Role `gorm:"size:36" json:"role_type"`
	Description string     `gorm:"size:36" json:"description"`
}
