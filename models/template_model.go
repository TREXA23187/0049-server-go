package models

type TemplateModel struct {
	MODEL
	Name    string `gorm:"size:36" json:"name"`
	Content string `gorm:"size:256" json:"content"`
}
