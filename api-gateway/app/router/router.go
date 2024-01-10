package router

import (
	"api-gateway/app/middleware"
	"api-gateway/config"

	"github.com/gin-gonic/gin"
)

func Init(init *config.Injector) *gin.Engine {
	router := gin.New()

	// gin Middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// insert custom middlewares here
	router.Use(middleware.Cors())

	auth := router.Group("/auth")
	auth.POST("/login", init.UserCtrl.Login)
	auth.POST("/logout", init.UserCtrl.Logout)
	auth.Use(middleware.RequiredAuth())
	auth.GET("/me", init.UserCtrl.Me)

	/** These API requests stay here and are handled by api-gateway */
	gatewayAPI := router.Group("/api/v1")
	// TODO: apiV1.Use(middleware.RequiredAuth())
	{
		gatewayAPI.GET("/ping", init.SystemCtrl.Ping)

		user := gatewayAPI.Group("/user")
		user.GET("", init.UserCtrl.GetAll)
		user.GET("/:userID", init.UserCtrl.Get)
		user.POST("", init.UserCtrl.Create)
		user.PUT("/:userID", init.UserCtrl.Update)
		user.DELETE("/:userID", init.UserCtrl.Delete)
		userPermission := user.Group("/:userID/permission")
		userPermission.POST("/:permissionId", init.UserCtrl.AddPermission)
		userPermission.DELETE("/:permissionId", init.UserCtrl.DeletePermission)

		department := gatewayAPI.Group("/department")
		department.GET("", init.DepartmentCtrl.GetAll)
		department.GET("/:departmentID", init.DepartmentCtrl.Get)

		department.POST("", init.DepartmentCtrl.Create)
		department.PUT("/:departmentID", init.DepartmentCtrl.Update)
		department.DELETE("/:departmentID", init.DepartmentCtrl.Delete)

		permission := gatewayAPI.Group("/permission")
		permission.GET("", init.PermissionCtrl.GetAll)
		permission.GET("/:permissionID", init.PermissionCtrl.Get)
		permission.POST("", init.PermissionCtrl.Create)
		permission.PUT("/:permissionID", init.PermissionCtrl.Update)
		permission.DELETE("/:permissionID", init.PermissionCtrl.Delete)
	}

	return router
}
