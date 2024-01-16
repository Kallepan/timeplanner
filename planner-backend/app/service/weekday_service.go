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
	AddWeekdayToTimeslot(c *gin.Context)
	DeleteWeekdayFromTimeslot(c *gin.Context)
}

type WeekdayServiceImpl struct {
	WeekdayRepository  repository.WeekdayRepository
	TimeslotRepository repository.TimeslotRepository
}

func (w WeekdayServiceImpl) AddWeekdayToTimeslot(c *gin.Context) {
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program add weekday to timeslot")

	departmentName := c.Param("departmentName")
	workplaceName := c.Param("workplaceName")
	timeslotName := c.Param("timeslotName")
	if departmentName == "" || workplaceName == "" || timeslotName == "" {
		pkg.PanicException(constant.InvalidRequest)
	}

	var weekdayRequest dco.WeekdayRequest
	if err := c.ShouldBindJSON(&weekdayRequest); err != nil {
		pkg.PanicException(constant.InvalidRequest)
	}
	if err := weekdayRequest.Validate(); err != nil {
		pkg.PanicException(constant.InvalidRequest)
	}

	timeslot, err := w.TimeslotRepository.FindTimeslotByName(departmentName, workplaceName, timeslotName)
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
	if err != nil {
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	data := mapWeekdayListToWeekdayResponseList(weekdays)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (w WeekdayServiceImpl) DeleteWeekdayFromTimeslot(c *gin.Context) {
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program delete weekday from timeslot")

	departmentName := c.Param("departmentName")
	workplaceName := c.Param("workplaceName")
	timeslotName := c.Param("timeslotName")
	if departmentName == "" || workplaceName == "" || timeslotName == "" {
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

	timeslot, err := w.TimeslotRepository.FindTimeslotByName(departmentName, workplaceName, timeslotName)
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

		departmentName := c.Param("departmentName")
		workplaceName := c.Param("workplaceName")
		timeslotName := c.Param("timeslotName")
		if departmentName == "" || workplaceName == "" || timeslotName == "" {
			pkg.PanicException(constant.InvalidRequest)
		}

		var weekdayRequest dco.WeekdayRequest
		if err := c.ShouldBindJSON(&weekdayRequest); err != nil {
			slog.Error("Error when binding json", "error", err)
			pkg.PanicException(constant.InvalidRequest)
		}

		timeslot, err := w.TimeslotRepository.FindTimeslotByName(departmentName, workplaceName, timeslotName)
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

func mapWeekdayToWeekdayResponse(weekday dao.OnWeekday) dco.WeekdayResponse {
	/* mapWeekdayToWeekdayResponse is a function to map weekday to weekday response
	 * @param weekday is dao.OnWeekday
	 * @return dco.WeekdayResponse
	 */

	return dco.WeekdayResponse{
		ID:        weekday.ID,
		Name:      weekday.Name,
		StartTime: weekday.StartTime.Format(constant.TimeFormat),
		EndTime:   weekday.EndTime.Format(constant.TimeFormat),
	}
}

func mapWeekdayListToWeekdayResponseList(weekdays []dao.OnWeekday) []dco.WeekdayResponse {
	/* mapWeekdayListToWeekdayResponseList is a function to map weekday list to weekday response list
	 * @param weekdays is []dao.OnWeekday
	 * @return []dco.WeekdayResponse
	 */

	weekdayResponseList := []dco.WeekdayResponse{}
	for _, weekday := range weekdays {
		weekdayResponseList = append(weekdayResponseList, mapWeekdayToWeekdayResponse(weekday))
	}

	return weekdayResponseList
}

var weekDayServiceSet = wire.NewSet(
	wire.Struct(new(WeekdayServiceImpl), "*"),
	wire.Bind(new(WeekdayService), new(*WeekdayServiceImpl)),
)
