package ctype

type Status string
type ImageStatus string

const (
	StatusRunning Status = "running"
	StatusExited  Status = "exited"
	StatusRemoved Status = "removed"
	StatusFailed  Status = "failed"
)

const (
	ImageStatusBeingBuilt ImageStatus = "being built"
	ImageStatusUnused     ImageStatus = "unused"
	ImageStatusInSse      ImageStatus = "in use"
)
