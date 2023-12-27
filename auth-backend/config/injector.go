package config

import (
	"auth-backend/app/controller"
)

type Injector struct {
	UserCtrl       controller.UserController
	DepartmentCtrl controller.DepartmentController
	PermissionCtrl controller.PermissionController
}
