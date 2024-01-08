package service

import (
	"log/slog"
	"net/http"
	"planner-backend/app/constant"
	"planner-backend/app/domain/dao"
	"planner-backend/app/domain/dco"
	"planner-backend/app/pkg"
	"planner-backend/app/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type DepartmentService interface {
	// Function Used by the controller
	GetAllDepartments(c *gin.Context)
	GetDepartmentByName(c *gin.Context)
	AddDepartment(c *gin.Context)
	UpdateDepartment(c *gin.Context)
	DeleteDepartment(c *gin.Context)
}

type DepartmentServiceImpl struct {
	departmentRepository repository.DepartmentRepository
}

func (d DepartmentServiceImpl) GetAllDepartments(c *gin.Context) {
	/* GetAllDepartments is a function to get all departments
	 * @param c is gin context
	 * @return void
	 */

	defer pkg.PanicHandler(c)
	slog.Info("start to execute program get all departments")

	rawData, err := d.departmentRepository.FindAllDepartments()
	if err != nil {
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	data := mapDepartmentListToDepartmentResponseList(rawData)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (d DepartmentServiceImpl) GetDepartmentByName(c *gin.Context) {
	/* GetDepartmentById is a function to get department by name
	 * @param c is gin context
	 * @return void
	 */

	defer pkg.PanicHandler(c)
	slog.Info("start to execute program get department by name")

	name := c.Param("departmentName")
	rawData, err := d.departmentRepository.FindDepartmentByName(name)
	switch err {
	case nil:
		break
	case pkg.ErrNoRows:
		pkg.PanicException(constant.DataNotFound)
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	data := mapDepartmentToDepartmentResponse(rawData)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (d DepartmentServiceImpl) AddDepartment(c *gin.Context) {
	/* AddDepartment is a function to add department
	 * @param c is gin context
	 * @return void
	 */

	defer pkg.PanicHandler(c)
	slog.Info("start to execute program add department")

	var departmentRequest dco.DepartmentRequest
	if err := c.ShouldBindJSON(&departmentRequest); err != nil {
		slog.Error("Error when binding json", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	department := mapDepartmentRequestToDepartment(departmentRequest)

	_, err := d.departmentRepository.FindDepartmentByName(department.Name)
	switch err {
	case nil:
		pkg.PanicException(constant.Conflict)
	case pkg.ErrNoRows:
		break
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	rawData, err := d.departmentRepository.Save(&department)
	if err != nil {
		slog.Error("Error when saving data to database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	data := mapDepartmentToDepartmentResponse(rawData)

	c.JSON(http.StatusCreated, pkg.BuildResponse(constant.Success, data))
}

func (d DepartmentServiceImpl) UpdateDepartment(c *gin.Context) {
	/* UpdateDepartment is a function to update department by name
	 * @param c is gin context
	 * @return void
	 */
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program update department")

	name := c.Param("departmentName")
	department, err := d.departmentRepository.FindDepartmentByName(name)
	switch err {
	case nil:
		break
	case pkg.ErrNoRows:
		pkg.PanicException(constant.DataNotFound)
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	var departmentRequest dco.DepartmentRequest
	if err := c.ShouldBindJSON(&departmentRequest); err != nil {
		slog.Error("Error when binding json", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	department.Name = departmentRequest.Name
	rawData, err := d.departmentRepository.Save(&department)
	if err != nil {
		slog.Error("Error when updating data to database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	data := mapDepartmentToDepartmentResponse(rawData)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (d DepartmentServiceImpl) DeleteDepartment(c *gin.Context) {
	/* DeleteDepartment is a function to delete department by name
	 * @param c is gin context
	 * @return void
	 */

	defer pkg.PanicHandler(c)
	slog.Info("start to execute program delete department")

	name := c.Param("departmentName")
	department, err := d.departmentRepository.FindDepartmentByName(name)
	switch err {
	case nil:
		break
	case pkg.ErrNoRows:
		pkg.PanicException(constant.DataNotFound)
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	err = d.departmentRepository.Delete(&department)
	switch err {
	case nil:
		break
	case pkg.ErrNoRows:
		break
	default:
		slog.Error("Error when updating data to database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}

func mapDepartmentToDepartmentResponse(department dao.Department) dco.DepartmentResponse {
	/* mapDepartmentToDepartmentResponse is a function to map department to department response
	 * @param department is a department
	 * @return dco.DepartmentResponse
	 */
	return dco.DepartmentResponse{
		Base: dco.Base{
			CreatedAt: department.CreatedAt,
			UpdatedAt: department.UpdatedAt,
		},
		Name: department.Name,
	}

}
func mapDepartmentListToDepartmentResponseList(departments []dao.Department) []dco.DepartmentResponse {
	/* mapDepartmentListToDepartmentResponseList is a function to map department list to department response list
	 * @param departments is a list of department
	 * @return []dco.DepartmentResponse
	 */

	var departmentResponseList []dco.DepartmentResponse
	for _, department := range departments {
		departmentResponseList = append(departmentResponseList, mapDepartmentToDepartmentResponse(department))
	}
	return departmentResponseList
}

func mapDepartmentRequestToDepartment(departmentRequest dco.DepartmentRequest) dao.Department {
	/* mapDepartmentRequestToDepartment is a function to map department request to department
	 * @param departmentRequest is a department request
	 * @return dao.Department
	 */
	return dao.Department{
		Name: departmentRequest.Name,
	}
}

var departmentServiceSet = wire.NewSet(
	wire.Struct(new(DepartmentServiceImpl), "*"),
	wire.Bind(new(DepartmentService), new(*DepartmentServiceImpl)),
)
