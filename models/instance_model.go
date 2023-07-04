package models

type InstanceModel struct {
	MODEL
	Title       string `gorm:"size:36" json:"title"`
	Description string `gorm:"size:128" json:"description"`
	TemplateID  string `gorm:"size:18" json:"template_id"`
	ModelID     string `gorm:"size:18" json:"model_id"`
	URL         string `gorm:"size:36" json:"url"`
	DataFile    string `gorm:"size:18" json:"data_file"`
}
