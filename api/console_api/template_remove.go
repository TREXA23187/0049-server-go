package console_api

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/res"
	"github.com/gin-gonic/gin"
)

func (ConsoleApi) TemplateRemoveView(ctx *gin.Context) {
	id := ctx.Query("id")

	global.DB.Delete(&models.TemplateModel{}, "id = ?", id)

	res.OkWithMessage("template deleted", ctx)
}
