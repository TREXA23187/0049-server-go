package ctype

type Status string
type TaskStatus string
type ImageStatus string

const (
	StatusRunning Status = "running"
	StatusExited  Status = "exited"
	StatusRemoved Status = "removed"
	StatusFailed  Status = "failed"
)

const (
	TaskStatusPending   TaskStatus = "pending"
	TaskStatusRunning   TaskStatus = "running"
	TaskStatusCompleted TaskStatus = "completed"
	TaskStatusFailed    TaskStatus = "failed"
)

const (
	ImageStatusBeingBuilt ImageStatus = "being built"
	ImageStatusUnused     ImageStatus = "unused"
	ImageStatusInUse      ImageStatus = "in use"
)
