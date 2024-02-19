package controller

import (
	"planner-backend/app/service"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type AbsenceController interface {
	GetAll(ctx *gin.Context)
}

type AbsenceControllerImpl struct {
	AbsencyService service.AbsenceService
}

func (u AbsenceControllerImpl) GetAll(ctx *gin.Context) {
	u.AbsencyService.GetAllAbsencies(ctx)
}

var absenceControllerSet = wire.NewSet(
	wire.Struct(new(AbsenceControllerImpl), "*"),
	wire.Bind(new(AbsenceController), new(*AbsenceControllerImpl)),
)
