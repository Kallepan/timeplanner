package controller

import (
	"auth-backend/app/service"

	"github.com/gin-gonic/gin"
)

type DepartmentController interface {
	GetAll(ctx *gin.Context)
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type DepartmentControllerImpl struct {
	svc service.DepartmentService
}

func (u DepartmentControllerImpl) GetAll(ctx *gin.Context) {
	u.svc.GetAllDepartments(ctx)
}

func (u DepartmentControllerImpl) Get(ctx *gin.Context) {
	u.svc.GetDepartmentById(ctx)
}

func (u DepartmentControllerImpl) Create(ctx *gin.Context) {
	u.svc.AddDepartment(ctx)
}

func (u DepartmentControllerImpl) Update(ctx *gin.Context) {
	u.svc.UpdateDepartment(ctx)
}

func (u DepartmentControllerImpl) Delete(ctx *gin.Context) {
	u.svc.DeleteDepartment(ctx)
}

func DepartmentControllerInit(departmentService service.DepartmentService) *DepartmentControllerImpl {
	return &DepartmentControllerImpl{
		svc: departmentService,
	}
}