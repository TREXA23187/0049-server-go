package users_api

import (
	"0049-server-go/global"
	"0049-server-go/models/ctype"
	"0049-server-go/models/res"
	"0049-server-go/services/user_service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type PermissionRequest struct {
	URL         string `json:"url" binding:"required" msg:"Please enter URL"`
	Method      string `json:"method" binding:"required" msg:"Please enter method"`
	Description string `json:"description" binding:"required" msg:"Please enter description"`
}

func (UserApi) PermissionCreateView(ctx *gin.Context) {
	var cr PermissionRequest
	if err := ctx.ShouldBindJSON(&cr); err != nil {
		res.FailWithValidMsg(err, &cr, ctx)
		return
	}
	fmt.Println(cr.Method)
	err := user_service.CreatePermission(cr.URL, ctype.StringToPermissionMethod(cr.Method), cr.Description)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
		return
	}

	res.OkWithMessage(fmt.Sprintf("Permission %s (%s) was created successfully", cr.URL, cr.Method), ctx)
}
