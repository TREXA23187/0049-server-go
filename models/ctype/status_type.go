package ctype

type Status string

const (
	StatusRunning Status = "running"
	StatusExited  Status = "exited"
	StatusRemoved Status = "removed"
	StatusFailed  Status = "failed"
)
