package main

import (
	"api-gateway/app"
	"api-gateway/app/router"
	"api-gateway/config"
	"os"
)

func main() {
	port := os.Getenv("AUTH_BACKEND_PORT")
	if port == "" {
		port = "8081"
	}

	config.InitLogger()

	init, _, _ := app.BuildInjector()
	router := router.Init(init)

	// run migration
	config.Migrate(init.DB)

	router.Run(":" + port)
}
