//go:build wireinject
// +build wireinject

package app

import (
	"context"
	"planner-backend/app/controller"
	"planner-backend/app/repository"
	"planner-backend/app/service"
	"planner-backend/config"

	"github.com/google/wire"
)

var db = wire.NewSet(config.ConnectToDB)

var InjectorSet = wire.NewSet(
	wire.Struct(new(config.Injector), "*"))

func BuildInjector(ctx context.Context) (*config.Injector, func(), error) {
	wire.Build(
		db,
		repository.RepositorySet,
		service.ServiceSet,
		controller.ControllerSet,
		InjectorSet,
	)

	return new(config.Injector), nil, nil
}
