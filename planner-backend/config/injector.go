package config

import (
	"planner-backend/app/controller"
	"planner-backend/app/repository"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Injector struct {
	DB              *neo4j.DriverWithContext
	SystemCtrl      controller.SystemController
	DepartmentCtrl  controller.DepartmentController
	WorkplaceCtrl   controller.WorkplaceController
	TimeslotCtrl    controller.TimeslotController
	WeekdayCtrl     controller.WeekdayController
	PersonCtrl      controller.PersonController
	PersonRelCtrl   controller.PersonRelController
	WorkdayCtrl     controller.WorkdayController
	AbsenceCtrl     controller.AbsenceController
	SynchronizeRepo repository.SynchronizeRepository
}
