package main

import (
	"auth-backend/app"
	"auth-backend/app/router"
	"auth-backend/config"
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

	router.Run(":" + port)
}
