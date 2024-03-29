package mock

import (
	"api-gateway/app/domain/dao"

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
	if _, ok := mockData.([]dao.User); ok {
		formattedData := mockData.([]dao.User)
		for i := range formattedData {
			formattedData[i].Department = dao.Department{
				BaseModel: dao.BaseModel{
					ID: mockData.([]dao.User)[i].DepartmentID,
				},
			}
		}
	} else if _, ok := mockData.(dao.User); ok {
		formattedData := mockData.(dao.User)
		formattedData.Department = dao.Department{
			BaseModel: dao.BaseModel{
				ID: mockData.(dao.User).DepartmentID,
			},
			Name: mockData.(dao.User).Department.Name,
		}
		mockData = formattedData
	}

	r.dataContainer[r.primedFunctionName] = mockData
	r.errorContainer[r.primedFunctionName] = errorData

	return r
}

/* Repostory interface implementations */
func (r *UserRepositoryMock) FindAllUsers() ([]dao.User, error) {
	if r.dataContainer["FindAllUsers"] == nil {
		return nil, r.errorContainer["FindAllUsers"]
	}

	return r.dataContainer["FindAllUsers"].([]dao.User), r.errorContainer["FindAllUsers"]
}

func (r *UserRepositoryMock) FindUserByUsername(username string) (dao.User, error) {
	if r.dataContainer["FindUserByUsername"] == nil {
		return dao.User{}, r.errorContainer["FindUserByUsername"]
	}

	return r.dataContainer["FindUserByUsername"].(dao.User), r.errorContainer["FindUserByUsername"]
}

func (r *UserRepositoryMock) FindUserById(id uuid.UUID) (dao.User, error) {
	if r.dataContainer["FindUserById"] == nil {
		return dao.User{}, r.errorContainer["FindUserById"]
	}

	return r.dataContainer["FindUserById"].(dao.User), r.errorContainer["FindUserById"]
}

func (r *UserRepositoryMock) Save(user *dao.User) (dao.User, error) {
	if r.dataContainer["Save"] == nil {
		return dao.User{}, r.errorContainer["Save"]
	}

	return r.dataContainer["Save"].(dao.User), r.errorContainer["Save"]
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
