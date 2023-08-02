package console_api

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/res"
	"github.com/gin-gonic/gin"
)

func (ConsoleApi) TemplateUpdateView(ctx *gin.Context) {
	var template models.TemplateModel
	if err := ctx.ShouldBindJSON(&template); err != nil {
		res.FailWithValidMsg(err, &template, ctx)
		return
	}

	err := global.DB.Model(&template).Where("id = ?", template.ID).Updates(template).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
		return
	}

	res.OkWithMessage("template updated", ctx)
}
