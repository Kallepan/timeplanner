//go:build wireinject
// +build wireinject

package app

import (
	"auth-backend/app/controller"
	"auth-backend/app/repository"
	"auth-backend/app/service"
	"auth-backend/config"

	"github.com/google/wire"
)

var db = wire.NewSet(config.ConnectToDB)

var InjectorSet = wire.NewSet(
	wire.Struct(new(config.Injector), "*"))

func BuildInjector() (*config.Injector, func(), error) {
	wire.Build(
		db,
		repository.RepositorySet,
		service.ServiceSet,
		controller.ControllerSet,
		InjectorSet,
	)

	return new(config.Injector), nil, nil
}
