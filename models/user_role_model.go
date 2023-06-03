package models

import "time"

type UserRoleModel struct {
	UserID    uint      `json:"user_id"`
	UserModel UserModel `gorm:"foreignKey:UserID"`
	RoleID    uint      `json:"role_id"`
	RoleModel RoleModel `gorm:"foreignKey:RoleID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
}
