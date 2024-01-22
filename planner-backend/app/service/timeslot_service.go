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

type TimeslotService interface {
	GetAllTimeslots(c *gin.Context)
	GetTimeslotByName(c *gin.Context)
	AddTimeslot(c *gin.Context)
	UpdateTimeslot(c *gin.Context)
	DeleteTimeslot(c *gin.Context)
}

type TimeslotServiceImpl struct {
	TimeslotRepository repository.TimeslotRepository
}

func (t TimeslotServiceImpl) GetAllTimeslots(c *gin.Context) {
	/* GetAllTimeslots is a function to get all timeslots
	 * @param c is gin context
	 * @return void
	 */

	defer pkg.PanicHandler(c)
	slog.Info("start to execute program get all timeslots")

	departmentID := c.Param("departmentID")
	if departmentID == "" {
		pkg.PanicException(constant.InvalidRequest)
	}
	workplaceID := c.Param("workplaceID")
	if workplaceID == "" {
		pkg.PanicException(constant.InvalidRequest)
	}

	rawData, err := t.TimeslotRepository.FindAllTimeslots(departmentID, workplaceID)
	switch err {
	case nil:
		break
	case pkg.ErrNoRows:
		break
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	data := mapTimeslotListToTimeslotResponseList(rawData)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (t TimeslotServiceImpl) GetTimeslotByName(c *gin.Context) {
	/* GetTimeslotByName is a function to get timeslot by name
	 * @param c is gin context
	 * @return void
	 */

	defer pkg.PanicHandler(c)
	slog.Info("start to execute program get timeslot by name")

	departmentID := c.Param("departmentID")
	workplaceID := c.Param("workplaceID")
	timeslotName := c.Param("timeslotName")
	if departmentID == "" || workplaceID == "" || timeslotName == "" {
		pkg.PanicException(constant.InvalidRequest)
	}

	rawData, err := t.TimeslotRepository.FindTimeslotByName(departmentID, workplaceID, timeslotName)
	switch err {
	case nil:
		break
	case pkg.ErrNoRows:
		pkg.PanicException(constant.DataNotFound)
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	data := mapTimeslotToTimeslotResponse(rawData)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (t TimeslotServiceImpl) AddTimeslot(c *gin.Context) {
	/* AddTimeslot is a function to add a timeslot
	 * @param c is gin context
	 * @return void
	 */

	defer pkg.PanicHandler(c)
	slog.Info("start to execute program add timeslot")

	timeslotRequest := dco.TimeslotRequest{}
	if err := c.ShouldBindJSON(&timeslotRequest); err != nil {
		slog.Error("Error when binding json", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}
	departmentID := c.Param("departmentID")
	workplaceID := c.Param("workplaceID")
	if departmentID == "" || workplaceID == "" {
		pkg.PanicException(constant.InvalidRequest)
	}
	_, err := t.TimeslotRepository.FindTimeslotByName(departmentID, workplaceID, timeslotRequest.Name)
	switch err {
	case nil:
		pkg.PanicException(constant.Conflict)
	case pkg.ErrNoRows:
		break
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	timeslot := mapTimeslotRequestToTimeslot(timeslotRequest)
	rawData, err := t.TimeslotRepository.Save(departmentID, workplaceID, &timeslot)
	switch err {
	case nil:
		break
	case pkg.ErrNoRows:
		pkg.PanicException(constant.DataNotFound)
	default:
		slog.Error("Error when saving data to database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	data := mapTimeslotToTimeslotResponse(rawData)

	c.JSON(http.StatusCreated, pkg.BuildResponse(constant.Success, data))
}

func (t TimeslotServiceImpl) UpdateTimeslot(c *gin.Context) {
	/* UpdateTimeslot is a function to update a timeslot
	 * @param c is gin context
	 * @return void
	 */

	defer pkg.PanicHandler(c)
	slog.Info("start to execute program update timeslot")

	departmentID := c.Param("departmentID")
	workplaceID := c.Param("workplaceID")
	timeslotName := c.Param("timeslotName")
	if timeslotName == "" || departmentID == "" || workplaceID == "" {
		pkg.PanicException(constant.InvalidRequest)
	}

	timeslot, err := t.TimeslotRepository.FindTimeslotByName(departmentID, workplaceID, timeslotName)
	switch err {
	case nil:
		break
	case pkg.ErrNoRows:
		pkg.PanicException(constant.DataNotFound)
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	var timeslotRequest dco.TimeslotRequest
	if err := c.ShouldBindJSON(&timeslotRequest); err != nil {
		slog.Error("Error when binding json", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	timeslot.Name = timeslotRequest.Name
	timeslot.Active = *timeslotRequest.Active

	rawData, err := t.TimeslotRepository.Save(departmentID, workplaceID, &timeslot)
	switch err {
	case nil:
		break
	case pkg.ErrNoRows:
		pkg.PanicException(constant.DataNotFound)
	default:
		slog.Error("Error when saving data to database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	data := mapTimeslotToTimeslotResponse(rawData)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (t TimeslotServiceImpl) DeleteTimeslot(c *gin.Context) {
	/* DeleteTimeslot is a function to delete a timeslot
	 * @param c is gin context
	 * @return void
	 */

	defer pkg.PanicHandler(c)
	slog.Info("start to execute program delete timeslot")

	departmentID := c.Param("departmentID")
	workplaceID := c.Param("workplaceID")
	timeslotName := c.Param("timeslotName")
	if departmentID == "" || workplaceID == "" || timeslotName == "" {
		pkg.PanicException(constant.InvalidRequest)
	}

	timeslot, err := t.TimeslotRepository.FindTimeslotByName(departmentID, workplaceID, timeslotName)
	switch err {
	case nil:
		break
	case pkg.ErrNoRows:
		pkg.PanicException(constant.DataNotFound)
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	if err := t.TimeslotRepository.Delete(departmentID, workplaceID, &timeslot); err != nil {
		slog.Error("Error when deleting data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}

func mapTimeslotToTimeslotResponse(timeslot dao.Timeslot) dco.TimeslotResponse {
	/* mapTimeslotToTimeslotResponse is a function to map timeslot to timeslot response
	 * @param timeslot is dao.Timeslot
	 * @return dco.TimeslotResponse
	 */

	return dco.TimeslotResponse{
		Name:         timeslot.Name,
		Active:       &timeslot.Active,
		DepartmentID: timeslot.DepartmentID,
		WorkplaceID:  timeslot.WorkplaceID,
		Weekdays:     mapOnWeekdayListToWeekdayResponseList(timeslot.Weekdays),
		Base: dco.Base{
			CreatedAt: timeslot.Base.CreatedAt,
			UpdatedAt: timeslot.Base.UpdatedAt,
			DeletedAt: timeslot.Base.DeletedAt,
		},
	}
}

func mapTimeslotListToTimeslotResponseList(timeslots []dao.Timeslot) []dco.TimeslotResponse {
	/* mapTimeslotListToTimeslotResponseList is a function to map timeslot list to timeslot response list
	 * @param timeslots is []dao.Timeslot
	 * @return []dco.TimeslotResponse
	 */

	timeslotResponseList := []dco.TimeslotResponse{}
	for _, timeslot := range timeslots {
		timeslotResponseList = append(timeslotResponseList, mapTimeslotToTimeslotResponse(timeslot))
	}

	return timeslotResponseList
}

func mapTimeslotRequestToTimeslot(timeslot dco.TimeslotRequest) dao.Timeslot {
	/* mapTimeslotRequestToTimeslot is a function to map timeslot request to timeslot
	 * @param timeslot is dao.Timeslot
	 * @return dco.TimeslotRequest
	 */

	var active bool
	if timeslot.Active != nil {
		active = *timeslot.Active
	}

	return dao.Timeslot{
		Name:   timeslot.Name,
		Active: active,
	}
}

var timeslotServiceSet = wire.NewSet(
	wire.Struct(new(TimeslotServiceImpl), "*"),
	wire.Bind(new(TimeslotService), new(*TimeslotServiceImpl)),
)
