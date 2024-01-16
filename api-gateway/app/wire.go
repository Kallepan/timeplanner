//go:build wireinject
// +build wireinject

package app

import (
	"api-gateway/app/controller"
	"api-gateway/app/repository"
	"api-gateway/app/service"
	"api-gateway/config"

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
