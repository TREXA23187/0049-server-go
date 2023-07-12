package models

import "0049-server-go/models/ctype"

type ImageModel struct {
	MODEL
	Repository string            `gorm:"size:36" json:"repository"`
	Tag        string            `gorm:"size:36" json:"tag"`
	Status     ctype.ImageStatus `gorm:"size:36" json:"status"`
	Task       string            `gorm:"size:36" json:"task"`
	ImageID    string            `gorm:"size:128" json:"image_id"`
	Size       int64             `gorm:"size:36" json:"size"`
}
