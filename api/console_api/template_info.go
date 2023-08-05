package console_api

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/res"
	"0049-server-go/utils/jwts"
	"github.com/gin-gonic/gin"
)

func (ConsoleApi) TemplateInfoView(ctx *gin.Context) {
	_claims, _ := ctx.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	id := ctx.Query("id")

	var template models.TemplateModel
	err := global.DB.Take(&template, "id = ?", id).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
		return
	}

	if template.CreateUser != claims.UserID {
		res.FailWithCode(res.PermissionError, ctx)
		return
	}

	res.OkWithData(template, ctx)
}
