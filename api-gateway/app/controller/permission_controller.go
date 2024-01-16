package controller

import (
	"api-gateway/app/service"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type PermissionController interface {
	GetAll(ctx *gin.Context)
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type PermissionControllerImpl struct {
	PermissionService service.PermissionService
}

func (u PermissionControllerImpl) GetAll(ctx *gin.Context) {
	u.PermissionService.GetAllPermissions(ctx)
}

func (u PermissionControllerImpl) Get(ctx *gin.Context) {
	u.PermissionService.GetPermissionById(ctx)
}

func (u PermissionControllerImpl) Create(ctx *gin.Context) {
	u.PermissionService.AddPermission(ctx)
}

func (u PermissionControllerImpl) Update(ctx *gin.Context) {
	u.PermissionService.UpdatePermission(ctx)
}

func (u PermissionControllerImpl) Delete(ctx *gin.Context) {
	u.PermissionService.DeletePermission(ctx)
}

var permissionControllerSet = wire.NewSet(
	wire.Struct(new(PermissionControllerImpl), "*"),
	wire.Bind(new(PermissionController), new(*PermissionControllerImpl)),
)
