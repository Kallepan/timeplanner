package mock

import (
	"planner-backend/app/domain/dao"
)

type WeekdayRepositoryMock struct {
	dataContainer      map[string]interface{}
	errorContainer     map[string]error
	primedFunctionName string
}

/* Mock interface implementations */
func (r *WeekdayRepositoryMock) On(functionName string) Mock {
	// set default value
	r.dataContainer[functionName] = nil
	r.errorContainer[functionName] = nil

	// Set primed function name
	r.primedFunctionName = functionName

	return r
}

func (r *WeekdayRepositoryMock) Return(mockData interface{}, errorData error) Mock {
	r.dataContainer[r.primedFunctionName] = mockData
	r.errorContainer[r.primedFunctionName] = errorData

	return r
}

/* Repository interface implementations */
func (r *WeekdayRepositoryMock) AddWeekdayToTimeslot(timeslot *dao.Timeslot, weekday *dao.Weekday) ([]dao.Weekday, error) {
	if r.dataContainer["AddWeekdayToTimeslot"] == nil {
		return nil, r.errorContainer["AddWeekdayToTimeslot"]
	}
	return r.dataContainer["AddWeekdayToTimeslot"].([]dao.Weekday), r.errorContainer["AddWeekdayToTimeslot"]
}

func (r *WeekdayRepositoryMock) DeleteWeekdayFromTimeslot(timeslot *dao.Timeslot, weekday *dao.Weekday) error {
	return r.errorContainer["DeleteWeekdayFromTimeslot"]
}

/**
* Function to create new WeekdayRepositoryMock
 */
func NewWeekdayRepositoryMock() *WeekdayRepositoryMock {
	return &WeekdayRepositoryMock{
		dataContainer:  make(map[string]interface{}),
		errorContainer: make(map[string]error),
	}
}
