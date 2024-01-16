package mock

import (
	"planner-backend/app/domain/dao"
)

type TimeslotRepositoryMock struct {
	dataContainer      map[string]interface{}
	errorContainer     map[string]error
	primedFunctionName string
}

/* Mock interface implementations */
func (r *TimeslotRepositoryMock) On(functionName string) Mock {
	// set default value
	r.dataContainer[functionName] = nil
	r.errorContainer[functionName] = nil

	// Set primed function name
	r.primedFunctionName = functionName

	return r
}

func (r *TimeslotRepositoryMock) Return(mockData interface{}, errorData error) Mock {
	r.dataContainer[r.primedFunctionName] = mockData
	r.errorContainer[r.primedFunctionName] = errorData

	return r
}

/* Repository interface implementations */
func (r *TimeslotRepositoryMock) FindAllTimeslots(departmentName string, workplaceName string) ([]dao.Timeslot, error) {
	if r.dataContainer["FindAllTimeslots"] == nil {
		return nil, r.errorContainer["FindAllTimeslots"]
	}
	return r.dataContainer["FindAllTimeslots"].([]dao.Timeslot), r.errorContainer["FindAllTimeslots"]
}

func (r *TimeslotRepositoryMock) FindTimeslotByName(departmentName string, workplaceName string, timeslotName string) (dao.Timeslot, error) {
	if r.dataContainer["FindTimeslotByName"] == nil {
		return dao.Timeslot{}, r.errorContainer["FindTimeslotByName"]
	}

	return r.dataContainer["FindTimeslotByName"].(dao.Timeslot), r.errorContainer["FindTimeslotByName"]
}

func (r *TimeslotRepositoryMock) Save(departmentName string, workplaceName string, timeslot *dao.Timeslot) (dao.Timeslot, error) {
	if r.dataContainer["Save"] == nil {
		return dao.Timeslot{}, r.errorContainer["Save"]
	}
	return r.dataContainer["Save"].(dao.Timeslot), r.errorContainer["Save"]
}

func (r *TimeslotRepositoryMock) Delete(departmentName string, workplaceName string, timeslot *dao.Timeslot) error {
	return r.errorContainer["Delete"]
}

/**
* Function to create new TimeslotRepositoryMock
 */
func NewTimeslotRepositoryMock() *TimeslotRepositoryMock {
	return &TimeslotRepositoryMock{
		dataContainer:  make(map[string]interface{}),
		errorContainer: make(map[string]error),
	}
}
