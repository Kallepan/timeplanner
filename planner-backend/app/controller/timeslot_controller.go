package controller

import (
	"planner-backend/app/service"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type TimeslotController interface {
	GetAll(ctx *gin.Context)
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type TimeslotControllerImpl struct {
	TimeslotService service.TimeslotService
}

func (u TimeslotControllerImpl) GetAll(ctx *gin.Context) {
	u.TimeslotService.GetAllTimeslots(ctx)
}

func (u TimeslotControllerImpl) Get(ctx *gin.Context) {
	u.TimeslotService.GetTimeslotByID(ctx)
}

func (u TimeslotControllerImpl) Create(ctx *gin.Context) {
	u.TimeslotService.AddTimeslot(ctx)
}

func (u TimeslotControllerImpl) Update(ctx *gin.Context) {
	u.TimeslotService.UpdateTimeslot(ctx)
}

func (u TimeslotControllerImpl) Delete(ctx *gin.Context) {
	u.TimeslotService.DeleteTimeslot(ctx)
}

var timeslotControllerSet = wire.NewSet(
	wire.Struct(new(TimeslotControllerImpl), "*"),
	wire.Bind(new(TimeslotController), new(*TimeslotControllerImpl)),
)
