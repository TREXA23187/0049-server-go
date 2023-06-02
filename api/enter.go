package api

import "0049-server-go/api/users_api"

type ApiGroup struct {
	UserApi users_api.UserApi
}

var ApiGroupApp = new(ApiGroup)
