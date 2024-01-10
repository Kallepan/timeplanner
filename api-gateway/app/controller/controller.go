package controller

import "github.com/google/wire"

var ControllerSet = wire.NewSet(
	userControllerSet,
	permissionControllerSet,
	departmentControllerSet,
	systemControllerSet,
)
