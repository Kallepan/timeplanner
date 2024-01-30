package mock

import (
	"planner-backend/app/domain/dao"
)

type PersonRepositoryMock struct {
	dataContainer      map[string]interface{}
	errorContainer     map[string]error
	primedFunctionName string
}

/* Mock interface implementations */
func (r *PersonRepositoryMock) On(functionName string) Mock {
	// set default value
	r.dataContainer[functionName] = nil
	r.errorContainer[functionName] = nil

	// Set primed function name
	r.primedFunctionName = functionName

	return r
}

func (r *PersonRepositoryMock) Return(mockData interface{}, errorData error) Mock {
	r.dataContainer[r.primedFunctionName] = mockData
	r.errorContainer[r.primedFunctionName] = errorData

	return r
}

/* Repository interface implementations */
func (r *PersonRepositoryMock) FindAllPersons(departmentID string) ([]dao.Person, error) {
	if r.dataContainer["FindAllPersons"] == nil {
		return nil, r.errorContainer["FindAllPersons"]
	}
	return r.dataContainer["FindAllPersons"].([]dao.Person), r.errorContainer["FindAllPersons"]
}

func (r *PersonRepositoryMock) FindAllPersonsBy(departmentID string, workplaceID string, weekdayID string, notAbsentOn string) ([]dao.Person, error) {
	if r.dataContainer["FindAllPersonsBy"] == nil {
		return nil, r.errorContainer["FindAllPersonsBy"]
	}
	return r.dataContainer["FindAllPersonsBy"].([]dao.Person), r.errorContainer["FindAllPersonsBy"]
}

func (r *PersonRepositoryMock) FindPersonByID(id string) (dao.Person, error) {
	if r.dataContainer["FindPersonByID"] == nil {
		return dao.Person{}, r.errorContainer["FindPersonByID"]
	}

	return r.dataContainer["FindPersonByID"].(dao.Person), r.errorContainer["FindPersonByID"]
}

func (r *PersonRepositoryMock) Save(person *dao.Person) (dao.Person, error) {
	if r.dataContainer["Save"] == nil {
		return dao.Person{}, r.errorContainer["Save"]
	}
	return r.dataContainer["Save"].(dao.Person), r.errorContainer["Save"]
}

func (r *PersonRepositoryMock) Delete(person *dao.Person) error {
	return r.errorContainer["Delete"]
}

/**
* Function to create new PersonRepositoryMock
 */
func NewPersonRepositoryMock() *PersonRepositoryMock {
	return &PersonRepositoryMock{
		dataContainer:  make(map[string]interface{}),
		errorContainer: make(map[string]error),
	}
}
