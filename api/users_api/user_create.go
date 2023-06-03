package users_api

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/ctype"
	"0049-server-go/models/res"
	"0049-server-go/services/user_service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserCreateRequest struct {
	NickName string     `json:"nick_name" binding:"required" msg:"Please enter nickname"`
	UserName string     `json:"user_name" binding:"required" msg:"Please enter username"`
	Password string     `json:"password" binding:"required" msg:"Please enter password"`
	Role     ctype.Role `json:"role" binding:"required" msg:"Please enter role"`
}

func (UserApi) UserCreateView(ctx *gin.Context) {
	var cr UserCreateRequest
	if err := ctx.ShouldBindJSON(&cr); err != nil {
		res.FailWithValidMsg(err, &cr, ctx)
		return
	}

	var roleModel models.RoleModel
	err := global.DB.Take(&roleModel, "role_type = ?", cr.Role).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
		return
	}

	err = user_service.UserService{}.CreateUser(cr.UserName, cr.NickName, cr.Password, "", ctx.ClientIP(), roleModel.ID)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
		return
	}

	res.OkWithMessage(fmt.Sprintf("User %s was created successfully", cr.UserName), ctx)
}
