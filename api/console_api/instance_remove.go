package console_api

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/res"
	"github.com/gin-gonic/gin"
)

func (ConsoleApi) InstanceRemoveView(ctx *gin.Context) {
	instanceId := ctx.Query("id")

	global.DB.Delete(&models.InstanceModel{}, "instance_id = ?", instanceId)

	res.OkWithMessage("instance deleted", ctx)
}
