package mock

import (
	"auth-backend/app/domain/dao"

	"github.com/google/uuid"
)

type UserRepositoryMock struct {
	dataContainer      map[string]interface{}
	errorContainer     map[string]error
	primedFunctionName string
}

/* Mock interface implementations */
func (r *UserRepositoryMock) On(functionName string) Mock {
	// set default value
	r.dataContainer[functionName] = nil
	r.errorContainer[functionName] = nil

	// Set primed function name
	r.primedFunctionName = functionName

	return r
}

func (r *UserRepositoryMock) Return(mockData interface{}, errorData error) Mock {
	r.dataContainer[r.primedFunctionName] = mockData
	r.errorContainer[r.primedFunctionName] = errorData

	return r
}

/* Repostory interface implementations */
func (r *UserRepositoryMock) FindAllUsers() ([]dao.User, error) {
	return r.dataContainer["FindAllUsers"].([]dao.User), r.errorContainer[r.primedFunctionName]
}

func (r *UserRepositoryMock) FindUserById(id uuid.UUID) (dao.User, error) {
	return r.dataContainer["FindUserById"].(dao.User), r.errorContainer[r.primedFunctionName]
}

func (r *UserRepositoryMock) Save(user *dao.User) (dao.User, error) {
	return r.dataContainer["Save"].(dao.User), r.errorContainer[r.primedFunctionName]
}

func (r *UserRepositoryMock) DeleteUser(id uuid.UUID) error {
	return r.errorContainer["DeleteUser"]
}

func (r *UserRepositoryMock) AddPermissionToUser(userID uuid.UUID, permissionID uuid.UUID) error {
	return r.errorContainer["AddPermissionToUser"]
}

func (r *UserRepositoryMock) DeletePermissionFromUser(userID uuid.UUID, permissionID uuid.UUID) error {
	return r.errorContainer["DeletePermissionFromUser"]
}

/**
 * Function to create new UserRepositoryMock
 * @param void
 * @return UserRepositoryMock
 */

func NewUserRepositoryMock() UserRepositoryMock {
	return UserRepositoryMock{
		dataContainer:  make(map[string]interface{}),
		errorContainer: make(map[string]error),
	}
}
