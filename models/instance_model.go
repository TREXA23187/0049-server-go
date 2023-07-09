package models

type InstanceModel struct {
	MODEL
	InstanceID   string `gorm:"size:128" json:"instance_id"`
	InstanceName string `gorm:"size:128" json:"instance_name"`
	Title        string `gorm:"size:36" json:"title"`
	Description  string `gorm:"size:128" json:"description"`
	TemplateID   string `gorm:"size:18" json:"template_id"`
	ModelID      string `gorm:"size:18" json:"model_id"`
	URL          string `gorm:"size:36" json:"url"`
	IP           string `gorm:"size:18" json:"ip"`
	Port         int    `gorm:"size:18" json:"port"`
	Status       string `gorm:"size:36" json:"status"`
	DataFileName string `gorm:"size:128" json:"data_file_name"`
	DataFilePath string `gorm:"size:128" json:"data_file_path"`
}
