package controller

import (
	"api-gateway/app/service"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type DepartmentController interface {
	GetAll(ctx *gin.Context)
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type DepartmentControllerImpl struct {
	DepartmentService service.DepartmentService
}

func (u DepartmentControllerImpl) GetAll(ctx *gin.Context) {
	u.DepartmentService.GetAllDepartments(ctx)
}

func (u DepartmentControllerImpl) Get(ctx *gin.Context) {
	u.DepartmentService.GetDepartmentById(ctx)
}

func (u DepartmentControllerImpl) Create(ctx *gin.Context) {
	u.DepartmentService.AddDepartment(ctx)
}

func (u DepartmentControllerImpl) Update(ctx *gin.Context) {
	u.DepartmentService.UpdateDepartment(ctx)
}

func (u DepartmentControllerImpl) Delete(ctx *gin.Context) {
	u.DepartmentService.DeleteDepartment(ctx)
}

var departmentControllerSet = wire.NewSet(
	wire.Struct(new(DepartmentControllerImpl), "*"),
	wire.Bind(new(DepartmentController), new(*DepartmentControllerImpl)),
)
