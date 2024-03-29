package routers

import (
	"0049-server-go/global"
	"0049-server-go/middleware"
	"0049-server-go/models/res"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	//Global authentication middleware
	router.Use(middleware.JwtAuth())

	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	res.InitErrorCode()

	apiRouterGroup := router.Group("/api/v1")

	routerGroupApp := RouterGroup{
		apiRouterGroup,
	}

	routerGroupApp.UsersRouter()
	routerGroupApp.ConsoleRouter()
	routerGroupApp.FileRouter()

	return router
}
