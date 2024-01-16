package config

import (
	"api-gateway/app/controller"

	"gorm.io/gorm"
)

type Injector struct {
	DB             *gorm.DB
	SystemCtrl     controller.SystemController
	UserCtrl       controller.UserController
	DepartmentCtrl controller.DepartmentController
	PermissionCtrl controller.PermissionController
}
