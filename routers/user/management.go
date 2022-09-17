package user

import (
	"github.com/gin-gonic/gin"
)

type ManagementRouter struct {
}

func (m ManagementRouter) InitUserRouters(router *gin.RouterGroup) {
	group := router.Group("user")

	{
		// 注册
		group.POST("register", userApi.RegisterUser)
		// 登录
		group.POST("login", userApi.Login)
		// 获取用户信息
		group.GET("details", userApi.GetUserInfo)
		// 修该用户信息
		group.PUT("update", userApi.UpdateUser)
		// 修改用户密码
		group.PUT("update_pass", userApi.UpdatePass)
		// 删除用户
		group.DELETE("del", userApi.DelUser)
	}
}
