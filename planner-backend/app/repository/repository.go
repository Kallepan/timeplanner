package repository

import "github.com/google/wire"

var RepositorySet = wire.NewSet(
	DepartmentRepositorySet,
	WorkplaceRepositorySet,
	timeslotRepositorySet,
	weekdayRepositorySet,
	personRepositorySet,
	personRelRepositorySet,
)
