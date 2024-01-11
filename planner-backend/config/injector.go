package config

import (
	"planner-backend/app/controller"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Injector struct {
	DB             *neo4j.DriverWithContext
	SystemCtrl     controller.SystemController
	DepartmentCtrl controller.DepartmentController
	WorkplaceCtrl  controller.WorkplaceController
	TimeslotCtrl   controller.TimeslotController
	WeekdayCtrl    controller.WeekdayController
}
