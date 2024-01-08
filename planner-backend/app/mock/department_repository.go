package mock

import (
	"planner-backend/app/domain/dao"
)

type DepartmentRepositoryMock struct {
	dataContainer      map[string]interface{}
	errorContainer     map[string]error
	primedFunctionName string
}

/* Mock interface implementations */
func (r *DepartmentRepositoryMock) On(functionName string) Mock {
	// set default value
	r.dataContainer[functionName] = nil
	r.errorContainer[functionName] = nil

	// Set primed function name
	r.primedFunctionName = functionName

	return r
}

func (r *DepartmentRepositoryMock) Return(mockData interface{}, errorData error) Mock {
	r.dataContainer[r.primedFunctionName] = mockData
	r.errorContainer[r.primedFunctionName] = errorData

	return r
}

/* Repostory interface implementations */
func (r *DepartmentRepositoryMock) FindAllDepartments() ([]dao.Department, error) {
	if r.dataContainer["FindAllDepartments"] == nil {
		return nil, r.errorContainer["FindAllDepartments"]
	}
	return r.dataContainer["FindAllDepartments"].([]dao.Department), r.errorContainer["FindAllDepartments"]
}

func (r *DepartmentRepositoryMock) FindDepartmentByName(name string) (dao.Department, error) {
	if r.dataContainer["FindDepartmentByName"] == nil {
		return dao.Department{}, r.errorContainer["FindDepartmentByName"]
	}

	return r.dataContainer["FindDepartmentByName"].(dao.Department), r.errorContainer["FindDepartmentByName"]
}

func (r *DepartmentRepositoryMock) Save(Department *dao.Department) (dao.Department, error) {
	if r.dataContainer["Save"] == nil {
		return dao.Department{}, r.errorContainer["Save"]
	}
	return r.dataContainer["Save"].(dao.Department), r.errorContainer["Save"]
}

func (r *DepartmentRepositoryMock) Delete(Department *dao.Department) error {
	return r.errorContainer["Delete"]
}

/**
* Function to create new DepartmentRepositoryMock
 */
func NewDepartmentRepositoryMock() *DepartmentRepositoryMock {
	return &DepartmentRepositoryMock{
		dataContainer:  make(map[string]interface{}),
		errorContainer: make(map[string]error),
	}
}
