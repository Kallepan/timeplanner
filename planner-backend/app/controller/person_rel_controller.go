package controller

import (
	"planner-backend/app/constant"
	"planner-backend/app/pkg"
	"planner-backend/app/service"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type PersonRelController interface {
	AddAbsency(ctx *gin.Context)
	RemoveAbsency(ctx *gin.Context)
	FindAbsencyForPerson(ctx *gin.Context)

	AddDepartment(ctx *gin.Context)
	RemoveDepartment(ctx *gin.Context)

	AddWorkplace(ctx *gin.Context)
	RemoveWorkplace(ctx *gin.Context)

	AddWeekday(ctx *gin.Context)
	RemoveWeekday(ctx *gin.Context)
}

type PersonRelControllerImpl struct {
	PersonRelService service.PersonRelService
}

func (u PersonRelControllerImpl) AddAbsency(ctx *gin.Context) {
	u.PersonRelService.AddAbsencyToPerson(ctx)
}

func (u PersonRelControllerImpl) RemoveAbsency(ctx *gin.Context) {
	u.PersonRelService.RemoveAbsencyFromPerson(ctx)
}

func (u PersonRelControllerImpl) FindAbsencyForPerson(ctx *gin.Context) {
	/** Find absency for a person
	* if startDate and endDate are present, call FindAbsencyForPersonInRange
	* else call FindAbsencyForPerson
	 */
	if ctx.Query("start_date") != "" && ctx.Query("end_date") != "" {
		u.PersonRelService.FindAbsencyForPersonInRange(ctx)
		return
	}

	if ctx.Query("date") != "" {
		u.PersonRelService.FindAbsencyForPerson(ctx)
		return
	}

	pkg.SendResponse(ctx, constant.InvalidRequest, pkg.Null())
}

func (u PersonRelControllerImpl) AddDepartment(ctx *gin.Context) {
	u.PersonRelService.AddDepartmentToPerson(ctx)
}

func (u PersonRelControllerImpl) RemoveDepartment(ctx *gin.Context) {
	u.PersonRelService.RemoveDepartmentFromPerson(ctx)
}

func (u PersonRelControllerImpl) AddWorkplace(ctx *gin.Context) {
	u.PersonRelService.AddWorkplaceToPerson(ctx)
}

func (u PersonRelControllerImpl) RemoveWorkplace(ctx *gin.Context) {
	u.PersonRelService.RemoveWorkplaceFromPerson(ctx)
}

func (u PersonRelControllerImpl) AddWeekday(ctx *gin.Context) {
	u.PersonRelService.AddWeekdayToPerson(ctx)
}

func (u PersonRelControllerImpl) RemoveWeekday(ctx *gin.Context) {
	u.PersonRelService.RemoveWeekdayFromPerson(ctx)
}

var personRelControllerSet = wire.NewSet(
	wire.Struct(new(PersonRelControllerImpl), "*"),
	wire.Bind(new(PersonRelController), new(*PersonRelControllerImpl)),
)
