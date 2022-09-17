package user

import "base_frame/logic"

type ApiGroup struct {
	ManagementApi
}

var userLogic = logic.LogicGroupApp.UserLogic.ManagementLogic
