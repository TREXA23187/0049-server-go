package routers

import "0049-server-go/api"

func (router RouterGroup) FileRouter() {
	fileApi := api.ApiGroupApp.FileApi
	router.POST("/file/upload", fileApi.FileUploadView)
}
