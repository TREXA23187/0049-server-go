package models

import "0049-server-go/models/ctype"

type TaskModel struct {
	MODEL
	Name       string         `gorm:"size:36" json:"name"`
	Type       ctype.TaskType `gorm:"size:36" json:"type"`
	InstanceID uint           `gorm:"size:128" json:"instance_id"`
	Status     ctype.Status   `gorm:"size:36" json:"status"`

	TrainingLabel string `gorm:"size:36" json:"training_label"`
}
