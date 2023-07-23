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
	UserName string `json:"user_name" binding:"required" msg:"Please enter username"`
	NickName string `json:"nick_name"`
	Password string `json:"password" binding:"required" msg:"Please enter password"`
	Role     string `json:"role"`
}

func (UserApi) UserCreateView(ctx *gin.Context) {
	var cr UserCreateRequest
	if err := ctx.ShouldBindJSON(&cr); err != nil {
		res.FailWithValidMsg(err, &cr, ctx)
		return
	}

	var nickName string
	if cr.NickName == "" {
		nickName = cr.UserName
	} else {
		nickName = cr.NickName
	}

	fmt.Println(33, cr.Role)
	var roleModel models.RoleModel
	err := global.DB.Take(&roleModel, "role_type = ?", ctype.StringToRoleType(cr.Role)).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
		return
	}

	err = user_service.UserService{}.CreateUser(cr.UserName, nickName, cr.Password, "", ctx.ClientIP(), roleModel.ID)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
		return
	}

	res.OkWithMessage(fmt.Sprintf("User %s was created successfully", cr.UserName), ctx)
}
