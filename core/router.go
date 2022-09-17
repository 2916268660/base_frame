package core

import (
	"base_frame/middleware"
	"base_frame/routers"
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	var router = gin.Default()
	// 配置跨域中间件
	router.Use(middleware.Cors())

	v1Router := router.Group("v1")
	// 获取用户路由组实例
	userRouters := routers.RoutersGroupApp.UserRouterGroup

	{
		userRouters.InitUserRouters(v1Router)
	}
	return router
}
