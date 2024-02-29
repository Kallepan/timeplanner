package router

import (
	"os"
	"planner-backend/config"

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
	// router.Use()

	plannerAPI := router.Group("/api/v1/planner")
	{
		plannerAPI.GET("/ping", init.SystemCtrl.Ping)

		department := plannerAPI.Group("/department")
		{
			department.GET("/", init.DepartmentCtrl.GetAll)
			department.GET("/:departmentID", init.DepartmentCtrl.Get)

			workplace := department.Group("/:departmentID/workplace")
			workplace.GET("/", init.WorkplaceCtrl.GetAll)
			workplace.GET("/:workplaceID", init.WorkplaceCtrl.Get)

			timeslot := workplace.Group("/:workplaceID/timeslot")
			timeslot.GET("/", init.TimeslotCtrl.GetAll)
			timeslot.GET("/:timeslotID", init.TimeslotCtrl.Get)

			absency := department.Group("/:departmentID/absency")
			absency.GET("/", init.AbsenceCtrl.GetAll) // ?date=...
		}
		// secured routes
		departmentSecured := plannerAPI.Group("/department")
		//departmentSecured.Use(middleware.RequiredAuth())
		{

			departmentSecured.POST("/", init.DepartmentCtrl.Create)
			departmentSecured.PUT("/:departmentID", init.DepartmentCtrl.Update)
			departmentSecured.DELETE("/:departmentID", init.DepartmentCtrl.Delete)

			workplaceSecured := departmentSecured.Group("/:departmentID/workplace")
			workplaceSecured.POST("/", init.WorkplaceCtrl.Create)
			workplaceSecured.PUT("/:workplaceID", init.WorkplaceCtrl.Update)
			workplaceSecured.DELETE("/:workplaceID", init.WorkplaceCtrl.Delete)

			timeslotSecured := workplaceSecured.Group("/:workplaceID/timeslot")
			timeslotSecured.POST("/", init.TimeslotCtrl.Create)
			timeslotSecured.PUT("/:timeslotID", init.TimeslotCtrl.Update)
			timeslotSecured.DELETE("/:timeslotID", init.TimeslotCtrl.Delete)

			weekdaySecured := timeslotSecured.Group("/:timeslotID/weekday")
			weekdaySecured.POST("/", init.WeekdayCtrl.AddWeekdayToTimeslot)
			weekdaySecured.DELETE("/", init.WeekdayCtrl.RemoveWeekdayFromTimeslot)
			weekdaySecured.POST("/bulk", init.WeekdayCtrl.BulkUpdateWeekdaysForTimeslot)

		}

		person := plannerAPI.Group("/person")
		{
			person.GET("/", init.PersonCtrl.GetAll) // ?departmentID=...
			person.GET("/:personID", init.PersonCtrl.Get)

			personRel := person.Group("/:personID")
			{
				personRel.GET("/absency", init.PersonRelCtrl.FindAbsencyForPerson) // ?date=... or ?start_date=...&end_date=...
			}
		}
		// secured routes
		personSecured := plannerAPI.Group("/person")
		//personSecured.Use(middleware.RequiredAuth())
		{
			personSecured.POST("/", init.PersonCtrl.Create)
			personSecured.PUT("/:personID", init.PersonCtrl.Update)
			personSecured.DELETE("/:personID", init.PersonCtrl.Delete)

			personRelSecured := personSecured.Group("/:personID")
			{
				personRelSecured.POST("/absency", init.PersonRelCtrl.AddAbsency)
				personRelSecured.DELETE("/absency/:date", init.PersonRelCtrl.RemoveAbsency)

				personRelSecured.POST("/department", init.PersonRelCtrl.AddDepartment)
				personRelSecured.DELETE("/department/:departmentID", init.PersonRelCtrl.RemoveDepartment)

				personRelSecured.POST("/workplace", init.PersonRelCtrl.AddWorkplace)
				personRelSecured.DELETE("/workplace", init.PersonRelCtrl.RemoveWorkplace)

				personRelSecured.POST("/weekday", init.PersonRelCtrl.AddWeekday)
				personRelSecured.DELETE("/weekday/:weekdayID", init.PersonRelCtrl.RemoveWeekday)
			}

		}

		workday := plannerAPI.Group("/workday")
		{
			workday.GET("/", init.WorkdayCtrl.GetWorkdaysForDepartmentAndDate) // ?departmentID=...&date=...
			workday.GET("/detail", init.WorkdayCtrl.GetWorkday)                // ?departmentID=...&date=...&workplaceID=...&timeslotID=...
		}

		// secured routes
		workdaySecured := plannerAPI.Group("/workday")
		//workdaySecured.Use(middleware.RequiredAuth())
		{
			workdaySecured.PUT("/", init.WorkdayCtrl.UpdateWorkday)
			workdaySecured.POST("/assign", init.WorkdayCtrl.AssignPersonToWorkday)
			workdaySecured.DELETE("/assign", init.WorkdayCtrl.UnassignPersonFromWorkday)
		}
	}

	return router
}
