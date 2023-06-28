package routers

import (
	"0049-server-go/api"
)

func (router RouterGroup) UsersRouter() {
	userApi := api.ApiGroupApp.UserApi
	router.GET("/users/list", userApi.UserListView)
	router.GET("/users/perm/list", userApi.PermissionListView)
	router.POST("/users/login", userApi.UserLoginView)
	router.POST("/users/user", userApi.UserCreateView)
	router.POST("/users/role", userApi.RoleCreateView)
	router.POST("/users/perm", userApi.PermissionCreateView)
	router.POST("/users/role_perm", userApi.RolePermissionCreateView)
	router.POST("/users/logout", userApi.UserLogoutView)

	router.GET("/users/pb", userApi.PbListView)
	router.POST("/users/iris", userApi.IrisPredictView)
}
