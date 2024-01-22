package service

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"planner-backend/app/domain/dao"
	"planner-backend/app/mock"
	"planner-backend/app/pkg"
	"testing"
)

type serviceTestPersonRel struct {
	params      map[string]string
	mockRequest map[string]interface{}
	mockValue   interface{}
	mockError   error

	expectedStatusCode int

	// find fields
	findValue interface{}
	findError error

	// extra fields
	additionalFindValue interface{}
	additionalFindError error
}

func TestAddAbsencyToPerson(t *testing.T) {
	PersonRepository := mock.NewPersonRepositoryMock()
	PersonRelRepository := mock.NewPersonRelRepositoryMock()
	personRelService := PersonRelServiceImpl{
		PersonRepository:    PersonRepository,
		PersonRelRepository: PersonRelRepository,
	}

	testSteps := []serviceTestPersonRel{
		{
			// capital letters personID
			mockRequest: map[string]interface{}{
				"reason": "reason1",
				"date":   "2020-01-01",
			},
			findValue: dao.Person{
				ID: "TEST",
			},
			findError: nil,
			params: map[string]string{
				"personID": "test",
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			// small letters personID
			mockRequest: map[string]interface{}{
				"reason": "reason1",
				"date":   "2020-01-01",
			},
			findValue: dao.Person{
				ID: "test",
			},
			findError: nil,
			params: map[string]string{
				"personID": "test",
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			// no personID
			mockRequest: map[string]interface{}{
				"reason": "reason1",
				"date":   "2020-01-01",
			},
			findValue: dao.Person{
				ID: "test",
			},
			findError:          nil,
			params:             map[string]string{},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			// no date
			findValue: nil,
			findError: pkg.ErrNoRows,
			params: map[string]string{
				"personID": "test",
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			// no data in request
			findValue:          nil,
			findError:          nil,
			params:             map[string]string{},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			// error in dao
			mockRequest: map[string]interface{}{
				"reason": "reason1",
				"date":   "2020-01-01",
			},
			findValue: dao.Person{
				ID: "test",
			},
			findError: nil,
			params: map[string]string{
				"personID": "test",
			},
			mockError:          errors.New("test"),
			expectedStatusCode: 500,
		},
	}

	for i, testStep := range testSteps {
		t.Run("Test Add Absency To Person", func(t *testing.T) {
			PersonRepository.On("FindPersonByID").Return(testStep.findValue, testStep.findError)
			PersonRelRepository.On("AddAbsencyToPerson").Return(testStep.mockValue, testStep.mockError)

			// get GIN context
			w := httptest.NewRecorder()
			c, err := mock.NewTestContextBuilder(w).WithMethod("POST").WithBody(testStep.mockRequest).WithParamsRaw(testStep.params).Build()
			if err != nil {
				t.Errorf("Test Step %d: Error while building context: %s", i, err)
			}

			personRelService.AddAbsencyToPerson(c)
			response := w.Result()

			if response.StatusCode != testStep.expectedStatusCode {
				t.Errorf("Test Step %d: Expected status code %d, got %d", i, testStep.expectedStatusCode, response.StatusCode)
			}
		})
	}
}

func TestRemoveAbsencyFromPerson(t *testing.T) {
	PersonRepository := mock.NewPersonRepositoryMock()
	PersonRelRepository := mock.NewPersonRelRepositoryMock()
	personRelService := PersonRelServiceImpl{
		PersonRepository:    PersonRepository,
		PersonRelRepository: PersonRelRepository,
	}

	testSteps := []serviceTestPersonRel{
		{
			// capital letters personID
			findValue: dao.Person{
				ID: "TEST",
			},
			findError: nil,
			params: map[string]string{
				"personID": "test",
				"date":     "2020-01-01",
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			// small letters personID
			findValue: dao.Person{
				ID: "test",
			},
			findError: nil,
			params: map[string]string{
				"personID": "test",
				"date":     "2020-01-01",
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			// no date
			findValue: dao.Person{
				ID: "test",
			},
			findError: nil,
			params: map[string]string{
				"personID": "test",
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			// no person found
			findValue: nil,
			findError: pkg.ErrNoRows,
			params: map[string]string{
				"personID": "test",
				"date":     "2020-01-01",
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			// no personID
			findValue: nil,
			findError: nil,
			params: map[string]string{
				"date": "2020-01-01",
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			// error in dao
			findValue: dao.Person{
				ID: "test",
			},
			findError: nil,
			params: map[string]string{
				"personID": "test",
				"date":     "2020-01-01",
			},
			mockError:          errors.New("test"),
			expectedStatusCode: 500,
		},
	}

	for i, testStep := range testSteps {
		t.Run("Test Remove Absency From Person", func(t *testing.T) {
			PersonRepository.On("FindPersonByID").Return(testStep.findValue, testStep.findError)
			PersonRelRepository.On("RemoveAbsencyFromPerson").Return(testStep.mockValue, testStep.mockError)

			// get GIN context
			w := httptest.NewRecorder()
			c, err := mock.NewTestContextBuilder(w).WithMethod("POST").WithBody(testStep.mockRequest).WithParamsRaw(testStep.params).Build()
			if err != nil {
				t.Errorf("Test Step %d: Error while building context: %s", i, err)
			}

			personRelService.RemoveAbsencyFromPerson(c)
			response := w.Result()

			if response.StatusCode != testStep.expectedStatusCode {
				t.Errorf("Test Step %d: Expected status code %d, got %d", i, testStep.expectedStatusCode, response.StatusCode)
			}
		})
	}
}

func TestFindAbsencyForPerson(t *testing.T) {
	PersonRepository := mock.NewPersonRepositoryMock()
	PersonRelRepository := mock.NewPersonRelRepositoryMock()
	personRelService := PersonRelServiceImpl{
		PersonRepository:    PersonRepository,
		PersonRelRepository: PersonRelRepository,
	}

	testSteps := []serviceTestPersonRel{
		{
			// capital letters personID
			params: map[string]string{
				"personID": "test",
				"date":     "2020-01-01",
			},
			findValue: dao.Absence{
				Reason: "reason1",
				Date:   "2020-01-01",
			},
			findError:          nil,
			expectedStatusCode: http.StatusOK,
		},
		{
			// small letters personID
			params: map[string]string{
				"date": "2020-01-01",
			},
			findValue:          nil,
			findError:          pkg.ErrNoRows,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			// no date
			params: map[string]string{
				"personID": "test",
			},
			findValue:          nil,
			findError:          pkg.ErrNoRows,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			// error in dao
			params: map[string]string{
				"personID": "test",
				"date":     "2020-01-01",
			},
			findValue:          nil,
			findError:          errors.New("test"),
			expectedStatusCode: 500,
		},
	}

	for i, testStep := range testSteps {
		t.Run("Test Find Absency For Person", func(t *testing.T) {
			PersonRelRepository.On("FindAbsencyForPerson").Return(testStep.findValue, testStep.findError)
			// get GIN context
			w := httptest.NewRecorder()
			c, err := mock.NewTestContextBuilder(w).WithMethod("GET").WithParamsRaw(testStep.params).Build()
			if err != nil {
				t.Errorf("Test Step %d: Error while building context: %s", i, err)
			}

			personRelService.FindAbsencyForPerson(c)
			response := w.Result()

			if response.StatusCode != testStep.expectedStatusCode {
				t.Errorf("Test Step %d: Expected status code %d, got %d", i, testStep.expectedStatusCode, response.StatusCode)
			}
		})
	}
}

func TestAddDepartmentToPerson(t *testing.T) {
	PersonRepository := mock.NewPersonRepositoryMock()
	PersonRelRepository := mock.NewPersonRelRepositoryMock()
	DepartmentRepository := mock.NewDepartmentRepositoryMock()
	personRelService := PersonRelServiceImpl{
		PersonRepository:     PersonRepository,
		PersonRelRepository:  PersonRelRepository,
		DepartmentRepository: DepartmentRepository,
	}

	testSteps := []serviceTestPersonRel{
		{
			// no department found
			mockRequest: map[string]interface{}{
				"department_name": "department1",
			},
			findValue: dao.Person{
				ID: "TEST",
			},
			findError: nil,
			params: map[string]string{
				"personID": "test",
			},
			expectedStatusCode:  http.StatusBadRequest,
			additionalFindError: pkg.ErrNoRows,
		},
		{
			// capital letters personID
			mockRequest: map[string]interface{}{
				"department_name": "department1",
			},
			findValue: dao.Person{
				ID: "TEST",
			},
			findError: nil,
			params: map[string]string{
				"personID": "test",
			},
			expectedStatusCode: http.StatusOK,
			additionalFindValue: dao.Department{
				Name: "department1",
			},
			additionalFindError: nil,
		},
		{
			// small letters personID
			mockRequest: map[string]interface{}{
				"department_name": "department1",
			},
			findValue: dao.Person{
				ID: "test",
			},
			findError: nil,
			params: map[string]string{
				"personID": "test",
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			// no request body
			findValue: dao.Person{
				ID: "test",
			},
			findError: nil,
			params: map[string]string{
				"personID": "test",
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			// no person found
			mockRequest: map[string]interface{}{
				"department_name": "department1",
			},
			findValue: nil,
			findError: pkg.ErrNoRows,
			params: map[string]string{
				"personID": "test",
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			// incorrect request params
			mockRequest: map[string]interface{}{
				"department_name": "department1",
			},
			findValue:          nil,
			findError:          nil,
			params:             map[string]string{},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			// error in dao
			mockRequest: map[string]interface{}{
				"department_name": "department1",
			},
			findValue: dao.Person{
				ID: "test",
			},
			findError: nil,
			params: map[string]string{
				"personID": "test",
			},
			mockError:          errors.New("test"),
			expectedStatusCode: 500,
		},
	}

	for i, testStep := range testSteps {
		t.Run("Test Add Department To Person", func(t *testing.T) {
			PersonRepository.On("FindPersonByID").Return(testStep.findValue, testStep.findError)
			PersonRelRepository.On("AddDepartmentToPerson").Return(testStep.mockValue, testStep.mockError)
			DepartmentRepository.On("FindDepartmentByID").Return(testStep.additionalFindValue, testStep.additionalFindError)

			// get GIN context
			w := httptest.NewRecorder()
			c, err := mock.NewTestContextBuilder(w).WithMethod("POST").WithBody(testStep.mockRequest).WithParamsRaw(testStep.params).Build()
			if err != nil {
				t.Errorf("Test Step %d: Error while building context: %s", i, err)
			}

			personRelService.AddDepartmentToPerson(c)
			response := w.Result()

			if response.StatusCode != testStep.expectedStatusCode {
				t.Errorf("Test Step %d: Expected status code %d, got %d", i, testStep.expectedStatusCode, response.StatusCode)
			}
		})
	}
}

func TestRemoveDepartmentFromPerson(t *testing.T) {
	PersonRepository := mock.NewPersonRepositoryMock()
	PersonRelRepository := mock.NewPersonRelRepositoryMock()
	personRelService := PersonRelServiceImpl{
		PersonRepository:    PersonRepository,
		PersonRelRepository: PersonRelRepository,
	}

	testSteps := []serviceTestPersonRel{
		{
			// capital letters personID
			findValue: dao.Person{
				ID: "TEST",
			},
			findError: nil,
			params: map[string]string{
				"personID":     "test",
				"departmentID": "department1",
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			// small letters personID
			findValue: dao.Person{
				ID: "test",
			},
			findError: nil,
			params: map[string]string{
				"personID":     "test",
				"departmentID": "department1",
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			findValue: dao.Person{
				ID: "test",
			},
			findError: nil,
			params: map[string]string{
				"personID": "test",
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			findValue: nil,
			findError: pkg.ErrNoRows,
			params: map[string]string{
				"personID":     "test",
				"departmentID": "department1",
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			findValue: nil,
			findError: nil,
			params: map[string]string{
				"departmentID": "department1",
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			findValue: dao.Person{
				ID: "test",
			},
			findError: nil,
			params: map[string]string{
				"personID":     "test",
				"departmentID": "department1",
			},
			mockError:          errors.New("test"),
			expectedStatusCode: 500,
		},
	}

	for i, testStep := range testSteps {
		t.Run("Test Remove Department From Person", func(t *testing.T) {
			PersonRepository.On("FindPersonByID").Return(testStep.findValue, testStep.findError)
			PersonRelRepository.On("RemoveDepartmentFromPerson").Return(testStep.mockValue, testStep.mockError)

			// get GIN context
			w := httptest.NewRecorder()
			c, err := mock.NewTestContextBuilder(w).WithMethod("DELETE").WithParamsRaw(testStep.params).Build()
			if err != nil {
				t.Errorf("Test Step %d: Error while building context: %s", i, err)
			}

			personRelService.RemoveDepartmentFromPerson(c)
			response := w.Result()

			if response.StatusCode != testStep.expectedStatusCode {
				t.Errorf("Test Step %d: Expected status code %d, got %d", i, testStep.expectedStatusCode, response.StatusCode)
			}
		})
	}
}

func TestAddWorkplaceToPerson(t *testing.T) {
	PersonRepository := mock.NewPersonRepositoryMock()
	PersonRelRepository := mock.NewPersonRelRepositoryMock()
	WorkplaceRepository := mock.NewWorkplaceRepositoryMock()
	personRelService := PersonRelServiceImpl{
		PersonRepository:    PersonRepository,
		PersonRelRepository: PersonRelRepository,
		WorkplaceRepository: WorkplaceRepository,
	}

	testSteps := []serviceTestPersonRel{
		{
			// not found workplace
			mockRequest: map[string]interface{}{
				"department_name": "department1",
				"workplace_name":  "workplace1",
			},
			findValue: dao.Person{
				ID: "TEST",
			},
			findError: nil,
			params: map[string]string{
				"personID": "test",
			},
			expectedStatusCode:  http.StatusBadRequest,
			additionalFindValue: nil,
			additionalFindError: pkg.ErrNoRows,
		},
		{
			mockRequest: map[string]interface{}{
				"department_name": "department1",
				"workplace_name":  "workplace1",
			},
			findValue: dao.Person{
				ID: "TEST",
			},
			findError: nil,
			params: map[string]string{
				"personID": "test",
			},
			expectedStatusCode: http.StatusOK,
			additionalFindValue: dao.Workplace{
				Name: "workplace1",
			},
			additionalFindError: nil,
		},
		{
			mockRequest: map[string]interface{}{
				"department_name": "department1",
				"workplace_name":  "workplace1",
			},
			findValue: dao.Person{
				ID: "test",
			},
			findError: nil,
			params: map[string]string{
				"personID": "test",
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			findValue: dao.Person{
				ID: "test",
			},
			findError: nil,
			params: map[string]string{
				"personID": "test",
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			mockRequest: map[string]interface{}{
				"department_name": "department1",
				"workplace_name":  "workplace1",
			},
			findValue: nil,
			findError: pkg.ErrNoRows,
			params: map[string]string{
				"personID": "test",
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			mockRequest: map[string]interface{}{
				"department_name": "department1",
				"workplace_name":  "workplace1",
			},
			findValue:          nil,
			findError:          nil,
			params:             map[string]string{},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			mockRequest: map[string]interface{}{
				"department_name": "department1",
				"workplace_name":  "workplace1",
			},
			findValue: dao.Person{
				ID: "test",
			},
			findError: nil,
			params: map[string]string{
				"personID": "test",
			},
			mockError:          errors.New("test"),
			expectedStatusCode: 500,
		},
	}

	for i, testStep := range testSteps {
		t.Run("Test Add Workplace To Person", func(t *testing.T) {
			PersonRepository.On("FindPersonByID").Return(testStep.findValue, testStep.findError)
			PersonRelRepository.On("AddWorkplaceToPerson").Return(testStep.mockValue, testStep.mockError)
			WorkplaceRepository.On("FindWorkplaceByID").Return(testStep.additionalFindValue, testStep.additionalFindError)

			// get GIN context
			w := httptest.NewRecorder()
			c, err := mock.NewTestContextBuilder(w).WithMethod("POST").WithBody(testStep.mockRequest).WithParamsRaw(testStep.params).Build()
			if err != nil {
				t.Errorf("Test Step %d: Error while building context: %s", i, err)
			}

			personRelService.AddWorkplaceToPerson(c)
			response := w.Result()

			if response.StatusCode != testStep.expectedStatusCode {
				t.Errorf("Test Step %d: Expected status code %d, got %d", i, testStep.expectedStatusCode, response.StatusCode)
			}
		})
	}
}

func TestRemoveWorkplaceFromPerson(t *testing.T) {
	PersonRepository := mock.NewPersonRepositoryMock()
	PersonRelRepository := mock.NewPersonRelRepositoryMock()
	personRelService := PersonRelServiceImpl{
		PersonRepository:    PersonRepository,
		PersonRelRepository: PersonRelRepository,
	}

	testSteps := []serviceTestPersonRel{
		{
			mockRequest: map[string]interface{}{
				"department_name": "department1",
				"workplace_name":  "workplace1",
			},
			findValue: dao.Person{
				ID: "TEST",
			},
			findError: nil,
			params: map[string]string{
				"personID": "test",
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			mockRequest: map[string]interface{}{
				"department_name": "department1",
				"workplace_name":  "workplace1",
			},
			findValue: dao.Person{
				ID: "test",
			},
			findError: nil,
			params: map[string]string{
				"personID": "test",
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			mockRequest: map[string]interface{}{},
			findValue: dao.Person{
				ID: "test",
			},
			findError: nil,
			params: map[string]string{
				"personID": "test",
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			mockRequest: map[string]interface{}{
				"department_name": "department1",
				"workplace_name":  "workplace1",
			},
			findValue: nil,
			findError: pkg.ErrNoRows,
			params: map[string]string{
				"personID": "test",
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			mockRequest: map[string]interface{}{
				"department_name": "department1",
				"workplace_name":  "workplace1",
			},
			findValue:          nil,
			findError:          nil,
			params:             map[string]string{},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			mockRequest: map[string]interface{}{
				"department_name": "department1",
				"workplace_name":  "workplace1",
			},
			findValue: dao.Person{
				ID: "test",
			},
			findError: nil,
			params: map[string]string{
				"personID": "test",
			},
			mockError:          errors.New("test"),
			expectedStatusCode: 500,
		},
	}

	for i, testStep := range testSteps {
		t.Run("Test Remove Workplace From Person", func(t *testing.T) {
			PersonRepository.On("FindPersonByID").Return(testStep.findValue, testStep.findError)
			PersonRelRepository.On("RemoveWorkplaceFromPerson").Return(testStep.mockValue, testStep.mockError)

			// get GIN context
			w := httptest.NewRecorder()
			c, err := mock.NewTestContextBuilder(w).WithMethod("DELETE").WithParamsRaw(testStep.params).WithBody(testStep.mockRequest).Build()
			if err != nil {
				t.Errorf("Test Step %d: Error while building context: %s", i, err)
			}

			personRelService.RemoveWorkplaceFromPerson(c)
			response := w.Result()

			if response.StatusCode != testStep.expectedStatusCode {
				t.Errorf("Test Step %d: Expected status code %d, got %d", i, testStep.expectedStatusCode, response.StatusCode)
			}
		})
	}
}

func TestAddWeekdayToPerson(t *testing.T) {
	PersonRepository := mock.NewPersonRepositoryMock()
	PersonRelRepository := mock.NewPersonRelRepositoryMock()
	personRelService := PersonRelServiceImpl{
		PersonRepository:    PersonRepository,
		PersonRelRepository: PersonRelRepository,
	}

	testSteps := []serviceTestPersonRel{
		{
			mockRequest: map[string]interface{}{
				"weekday_id": "1",
			},
			findValue: dao.Person{
				ID: "TEST",
			},
			findError: nil,
			params: map[string]string{
				"personID": "test",
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			mockRequest: map[string]interface{}{
				"weekday_id": "MON",
			},
			findValue: dao.Person{
				ID: "TEST",
			},
			findError: nil,
			params: map[string]string{
				"personID": "test",
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			mockRequest: map[string]interface{}{
				"weekday_id": "MON",
			},
			findValue: dao.Person{
				ID: "test",
			},
			findError: nil,
			params: map[string]string{
				"personID": "test",
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			findValue: dao.Person{
				ID: "test",
			},
			findError: nil,
			params: map[string]string{
				"personID": "test",
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			mockRequest: map[string]interface{}{
				"weekday_id": "MON",
			},
			findValue: nil,
			findError: pkg.ErrNoRows,
			params: map[string]string{
				"personID": "test",
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			mockRequest: map[string]interface{}{
				"weekday_id": "MON",
			},
			findValue:          nil,
			findError:          nil,
			params:             map[string]string{},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			mockRequest: map[string]interface{}{
				"weekday_id": "MON",
			},
			findValue: dao.Person{
				ID: "test",
			},
			findError: nil,
			params: map[string]string{
				"personID": "test",
			},
			mockError:          errors.New("test"),
			expectedStatusCode: 500,
		},
	}

	for i, testStep := range testSteps {
		t.Run("Test Add Weekday To Person", func(t *testing.T) {
			PersonRepository.On("FindPersonByID").Return(testStep.findValue, testStep.findError)
			PersonRelRepository.On("AddWeekdayToPerson").Return(testStep.mockValue, testStep.mockError)

			// get GIN context
			w := httptest.NewRecorder()
			c, err := mock.NewTestContextBuilder(w).WithMethod("POST").WithBody(testStep.mockRequest).WithParamsRaw(testStep.params).Build()
			if err != nil {
				t.Errorf("Test Step %d: Error while building context: %s", i, err)
			}

			personRelService.AddWeekdayToPerson(c)
			response := w.Result()

			if response.StatusCode != testStep.expectedStatusCode {
				t.Errorf("Test Step %d: Expected status code %d, got %d", i, testStep.expectedStatusCode, response.StatusCode)
			}
		})
	}
}

func TestRemoveWeekdayFromPerson(t *testing.T) {
	PersonRepository := mock.NewPersonRepositoryMock()
	PersonRelRepository := mock.NewPersonRelRepositoryMock()
	personRelService := PersonRelServiceImpl{
		PersonRepository:    PersonRepository,
		PersonRelRepository: PersonRelRepository,
	}

	testSteps := []serviceTestPersonRel{
		{
			findValue: dao.Person{
				ID: "TEST",
			},
			findError: nil,
			params: map[string]string{
				"personID":  "test",
				"weekdayID": "INV",
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			findValue: dao.Person{
				ID: "TEST",
			},
			findError: nil,
			params: map[string]string{
				"personID":  "test",
				"weekdayID": "MON",
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			findValue: dao.Person{
				ID: "test",
			},
			findError: nil,
			params: map[string]string{
				"personID":  "test",
				"weekdayID": "MON",
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			findValue: dao.Person{
				ID: "test",
			},
			findError: nil,
			params: map[string]string{
				"personID": "test",
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			findValue: nil,
			findError: pkg.ErrNoRows,
			params: map[string]string{
				"personID":  "test",
				"weekdayID": "MON",
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			findValue: nil,
			findError: nil,
			params: map[string]string{
				"weekdayID": "MON",
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			findValue: dao.Person{
				ID: "test",
			},
			findError: nil,
			params: map[string]string{
				"personID":  "test",
				"weekdayID": "MON",
			},
			mockError:          errors.New("test"),
			expectedStatusCode: 500,
		},
	}

	for i, testStep := range testSteps {
		t.Run("Test Remove Weekday From Person", func(t *testing.T) {
			PersonRepository.On("FindPersonByID").Return(testStep.findValue, testStep.findError)
			PersonRelRepository.On("RemoveWeekdayFromPerson").Return(testStep.mockValue, testStep.mockError)

			// get GIN context
			w := httptest.NewRecorder()
			c, err := mock.NewTestContextBuilder(w).WithMethod("DELETE").WithParamsRaw(testStep.params).Build()
			if err != nil {
				t.Errorf("Test Step %d: Error while building context: %s", i, err)
			}

			personRelService.RemoveWeekdayFromPerson(c)
			response := w.Result()

			if response.StatusCode != testStep.expectedStatusCode {
				t.Errorf("Test Step %d: Expected status code %d, got %d", i, testStep.expectedStatusCode, response.StatusCode)
			}
		})
	}
}
