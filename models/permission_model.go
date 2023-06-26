package models

import "0049-server-go/models/ctype"

type PermissionModel struct {
	MODEL
	Category    string                 `gorm:"size:36" json:"category"`
	URL         string                 `gorm:"size:36" json:"url"`
	Method      ctype.PermissionMethod `gorm:"size:36" json:"method"`
	Description string                 `gorm:"size:36" json:"description"`
}
