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

type workplaceService interface {
	// Function Used by the controller
	GetAllWorkplaces(c *gin.Context)
	GetWorkplaceByName(c *gin.Context)
	AddWorkplace(c *gin.Context)
	UpdateWorkplace(c *gin.Context)
	DeleteWorkplace(c *gin.Context)
}

type workplaceServiceImpl struct {
	workplaceRepository repository.WorkplaceRepository
}

func (w workplaceServiceImpl) GetAllWorkplaces(c *gin.Context) {
	/* GetAllWorkplaces is a function to get all workplaces
	 * @param c is gin context
	 * @return void
	 */

	defer pkg.PanicHandler(c)
	slog.Info("start to execute program get all workplaces")

	departmentName := c.Param("departmentName")
	if departmentName == "" {
		pkg.PanicException(constant.InvalidRequest)
	}

	rawData, err := w.workplaceRepository.FindAllWorkplaces(departmentName)
	switch err {
	case nil:
		break
	case pkg.ErrNoRows:
		break
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	data := mapWorkplaceListToWorkplaceResponseList(rawData)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (w workplaceServiceImpl) GetWorkplaceByName(c *gin.Context) {
	/* GetWorkplaceById is a function to get workplace by name
	 * @param c is gin context
	 * @return void
	 */

	defer pkg.PanicHandler(c)
	slog.Info("start to execute program get workplace by name")

	departmentName := c.Param("departmentName")
	workplaceName := c.Param("workplaceName")
	if departmentName == "" || workplaceName == "" {
		pkg.PanicException(constant.InvalidRequest)
	}

	rawData, err := w.workplaceRepository.FindWorkplaceByName(departmentName, workplaceName)
	switch err {
	case nil:
		break
	case pkg.ErrNoRows:
		pkg.PanicException(constant.DataNotFound)
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	data := mapWorkplaceToWorkplaceResponse(rawData)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (w workplaceServiceImpl) AddWorkplace(c *gin.Context) {
	/* AddWorkplace is a function to add a workplace
	 * @param c is gin context
	 * @return void
	 */

	defer pkg.PanicHandler(c)
	slog.Info("start to execute program add workplace")

	workplaceRequest := dco.WorkplaceRequest{}
	if err := c.ShouldBindJSON(&workplaceRequest); err != nil {
		slog.Error("Error when binding json", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	workplace := mapWorkplaceRequestToWorkplace(workplaceRequest)

	departmentName := c.Param("departmentName")
	if departmentName == "" {
		pkg.PanicException(constant.InvalidRequest)
	}
	rawData, err := w.workplaceRepository.Save(departmentName, &workplace)
	if err != nil {
		slog.Error("Error when saving data to database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	data := mapWorkplaceToWorkplaceResponse(rawData)

	c.JSON(http.StatusCreated, pkg.BuildResponse(constant.Success, data))
}

func (w workplaceServiceImpl) UpdateWorkplace(c *gin.Context) {
	/* UpdateWorkplace is a function to update a workplace
	 * @param c is gin context
	 * @return void
	 */

	defer pkg.PanicHandler(c)
	slog.Info("start to execute program update workplace")

	departmentName := c.Param("departmentName")
	workplaceName := c.Param("workplaceName")
	if workplaceName == "" || departmentName == "" {
		pkg.PanicException(constant.InvalidRequest)
	}

	workplace, err := w.workplaceRepository.FindWorkplaceByName(departmentName, workplaceName)
	switch err {
	case nil:
		break
	case pkg.ErrNoRows:
		pkg.PanicException(constant.DataNotFound)
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	var workplaceRequest dco.WorkplaceRequest
	if err := c.ShouldBindJSON(&workplaceRequest); err != nil {
		slog.Error("Error when binding json", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	workplace.Name = workplaceRequest.Name
	rawData, err := w.workplaceRepository.Save(departmentName, &workplace)
	if err != nil {
		slog.Error("Error when saving data to database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	data := mapWorkplaceToWorkplaceResponse(rawData)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (w workplaceServiceImpl) DeleteWorkplace(c *gin.Context) {
	/* DeleteWorkplace is a function to delete a workplace
	 * @param c is gin context
	 * @return void
	 */

	defer pkg.PanicHandler(c)
	slog.Info("start to execute program delete workplace")

	departmentName := c.Param("departmentName")
	workplaceName := c.Param("workplaceName")
	if departmentName == "" || workplaceName == "" {
		pkg.PanicException(constant.InvalidRequest)
	}

	workplace, err := w.workplaceRepository.FindWorkplaceByName(departmentName, workplaceName)
	switch err {
	case nil:
		break
	case pkg.ErrNoRows:
		pkg.PanicException(constant.DataNotFound)
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}
	if err := w.workplaceRepository.Delete(departmentName, &workplace); err != nil {
		slog.Error("Error when deleting data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}

func mapWorkplaceToWorkplaceResponse(workplace dao.Workplace) dco.WorkplaceResponse {
	/* mapWorkplaceToWorkplaceResponse is a function to map workplace to workplace response
	 * @param workplace is dao.Workplace
	 * @return dco.WorkplaceResponse
	 */

	return dco.WorkplaceResponse{
		Name: workplace.Name,
		Base: dco.Base{
			CreatedAt: workplace.Base.CreatedAt,
			UpdatedAt: workplace.Base.UpdatedAt,
		},
	}
}

func mapWorkplaceListToWorkplaceResponseList(workplaces []dao.Workplace) []dco.WorkplaceResponse {
	/* mapWorkplaceListToWorkplaceResponseList is a function to map workplace list to workplace response list
	 * @param workplaces is []dao.Workplace
	 * @return []dco.WorkplaceResponse
	 */

	workplaceResponseList := []dco.WorkplaceResponse{}
	for _, workplace := range workplaces {
		workplaceResponseList = append(workplaceResponseList, mapWorkplaceToWorkplaceResponse(workplace))
	}

	return workplaceResponseList
}

func mapWorkplaceRequestToWorkplace(workplace dco.WorkplaceRequest) dao.Workplace {
	/* mapWorkplaceRequestToWorkplace is a function to map workplace request to workplace
	 * @param workplace is dao.Workplace
	 * @return dco.WorkplaceRequest
	 */

	return dao.Workplace{
		Name: workplace.Name,
	}
}

var workplaceServiceSet = wire.NewSet(
	wire.Struct(new(workplaceServiceImpl), "*"),
	wire.Bind(new(workplaceService), new(*workplaceServiceImpl)),
)
