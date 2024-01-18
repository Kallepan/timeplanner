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

			timeslot := workplace.Group("/:workplaceName/timeslot")
			timeslot.GET("/", init.TimeslotCtrl.GetAll)
			timeslot.GET("/:timeslotName", init.TimeslotCtrl.Get)
			timeslot.POST("/", init.TimeslotCtrl.Create)
			timeslot.PUT("/:timeslotName", init.TimeslotCtrl.Update)
			timeslot.DELETE("/:timeslotName", init.TimeslotCtrl.Delete)

			weekday := timeslot.Group("/:timeslotName/weekday")
			weekday.POST("/", init.WeekdayCtrl.AddWeekdayToTimeslot)
			weekday.DELETE("/", init.WeekdayCtrl.RemoveWeekdayFromTimeslot)
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
				personRel.GET("/absency/:date", init.PersonRelCtrl.FindAbsencyForPerson)

				personRel.POST("/department", init.PersonRelCtrl.AddDepartment)
				personRel.DELETE("/department/:departmentName", init.PersonRelCtrl.RemoveDepartment)

				personRel.POST("/workplace", init.PersonRelCtrl.AddWorkplace)
				personRel.DELETE("/workplace", init.PersonRelCtrl.RemoveWorkplace)

				personRel.POST("/weekday", init.PersonRelCtrl.AddWeekday)
				personRel.DELETE("/weekday/:weekdayID", init.PersonRelCtrl.RemoveWeekday)
			}
		}
		workday := plannerAPI.Group("/workday")
		{
			workday.GET("/", init.WorkdayCtrl.GetWorkdaysForDepartmentAndDate) // ?departmentName=...&date=...
			workday.GET("/detail", init.WorkdayCtrl.GetWorkday)                // ?departmentName=...&date=...&workplaceName=...&timeslotName=...
			workday.POST("/assign", init.WorkdayCtrl.AssignPersonToWorkday)
			workday.DELETE("/assign", init.WorkdayCtrl.UnassignPersonFromWorkday)
		}
	}

	return router
}
