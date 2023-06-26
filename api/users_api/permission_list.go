package users_api

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/res"
	"0049-server-go/services/common"
	"github.com/gin-gonic/gin"
)

type PermissionListResponse struct {
	models.PermissionModel
}

func (UserApi) PermissionListView(ctx *gin.Context) {
	var page models.PageInfo
	if err := ctx.ShouldBindQuery(&page); err != nil {
		res.FailWithCode(res.ArgumentError, ctx)
		return
	}

	var permissions []models.PermissionModel
	global.DB.Find(&permissions)

	list, count, _ := common.ComList(models.PermissionModel{}, common.Option{
		PageInfo: page,
	})

	res.OkWithList(list, count, ctx)
}
