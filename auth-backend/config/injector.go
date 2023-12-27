// go:build wireinject
//go:build wireinject
// +build wireinject

package config

import (
	"auth-backend/app/controller"
	"auth-backend/app/repository"
	"auth-backend/app/service"

	"github.com/google/wire"
)

var db = wire.NewSet(ConnectToDB)

/* User */
var userServiceSet = wire.NewSet(service.UserServiceInit,
	wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
)
var userRepoSet = wire.NewSet(repository.UserRepositoryInit,
	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
)
var userCtrlSet = wire.NewSet(controller.UserControllerInit,
	wire.Bind(new(controller.UserController), new(*controller.UserControllerImpl)),
)

/* Department */
var departmentServiceSet = wire.NewSet(service.DepartmentServiceInit,
	wire.Bind(new(service.DepartmentService), new(*service.DepartmentServiceImpl)),
)
var departmentRepoSet = wire.NewSet(repository.DepartmentRepositoryInit,
	wire.Bind(new(repository.DepartmentRepository), new(*repository.DepartmentRepositoryImpl)),
)
var departmentCtrlSet = wire.NewSet(controller.DepartmentControllerInit,
	wire.Bind(new(controller.DepartmentController), new(*controller.DepartmentControllerImpl)),
)

/* Permission */
var permissionServiceSet = wire.NewSet(service.PermissionServiceInit,
	wire.Bind(new(service.PermissionService), new(*service.PermissionServiceImpl)),
)
var permissionRepoSet = wire.NewSet(repository.PermissionRepositoryInit,
	wire.Bind(new(repository.PermissionRepository), new(*repository.PermissionRepositoryImpl)),
)
var permissionCtrlSet = wire.NewSet(controller.PermissionControllerInit,
	wire.Bind(new(controller.PermissionController), new(*controller.PermissionControllerImpl)),
)

func Init() *Initialization {
	wire.Build(
		NewInitialization, 
		db, 
		userCtrlSet, 
		userServiceSet, 
		userRepoSet,
		departmentCtrlSet,
		departmentServiceSet,
		departmentRepoSet,
		permissionCtrlSet,
		permissionServiceSet,
		permissionRepoSet,
	)
	
	return nil
}