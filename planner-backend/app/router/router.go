package router

import (
	"planner-backend/config"

	"github.com/gin-gonic/gin"
)

func Init(init *config.Injector) *gin.Engine {
	router := gin.New()

	// gin Middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// insert custom middlewares here
	// router.Use()

	apiV1 := router.Group("/api/v1")
	{
		apiV1.GET("/ping", init.SystemCtrl.Ping)
	}

	return router
}
