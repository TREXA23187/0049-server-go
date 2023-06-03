package routers

import "0049-server-go/api"

func (router RouterGroup) UsersRouter() {
	userApi := api.ApiGroupApp.UserApi
	router.POST("login", userApi.UserLoginView)
	router.POST("user", userApi.UserCreateView)
	router.GET("users", userApi.UserListView)
	router.POST("role", userApi.RoleCreateView)
	router.POST("perm", userApi.PermissionCreateView)
}
