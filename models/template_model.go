package models

type TemplateModel struct {
	MODEL
	Name       string `gorm:"size:36" json:"name"`
	Content    string `gorm:"type:text" json:"content"`
	CreateUser uint   `json:"create_user"`
}
