package mock

import (
	"planner-backend/app/domain/dao"
)

type WorkplaceRepositoryMock struct {
	dataContainer      map[string]interface{}
	errorContainer     map[string]error
	primedFunctionName string
}

/* Mock interface implementations */
func (r *WorkplaceRepositoryMock) On(functionName string) Mock {
	// set default value
	r.dataContainer[functionName] = nil
	r.errorContainer[functionName] = nil

	// Set primed function name
	r.primedFunctionName = functionName

	return r
}

func (r *WorkplaceRepositoryMock) Return(mockData interface{}, errorData error) Mock {
	r.dataContainer[r.primedFunctionName] = mockData
	r.errorContainer[r.primedFunctionName] = errorData

	return r
}

/* Repostory interface implementations */
func (r *WorkplaceRepositoryMock) FindAllWorkplaces(departmentID string) ([]dao.Workplace, error) {
	if r.dataContainer["FindAllWorkplaces"] == nil {
		return nil, r.errorContainer["FindAllWorkplaces"]
	}
	return r.dataContainer["FindAllWorkplaces"].([]dao.Workplace), r.errorContainer["FindAllWorkplaces"]
}

func (r *WorkplaceRepositoryMock) FindWorkplaceByName(departmentID string, workplaceID string) (dao.Workplace, error) {
	if r.dataContainer["FindWorkplaceByName"] == nil {
		return dao.Workplace{}, r.errorContainer["FindWorkplaceByName"]
	}

	return r.dataContainer["FindWorkplaceByName"].(dao.Workplace), r.errorContainer["FindWorkplaceByName"]
}

func (r *WorkplaceRepositoryMock) FindWorkplaceByID(departmentID string, workplaceID string) (dao.Workplace, error) {
	if r.dataContainer["FindWorkplaceByID"] == nil {
		return dao.Workplace{}, r.errorContainer["FindWorkplaceByID"]
	}

	return r.dataContainer["FindWorkplaceByID"].(dao.Workplace), r.errorContainer["FindWorkplaceByID"]
}

func (r *WorkplaceRepositoryMock) Save(departmentID string, Workplace *dao.Workplace) (dao.Workplace, error) {
	if r.dataContainer["Save"] == nil {
		return dao.Workplace{}, r.errorContainer["Save"]
	}
	return r.dataContainer["Save"].(dao.Workplace), r.errorContainer["Save"]
}

func (r *WorkplaceRepositoryMock) Delete(departmentID string, Workplace *dao.Workplace) error {
	return r.errorContainer["Delete"]
}

/**
* Function to create new WorkplaceRepositoryMock
 */
func NewWorkplaceRepositoryMock() *WorkplaceRepositoryMock {
	return &WorkplaceRepositoryMock{
		dataContainer:  make(map[string]interface{}),
		errorContainer: make(map[string]error),
	}
}
