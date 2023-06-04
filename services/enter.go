package services

import "0049-server-go/services/user_service"

type ServiceGroup struct {
	UserService user_service.UserService
}

var ServiceApp = new(ServiceGroup)
