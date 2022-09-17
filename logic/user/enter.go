package user

import (
	"base_frame/repository"
)

type LogicGroup struct {
	ManagementLogic
}

var userModel = repository.ModelGroupApp.UserModel
