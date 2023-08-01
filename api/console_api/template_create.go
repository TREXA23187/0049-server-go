package console_api

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/res"
	"0049-server-go/utils/jwts"
	"github.com/gin-gonic/gin"
)

type TemplateCreateRequest struct {
	Name    string `json:"name" binding:"required" msg:"Please enter name"`
	Content string `json:"content" binding:"required" msg:"Please enter content"`
}

func (ConsoleApi) TemplateCreateView(ctx *gin.Context) {
	_claims, _ := ctx.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var cr TemplateCreateRequest
	if err := ctx.ShouldBindJSON(&cr); err != nil {
		res.FailWithValidMsg(err, &cr, ctx)
		return
	}

	var templateModel models.TemplateModel
	err := global.DB.Take(&templateModel, "name = ?", cr.Name).Error
	if err == nil {
		global.Log.Error(err)
		res.FailWithMessage("template title already exists", ctx)
		return
	}

	templateModel.Name = cr.Name
	templateModel.Content = cr.Content
	templateModel.CreateUser = claims.UserID

	err = global.DB.Create(&templateModel).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
		return
	}

	res.OkWithMessage("create template successfully", ctx)
}
