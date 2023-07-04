package ctype

type Status int

const (
	StatusRunning Status = 1
	StatusStopped Status = 2
	StatusRemoved Status = 3
	StatusFailed  Status = 4
)
