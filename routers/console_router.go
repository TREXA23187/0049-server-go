package routers

import "0049-server-go/api"

func (router RouterGroup) ConsoleRouter() {
	consoleApi := api.ApiGroupApp.ConsoleApi
	router.POST("/console/instance", consoleApi.InstanceCreateView)
	router.GET("/console/instance/list", consoleApi.InstanceListView)
	router.GET("/console/instance/info", consoleApi.InstanceInfoView)
}
