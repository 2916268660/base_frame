package routers

import (
	"base_frame/routers/user"
)

type routersGroup struct {
	UserRouterGroup user.RouterGroup
}

var RoutersGroupApp = new(routersGroup)
