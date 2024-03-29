package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"planner-backend/app/domain/dao"
	"planner-backend/app/domain/dco"
	"planner-backend/app/domain/dto"
	"planner-backend/app/mock"
	"planner-backend/app/pkg"
	"testing"
)

func TestDeletePerson(t *testing.T) {
	PersonRepository := mock.NewPersonRepositoryMock()
	personService := PersonServiceImpl{
		PersonRepository: PersonRepository,
	}

	testSteps := []ServiceTestDELETE{
		{
			mockValue: dao.Person{
				ID: "TEST",
			},
			params: map[string]string{
				"personID": "test",
			},
			mockError:          nil,
			expectedStatusCode: http.StatusOK,
		},
		{
			mockValue: dao.Person{
				ID: "test",
			},
			params: map[string]string{
				"personID": "test",
			},
			mockError:          nil,
			expectedStatusCode: http.StatusOK,
		},
		{
			mockValue: dao.Person{
				ID: "test",
			},
			params:             map[string]string{},
			mockError:          pkg.ErrNoRows,
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for i, testStep := range testSteps {
		t.Run("Test Delete Person", func(t *testing.T) {
			PersonRepository.On("FindPersonByID").Return(testStep.mockValue, testStep.mockError)
			PersonRepository.On("Delete").Return(nil, nil)

			// get GIN context
			w := httptest.NewRecorder()
			c, err := mock.NewTestContextBuilder(w).WithMethod("DELETE").WithMapParams(testStep.params).WithBody(nil).Build()
			if err != nil {
				t.Errorf("Test Step %d: Error when building context", i)
			}

			personService.DeletePerson(c)
			response := w.Result()

			if response.StatusCode != testStep.expectedStatusCode {
				t.Errorf("Test Step %d: Expected status code %d, got %d", i, testStep.expectedStatusCode, response.StatusCode)
			}
		})
	}
}

func TestUpdatePerson(t *testing.T) {
	PersonRepository := mock.NewPersonRepositoryMock()
	personService := PersonServiceImpl{
		PersonRepository: PersonRepository,
	}

	testSteps := []ServiceTestPUT{
		{
			mockRequestData: map[string]interface{}{
				"id":            "NTES",
				"first_name":    "newFirstName",
				"last_name":     "newLastName",
				"email":         "newEmail@example.com",
				"active":        true,
				"working_hours": 8,
			},
			findValue: dao.Person{
				ID:           "NTES",
				FirstName:    "oldFirstName",
				LastName:     "oldLastName",
				Email:        "oldEmail@example.com",
				Active:       true,
				WorkingHours: 8,
			},
			saveValue: dao.Person{
				ID:           "TEST",
				FirstName:    "newFirstName",
				LastName:     "newLastName",
				Email:        "newEmail@example.com",
				Active:       true,
				WorkingHours: 8,
			},
			expectedStatusCode: http.StatusOK,
			findError:          nil,
			saveError:          nil,
			params: map[string]string{
				"personID": "ntes",
			},
		},
		{
			mockRequestData: map[string]interface{}{
				"id":            "NTES",
				"first_name":    "newFirstName",
				"last_name":     "newLastName",
				"email":         "newEmail@example.com",
				"active":        true,
				"working_hours": 8,
			},
			findValue: dao.Person{
				ID:           "NTES",
				FirstName:    "oldFirstName",
				LastName:     "oldLastName",
				Email:        "oldEmail@example.com",
				Active:       true,
				WorkingHours: 8,
			},
			saveValue: dao.Person{
				ID:           "TEST",
				FirstName:    "newFirstName",
				LastName:     "newLastName",
				Email:        "newEmail@example.com",
				Active:       true,
				WorkingHours: 8,
			},
			expectedStatusCode: http.StatusOK,
			findError:          nil,
			saveError:          nil,
			params: map[string]string{
				"personID": "NTES",
			},
		},
		{
			mockRequestData: map[string]interface{}{
				"id":            "TEST",
				"first_name":    "newFirstName",
				"last_name":     "newLastName",
				"email":         "newEmail@example.com",
				"active":        true,
				"working_hours": 8,
			},
			findValue: dao.Person{
				ID:           "test",
				FirstName:    "oldFirstName",
				LastName:     "oldLastName",
				Email:        "oldEmail@example.com",
				Active:       true,
				WorkingHours: 8,
			},
			saveValue: dao.Person{
				ID:           "test",
				FirstName:    "newFirstName",
				LastName:     "newLastName",
				Email:        "newEmail@example.com",
				Active:       true,
				WorkingHours: 8,
			},
			expectedStatusCode: http.StatusOK,
			findError:          nil,
			saveError:          nil,
			params: map[string]string{
				"personID": "test",
			},
		},
		// ... add more test steps here ...
	}

	for i, testStep := range testSteps {
		t.Run(fmt.Sprintf("Test Update Person - Step %d", i), func(t *testing.T) {
			PersonRepository.On("FindPersonByID").Return(testStep.findValue, testStep.findError)
			PersonRepository.On("Save").Return(testStep.saveValue, testStep.saveError)

			// get GIN context
			w := httptest.NewRecorder()
			c, err := mock.NewTestContextBuilder(w).WithMethod("PUT").WithMapParams(testStep.params).WithBody(testStep.mockRequestData).Build()
			if err != nil {
				t.Errorf("Test Step %d: Error when building context", i)
			}

			personService.UpdatePerson(c)
			response := w.Result()

			if response.StatusCode != testStep.expectedStatusCode {
				t.Errorf("Test Step %d: Expected status code %d, got %d", i, testStep.expectedStatusCode, response.StatusCode)
			}

			if response.StatusCode != http.StatusOK {
				return
			}

			var responseBody dto.APIResponse[dco.PersonResponse]
			if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
				t.Errorf("Test Step %d: Error when decoding response body", i)
			}

			if responseBody.Data.FirstName != testStep.saveValue.(dao.Person).FirstName {
				t.Errorf("Test Step %d: Expected response body %v, got %v", i, testStep.saveValue, responseBody.Data)
			}

			if responseBody.Data.LastName != testStep.saveValue.(dao.Person).LastName {
				t.Errorf("Test Step %d: Expected response body %v, got %v", i, testStep.saveValue, responseBody.Data)
			}

			if responseBody.Data.Email != testStep.saveValue.(dao.Person).Email {
				t.Errorf("Test Step %d: Expected response body %v, got %v", i, testStep.saveValue, responseBody.Data)
			}

			if *responseBody.Data.Active != testStep.saveValue.(dao.Person).Active {
				t.Errorf("Test Step %d: Expected response body %v, got %v", i, testStep.saveValue, responseBody.Data)
			}

			if responseBody.Data.WorkingHours != testStep.saveValue.(dao.Person).WorkingHours {
				t.Errorf("Test Step %d: Expected response body %v, got %v", i, testStep.saveValue, responseBody.Data)
			}
		})
	}
}

func TestAddPerson(t *testing.T) {
	PersonRepository := mock.NewPersonRepositoryMock()
	personService := PersonServiceImpl{
		PersonRepository: PersonRepository,
	}

	testSteps := []ServiceTestPOST{
		{
			mockRequestData: map[string]interface{}{
				"id":            "TEST",
				"first_name":    "John",
				"last_name":     "Doe",
				"email":         "john.doe@example.com",
				"active":        true,
				"working_hours": 8,
			},
			findValue: nil,
			saveValue: dao.Person{
				ID:           "TEST",
				FirstName:    "John",
				LastName:     "Doe",
				Email:        "john.doe@example.com",
				Active:       true,
				WorkingHours: 8,
			},
			expectedStatusCode: http.StatusCreated,
			findError:          pkg.ErrNoRows,
			saveError:          nil,
		},
		// ... add more test steps here ...
	}

	for i, testStep := range testSteps {
		t.Run(fmt.Sprintf("Test Add Person - Step %d", i), func(t *testing.T) {

			PersonRepository.On("FindPersonByID").Return(testStep.findValue, testStep.findError)
			PersonRepository.On("Save").Return(testStep.saveValue, testStep.saveError)

			// get GIN context
			w := httptest.NewRecorder()
			c, err := mock.NewTestContextBuilder(w).WithMethod("POST").WithBody(testStep.mockRequestData).Build()
			if err != nil {
				t.Errorf("Test Step %d: Error when building context", i)
			}

			personService.AddPerson(c)
			response := w.Result()

			if response.StatusCode != testStep.expectedStatusCode {
				t.Errorf("Test Step %d: Expected status code %d, got %d", i, testStep.expectedStatusCode, response.StatusCode)
			}

			if response.StatusCode != http.StatusOK {
				return
			}

			var responseBody dto.APIResponse[dco.PersonResponse]
			if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
				t.Errorf("Test Step %d: Error when decoding response body", i)
			}

			if responseBody.Data.FirstName != testStep.saveValue.(dao.Person).FirstName {
				t.Errorf("Test Step %d: Expected response body %v, got %v", i, testStep.saveValue, responseBody.Data)
			}

			if responseBody.Data.LastName != testStep.saveValue.(dao.Person).LastName {
				t.Errorf("Test Step %d: Expected response body %v, got %v", i, testStep.saveValue, responseBody.Data)
			}

			if responseBody.Data.Email != testStep.saveValue.(dao.Person).Email {
				t.Errorf("Test Step %d: Expected response body %v, got %v", i, testStep.saveValue, responseBody.Data)
			}

			if *responseBody.Data.Active != testStep.saveValue.(dao.Person).Active {
				t.Errorf("Test Step %d: Expected response body %v, got %v", i, testStep.saveValue, responseBody.Data)
			}

			if responseBody.Data.WorkingHours != testStep.saveValue.(dao.Person).WorkingHours {
				t.Errorf("Test Step %d: Expected response body %v, got %v", i, testStep.saveValue, responseBody.Data)
			}
		})
	}
}

func TestGetAllPersons(t *testing.T) {
	PersonRepository := mock.NewPersonRepositoryMock()
	personService := PersonServiceImpl{
		PersonRepository: PersonRepository,
	}

	var trueValue = true
	var falseValue = false
	testSteps := []ServiceTestGET{
		{
			mockValue: []dao.Person{
				{
					ID:           "test",
					FirstName:    "test",
					LastName:     "test",
					Email:        "test",
					Active:       false,
					WorkingHours: 8,
				},
			},
			expectedResponse: []dco.PersonResponse{
				{
					ID:           "test",
					FirstName:    "test",
					LastName:     "test",
					Email:        "test",
					Active:       &falseValue,
					WorkingHours: 8,
				},
			},
			expectedStatusCode: http.StatusOK,
			mockError:          nil,
			queries: map[string]string{
				"department": "test",
			},
		},
		{
			mockValue: []dao.Person{
				{
					ID:           "test",
					FirstName:    "test",
					LastName:     "test",
					Email:        "test",
					Active:       true,
					WorkingHours: 8,
				},
			},
			expectedResponse: []dco.PersonResponse{
				{
					ID:           "test",
					FirstName:    "test",
					LastName:     "test",
					Email:        "test",
					Active:       &trueValue,
					WorkingHours: 8,
				},
			},
			expectedStatusCode: http.StatusOK,
			mockError:          nil,
			queries: map[string]string{
				"department": "test",
			},
		},
		{
			mockValue:          nil,
			expectedResponse:   nil,
			expectedStatusCode: http.StatusBadRequest,
			mockError:          nil,
		},
		{
			mockValue:          nil,
			expectedResponse:   nil,
			expectedStatusCode: http.StatusBadRequest,
			mockError:          pkg.ErrNoRows,
			queries:            map[string]string{},
		},
		{
			mockValue:          nil,
			expectedResponse:   nil,
			expectedStatusCode: http.StatusOK,
			mockError:          pkg.ErrNoRows,
			queries: map[string]string{
				"department": "test",
			},
		},
	}

	for i, testStep := range testSteps {
		t.Run("Test Get All Persons", func(t *testing.T) {
			PersonRepository.On("FindAllPersons").Return(testStep.mockValue, testStep.mockError)

			// get GIN context
			w := httptest.NewRecorder()
			contextBuilder := mock.NewTestContextBuilder(w)
			contextBuilder.WithMethod("GET")
			contextBuilder.WithGinParams(testStep.ParamsToGinParams())
			contextBuilder.WithBody(nil)
			contextBuilder.WithQueries(testStep.queries)
			c, err := contextBuilder.Build()
			if err != nil {
				t.Errorf("Test Step %d: Error when building context", i)
			}

			personService.GetAllPersons(c)
			response := w.Result()

			if response.StatusCode != testStep.expectedStatusCode {
				t.Errorf("Test Step %d: Expected status code %d, got %d", i, testStep.expectedStatusCode, response.StatusCode)
			}

			if response.StatusCode != http.StatusOK {
				return
			}

			var responseBody dto.APIResponse[[]dco.PersonResponse]
			if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
				t.Errorf("Test Step %d: Error when decoding response body", i)
			}

			for i, person := range responseBody.Data {
				expectedPerson := testStep.expectedResponse.([]dco.PersonResponse)[i]
				if person.ID != expectedPerson.ID {
					t.Errorf("Test Step %d: Expected ID %v, got %v", i, expectedPerson.ID, person.ID)
				}
				if person.FirstName != expectedPerson.FirstName {
					t.Errorf("Test Step %d: Expected FirstName %v, got %v", i, expectedPerson.FirstName, person.FirstName)
				}
				if person.LastName != expectedPerson.LastName {
					t.Errorf("Test Step %d: Expected LastName %v, got %v", i, expectedPerson.LastName, person.LastName)
				}
				if person.Email != expectedPerson.Email {
					t.Errorf("Test Step %d: Expected Email %v, got %v", i, expectedPerson.Email, person.Email)
				}
				if *person.Active != *expectedPerson.Active {
					t.Errorf("Test Step %d: Expected Active %v, got %v", i, expectedPerson.Active, person.Active)
				}
				if person.WorkingHours != expectedPerson.WorkingHours {
					t.Errorf("Test Step %d: Expected WorkingHours %v, got %v", i, expectedPerson.WorkingHours, person.WorkingHours)
				}
			}
		})
	}
}

func TestGetPerson(t *testing.T) {
	PersonRepository := mock.NewPersonRepositoryMock()
	personService := PersonServiceImpl{
		PersonRepository: PersonRepository,
	}

	var trueValue = true
	var falseValue = false

	testSteps := []ServiceTestGET{
		{
			mockValue: dao.Person{
				ID:           "test",
				FirstName:    "test",
				LastName:     "test",
				Email:        "test",
				Active:       false,
				WorkingHours: 8,
			},
			expectedResponse: dco.PersonResponse{
				ID:           "test",
				FirstName:    "test",
				LastName:     "test",
				Email:        "test",
				Active:       &falseValue,
				WorkingHours: 8,
			},
			expectedStatusCode: http.StatusOK,
			mockError:          nil,
			params: map[string]string{
				"personID": "test",
			},
		},
		{
			mockValue: dao.Person{
				ID:           "test",
				FirstName:    "test",
				LastName:     "test",
				Email:        "test",
				Active:       true,
				WorkingHours: 8,
			},
			expectedResponse: dco.PersonResponse{
				ID:        "test",
				FirstName: "test",
				LastName:  "test",
				Email:     "test",
				Active:    &trueValue,

				WorkingHours: 8,
			},
			expectedStatusCode: http.StatusOK,
			mockError:          nil,
			params: map[string]string{
				"personID": "test",
			},
		},
		{
			mockValue:          nil,
			expectedResponse:   nil,
			expectedStatusCode: http.StatusBadRequest,
			mockError:          nil,
		},
		{
			mockValue:          nil,
			expectedResponse:   nil,
			expectedStatusCode: http.StatusNotFound,
			mockError:          pkg.ErrNoRows,
			params: map[string]string{
				"personID": "test",
			},
		},
	}

	for i, testStep := range testSteps {
		t.Run(fmt.Sprintf("Test Get Person - Step %d", i), func(t *testing.T) {
			PersonRepository.On("FindPersonByID").Return(testStep.mockValue, testStep.mockError)

			// get GIN context
			w := httptest.NewRecorder()
			c, err := mock.NewTestContextBuilder(w).WithMethod("GET").WithMapParams(testStep.params).WithBody(nil).Build()
			if err != nil {
				t.Errorf("Test Step %d: Error when building context", i)
			}

			personService.GetPersonByID(c)
			response := w.Result()

			if response.StatusCode != testStep.expectedStatusCode {
				t.Errorf("Test Step %d: Expected status code %d, got %d", i, testStep.expectedStatusCode, response.StatusCode)
			}

			if response.StatusCode != http.StatusOK {
				return
			}

			var responseBody dto.APIResponse[dco.PersonResponse]
			if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
				t.Errorf("Test Step %d: Error when decoding response body", i)
			}

			expectedPerson := testStep.expectedResponse.(dco.PersonResponse)
			person := responseBody.Data
			if person.ID != expectedPerson.ID {
				t.Errorf("Test Step %d: Expected ID %v, got %v", i, expectedPerson.ID, person.ID)
			}

			if person.FirstName != expectedPerson.FirstName {
				t.Errorf("Test Step %d: Expected FirstName %v, got %v", i, expectedPerson.FirstName, person.FirstName)
			}

			if person.LastName != expectedPerson.LastName {
				t.Errorf("Test Step %d: Expected LastName %v, got %v", i, expectedPerson.LastName, person.LastName)
			}

			if person.Email != expectedPerson.Email {
				t.Errorf("Test Step %d: Expected Email %v, got %v", i, expectedPerson.Email, person.Email)
			}

			if *person.Active != *expectedPerson.Active {
				t.Errorf("Test Step %d: Expected Active %v, got %v", i, expectedPerson.Active, person.Active)
			}

			if person.WorkingHours != expectedPerson.WorkingHours {
				t.Errorf("Test Step %d: Expected WorkingHours %v, got %v", i, expectedPerson.WorkingHours, person.WorkingHours)
			}
		})
	}
}
