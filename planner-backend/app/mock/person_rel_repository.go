package mock

import "planner-backend/app/domain/dao"

type PersonRelRepositoryMock struct {
	dataContainer      map[string]interface{}
	errorContainer     map[string]error
	primedFunctionName string
}

/* Mock interface implementations */
func (r *PersonRelRepositoryMock) On(functionName string) Mock {
	// set default value
	r.dataContainer[functionName] = nil
	r.errorContainer[functionName] = nil

	// Set primed function name
	r.primedFunctionName = functionName

	return r
}

func (r *PersonRelRepositoryMock) Return(mockData interface{}, errorData error) Mock {
	r.dataContainer[r.primedFunctionName] = mockData
	r.errorContainer[r.primedFunctionName] = errorData

	return r
}

/* Repository interface implementations */
func (r *PersonRelRepositoryMock) AddAbsencyToPerson(person dao.Person, absence dao.Absence) error {
	return r.errorContainer["AddAbsencyToPerson"]
}
func (r *PersonRelRepositoryMock) RemoveAbsencyFromPerson(person dao.Person, absency dao.Absence) error {
	return r.errorContainer["RemoveAbsencyFromPerson"]
}
func (r *PersonRelRepositoryMock) FindAbsencyForPerson(personID string, date string) (dao.Absence, error) {
	if r.dataContainer["FindAbsencyForPerson"] == nil {
		return dao.Absence{}, r.errorContainer["FindAbsencyForPerson"]
	}
	return r.dataContainer["FindAbsencyForPerson"].(dao.Absence), r.errorContainer["FindAbsencyForPerson"]
}

func (r *PersonRelRepositoryMock) AddDepartmentToPerson(person dao.Person, departmentName string) error {
	return r.errorContainer["AddDepartmentToPerson"]
}
func (r *PersonRelRepositoryMock) RemoveDepartmentFromPerson(person dao.Person, departmentName string) error {
	return r.errorContainer["RemoveDepartmentFromPerson"]
}
func (r *PersonRelRepositoryMock) AddWorkplaceToPerson(person dao.Person, workplaceName string) error {
	return r.errorContainer["AddWorkplaceToPerson"]
}
func (r *PersonRelRepositoryMock) RemoveWorkplaceFromPerson(person dao.Person, workplaceName string) error {
	return r.errorContainer["RemoveWorkplaceFromPerson"]
}
func (r *PersonRelRepositoryMock) AddWeekdayToPerson(person dao.Person, weekdayID string) error {
	return r.errorContainer["AddWeekdayToPerson"]
}
func (r *PersonRelRepositoryMock) RemoveWeekdayFromPerson(person dao.Person, weekdayID string) error {
	return r.errorContainer["RemoveWeekdayFromPerson"]
}

/** Function to create new PersonRelRepositoryMock */
func NewPersonRelRepositoryMock() *PersonRelRepositoryMock {
	return &PersonRelRepositoryMock{
		dataContainer:  make(map[string]interface{}),
		errorContainer: make(map[string]error),
	}
}
