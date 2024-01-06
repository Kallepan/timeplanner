package service

import (
	"auth-backend/app/constant"
	"auth-backend/app/domain/dao"
	"auth-backend/app/domain/dco"
	"auth-backend/app/pkg"
	"auth-backend/app/repository"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/google/wire"
	"gorm.io/gorm"
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

	rawData, err := d.DepartmentRepository.FindAllDepartments()
	if err != nil {
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	data := mapDepartmentListToDepartmentResponseList(rawData)

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

	rawData, err := d.DepartmentRepository.FindDepartmentById(departmentID)
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.DataNotFound)
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	data := mapDepartmentToDepartmentResponse(rawData)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (d DepartmentServiceImpl) AddDepartment(c *gin.Context) {
	/* AddDepartment is a function to add new department to database
	 * @param c is gin context
	 * @return void
	 */

	defer pkg.PanicHandler(c)
	slog.Info("start to execute program add department")

	var rawRequest dco.DepartmentRequest
	if err := c.ShouldBindJSON(&rawRequest); err != nil {
		slog.Error("Error happened: when mapping request from FE. Error", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	request := mapDepartmentRequestToDepartment(rawRequest)

	rawData, err := d.DepartmentRepository.Save(&request)
	if err != nil {
		slog.Error("Error when saving data to database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	data := mapDepartmentToDepartmentResponse(rawData)

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

	var rawRequest dco.DepartmentRequest
	if err := c.ShouldBindJSON(&rawRequest); err != nil {
		slog.Error("Error happened: when mapping request from FE. Error", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}
	request := mapDepartmentRequestToDepartment(rawRequest)

	oldData, err := d.DepartmentRepository.FindDepartmentById(departmentID)
	if err != nil {
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	oldData.Name = request.Name
	rawData, err := d.DepartmentRepository.Save(&oldData)
	if err != nil {
		slog.Error("Error when updating data to database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	data := mapDepartmentToDepartmentResponse(rawData)

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
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		pkg.PanicException(constant.DataNotFound)
	default:
		slog.Error("Error when deleting data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}

var departmentServiceSet = wire.NewSet(
	wire.Struct(new(DepartmentServiceImpl), "*"),
	wire.Bind(new(DepartmentService), new(*DepartmentServiceImpl)),
)

func mapDepartmentToDepartmentResponse(department dao.Department) dco.DepartmentResponse {
	/* mapDepartmentToDepartmentResponse is a function to map department to department response
	 * @param department is dao.Department
	 * @return dco.DepartmentResponse
	 */
	return dco.DepartmentResponse{
		BaseModel: dco.BaseModel{
			ID:        department.BaseModel.ID,
			CreatedAt: department.CreatedAt,
			UpdatedAt: department.UpdatedAt,
		},
		Name: department.Name,
	}
}

func mapDepartmentListToDepartmentResponseList(departments []dao.Department) []dco.DepartmentResponse {
	/* mapDepartmentListToDepartmentResponseList is a function to map department list to department response list
	 * @param departments is []dao.Department
	 * @return []dco.DepartmentResponse
	 */
	var data []dco.DepartmentResponse
	for _, v := range departments {
		data = append(data, mapDepartmentToDepartmentResponse(v))
	}
	return data
}

func mapDepartmentRequestToDepartment(req dco.DepartmentRequest) dao.Department {
	/* mapDepartmentRequestToDepartment is a function to map department request to department
	 * @param req is dco.DepartmentRequest
	 * @return dao.Department
	 */
	return dao.Department{
		Name: req.Name,
	}
}
