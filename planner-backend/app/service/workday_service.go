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
	timeslotName := c.Query("timeslot")
	if timeslotName == "" {
		pkg.PanicException(constant.InvalidRequest)
	}
	date := c.Query("date")
	if date == "" {
		pkg.PanicException(constant.InvalidRequest)
	}

	rawData, err := w.WorkdayRepository.GetWorkday(departmentID, workplaceID, timeslotName, date)
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

func (w WorkdayServiceImpl) AssignPersonToWorkday(c *gin.Context) {
	/*
	 * Assigns a person to a workday
	 */
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program assign person to workday")

	var request dco.AssignPersonToWorkdayRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		slog.Error("Error when binding request", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	if err := request.Validate(); err != nil {
		slog.Error("Error when validating request", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	if err := w.WorkdayRepository.AssignPersonToWorkday(
		request.PersonID,
		request.DepartmentID,
		request.WorkplaceID,
		request.TimeslotName,
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
		slog.Error("Error when binding request", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	if err := request.Validate(); err != nil {
		slog.Error("Error when validating request", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	if err := w.WorkdayRepository.UnassignPersonFromWorkday(
		request.PersonID,
		request.DepartmentID,
		request.WorkplaceID,
		request.TimeslotName,
		request.Date,
	); err != nil {
		slog.Error("Error when unassigning person from workday", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}

func mapWorkdayPersonToWorkdayPersonResponse(workdayPerson *dao.Person) *dco.WorkdayPersonResponse {
	/*
	 * Maps a WorkdayPerson to a WorkdayPersonResponse
	 */
	if workdayPerson == nil {
		return nil
	}
	return &dco.WorkdayPersonResponse{
		ID:           workdayPerson.ID,
		FirstName:    workdayPerson.FirstName,
		LastName:     workdayPerson.LastName,
		Email:        workdayPerson.Email,
		WorkingHours: workdayPerson.WorkingHours,
	}
}

func mapWorkdayToWorkdayResponse(workday dao.Workday) dco.WorkdayResponse {
	/*
	 * Maps a Workday to a WorkdayResponse
	 */

	return dco.WorkdayResponse{
		Department: mapDepartmentToDepartmentResponse(workday.Department),
		Workplace:  mapWorkplaceToWorkplaceResponse(workday.Workplace),
		Timeslot:   mapTimeslotToTimeslotResponse(workday.Timeslot),
		Date:       workday.Date,
		StartTime:  workday.StartTime,
		EndTime:    workday.EndTime,
		Person:     mapWorkdayPersonToWorkdayPersonResponse(workday.Person),
		Weekday:    workday.Weekday,
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
