package repository

import (
	"base_frame/repository/user"
)

type modeGroup struct {
	UserModel user.ManagementModel
}

var ModelGroupApp = new(modeGroup)
