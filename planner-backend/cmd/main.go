package main

import (
	"os"
	"planner-backend/app"
	"planner-backend/app/router"
	"planner-backend/config"
)

func main() {
	port := os.Getenv("PLANNER_BACKEND_PORT")
	if port == "" {
		port = "8081"
	}

	config.InitLogger()

	init, _, _ := app.BuildInjector()

	router := router.Init(init)

	router.Run(":" + port)
}
