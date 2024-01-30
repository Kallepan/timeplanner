package controller

import (
	"planner-backend/app/service"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type WeekdayController interface {
	AddWeekdayToTimeslot(ctx *gin.Context)
	RemoveWeekdayFromTimeslot(ctx *gin.Context)
	BulkUpdateWeekdaysForTimeslot(ctx *gin.Context)
}

type WeekdayControllerImpl struct {
	WeekdayService service.WeekdayService
}

func (u WeekdayControllerImpl) AddWeekdayToTimeslot(ctx *gin.Context) {
	u.WeekdayService.AddWeekdayToTimeslot(ctx)
}

func (u WeekdayControllerImpl) RemoveWeekdayFromTimeslot(ctx *gin.Context) {
	u.WeekdayService.DeleteWeekdayFromTimeslot(ctx)
}

func (u WeekdayControllerImpl) BulkUpdateWeekdaysForTimeslot(ctx *gin.Context) {
	u.WeekdayService.BulkUpdateWeekdaysForTimeslot(ctx)
}

var weekdayControllerSet = wire.NewSet(
	wire.Struct(new(WeekdayControllerImpl), "*"),
	wire.Bind(new(WeekdayController), new(*WeekdayControllerImpl)),
)
