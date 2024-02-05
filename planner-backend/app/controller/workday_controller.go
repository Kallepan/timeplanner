package controller

import (
	"planner-backend/app/service"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type WorkdayController interface {
	GetWorkdaysForDepartmentAndDate(ctx *gin.Context)
	GetWorkday(ctx *gin.Context)
	UpdateWorkday(ctx *gin.Context)

	AssignPersonToWorkday(ctx *gin.Context)
	UnassignPersonFromWorkday(ctx *gin.Context)
}

type WorkdayControllerImpl struct {
	WorkdayService service.WorkdayService
}

func (w WorkdayControllerImpl) GetWorkdaysForDepartmentAndDate(ctx *gin.Context) {
	w.WorkdayService.GetWorkdaysForDepartmentAndDate(ctx)
}

func (w WorkdayControllerImpl) GetWorkday(ctx *gin.Context) {
	w.WorkdayService.GetWorkday(ctx)
}

func (w WorkdayControllerImpl) UpdateWorkday(ctx *gin.Context) {
	w.WorkdayService.UpdateWorkday(ctx)
}

func (w WorkdayControllerImpl) AssignPersonToWorkday(ctx *gin.Context) {
	w.WorkdayService.AssignPersonToWorkday(ctx)
}

func (w WorkdayControllerImpl) UnassignPersonFromWorkday(ctx *gin.Context) {
	w.WorkdayService.UnassignPersonFromWorkday(ctx)
}

var workdayControllerSet = wire.NewSet(
	wire.Struct(new(WorkdayControllerImpl), "*"),
	wire.Bind(new(WorkdayController), new(*WorkdayControllerImpl)),
)
