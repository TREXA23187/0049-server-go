package api

import (
	"0049-server-go/api/console_api"
	"0049-server-go/api/file_api"
	"0049-server-go/api/users_api"
)

type ApiGroup struct {
	UserApi    users_api.UserApi
	ConsoleApi console_api.ConsoleApi
	FileApi    file_api.FileApi
}

var ApiGroupApp = new(ApiGroup)
