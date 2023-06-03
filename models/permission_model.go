package models

type PermissionModel struct {
	MODEL
	URL         string `gorm:"size:36" json:"url"`
	Method      string `gorm:"size:36" json:"method"`
	Description string `gorm:"size:36" json:"description"`
}
