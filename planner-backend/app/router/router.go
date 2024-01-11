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

	apiV1 := router.Group("/api/v1/planner")
	{
		apiV1.GET("/ping", init.SystemCtrl.Ping)

		department := apiV1.Group("/department")
		{
			department.GET("/", init.DepartmentCtrl.GetAll)
			department.GET("/:departmentName", init.DepartmentCtrl.Get)
			department.POST("/", init.DepartmentCtrl.Create)
			department.PUT("/:departmentName", init.DepartmentCtrl.Update)
			department.DELETE("/:departmentName", init.DepartmentCtrl.Delete)

			workplace := department.Group("/:departmentName/workplace")
			workplace.GET("/", init.WorkplaceCtrl.GetAll)
			workplace.GET("/:workplaceName", init.WorkplaceCtrl.Get)
			workplace.POST("/", init.WorkplaceCtrl.Create)
			workplace.PUT("/:workplaceName", init.WorkplaceCtrl.Update)
			workplace.DELETE("/:workplaceName", init.WorkplaceCtrl.Delete)
		}
	}

	return router
}
