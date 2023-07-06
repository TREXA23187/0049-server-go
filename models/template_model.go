package models

type TemplateModel struct {
	MODEL
	Title   string `gorm:"size:36" json:"title"`
	Content string `gorm:"size:1024" json:"content"`
}
