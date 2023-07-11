package console_api

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/res"
	"github.com/gin-gonic/gin"
)

func (ConsoleApi) TaskRemoveView(ctx *gin.Context) {
	id := ctx.Query("id")

	global.DB.Delete(&models.TaskModel{}, "id = ?", id)

	res.OkWithMessage("taska deleted", ctx)
}
