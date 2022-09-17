package user

import (
	"base_frame/global"
	"github.com/gin-gonic/gin"
)

type ManagementApi struct {
}

// Login 登录
func (m *ManagementApi) Login(ctx *gin.Context) {
	global.OkWithDetails(ctx, "登录成功", map[string]string{"token": "Bearer "})
}

// RegisterUser 注册用户
func (m *ManagementApi) RegisterUser(ctx *gin.Context) {
	global.OkWithMsg(ctx, "注册成功")
}

// GetUserInfo 获取用户相信信息
func (m *ManagementApi) GetUserInfo(ctx *gin.Context) {
	global.OkWithMsg(ctx, "获取成功")
}

// UpdateUser 更新用户信息
func (m *ManagementApi) UpdateUser(ctx *gin.Context) {
	global.OkWithMsg(ctx, "更改成功")
}

// UpdatePass 修改密码
func (m *ManagementApi) UpdatePass(ctx *gin.Context) {
	global.OkWithMsg(ctx, "修改成功")
}

func (m *ManagementApi) DelUser(ctx *gin.Context) {

	global.OkWithMsg(ctx, "删除成功")
}
