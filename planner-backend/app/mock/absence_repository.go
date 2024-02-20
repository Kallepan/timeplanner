package mock

import (
	"planner-backend/app/domain/dao"
)

type AbsenceRepositoryMock struct {
	dataContainer      map[string]interface{}
	errorContainer     map[string]error
	primedFunctionName string
}

/* Mock interface implementations */
func (r *AbsenceRepositoryMock) On(functionName string) Mock {
	// set default value
	r.dataContainer[functionName] = nil
	r.errorContainer[functionName] = nil

	// Set primed function name
	r.primedFunctionName = functionName

	return r
}

func (r *AbsenceRepositoryMock) Return(mockData interface{}, errorData error) Mock {
	r.dataContainer[r.primedFunctionName] = mockData
	r.errorContainer[r.primedFunctionName] = errorData

	return r
}

/* Repostory interface implementations */
func (r *AbsenceRepositoryMock) FindAllAbsencies(departmentID string, date string) ([]dao.Absence, error) {
	if r.dataContainer["FindAllAbsencies"] == nil {
		return nil, r.errorContainer["FindAllAbsencies"]
	}
	return r.dataContainer["FindAllAbsencies"].([]dao.Absence), r.errorContainer["FindAllAbsencies"]
}

/**
* Function to create new AbsenceRepositoryMock
* @return AbsenceRepositoryMock
 */

func NewAbsenceRepositoryMock() *AbsenceRepositoryMock {
	return &AbsenceRepositoryMock{
		dataContainer:  make(map[string]interface{}),
		errorContainer: make(map[string]error),
	}
}
