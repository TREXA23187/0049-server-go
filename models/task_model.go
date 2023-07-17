package models

import "0049-server-go/models/ctype"

type TaskModel struct {
	MODEL
	Name          string         `gorm:"size:36" json:"name"`
	Type          ctype.TaskType `gorm:"size:36" json:"type"`
	Status        ctype.Status   `gorm:"size:36" json:"status"`
	Model         string         `gorm:"size:36" json:"model"`
	DataFileName  string         `gorm:"size:128" json:"data_file_name"`
	DataFilePath  string         `gorm:"size:128" json:"data_file_path"`
	TrainingLabel string         `gorm:"size:36" json:"training_label"`
}
