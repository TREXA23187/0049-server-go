package console_api

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/res"
	"github.com/gin-gonic/gin"
)

type TemplateCreateRequest struct {
	Title   string `json:"title" binding:"required" msg:"Please enter title"`
	Content string `json:"content" binding:"required" msg:"Please enter content"`
}

func (ConsoleApi) TemplateCreateView(ctx *gin.Context) {
	var cr TemplateCreateRequest
	if err := ctx.ShouldBindJSON(&cr); err != nil {
		res.FailWithValidMsg(err, &cr, ctx)
		return
	}

	var templateModel models.TemplateModel
	err := global.DB.Take(&templateModel, "title = ?", cr.Title).Error
	if err == nil {
		global.Log.Error(err)
		res.FailWithMessage("template title already exists", ctx)
		return
	}

	templateModel.Title = cr.Title
	templateModel.Content = cr.Content

	err = global.DB.Create(&templateModel).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
		return
	}

	res.OkWithMessage("create template successfully", ctx)
}
