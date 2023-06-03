package routers

import "0049-server-go/api"

func (router RouterGroup) UsersRouter() {
	userApi := api.ApiGroupApp.UserApi
	router.POST("login", userApi.UserLoginView)
	router.POST("users", userApi.UserCreateView)
	router.GET("users", userApi.UserListView)
}
