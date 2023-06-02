package log_stash

import "encoding/json"

type Level int

const (
	DebugLevel   Level = 1
	InfoLevel    Level = 2
	WarningLevel Level = 3
	ErrorLevel   Level = 4
)

func (level Level) MarshalJSON() ([]byte, error) {
	return json.Marshal(level.String())
}

func (level Level) String() string {
	var str string
	switch level {
	case DebugLevel:
		str = "debug"
	case InfoLevel:
		str = "info"
	case WarningLevel:
		str = "warning"
	case ErrorLevel:
		str = "error"
	default:
		str = "other"
	}
	return str
}
