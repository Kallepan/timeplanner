package repository

import (
	"github.com/google/wire"
)

var RepositorySet = wire.NewSet(
	departmentRepositorySet,
	workplaceRepositorySet,
	timeslotRepositorySet,
	weekdayRepositorySet,
	personRepositorySet,
	personRelRepositorySet,
	synchronizeRepositorySet,
	workdayRepositorySet,
	absenceRepositorySet,
)
