package router

import (
	"auth-backend/config"

	"github.com/gin-gonic/gin"
)

func Init(init *config.Initialization) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api/v1")
	{
		user := api.Group("/user")
		user.GET("", init.UserCtrl.GetAll)
		user.GET("/:id", init.UserCtrl.Get)
		user.POST("", init.UserCtrl.Create)
		user.PUT("/:id", init.UserCtrl.Update)
		user.DELETE("/:id", init.UserCtrl.Delete)
		userPermission := user.Group("/:id/permission")
		userPermission.POST("/:userId/permission/:permissionId", init.UserCtrl.AddPermission)
		userPermission.DELETE("/:userId/permission/:permissionId", init.UserCtrl.DeletePermission)

		department := api.Group("/department")
		department.GET("", init.DepartmentCtrl.GetAll)
		department.GET("/:id", init.DepartmentCtrl.Get)
		department.POST("", init.DepartmentCtrl.Create)
		department.PUT("/:id", init.DepartmentCtrl.Update)
		department.DELETE("/:id", init.DepartmentCtrl.Delete)


		permission := api.Group("/permission")
		permission.GET("", init.PermissionCtrl.GetAll)
		permission.GET("/:id", init.PermissionCtrl.Get)
		permission.POST("", init.PermissionCtrl.Create)
		permission.PUT("/:id", init.PermissionCtrl.Update)
		permission.DELETE("/:id", init.PermissionCtrl.Delete)
	}
	
	return router
}