package main

import (
	"context"
	"os"
	"planner-backend/app"
	"planner-backend/app/repository"
	"planner-backend/app/router"
	"planner-backend/config"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	port := os.Getenv("PLANNER_BACKEND_PORT")
	if port == "" {
		port = "8081"
	}

	config.InitLogger()

	init, _, _ := app.BuildInjector(ctx)

	// I dont know where else to fit this in
	repository.Migrate(ctx, init.DB)

	router := router.Init(init)

	app.InitalizeSynchronization(init)

	router.Run(":" + port)
}
