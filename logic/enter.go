package logic

import (
	"base_frame/logic/user"
)

type logicGroup struct {
	UserLogic user.LogicGroup
}

var LogicGroupApp = new(logicGroup)
