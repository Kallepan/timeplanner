package service

import (
	"auth-backend/app/constant"
	"auth-backend/app/domain/dao"
	"auth-backend/app/pkg"
	"auth-backend/app/repository"
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/google/wire"
)

type DepartmentService interface {
	GetAllDepartments(c *gin.Context)
	GetDepartmentById(c *gin.Context)
	AddDepartment(c *gin.Context)
	UpdateDepartment(c *gin.Context)
	DeleteDepartment(c *gin.Context)
}

type DepartmentServiceImpl struct {
	DepartmentRepository repository.DepartmentRepository
}

func (d DepartmentServiceImpl) GetAllDepartments(c *gin.Context) {
	/* GetAllDepartments is a function to get all departments
	 * @param c is gin context
	 * @return void
	 */
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program get all departments")

	data, err := d.DepartmentRepository.FindAllDepartments()
	if err != nil {
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (d DepartmentServiceImpl) GetDepartmentById(c *gin.Context) {
	/* GetDepartmentById is a function to get department by id
	 * @param c is gin context
	 * @return void
	 */
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program get department by id")

	id := c.Param("departmentID")
	departmentID, err := uuid.Parse(id)
	if err != nil {
		slog.Error("Error when parsing uuid. Error", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := d.DepartmentRepository.FindDepartmentById(departmentID)
	switch err {
		case nil:
			break
		case sql.ErrNoRows:
			slog.Error("Error when fetching data from database", "error", err)
			pkg.PanicException(constant.DataNotFound)
		default:
			slog.Error("Error when fetching data from database", "error", err)
			pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (d DepartmentServiceImpl) AddDepartment(c *gin.Context) {
	/* AddDepartment is a function to add new department to database
	 * @param c is gin context
	 * @return void
	 */

	defer pkg.PanicHandler(c)
	slog.Info("start to execute program add department")

	var request dao.Department
	if err := c.ShouldBindJSON(&request); err != nil {
		slog.Error("Error happened: when mapping request from FE. Error", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := d.DepartmentRepository.Save(&request)
	if err != nil {
		slog.Error("Error when saving data to database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusCreated, pkg.BuildResponse(constant.Success, data))
}

func (d DepartmentServiceImpl) UpdateDepartment(c *gin.Context) {
	/* UpdateDepartment is a function to update department by id
	 * @param c is gin context
	 * @return void
	 */
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
		slog.Error("Error happened: when mapping request from FE. Error", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := d.DepartmentRepository.FindDepartmentById(departmentID)
	if err != nil {
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	data.Name = request.Name
	data, err = d.DepartmentRepository.Save(&data)
	if err != nil {
		slog.Error("Error when updating data to database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))

}

func (d DepartmentServiceImpl) DeleteDepartment(c *gin.Context) {
	/* DeleteDepartment is a function to delete department by id
	 * @param c is gin context
	 * @return void
	 */
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program delete department")

	id := c.Param("departmentID")
	departmentID, err := uuid.Parse(id)
	if err != nil {
		slog.Error("Error when parsing uuid. Error", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	err = d.DepartmentRepository.DeleteDepartmentById(departmentID)
	if err != nil {
		slog.Error("Error when deleting data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}

var departmentServiceSet = wire.NewSet(
	wire.Struct(new(DepartmentServiceImpl), "*"),
	wire.Bind(new(DepartmentService), new(*DepartmentServiceImpl)),
)
