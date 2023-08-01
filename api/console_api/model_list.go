package console_api

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/res"
	"0049-server-go/services/common"
	"0049-server-go/utils/jwts"
	"github.com/gin-gonic/gin"
)

func (ConsoleApi) ModelListView(ctx *gin.Context) {
	_claims, _ := ctx.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var page models.PageInfo
	if err := ctx.ShouldBindQuery(&page); err != nil {
		res.FailWithCode(res.ArgumentError, ctx)
		return
	}

	var modelList []models.ModelModel
	global.DB.Find(&modelList)

	list, count, _ := common.ComList(models.ModelModel{}, common.Option{
		PageInfo:   page,
		CreateUser: claims.UserID,
	})

	res.OkWithList(list, count, ctx)
}
