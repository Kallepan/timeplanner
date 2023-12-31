package config

import (
	"auth-backend/app/controller"

	"gorm.io/gorm"
)

type Injector struct {
	DB             *gorm.DB
	UserCtrl       controller.UserController
	DepartmentCtrl controller.DepartmentController
	PermissionCtrl controller.PermissionController
}
