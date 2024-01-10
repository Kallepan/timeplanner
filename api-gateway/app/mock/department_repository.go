/* Mock file for department repository */
package mock

import (
	"api-gateway/app/domain/dao"

	"github.com/google/uuid"
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

func (r *DepartmentRepositoryMock) FindDepartmentById(id uuid.UUID) (dao.Department, error) {
	if r.dataContainer["FindDepartmentById"] == nil {
		return dao.Department{}, r.errorContainer["FindDepartmentById"]
	}

	return r.dataContainer["FindDepartmentById"].(dao.Department), r.errorContainer["FindDepartmentById"]
}

func (r *DepartmentRepositoryMock) Save(Department *dao.Department) (dao.Department, error) {
	if r.dataContainer["Save"] == nil {
		return dao.Department{}, r.errorContainer["Save"]
	}
	return r.dataContainer["Save"].(dao.Department), r.errorContainer["Save"]
}

func (r *DepartmentRepositoryMock) DeleteDepartmentById(id uuid.UUID) error {
	return r.errorContainer["DeleteDepartmentById"]
}

/**
 * Function to create new DepartmentRepositoryMock
 * @param void
 * @return DepartmentRepositoryMock
 */
func NewDepartmentRepositoryMock() DepartmentRepositoryMock {
	return DepartmentRepositoryMock{
		dataContainer:  make(map[string]interface{}),
		errorContainer: make(map[string]error),
	}
}
