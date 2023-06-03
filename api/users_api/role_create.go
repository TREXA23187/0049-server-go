package users_api

import (
	"0049-server-go/global"
	"0049-server-go/models/ctype"
	"0049-server-go/models/res"
	"0049-server-go/services/user_service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type RoleCreateRequest struct {
	Role        ctype.Role `json:"role" binding:"required" msg:"Please enter role"`
	Description string     `json:"description" binding:"required" msg:"Please enter description"`
}

func (UserApi) RoleCreateView(ctx *gin.Context) {
	var cr RoleCreateRequest
	if err := ctx.ShouldBindJSON(&cr); err != nil {
		res.FailWithValidMsg(err, &cr, ctx)
		return
	}

	err := user_service.CreateRole(cr.Role, cr.Description)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
		return
	}

	res.OkWithMessage(fmt.Sprintf("Role %s was created successfully", cr.Role.String()), ctx)
}
