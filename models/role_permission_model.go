package models

import "time"

type RolePermissionModel struct {
	RoleID          uint            `json:"role_id"`
	RoleModel       RoleModel       `gorm:"foreignKey:RoleID"`
	PermissionID    uint            `json:"permission_id"`
	PermissionModel PermissionModel `gorm:"foreignKey:PermissionID"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"-"`
}
