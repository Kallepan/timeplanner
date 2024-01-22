/** Note that in contrast to the other services, the weekday service does not have a
 * Update function. Furthermore, the weekday service always requires a request body
 * to be sent to the server. This is because the weekday service is used to add or
 * delete a weekday from a timeslot. The weekday service is not used to update a
 * weekday for a timeslot.
 */
package service

import (
	"log/slog"
	"net/http"
	"planner-backend/app/constant"
	"planner-backend/app/domain/dao"
	"planner-backend/app/domain/dco"
	"planner-backend/app/pkg"
	"planner-backend/app/repository"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type WeekdayService interface {
	BulkUpdateWeekdaysForTimeslot(c *gin.Context)

	AddWeekdayToTimeslot(c *gin.Context)
	DeleteWeekdayFromTimeslot(c *gin.Context)
}

type WeekdayServiceImpl struct {
	WeekdayRepository  repository.WeekdayRepository
	TimeslotRepository repository.TimeslotRepository
}

func (w WeekdayServiceImpl) BulkUpdateWeekdaysForTimeslot(c *gin.Context) {
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program bulk update weekdays for timeslot")

	departmentID := c.Param("departmentID")
	workplaceID := c.Param("workplaceID")
	timeslotName := c.Param("timeslotName")
	if departmentID == "" || workplaceID == "" || timeslotName == "" {
		pkg.PanicException(constant.InvalidRequest)
	}

	var weekdaysRequest dco.WeekdaysRequest
	if err := c.ShouldBindJSON(&weekdaysRequest); err != nil {
		slog.Error("Error when binding json", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	timeslot, err := w.TimeslotRepository.FindTimeslotByName(departmentID, workplaceID, timeslotName)
	switch err {
	case nil:
		break
	case pkg.ErrNoRows:
		pkg.PanicException(constant.InvalidRequest)
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	if err := w.WeekdayRepository.DeleteAllWeekdaysFromTimeslot(&timeslot); err != nil {
		slog.Error("Error when deleting data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	weekdaysToBeAdded, err := mapWeekdaysRequestToWeekdayList(weekdaysRequest)
	if err != nil {
		slog.Error("Error when mapping weekdays request to weekday list", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	weekdays, err := w.WeekdayRepository.AddWeekdaysToTimeslot(&timeslot, weekdaysToBeAdded)
	if err != nil {
		slog.Error("Error when adding weekdays to timeslot", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	data := mapOnWeekdayListToWeekdayResponseList(weekdays)

	c.JSON(http.StatusCreated, pkg.BuildResponse(constant.Success, data))
}

func (w WeekdayServiceImpl) AddWeekdayToTimeslot(c *gin.Context) {
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program add weekday to timeslot")

	departmentID := c.Param("departmentID")
	workplaceID := c.Param("workplaceID")
	timeslotName := c.Param("timeslotName")
	if departmentID == "" || workplaceID == "" || timeslotName == "" {
		pkg.PanicException(constant.InvalidRequest)
	}

	var weekdayRequest dco.WeekdayRequest
	if err := c.ShouldBindJSON(&weekdayRequest); err != nil {
		pkg.PanicException(constant.InvalidRequest)
	}
	if err := weekdayRequest.Validate(); err != nil {
		pkg.PanicException(constant.InvalidRequest)
	}

	timeslot, err := w.TimeslotRepository.FindTimeslotByName(departmentID, workplaceID, timeslotName)
	switch err {
	case nil:
		break
	case pkg.ErrNoRows:
		pkg.PanicException(constant.InvalidRequest)
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	weekday, err := mapWeekdayRequestToWeekday(weekdayRequest)
	if err != nil {
		pkg.PanicException(constant.InvalidRequest)
	}
	weekdays, err := w.WeekdayRepository.AddWeekdayToTimeslot(&timeslot, weekday)
	switch err {
	case nil:
		break
	case pkg.ErrNoRows:
		pkg.PanicException(constant.InvalidRequest)
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	data := mapOnWeekdayListToWeekdayResponseList(weekdays)

	c.JSON(http.StatusCreated, pkg.BuildResponse(constant.Success, data))
}

func (w WeekdayServiceImpl) DeleteWeekdayFromTimeslot(c *gin.Context) {
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program delete weekday from timeslot")

	departmentID := c.Param("departmentID")
	workplaceID := c.Param("workplaceID")
	timeslotName := c.Param("timeslotName")
	if departmentID == "" || workplaceID == "" || timeslotName == "" {
		pkg.PanicException(constant.InvalidRequest)
	}

	var weekdayRequest dco.WeekdayRequest
	if err := c.ShouldBindJSON(&weekdayRequest); err != nil {
		slog.Error("Error when binding json", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}
	if err := weekdayRequest.Validate(); err != nil {
		pkg.PanicException(constant.InvalidRequest)
	}

	timeslot, err := w.TimeslotRepository.FindTimeslotByName(departmentID, workplaceID, timeslotName)
	switch err {
	case nil:
		break
	case pkg.ErrNoRows:
		pkg.PanicException(constant.InvalidRequest)
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	weekday, err := mapWeekdayRequestToWeekday(weekdayRequest)
	if err != nil {
		pkg.PanicException(constant.InvalidRequest)
	}
	err = w.WeekdayRepository.DeleteWeekdayFromTimeslot(&timeslot, weekday)
	if err != nil {
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}

/*
*

	func (w WeekdayServiceImpl) UpdateWeekdayForTimeslot(c *gin.Context) {
		defer pkg.PanicHandler(c)
		slog.Info("start to execute program update weekday for timeslot")

		departmentID := c.Param("departmentID")
		workplaceID := c.Param("workplaceID")
		timeslotName := c.Param("timeslotName")
		if departmentID == "" || workplaceID == "" || timeslotName == "" {
			pkg.PanicException(constant.InvalidRequest)
		}

		var weekdayRequest dco.WeekdayRequest
		if err := c.ShouldBindJSON(&weekdayRequest); err != nil {
			slog.Error("Error when binding json", "error", err)
			pkg.PanicException(constant.InvalidRequest)
		}

		timeslot, err := w.TimeslotRepository.FindTimeslotByName(departmentID, workplaceID, timeslotName)
		switch err {
		case nil:
			break
		case pkg.ErrNoRows:
			pkg.PanicException(constant.InvalidRequest)
		default:
			slog.Error("Error when fetching data from database", "error", err)
			pkg.PanicException(constant.UnknownError)
		}

		weekday, err := mapWeekdayRequestToWeekday(weekdayRequest)
		if err != nil {
			pkg.PanicException(constant.InvalidRequest)
		}
		weekdays, err := w.WeekdayRepository.UpdateWeekdayForTimeslot(&timeslot, weekday)
		if err != nil {
			slog.Error("Error when fetching data from database", "error", err)
			pkg.PanicException(constant.UnknownError)
		}

		data := mapWeekdayListToWeekdayResponseList(weekdays)

		c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
	}

*
*/

func mapWeekdaysRequestToWeekdayList(weekdaysRequest dco.WeekdaysRequest) ([]dao.OnWeekday, error) {
	/* Maps a WeekdaysRequest to a list of Weekday */

	weekdays := []dao.OnWeekday{}
	for _, weekdayRequest := range weekdaysRequest.Weekdays {
		weekday, err := mapWeekdayRequestToWeekday(weekdayRequest)
		if err != nil {
			return nil, err
		}
		weekdays = append(weekdays, *weekday)
	}

	return weekdays, nil
}

func mapWeekdayRequestToWeekday(weekdayRequest dco.WeekdayRequest) (*dao.OnWeekday, error) {
	/* Maps a WeekdayRequest to a Weekday */

	// Convert the start and end time (string like: "08:00") to time.Time
	var unParsedStartTime, unParsedEndTime string
	if weekdayRequest.StartTime == nil {
		unParsedStartTime = "08:00"
	} else {
		unParsedStartTime = *weekdayRequest.StartTime
	}
	startTime, err := time.Parse("15:04", unParsedStartTime)
	if err != nil {
		return nil, err
	}

	if weekdayRequest.EndTime == nil {
		unParsedEndTime = "16:45"
	} else {
		unParsedEndTime = *weekdayRequest.EndTime
	}
	endTime, err := time.Parse("15:04", unParsedEndTime)
	if err != nil {
		return nil, err
	}

	return &dao.OnWeekday{
		ID:        weekdayRequest.ID,
		StartTime: startTime,
		EndTime:   endTime,
	}, nil
}

func mapOnWeekdayToWeekdayResponse(weekday dao.OnWeekday) dco.OnWeekdayResponse {
	/* mapWeekdayToWeekdayResponse is a function to map weekday to weekday response
	 * @param weekday is dao.OnWeekday
	 * @return dco.WeekdayResponse
	 */

	return dco.OnWeekdayResponse{
		ID:        weekday.ID,
		Name:      weekday.Name,
		StartTime: weekday.StartTime.Format(constant.TimeFormat),
		EndTime:   weekday.EndTime.Format(constant.TimeFormat),
	}
}

func mapOnWeekdayListToWeekdayResponseList(weekdays []dao.OnWeekday) []dco.OnWeekdayResponse {
	/* mapWeekdayListToWeekdayResponseList is a function to map weekday list to weekday response list
	 * @param weekdays is []dao.OnWeekday
	 * @return []dco.WeekdayResponse
	 */

	weekdayResponseList := []dco.OnWeekdayResponse{}
	for _, weekday := range weekdays {
		weekdayResponseList = append(weekdayResponseList, mapOnWeekdayToWeekdayResponse(weekday))
	}

	return weekdayResponseList
}

var weekDayServiceSet = wire.NewSet(
	wire.Struct(new(WeekdayServiceImpl), "*"),
	wire.Bind(new(WeekdayService), new(*WeekdayServiceImpl)),
)
