package console_api

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/res"
	"github.com/gin-gonic/gin"
)

func (ConsoleApi) TemplateInfoView(ctx *gin.Context) {
	id := ctx.Query("id")

	var template models.TemplateModel
	err := global.DB.Take(&template, "id = ?", id).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
		return
	}

	res.OkWithData(template, ctx)
}
