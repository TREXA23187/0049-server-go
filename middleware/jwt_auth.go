package middleware

import (
	"0049-server-go/models/res"
	"0049-server-go/utils/jwts"
	"fmt"
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

		// whether in redis
		//if redis_service.CheckLogout(token) {
		//	res.FailWithMessage("Token has expired", ctx)
		//	ctx.Abort()
		//	return
		//}

		fmt.Println(claims.Role)
		fmt.Println(ctx.Request.Method)
		fmt.Println(ctx.Request.URL)

		ctx.Set("claims", claims)
	}
}
