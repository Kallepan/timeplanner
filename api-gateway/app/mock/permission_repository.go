/* Mock file for department repository */
package mock

import (
	"api-gateway/app/domain/dao"

	"github.com/google/uuid"
)

type PermissionRepositoryMock struct {
	dataContainer      map[string]interface{}
	errorContainer     map[string]error
	primedFunctionName string
}

/* Mock interface implementations */
func (r *PermissionRepositoryMock) On(functionName string) Mock {
	// set default value
	r.dataContainer[functionName] = nil
	r.errorContainer[functionName] = nil

	// Set primed function name
	r.primedFunctionName = functionName

	return r
}

func (r *PermissionRepositoryMock) Return(mockData interface{}, errorData error) Mock {
	r.dataContainer[r.primedFunctionName] = mockData
	r.errorContainer[r.primedFunctionName] = errorData

	return r
}

/* Repostory interface implementations */
func (r *PermissionRepositoryMock) FindAllPermissions() ([]dao.Permission, error) {
	if r.dataContainer["FindAllPermissions"] == nil {
		return nil, r.errorContainer["FindAllPermissions"]
	}

	return r.dataContainer["FindAllPermissions"].([]dao.Permission), r.errorContainer["FindAllPermissions"]
}

func (r *PermissionRepositoryMock) FindPermissionByName(name string) (dao.Permission, error) {
	if r.dataContainer["FindPermissionByName"] == nil {
		return dao.Permission{}, r.errorContainer["FindPermissionByName"]
	}

	return r.dataContainer["FindPermissionByName"].(dao.Permission), r.errorContainer["FindPermissionByName"]
}

func (r *PermissionRepositoryMock) FindPermissionById(id uuid.UUID) (dao.Permission, error) {
	if r.dataContainer["FindPermissionById"] == nil {
		return dao.Permission{}, r.errorContainer["FindPermissionById"]
	}

	return r.dataContainer["FindPermissionById"].(dao.Permission), r.errorContainer["FindPermissionById"]
}

func (r *PermissionRepositoryMock) Save(Permission *dao.Permission) (dao.Permission, error) {
	if r.dataContainer["Save"] == nil {
		return dao.Permission{}, r.errorContainer["Save"]
	}
	return r.dataContainer["Save"].(dao.Permission), r.errorContainer["Save"]
}

func (r *PermissionRepositoryMock) DeletePermissionById(id uuid.UUID) error {
	return r.errorContainer["DeletePermissionById"]
}

/**
 * Function to create new PermissionRepositoryMock
 * @param void
 * @return PermissionRepositoryMock
 */
func NewPermissionRepositoryMock() PermissionRepositoryMock {
	return PermissionRepositoryMock{
		dataContainer:  make(map[string]interface{}),
		errorContainer: make(map[string]error),
	}
}
