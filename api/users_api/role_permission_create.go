package users_api

import (
	"0049-server-go/global"
	"0049-server-go/models/ctype"
	"0049-server-go/models/res"
	"0049-server-go/services/user_service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type RolePermissionCreateRequest struct {
	Role   ctype.Role `json:"role" binding:"required" msg:"Please enter role"`
	URL    string     `json:"url" binding:"required" msg:"Please enter URL"`
	Method string     `json:"method" binding:"required" msg:"Please enter method"`
}

func (UserApi) RolePermissionCreateView(ctx *gin.Context) {
	var cr RolePermissionCreateRequest
	if err := ctx.ShouldBindJSON(&cr); err != nil {
		res.FailWithValidMsg(err, &cr, ctx)
		return
	}

	if cr.Role == ctype.RoleSuperAdmin {
		res.OkWithMessage("This user is the super admin", ctx)
	}

	err := user_service.GrantPermissionToRole(cr.Role, cr.URL, ctype.StringToPermissionMethod(cr.Method))
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
		return
	}

	res.OkWithMessage(fmt.Sprintf("Role '%s' has been granted permission '%s (%s)' successfully", cr.Role.String(), cr.URL, cr.Method), ctx)
}
