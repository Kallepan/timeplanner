package router

import (
	"api-gateway/app/middleware"
	"api-gateway/config"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

func Init(init *config.Injector) *gin.Engine {
	// set gin to release mode
	if os.Getenv("MODE") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	// gin Middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// insert custom middlewares here
	router.Use(middleware.Cors())

	auth := router.Group("/auth")
	{
		auth.POST("/login", init.UserCtrl.Login)
		auth.POST("/logout", init.UserCtrl.Logout)
		auth.Use(middleware.RequiredAuth())
		auth.GET("/me", init.UserCtrl.Me) // ?department=XXX
		auth.GET("/check-admin", init.UserCtrl.CheckAdmin)
	}

	/** These API requests stay here and are handled by api-gateway */
	gatewayAPI := router.Group("/api/v1")
	{
		gatewayAPI.GET("/ping", init.SystemCtrl.Ping)

		// Secured routes
		gatewayAPI.Use(middleware.RequiredAuth())
		user := gatewayAPI.Group("/user")
		{
			user.GET("", init.UserCtrl.GetAll)
			user.GET("/:userID", init.UserCtrl.Get)
			user.POST("", init.UserCtrl.Create)
			user.PUT("/:userID", init.UserCtrl.Update)
			user.DELETE("/:userID", init.UserCtrl.Delete)
		}
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

	/** These API requests are forwarded to planner-backend service */
	// TODO: Add middleware to check if the user has permission to access the planner-backend
	// The problem is that GET requests should be allowed for everyone, but POST, PUT, DELETE should be allowed only for users with the right permissions
	// The middleware should check the user's permissions and the request method
	plannerAPI := router.Group("/api/v1/planner")
	{
		targetStr := os.Getenv("PLANNER_BACKEND_TARGET")
		url, _ := url.Parse(targetStr)
		proxy := httputil.NewSingleHostReverseProxy(url)
		plannerAPI.Any("/*any", func(c *gin.Context) {
			proxy.ServeHTTP(c.Writer, c.Request)
		})
	}

	return router
}
