package middleware

import (
	"0049-server-go/models/ctype"
	"0049-server-go/models/res"
	"0049-server-go/services/redis_service"
	"0049-server-go/utils/jwts"
	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token == "" {
			res.FailWithMessage("No token carried", ctx)
			ctx.Abort()
			return
		}
		claims, err := jwts.ParseToken(token)
		if err != nil {
			res.FailWithMessage("Token error", ctx)
			ctx.Abort()
			return
		}

		//是否在redis中
		if redis_service.CheckLogout(token) {
			res.FailWithMessage("Token has expired", ctx)
			ctx.Abort()
			return
		}

		ctx.Set("claims", claims)
	}
}

func JwtAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token == "" {
			res.FailWithMessage("No token carried", ctx)
			ctx.Abort()
			return
		}
		claims, err := jwts.ParseToken(token)
		if err != nil {
			res.FailWithMessage("Token error", ctx)
			ctx.Abort()
			return
		}

		//是否在redis中
		if redis_service.CheckLogout(token) {
			res.FailWithMessage("Token has expired", ctx)
			ctx.Abort()
			return
		}

		// 登陆的用户
		if claims.Role != int(ctype.PermissionAdmin) {
			res.FailWithMessage("Permission error", ctx)
			ctx.Abort()
			return
		}
		ctx.Set("claims", claims)
	}
}
