package models

import "0049-server-go/models/ctype"

type TaskModel struct {
	MODEL
	Name                 string           `gorm:"size:36" json:"name"`
	Type                 ctype.TaskType   `gorm:"size:36" json:"type"`
	Status               ctype.TaskStatus `gorm:"size:36" json:"status"`
	Model                string           `gorm:"size:36" json:"model"`
	DataFileName         string           `gorm:"size:128" json:"data_file_name"`
	DataFilePath         string           `gorm:"size:128" json:"data_file_path"`
	TrainingLabel        string           `gorm:"size:36" json:"training_label"`
	TrainingDataLabel    string           `gorm:"size:128" json:"training_data_label"`
	TrainedModelFileName string           `gorm:"size:128" json:"trained_model_file_name"`
	TrainedModelFilePath string           `gorm:"size:128" json:"trained_model_file_path"`
	CreateUser           uint             `json:"create_user"`
}
