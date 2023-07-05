package middleware

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/ctype"
	"0049-server-go/models/res"
	"0049-server-go/services/redis_service"
	"0049-server-go/services/user_service"
	"0049-server-go/utils"
	"0049-server-go/utils/jwts"
	"fmt"
	"github.com/gin-gonic/gin"
)

var authWhiteList = []string{
	"/api/v1/users/login-POST",
}

func JwtAuth() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		//whether this request is on the whitelist
		inWhiteList := utils.InList(fmt.Sprintf("%s-%s", ctx.Request.URL, ctx.Request.Method), authWhiteList)
		if inWhiteList {
			ctx.Next()
			return
		}

		token := ctx.Request.Header.Get("token")
		if token == "" {
			global.Log.Error("No token carried")
			res.FailWithCode(res.PermissionError, ctx)
			ctx.Abort()
			return
		}

		claims, err := jwts.ParseToken(token)
		if err != nil {
			global.Log.Error("Token error: ", err)
			res.FailWithCode(res.PermissionError, ctx)
			ctx.Abort()
			return
		}

		// whether in redis
		if redis_service.CheckLogout(token) {
			global.Log.Error("Token has expired:  ", err)
			res.FailWithCode(res.PermissionError, ctx)
			ctx.Abort()
			return
		}

		// super admin
		if ctype.Role(claims.Role) == ctype.RoleSuperAdmin {
			ctx.Set("claims", claims)
			ctx.Next()
			return
		}

		rolePermissionModels, err := user_service.GetPermissionByRole(ctype.Role(claims.Role))
		if err != nil {
			global.Log.Error(err.Error())
			res.FailWithCode(res.PermissionError, ctx)
			ctx.Abort()
			return
		}

		permissionList, err := GetPermissionList(rolePermissionModels)
		if err != nil {
			res.FailWithCode(res.PermissionError, ctx)
			ctx.Abort()
			return
		}

		inList := utils.InList(fmt.Sprintf("%s-%s", ctx.Request.URL, ctx.Request.Method), permissionList)
		if !inList {
			res.FailWithCode(res.PermissionError, ctx)
			ctx.Abort()
			return
		}

		ctx.Set("claims", claims)
	}
}

func GetPermissionList(rolePermissionModels []models.RolePermissionModel) ([]string, error) {
	result := make([]string, len(rolePermissionModels))
	for _, v := range rolePermissionModels {
		permissionModel, err := user_service.GetPermissionById(v.PermissionID)
		if err != nil {
			global.Log.Error(err)
			return nil, err
		}
		result = append(result, fmt.Sprintf("%s-%s", permissionModel.URL, permissionModel.Method))
	}

	return result, nil
}
