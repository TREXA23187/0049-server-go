package users_api

import (
	"0049-server-go/global"
	"0049-server-go/models/res"
	"0049-server-go/services"
	"0049-server-go/utils/jwts"
	"github.com/gin-gonic/gin"
)

func (UserApi) UserLogoutView(ctx *gin.Context) {
	_claims, _ := ctx.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	token := ctx.Request.Header.Get("token")

	// Calculate the expiration time from now
	err := services.ServiceApp.UserService.Logout(claims, token)

	if err != nil {
		global.Log.Error()
		res.FailWithMessage("Logout failed", ctx)
		return
	}

	res.OkWithMessage("Logged out successfully", ctx)
}
