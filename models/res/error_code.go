package res

import (
	"embed"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/fs"
	"log"
)

//go:embed error_code.json
var embedErrorCodeFiles embed.FS

type ErrorCode int

const (
	SettingsError   ErrorCode = 1001
	ArgumentError   ErrorCode = 1002
	PermissionError ErrorCode = 10001
)

var ErrorMap = map[ErrorCode]string{}

const file = "models/res/error_code.json"

func InitErrorCode() {
	//bytes, err := os.ReadFile(file)
	//if err != nil {
	//	logrus.Error(err)
	//	return
	//}

	bytes, err := fs.ReadFile(embedErrorCodeFiles, "error_code.json")
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	err = json.Unmarshal(bytes, &ErrorMap)
	if err != nil {
		logrus.Error(err)
		return
	}
}
