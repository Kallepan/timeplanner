package repository

import "github.com/google/wire"

var RepositorySet = wire.NewSet(
	userRepositorySet,
	permissionRepositorySet,
	departmentRepositorySet,
)
