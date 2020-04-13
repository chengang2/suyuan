package bll

import (
	"suyuan/service/menu_service"
	"suyuan/service/role_service"
	"suyuan/service/user_service"
)

type Common struct {
	UserAPI *user_service.User `inject:""`
	RoleAPI *role_service.Role `inject:""`
	MenuAPI *menu_service.Menu `inject:""`
}
