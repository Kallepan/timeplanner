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

type WorkplaceService interface {
	// Function Used by the controller
	GetAllWorkplaces(c *gin.Context)
	GetWorkplaceByName(c *gin.Context)
	AddWorkplace(c *gin.Context)
	UpdateWorkplace(c *gin.Context)
	DeleteWorkplace(c *gin.Context)
}

type WorkplaceServiceImpl struct {
	WorkplaceRepository repository.WorkplaceRepository
}

func (w WorkplaceServiceImpl) GetAllWorkplaces(c *gin.Context) {
	/* GetAllWorkplaces is a function to get all workplaces
	 * @param c is gin context
	 * @return void
	 */

	defer pkg.PanicHandler(c)
	slog.Info("start to execute program get all workplaces")

	departmentID := c.Param("departmentID")
	if departmentID == "" {
		pkg.PanicException(constant.InvalidRequest)
	}

	rawData, err := w.WorkplaceRepository.FindAllWorkplaces(departmentID)
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

func (w WorkplaceServiceImpl) GetWorkplaceByName(c *gin.Context) {
	/* GetWorkplaceById is a function to get workplace by name
	 * @param c is gin context
	 * @return void
	 */

	defer pkg.PanicHandler(c)
	slog.Info("start to execute program get workplace by name")

	departmentID := c.Param("departmentID")
	workplaceID := c.Param("workplaceID")
	if departmentID == "" || workplaceID == "" {
		pkg.PanicException(constant.InvalidRequest)
	}

	rawData, err := w.WorkplaceRepository.FindWorkplaceByID(departmentID, workplaceID)
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

func (w WorkplaceServiceImpl) AddWorkplace(c *gin.Context) {
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
	departmentID := c.Param("departmentID")
	if departmentID == "" {
		pkg.PanicException(constant.InvalidRequest)
	}
	_, err := w.WorkplaceRepository.FindWorkplaceByID(departmentID, workplaceRequest.Name)
	switch err {
	case nil:
		pkg.PanicException(constant.Conflict)
	case pkg.ErrNoRows:
		break
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	workplace := mapWorkplaceRequestToWorkplace(workplaceRequest)
	rawData, err := w.WorkplaceRepository.Save(departmentID, &workplace)
	switch err {
	case nil:
		break
	case pkg.ErrNoRows:
		pkg.PanicException(constant.DataNotFound)
	default:
		slog.Error("Error when saving data to database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	data := mapWorkplaceToWorkplaceResponse(rawData)

	c.JSON(http.StatusCreated, pkg.BuildResponse(constant.Success, data))
}

func (w WorkplaceServiceImpl) UpdateWorkplace(c *gin.Context) {
	/* UpdateWorkplace is a function to update a workplace
	 * @param c is gin context
	 * @return void
	 */

	defer pkg.PanicHandler(c)
	slog.Info("start to execute program update workplace")

	departmentID := c.Param("departmentID")
	workplaceID := c.Param("workplaceID")
	if workplaceID == "" || departmentID == "" {
		pkg.PanicException(constant.InvalidRequest)
	}

	workplace, err := w.WorkplaceRepository.FindWorkplaceByID(departmentID, workplaceID)
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
	rawData, err := w.WorkplaceRepository.Save(departmentID, &workplace)
	switch err {
	case nil:
		break
	case pkg.ErrNoRows:
		pkg.PanicException(constant.DataNotFound)
	default:
		slog.Error("Error when saving data to database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	data := mapWorkplaceToWorkplaceResponse(rawData)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (w WorkplaceServiceImpl) DeleteWorkplace(c *gin.Context) {
	/* DeleteWorkplace is a function to delete a workplace
	 * @param c is gin context
	 * @return void
	 */

	defer pkg.PanicHandler(c)
	slog.Info("start to execute program delete workplace")

	departmentID := c.Param("departmentID")
	workplaceID := c.Param("workplaceID")
	if departmentID == "" || workplaceID == "" {
		pkg.PanicException(constant.InvalidRequest)
	}

	workplace, err := w.WorkplaceRepository.FindWorkplaceByID(departmentID, workplaceID)
	switch err {
	case nil:
		break
	case pkg.ErrNoRows:
		pkg.PanicException(constant.DataNotFound)
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}
	if err := w.WorkplaceRepository.Delete(departmentID, &workplace); err != nil {
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
		ID:   workplace.ID,
		Base: dco.Base{
			CreatedAt: workplace.Base.CreatedAt,
			UpdatedAt: workplace.Base.UpdatedAt,
			DeletedAt: workplace.Base.DeletedAt,
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
		ID:   workplace.ID,
		Name: workplace.Name,
	}
}

var workplaceServiceSet = wire.NewSet(
	wire.Struct(new(WorkplaceServiceImpl), "*"),
	wire.Bind(new(WorkplaceService), new(*WorkplaceServiceImpl)),
)
