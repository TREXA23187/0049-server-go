package models

import "0049-server-go/models/ctype"

type InstanceModel struct {
	MODEL
	InstanceID   string       `gorm:"size:128" json:"instance_id"`
	InstanceName string       `gorm:"size:128" json:"instance_name"`
	Title        string       `gorm:"size:36" json:"title"`
	Description  string       `gorm:"size:128" json:"description"`
	TemplateID   string       `gorm:"size:18" json:"template_id"`
	ModelID      string       `gorm:"size:18" json:"model_id"`
	URL          string       `gorm:"size:36" json:"url"`
	Status       ctype.Status `gorm:"size:36" json:"status"`
	DataFile     string       `gorm:"size:18" json:"data_file"`
}
