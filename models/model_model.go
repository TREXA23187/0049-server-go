package models

type ModelModel struct {
	MODEL
	ModelKey string `gorm:"size:36" json:"model_key"`
	Name     string `gorm:"size:36" json:"name"`
	Code     string `gorm:"size:256" json:"code"`
}
