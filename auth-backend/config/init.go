package config

import (
	"auth-backend/app/controller"
	"auth-backend/app/repository"
	"auth-backend/app/service"
)

type Initialization struct {
	userRepo repository.UserRepository
	userSvc service.UserService
	UserCtrl controller.UserController
	departmentRepo repository.DepartmentRepository
	departmentSvc service.DepartmentService
	DepartmentCtrl controller.DepartmentController
	permissionRepo repository.PermissionRepository
	permissionSvc service.PermissionService
	PermissionCtrl controller.PermissionController
}

func NewInitialization(
	userRepo repository.UserRepository,
	userSvc service.UserService,
	userCtrl controller.UserController,
	departmentRepo repository.DepartmentRepository,
	departmentSvc service.DepartmentService,
	departmentCtrl controller.DepartmentController,
	permissionRepo repository.PermissionRepository,
	permissionSvc service.PermissionService,
	permissionCtrl controller.PermissionController,
) *Initialization {
	return &Initialization{
		userRepo: userRepo,
		userSvc: userSvc,
		UserCtrl: userCtrl,
		departmentRepo: departmentRepo,
		departmentSvc: departmentSvc,
		DepartmentCtrl: departmentCtrl,
		permissionRepo: permissionRepo,
		permissionSvc: permissionSvc,
		PermissionCtrl: permissionCtrl,
	}
}