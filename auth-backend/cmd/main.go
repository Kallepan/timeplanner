package main

import (
	"auth-backend/app/router"
	"auth-backend/config"
	"os"
)

func main() {
	port := os.Getenv("AUTH_BACKEND_PORT")
	if port == "" {
		port = "8080"
	}

	config.InitLogger()

	init := config.Init()	
	app := router.Init(init)

	app.Run(":" + port)
}