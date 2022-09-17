package v1

import (
	"base_frame/api/v1/user"
)

type apiGroup struct {
	UserApi user.ApiGroup
}

var ApiGroupApp = new(apiGroup)
