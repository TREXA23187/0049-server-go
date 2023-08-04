package models

import "0049-server-go/models/ctype"

type InstanceModel struct {
	MODEL
	Name         string       `gorm:"size:36" json:"name"`
	Description  string       `gorm:"size:128" json:"description"`
	URL          string       `gorm:"size:36" json:"url"`
	IP           string       `gorm:"size:18" json:"ip"`
	Port         int          `gorm:"size:18" json:"port"`
	Image        string       `gorm:"size:36" json:"image"`
	Status       ctype.Status `gorm:"size:36" json:"status"`
	InstanceID   string       `gorm:"size:128" json:"instance_id"`
	InstanceName string       `gorm:"size:128" json:"instance_name"`
	CreateUser   uint         `json:"create_user"`
	IsPublic     bool         `gorm:"size:16;default:false" json:"is_public"`
}
