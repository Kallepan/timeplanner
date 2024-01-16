package controller

import "github.com/google/wire"

var ControllerSet = wire.NewSet(
	systemControllerSet,
	workplaceControllerSet,
	departmentControllerSet,
	timeslotControllerSet,
	weekdayControllerSet,
	personControllerSet,
	personRelControllerSet,
	workdayControllerSet,
)
