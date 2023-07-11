package console_api

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/res"
	"github.com/gin-gonic/gin"
)

func (ConsoleApi) ModelRemoveView(ctx *gin.Context) {
	id := ctx.Query("id")

	global.DB.Delete(&models.ModelModel{}, "id = ?", id)

	res.OkWithMessage("model deleted", ctx)
}
