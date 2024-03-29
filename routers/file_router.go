package routers

import "0049-server-go/api"

func (router RouterGroup) FileRouter() {
	fileApi := api.ApiGroupApp.FileApi
	router.POST("/file/upload/:type", fileApi.FileUploadView)
	router.GET("/file/download", fileApi.FileDownloadView)
	router.GET("/file/info", fileApi.FileInfoView)
}
