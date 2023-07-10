package ctype

import "encoding/json"

type TaskType string

const (
	TrainingTask   TaskType = "training"
	DeploymentTask TaskType = "deployment"
)

func (t TaskType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t)
}
