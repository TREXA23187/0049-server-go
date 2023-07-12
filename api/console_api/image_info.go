package console_api

import (
	"0049-server-go/models/res"
	"0049-server-go/services/redis_service"
	"github.com/gin-gonic/gin"
)

func (ConsoleApi) ImageInfoView(ctx *gin.Context) {
	repository := ctx.Query("repository")
	inBuilding := redis_service.CheckIsImageInBuilding(repository)

	res.OkWithData(inBuilding, ctx)
}
