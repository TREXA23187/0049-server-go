package users_api

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/res"
	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	models.UserModel
}

func (UserApi) UserListView(ctx *gin.Context) {
	//var page models.PageInfo
	//if err := ctx.ShouldBindQuery(&page); err != nil {
	//	res.FailWithCode(res.ArgumentError, ctx)
	//	return
	//}

	var users []models.UserModel
	global.DB.Find(&users)

	//list, count, _ := common.ComList(models.UserModel{}, common.Option{
	//	PageInfo: page,
	//})

	res.OkWithData(users, ctx)
}
