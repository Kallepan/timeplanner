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

type WorkdayService interface {
	/*
	 * Gets all Workdays along with the (if present) assigned user
	 * for a given date
	 */
	GetWorkdaysForDepartmentAndDate(c *gin.Context)
	GetWorkday(c *gin.Context)
	UpdateWorkday(c *gin.Context)

	AssignPersonToWorkday(c *gin.Context)
	UnassignPersonFromWorkday(c *gin.Context)
}

type WorkdayServiceImpl struct {
	WorkdayRepository repository.WorkdayRepository
}

func (w WorkdayServiceImpl) GetWorkdaysForDepartmentAndDate(c *gin.Context) {
	/*
	 * Gets all Workdays along with the (if present) assigned user
	 * for a given date
	 */
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program get all workdays")

	departmentID := c.Query("department")
	if departmentID == "" {
		pkg.PanicException(constant.InvalidRequest)
	}

	date := c.Query("date")
	if date == "" {
		pkg.PanicException(constant.InvalidRequest)
	}

	rawData, err := w.WorkdayRepository.GetWorkdaysForDepartmentAndDate(departmentID, date)
	switch err {
	case nil:
		break
	case pkg.ErrNoRows:
		break
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	data := mapWorkdayListToWorkdayResponseList(rawData)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (w WorkdayServiceImpl) GetWorkday(c *gin.Context) {
	/*
	 * Gets all Workdays along with the (if present) assigned user
	 * for a given date
	 **/
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program get all workdays")

	departmentID := c.Query("department")
	if departmentID == "" {
		pkg.PanicException(constant.InvalidRequest)
	}
	workplaceID := c.Query("workplace")
	if workplaceID == "" {
		pkg.PanicException(constant.InvalidRequest)
	}
	timeslotID := c.Query("timeslot")
	if timeslotID == "" {
		pkg.PanicException(constant.InvalidRequest)
	}
	date := c.Query("date")
	if date == "" {
		pkg.PanicException(constant.InvalidRequest)
	}

	rawData, err := w.WorkdayRepository.GetWorkday(departmentID, workplaceID, timeslotID, date)
	switch err {
	case nil:
		break
	case pkg.ErrNoRows:
		pkg.PanicException(constant.DataNotFound)
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	data := mapWorkdayToWorkdayResponse(rawData)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (w WorkdayServiceImpl) UpdateWorkday(c *gin.Context) {
	/*
	 * Updates a workday
	 */

	defer pkg.PanicHandler(c)
	slog.Info("start to execute program update workday")

	// get params from request
	departmentID := c.Query("departmentID")
	if departmentID == "" {
		pkg.PanicException(constant.InvalidRequest)
	}
	workplaceID := c.Query("workplaceID")
	if workplaceID == "" {
		pkg.PanicException(constant.InvalidRequest)
	}
	timeslotID := c.Query("timeslotID")
	if timeslotID == "" {
		pkg.PanicException(constant.InvalidRequest)
	}
	date := c.Query("date")
	if date == "" {
		pkg.PanicException(constant.InvalidRequest)
	}

	// validate params

	// get request body
	var request dco.UpdateWorkdayRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		pkg.PanicException(constant.InvalidRequest)
	}
	if err := request.Validate(); err != nil {
		pkg.PanicException(constant.InvalidRequest)
	}

	// get workday
	workday, err := w.WorkdayRepository.GetWorkday(departmentID, workplaceID, timeslotID, date)
	switch err {
	case nil:
		break
	case pkg.ErrNoRows:
		pkg.PanicException(constant.DataNotFound)
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	// map request to workday
	workday.StartTime = request.StartTime
	workday.EndTime = request.EndTime
	workday.Active = *request.Active
	workday.Comment = request.Comment

	// save workday
	if err := w.WorkdayRepository.Save(&workday); err != nil {
		slog.Error("Error when saving workday", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, mapWorkdayToWorkdayResponse(workday)))
}

func (w WorkdayServiceImpl) AssignPersonToWorkday(c *gin.Context) {
	/*
	 * Assigns a person to a workday
	 */
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program assign person to workday")

	var request dco.AssignPersonToWorkdayRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		pkg.PanicException(constant.InvalidRequest)
	}

	if err := request.Validate(); err != nil {
		pkg.PanicException(constant.InvalidRequest)
	}

	if err := w.WorkdayRepository.AssignPersonToWorkday(
		request.PersonID,
		request.DepartmentID,
		request.WorkplaceID,
		request.TimeslotID,
		request.Date,
	); err != nil {
		slog.Error("Error when assigning person to workday", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusCreated, pkg.BuildResponse(constant.Success, pkg.Null()))
}

func (w WorkdayServiceImpl) UnassignPersonFromWorkday(c *gin.Context) {
	/*
	 * Unassigns a person from a workday
	 */
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program unassign person from workday")

	var request dco.UnassignPersonFromWorkdayRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		pkg.PanicException(constant.InvalidRequest)
	}

	if err := request.Validate(); err != nil {
		pkg.PanicException(constant.InvalidRequest)
	}

	if err := w.WorkdayRepository.UnassignPersonFromWorkday(
		request.PersonID,
		request.DepartmentID,
		request.WorkplaceID,
		request.TimeslotID,
		request.Date,
	); err != nil {
		slog.Error("Error when unassigning person from workday", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}

func mapWorkdayPersonToWorkdayPersonResponse(person *dao.Person) *dco.PersonResponse {
	/*
	 * Maps a WorkdayPerson to a WorkdayPersonResponse
	 */
	if person == nil {
		return nil
	}
	p := mapPersonToPersonResponse(*person)

	return &p
}

func mapWorkdayToWorkdayResponse(workday dao.Workday) dco.WorkdayResponse {
	/*
	 * Maps a Workday to a WorkdayResponse
	 */

	return dco.WorkdayResponse{
		Department:        mapDepartmentToDepartmentResponse(workday.Department),
		Workplace:         mapWorkplaceToWorkplaceResponse(workday.Workplace),
		Timeslot:          mapTimeslotToTimeslotResponse(workday.Timeslot),
		Date:              workday.Date,
		StartTime:         workday.StartTime,
		EndTime:           workday.EndTime,
		DurationInMinutes: workday.DurationInMinutes,
		Person:            mapWorkdayPersonToWorkdayPersonResponse(workday.Person),
		Weekday:           workday.Weekday,
		Comment:           workday.Comment,
	}
}

func mapWorkdayListToWorkdayResponseList(workdays []dao.Workday) []dco.WorkdayResponse {
	/*
	 * Maps a Workday list to a WorkdayResponse list
	 */
	workdayResponseList := []dco.WorkdayResponse{}
	for _, workday := range workdays {
		workdayResponseList = append(workdayResponseList, mapWorkdayToWorkdayResponse(workday))
	}

	return workdayResponseList
}

var workDayServiceSet = wire.NewSet(
	wire.Struct(new(WorkdayServiceImpl), "*"),
	wire.Bind(new(WorkdayService), new(*WorkdayServiceImpl)),
)
