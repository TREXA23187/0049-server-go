package models

type TemplateModel struct {
	MODEL
	Name    string `gorm:"size:36" json:"name"`
	Content string `gorm:"size:1024" json:"content"`
}
