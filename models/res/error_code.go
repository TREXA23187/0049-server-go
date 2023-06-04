package res

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"os"
)

type ErrorCode int

const (
	SettingsError   ErrorCode = 1001
	ArgumentError   ErrorCode = 1002
	PermissionError ErrorCode = 1009
)

var ErrorMap = map[ErrorCode]string{}

const file = "models/res/error_code.json"

func InitErrorCode() {
	bytes, err := os.ReadFile(file)
	if err != nil {
		logrus.Error(err)
		return
	}

	err = json.Unmarshal(bytes, &ErrorMap)
	if err != nil {
		logrus.Error(err)
		return
	}
}
