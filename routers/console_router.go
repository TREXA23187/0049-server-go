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
	router.DELETE("/console/template", consoleApi.TemplateRemoveView)
	router.GET("/console/template/info", consoleApi.TemplateInfoView)

	router.POST("/console/model", consoleApi.ModelCreateView)
	router.GET("/console/model/list", consoleApi.ModelListView)
	router.DELETE("/console/model", consoleApi.ModelRemoveView)

	router.POST("/console/task", consoleApi.TaskCreateView)
	router.GET("/console/task/list", consoleApi.TaskListView)
	router.POST("/console/task/operate", consoleApi.TaskOperateView)
	router.DELETE("/console/task", consoleApi.TaskRemoveView)

	router.POST("/console/image", consoleApi.ImageCreateView)
	router.GET("/console/image/list", consoleApi.ImageListView)
	router.GET("/console/image/info", consoleApi.ImageInfoView)
	router.DELETE("/console/image", consoleApi.ImageRemoveView)
}
