package controller

import (
	"auth-backend/app/service"

	"github.com/gin-gonic/gin"
)

type PermissionController interface {
	GetAll(ctx *gin.Context)
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type PermissionControllerImpl struct {
	svc service.PermissionService
}

func (u PermissionControllerImpl) GetAll(ctx *gin.Context) {
	u.svc.GetAllPermissions(ctx)
}

func (u PermissionControllerImpl) Get(ctx *gin.Context) {
	u.svc.GetPermissionById(ctx)
}

func (u PermissionControllerImpl) Create(ctx *gin.Context) {
	u.svc.AddPermission(ctx)
}

func (u PermissionControllerImpl) Update(ctx *gin.Context) {
	u.svc.UpdatePermission(ctx)
}

func (u PermissionControllerImpl) Delete(ctx *gin.Context) {
	u.svc.DeletePermission(ctx)
}

func PermissionControllerInit(permissionService service.PermissionService) *PermissionControllerImpl {
	return &PermissionControllerImpl{
		svc: permissionService,
	}
}
