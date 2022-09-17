package middleware

import (
	"base_frame/global"
	"base_frame/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

func JWTAuthMiddleWare() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			global.FailWithMsg(ctx, "无权进行访问,请先登录")
			ctx.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			global.FailWithMsg(ctx, "身份失效,请重新登录")
			ctx.Abort()
			return
		}
		mc, err := utils.ParseToken(parts[1])
		if err != nil {
			if strings.Contains(err.Error(), "expired") {
				global.FailWithMsg(ctx, "身份失效,请重新登录")
				ctx.Abort()
				return
			}
			global.FailWithMsg(ctx, "无权进行访问,请先登录")
			ctx.Abort()
			return
		}
		ctx.Set("userId", mc.UserId)
		ctx.Next()
	}
}
