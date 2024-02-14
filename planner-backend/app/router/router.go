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

	plannerAPI := router.Group("/api/v1/planner")
	{
		plannerAPI.GET("/ping", init.SystemCtrl.Ping)

		department := plannerAPI.Group("/department")
		{
			department.GET("/", init.DepartmentCtrl.GetAll)
			department.GET("/:departmentID", init.DepartmentCtrl.Get)
			department.POST("/", init.DepartmentCtrl.Create)
			department.PUT("/:departmentID", init.DepartmentCtrl.Update)
			department.DELETE("/:departmentID", init.DepartmentCtrl.Delete)

			workplace := department.Group("/:departmentID/workplace")
			workplace.GET("/", init.WorkplaceCtrl.GetAll)
			workplace.GET("/:workplaceID", init.WorkplaceCtrl.Get)
			workplace.POST("/", init.WorkplaceCtrl.Create)
			workplace.PUT("/:workplaceID", init.WorkplaceCtrl.Update)
			workplace.DELETE("/:workplaceID", init.WorkplaceCtrl.Delete)

			timeslot := workplace.Group("/:workplaceID/timeslot")
			timeslot.GET("/", init.TimeslotCtrl.GetAll)
			timeslot.GET("/:timeslotID", init.TimeslotCtrl.Get)
			timeslot.POST("/", init.TimeslotCtrl.Create)
			timeslot.PUT("/:timeslotID", init.TimeslotCtrl.Update)
			timeslot.DELETE("/:timeslotID", init.TimeslotCtrl.Delete)

			weekday := timeslot.Group("/:timeslotID/weekday")
			weekday.POST("/", init.WeekdayCtrl.AddWeekdayToTimeslot)
			weekday.DELETE("/", init.WeekdayCtrl.RemoveWeekdayFromTimeslot)
			weekday.POST("/bulk", init.WeekdayCtrl.BulkUpdateWeekdaysForTimeslot)

		}
		person := plannerAPI.Group("/person")
		{
			person.GET("/", init.PersonCtrl.GetAll)
			person.GET("/:personID", init.PersonCtrl.Get)
			person.POST("/", init.PersonCtrl.Create)
			person.PUT("/:personID", init.PersonCtrl.Update)
			person.DELETE("/:personID", init.PersonCtrl.Delete)

			personRel := person.Group("/:personID")
			{
				personRel.POST("/absency", init.PersonRelCtrl.AddAbsency)
				personRel.DELETE("/absency/:date", init.PersonRelCtrl.RemoveAbsency)
				personRel.GET("/absency", init.PersonRelCtrl.FindAbsencyForPerson) // ?date=... or ?start_date=...&end_date=...

				personRel.POST("/department", init.PersonRelCtrl.AddDepartment)
				personRel.DELETE("/department/:departmentID", init.PersonRelCtrl.RemoveDepartment)

				personRel.POST("/workplace", init.PersonRelCtrl.AddWorkplace)
				personRel.DELETE("/workplace", init.PersonRelCtrl.RemoveWorkplace)

				personRel.POST("/weekday", init.PersonRelCtrl.AddWeekday)
				personRel.DELETE("/weekday/:weekdayID", init.PersonRelCtrl.RemoveWeekday)
			}
		}
		workday := plannerAPI.Group("/workday")
		{
			workday.GET("/", init.WorkdayCtrl.GetWorkdaysForDepartmentAndDate) // ?departmentID=...&date=...
			workday.GET("/detail", init.WorkdayCtrl.GetWorkday)                // ?departmentID=...&date=...&workplaceID=...&timeslotID=...
			workday.PUT("/", init.WorkdayCtrl.UpdateWorkday)
			workday.POST("/assign", init.WorkdayCtrl.AssignPersonToWorkday)
			workday.DELETE("/assign", init.WorkdayCtrl.UnassignPersonFromWorkday)
		}
	}

	return router
}
