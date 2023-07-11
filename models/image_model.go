package models

type ImageModel struct {
	MODEL
	Repository string `gorm:"size:36" json:"repository"`
	Tag        string `gorm:"size:36" json:"tag"`
	ImageID    string `gorm:"size:128" json:"image_id"`
	Size       string `gorm:"size:36" json:"size"`
}
