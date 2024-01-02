package router

import (
	"auth-backend/app/middleware"
	"auth-backend/config"

	"github.com/gin-gonic/gin"
)

func Init(init *config.Injector) *gin.Engine {
	router := gin.New()

	// gin Middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// insert middlewares here
	router.Use(middleware.Cors())

	api := router.Group("/api/v1")
	{
		user := api.Group("/user")
		user.GET("", init.UserCtrl.GetAll)
		user.GET("/:userID", init.UserCtrl.Get)
		user.POST("", init.UserCtrl.Create)
		user.PUT("/:userID", init.UserCtrl.Update)
		user.DELETE("/:userID", init.UserCtrl.Delete)
		userPermission := user.Group("/:userID/permission")
		userPermission.POST("/:userID/permission/:permissionId", init.UserCtrl.AddPermission)
		userPermission.DELETE("/:userID/permission/:permissionId", init.UserCtrl.DeletePermission)

		department := api.Group("/department")
		department.GET("", init.DepartmentCtrl.GetAll)
		department.GET("/:departmentID", init.DepartmentCtrl.Get)
		department.POST("", init.DepartmentCtrl.Create)
		department.PUT("/:departmentID", init.DepartmentCtrl.Update)
		department.DELETE("/:departmentID", init.DepartmentCtrl.Delete)

		permission := api.Group("/permission")
		permission.GET("", init.PermissionCtrl.GetAll)
		permission.GET("/:permissionID", init.PermissionCtrl.Get)
		permission.POST("", init.PermissionCtrl.Create)
		permission.PUT("/:permissionID", init.PermissionCtrl.Update)
		permission.DELETE("/:permissionID", init.PermissionCtrl.Delete)
	}

	return router
}
