package models

import "time"

type RolePermissionModel struct {
	RoleID       uint      `json:"role_id"`
	PermissionID uint      `json:"permission_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"-"`
}
