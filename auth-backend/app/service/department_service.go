package service

import (
	"auth-backend/app/constant"
	"auth-backend/app/domain/dao"
	"auth-backend/app/pkg"
	"auth-backend/app/repository"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


type DepartmentService interface {
	GetAllDepartments(c *gin.Context)
	GetDepartmentById(c *gin.Context)
	AddDepartment(c *gin.Context)
	UpdateDepartment(c *gin.Context)
	DeleteDepartment(c *gin.Context)
}

type DepartmentServiceImpl struct {
	departmentRepository repository.DepartmentRepository
}


func (d DepartmentServiceImpl) GetAllDepartments(c *gin.Context) {
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program get all departments")

	data, err := d.departmentRepository.FindAllDepartments()
	if err != nil {
		slog.Error("Error when fetching data from database", err)
		pkg.PanicException(constant.DatabaseError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (d DepartmentServiceImpl) GetDepartmentById(c *gin.Context) {
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program get department by id")

	id := c.Param("departmentID")
	departmentID, err := uuid.Parse(id)
	if err != nil {
		slog.Error("Error when parsing uuid. Error", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}
	
	data, err := d.departmentRepository.FindDepartmentById(departmentID)
	if err != nil {
		slog.Error("Error when fetching data from database", err)
		pkg.PanicException(constant.DatabaseError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (d DepartmentServiceImpl) AddDepartment(c *gin.Context) {
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program add department")

	var request dao.Department
	if err := c.ShouldBindJSON(&request); err != nil {
		slog.Error("Happened error when mapping request from FE. Error", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := d.departmentRepository.Save(&request)
	if err != nil {
		slog.Error("Error when saving data to database", err)
		pkg.PanicException(constant.DatabaseError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (d DepartmentServiceImpl) UpdateDepartment(c *gin.Context) {
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program update department")

	id := c.Param("departmentID")
	departmentID, err := uuid.Parse(id)
	if err != nil {
		slog.Error("Error when parsing uuid. Error", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	var request dao.Department
	if err := c.ShouldBindJSON(&request); err != nil {
		slog.Error("Happened error when mapping request from FE. Error", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := d.departmentRepository.FindDepartmentById(departmentID)
	if err != nil {
		slog.Error("Error when fetching data from database", err)
		pkg.PanicException(constant.DatabaseError)
	}

	data.Name = request.Name
	data, err = d.departmentRepository.Save(&data)
	if err != nil {
		slog.Error("Error when updating data to database", err)
		pkg.PanicException(constant.DatabaseError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))

}

func (d DepartmentServiceImpl) DeleteDepartment(c *gin.Context) {
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program delete department")

	id := c.Param("departmentID")
	departmentID, err := uuid.Parse(id)
	if err != nil {
		slog.Error("Error when parsing uuid. Error", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	err = d.departmentRepository.DeleteDepartmentById(departmentID)
	if err != nil {
		slog.Error("Error when deleting data from database", err)
		pkg.PanicException(constant.DatabaseError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}

func DepartmentServiceInit(departmentRepository repository.DepartmentRepository) *DepartmentServiceImpl {
	return &DepartmentServiceImpl{
		departmentRepository: departmentRepository,
	}
}