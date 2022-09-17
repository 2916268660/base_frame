package user

import v1 "base_frame/api/v1"

type RouterGroup struct {
	ManagementRouter
}

var userApi = v1.ApiGroupApp.UserApi
