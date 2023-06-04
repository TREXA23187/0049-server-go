package routers

import (
	"0049-server-go/api"
)

func (router RouterGroup) UsersRouter() {
	userApi := api.ApiGroupApp.UserApi
	router.POST("/users/login", userApi.UserLoginView)
	router.POST("/users/user", userApi.UserCreateView)
	router.GET("/users/list", userApi.UserListView)
	router.POST("/users/role", userApi.RoleCreateView)
	router.POST("/users/perm", userApi.PermissionCreateView)
	router.POST("/users/role_perm", userApi.RolePermissionCreateView)
	router.POST("/users/logout", userApi.UserLogoutView)
}
