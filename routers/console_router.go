package routers

import "0049-server-go/api"

func (router RouterGroup) ConsoleRouter() {
	consoleApi := api.ApiGroupApp.ConsoleApi
	router.POST("/console/instance", consoleApi.InstanceCreateView)
	router.GET("/console/instance/list", consoleApi.InstanceListView)
	router.GET("/console/instance/info", consoleApi.InstanceInfoView)
	router.POST("/console/instance/operate", consoleApi.InstanceOperateView)
	router.DELETE("/console/instance", consoleApi.InstanceRemoveView)

	router.POST("/console/template", consoleApi.TemplateCreateView)
	router.GET("/console/template/list", consoleApi.TemplateListView)

}
