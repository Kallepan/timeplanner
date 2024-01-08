//go:build wireinject
// +build wireinject

package app

import (
	"context"
	"planner-backend/app/controller"
	"planner-backend/config"

	"github.com/google/wire"
)

var db = wire.NewSet(config.ConnectToDB, context.Background)

var InjectorSet = wire.NewSet(
	wire.Struct(new(config.Injector), "*"))

func BuildInjector() (*config.Injector, func(), error) {
	wire.Build(
		db,
		controller.ControllerSet,
		InjectorSet,
	)

	return new(config.Injector), nil, nil
}
