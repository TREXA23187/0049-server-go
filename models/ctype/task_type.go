package ctype

type TaskType string

const (
	TrainingTask   TaskType = "training"
	DeploymentTask TaskType = "deployment"
)
