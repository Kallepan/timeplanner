package controller

import (
	"planner-backend/app/service"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type WorkplaceController interface {
	GetAll(ctx *gin.Context)
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type WorkplaceControllerImpl struct {
	WorkplaceService service.WorkplaceService
}

func (u WorkplaceControllerImpl) GetAll(ctx *gin.Context) {
	u.WorkplaceService.GetAllWorkplaces(ctx)
}

func (u WorkplaceControllerImpl) Get(ctx *gin.Context) {
	u.WorkplaceService.GetWorkplaceByName(ctx)
}

func (u WorkplaceControllerImpl) Create(ctx *gin.Context) {
	u.WorkplaceService.AddWorkplace(ctx)
}

func (u WorkplaceControllerImpl) Update(ctx *gin.Context) {
	u.WorkplaceService.UpdateWorkplace(ctx)
}

func (u WorkplaceControllerImpl) Delete(ctx *gin.Context) {
	u.WorkplaceService.DeleteWorkplace(ctx)
}

var workplaceControllerSet = wire.NewSet(
	wire.Struct(new(WorkplaceControllerImpl), "*"),
	wire.Bind(new(WorkplaceController), new(*WorkplaceControllerImpl)),
)
