package users_api

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/res"
	"0049-server-go/services/common"
	"0049-server-go/services/user_service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserListResponse struct {
	models.UserModel
	Role int `json:"role"`
}

func (UserApi) UserListView(ctx *gin.Context) {
	var page models.PageInfo
	if err := ctx.ShouldBindQuery(&page); err != nil {
		res.FailWithCode(res.ArgumentError, ctx)
		return
	}

	var users []models.UserModel
	global.DB.Find(&users)

	list, count, _ := common.ComList(models.UserModel{}, common.Option{
		PageInfo: page,
	})

	resultList := make([]UserListResponse, len(list))
	for i, l := range list {
		roleModel, _ := user_service.GetRoleByUserId(strconv.Itoa(int(l.ID)))
		resultList[i] = UserListResponse{
			list[i],
			int(roleModel.RoleType),
		}
	}

	res.OkWithList(resultList, count, ctx)
}
