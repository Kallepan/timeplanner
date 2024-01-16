package mock

import "planner-backend/app/domain/dao"

type WorkdayRepositoryMock struct {
	dataContainer      map[string]interface{}
	errorContainer     map[string]error
	primedFunctionName string
}

/* Mock interface implementations */
func (r *WorkdayRepositoryMock) On(functionName string) Mock {
	// set default value
	r.dataContainer[functionName] = nil
	r.errorContainer[functionName] = nil

	// Set primed function name
	r.primedFunctionName = functionName

	return r
}

func (r *WorkdayRepositoryMock) Return(mockData interface{}, errorData error) Mock {
	r.dataContainer[r.primedFunctionName] = mockData
	r.errorContainer[r.primedFunctionName] = errorData

	return r
}

/* Repostory interface implementations */
func (r *WorkdayRepositoryMock) GetWorkdaysForDepartmentAndDate(departmentName string, date string) ([]dao.Workday, error) {
	if r.dataContainer["GetWorkdaysForDepartmentAndDate"] == nil {
		return nil, r.errorContainer["GetWorkdaysForDepartmentAndDate"]
	}
	return r.dataContainer["GetWorkdaysForDepartmentAndDate"].([]dao.Workday), r.errorContainer["GetWorkdaysForDepartmentAndDate"]
}

func (r *WorkdayRepositoryMock) GetWorkday(departmentName string, workplaceName string, timeslotName string, date string) (dao.Workday, error) {
	if r.dataContainer["GetWorkday"] == nil {
		return dao.Workday{}, r.errorContainer["GetWorkday"]
	}
	return r.dataContainer["GetWorkday"].(dao.Workday), r.errorContainer["GetWorkday"]
}

func (r *WorkdayRepositoryMock) AssignPersonToWorkday(personID string, departmentName string, workplaceName string, timeslotName string, date string) error {
	return r.errorContainer["AssignPersonToWorkday"]
}

func (r *WorkdayRepositoryMock) UnassignPersonFromWorkday(personID string, departmentName string, workplaceName string, timeslotName string, date string) error {
	return r.errorContainer["UnassignPersonFromWorkday"]
}

/**
* Function to create new WorkdayRepositoryMock
**/
func NewWorkdayRepositoryMock() *WorkdayRepositoryMock {
	return &WorkdayRepositoryMock{
		dataContainer:  make(map[string]interface{}),
		errorContainer: make(map[string]error),
	}
}