package service

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"planner-backend/app/domain/dao"
	"planner-backend/app/domain/dco"
	"planner-backend/app/domain/dto"
	"planner-backend/app/mock"
	"planner-backend/app/pkg"
	"testing"
)

func TestGetWorkdaysForDepartmentAndDate(t *testing.T) {
	workdayRepository := mock.NewWorkdayRepositoryMock()
	workdayService := WorkdayServiceImpl{
		WorkdayRepository: workdayRepository,
	}

	mockPerson := dao.Person{
		ID:           "person1",
		FirstName:    "first1",
		LastName:     "last1",
		Email:        "email1",
		WorkingHours: 8,
	}
	testSteps := []ServiceTestGET{
		{
			queries: map[string]string{
				"department": "department1",
				"date":       "2021-01-01",
			},
			mockValue: []dao.Workday{
				{
					DepartmentID: "department1",
					WorkplaceID:  "workplace1",
					TimeslotName: "timeslot1",
					Date:         "2021-01-01",
					StartTime:    "08:00:00",
					EndTime:      "16:00:00",
					Person:       nil,
				},
			},
			expectedStatusCode: http.StatusOK,
			mockError:          nil,
		},
		{
			queries: map[string]string{
				"department": "department1",
				"date":       "2021-01-01",
			},
			expectedStatusCode: http.StatusOK,
			mockError:          nil,
		},
		{
			queries: map[string]string{
				"department": "department1",
				"date":       "2021-01-01",
			},
			expectedStatusCode: http.StatusOK,
			mockError:          nil,
		},
		{
			queries: map[string]string{
				"department": "department1",
				"date":       "2021-01-01",
			},
			mockValue: []dao.Workday{
				{
					DepartmentID: "department1",
					WorkplaceID:  "workplace1",
					TimeslotName: "timeslot1",
					Date:         "2021-01-01",
					StartTime:    "08:00:00",
					EndTime:      "16:00:00",
					Person:       &mockPerson,
				},
			},
			expectedStatusCode: http.StatusOK,
			mockError:          nil,
		},
		{
			queries: map[string]string{
				"department": "department1",
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			queries: map[string]string{
				"date": "2021-01-01",
			},
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for i, testStep := range testSteps {
		t.Run("Test Get Workdays For Department And Date ", func(t *testing.T) {
			workdayRepository.On("GetWorkdaysForDepartmentAndDate").Return(testStep.mockValue, testStep.mockError).Return(testStep.mockValue, testStep.mockError)

			// get GIN context
			w := httptest.NewRecorder()
			c, err := mock.NewTestContextBuilder(w).
				WithQueries(testStep.queries).
				WithMethod("POST").Build()
			if err != nil {
				t.Errorf("Test Step %d: Error while building context: %s", i, err)
			}

			workdayService.GetWorkdaysForDepartmentAndDate(c)
			response := w.Result()

			if response.StatusCode != testStep.expectedStatusCode {
				t.Errorf("Test Step %d: Expected status code %d, got %d", i, testStep.expectedStatusCode, response.StatusCode)
			}

			if testStep.mockValue == nil {
				return
			}

			// compare response body
			var responseBody dto.APIResponse[[]dco.WorkdayResponse]
			if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
				t.Errorf("Test Step %d: Error while decoding response body: %s", i, err)
			}

			if len(responseBody.Data) != len(testStep.mockValue.([]dao.Workday)) {
				t.Errorf("Test Step %d: Expected %d workdays, got %d", i, len(testStep.mockValue.([]dao.Workday)), len(responseBody.Data))
			}

			for i, workday := range responseBody.Data {
				if workday.Department != testStep.mockValue.([]dao.Workday)[i].DepartmentID {
					t.Errorf("Test Step %d: Expected department %s, got %s", i, testStep.mockValue.([]dao.Workday)[i].DepartmentID, workday.Department)
				}

				if workday.Workplace != testStep.mockValue.([]dao.Workday)[i].WorkplaceID {
					t.Errorf("Test Step %d: Expected workplace %s, got %s", i, testStep.mockValue.([]dao.Workday)[i].WorkplaceID, workday.Workplace)
				}

				if workday.Timeslot != testStep.mockValue.([]dao.Workday)[i].TimeslotName {
					t.Errorf("Test Step %d: Expected timeslot %s, got %s", i, testStep.mockValue.([]dao.Workday)[i].TimeslotName, workday.Timeslot)
				}

				// check if Person is nil
				if testStep.mockValue.([]dao.Workday)[i].Person == nil {
					if workday.Person != nil {
						t.Errorf("Test Step %d: Expected person to be nil. But results were different.", i)
					}
				} else {
					if workday.Person == nil {
						t.Errorf("Test Step %d: Expected person to not be nil. But results were different.", i)
					}
				}
			}
		})
	}
}

func TestGetWorkday(t *testing.T) {
	workdayRepository := mock.NewWorkdayRepositoryMock()
	workdayService := WorkdayServiceImpl{
		WorkdayRepository: workdayRepository,
	}

	mockPerson := dao.Person{
		ID:           "person1",
		FirstName:    "first1",
		LastName:     "last1",
		Email:        "email1",
		WorkingHours: 8,
	}
	testSteps := []ServiceTestGET{
		{
			queries: map[string]string{
				"department": "department1",
				"workplace":  "workplace1",
				"timeslot":   "timeslot1",
				"date":       "2021-01-01",
			},
			mockValue: dao.Workday{
				DepartmentID: "department1",
				WorkplaceID:  "workplace1",
				TimeslotName: "timeslot1",
				Date:         "2021-01-01",
				StartTime:    "08:00:00",
				EndTime:      "16:00:00",
				Person:       nil,
			},
			expectedStatusCode: http.StatusOK,
			mockError:          nil,
		},
		{
			queries: map[string]string{
				"department": "department1",
				"workplace":  "workplace1",
				"timeslot":   "timeslot1",
				"date":       "2021-01-01",
			},
			mockValue:          nil,
			expectedStatusCode: http.StatusNotFound,
			mockError:          pkg.ErrNoRows,
		},
		{
			queries: map[string]string{
				"department": "department1",
				"workplace":  "workplace1",
				"timeslot":   "timeslot1",
				"date":       "2021-01-01",
			},
			mockValue: dao.Workday{
				DepartmentID: "department1",
				WorkplaceID:  "workplace1",
				TimeslotName: "timeslot1",
				Date:         "2021-01-01",
				StartTime:    "08:00:00",
				EndTime:      "16:00:00",
				Person:       &mockPerson,
			},
			expectedStatusCode: http.StatusOK,
			mockError:          nil,
		},
		{
			queries: map[string]string{
				"department": "department1",
				"workplace":  "workplace1",
				"timeslot":   "timeslot1",
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			queries: map[string]string{
				"department": "department1",
				"workplace":  "workplace1",
				"date":       "2021-01-01",
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			queries: map[string]string{
				"department": "department1",
				"timeslot":   "timeslot1",
				"date":       "2021-01-01",
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			queries: map[string]string{
				"workplace": "workplace1",
				"timeslot":  "timeslot1",
				"date":      "2021-01-01",
			},
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for i, testStep := range testSteps {
		t.Run("Test Get Workday", func(t *testing.T) {
			workdayRepository.On("GetWorkday").Return(testStep.mockValue, testStep.mockError)

			// get GIN context
			w := httptest.NewRecorder()
			c, err := mock.NewTestContextBuilder(w).
				WithQueries(testStep.queries).
				WithMethod("GET").Build()
			if err != nil {
				t.Errorf("Test Step %d: Error while building context: %s", i, err)
			}

			workdayService.GetWorkday(c)
			response := w.Result()

			if response.StatusCode != testStep.expectedStatusCode {
				t.Errorf("Test Step %d: Expected status code %d, got %d", i, testStep.expectedStatusCode, response.StatusCode)
			}

			if testStep.mockValue == nil {
				return
			}

			// compare response body
			var responseBody dto.APIResponse[dco.WorkdayResponse]
			if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
				t.Errorf("Test Step %d: Error while decoding response body: %s", i, err)
			}

			workday := responseBody.Data
			if workday.Department != testStep.mockValue.(dao.Workday).DepartmentID {
				t.Errorf("Test Step %d: Expected department %s, got %s", i, testStep.mockValue.(dao.Workday).DepartmentID, workday.Department)
			}

			if workday.Workplace != testStep.mockValue.(dao.Workday).WorkplaceID {
				t.Errorf("Test Step %d: Expected workplace %s, got %s", i, testStep.mockValue.(dao.Workday).WorkplaceID, workday.Workplace)
			}

			if workday.Timeslot != testStep.mockValue.(dao.Workday).TimeslotName {
				t.Errorf("Test Step %d: Expected timeslot %s, got %s", i, testStep.mockValue.(dao.Workday).TimeslotName, workday.Timeslot)
			}

			// check if Person is nil
			if testStep.mockValue.(dao.Workday).Person == nil {
				if workday.Person != nil {
					t.Errorf("Test Step %d: Expected person to be nil. But results were different.", i)
				}
			} else {
				if workday.Person == nil {
					t.Errorf("Test Step %d: Expected person to not be nil. But results were different.", i)
				}
			}
		})
	}
}

func TestAssignPersonToWorkday(t *testing.T) {
	workdayRepository := mock.NewWorkdayRepositoryMock()
	workdayService := WorkdayServiceImpl{
		WorkdayRepository: workdayRepository,
	}

	testSteps := []ServiceTestPOST{
		{
			// valid request
			mockRequestData: map[string]interface{}{
				"person_id":     "person1",
				"department_id": "department1",
				"workplace_id":  "workplace1",
				"timeslot_name": "timeslot1",
				"date":          "2021-01-01",
			},
			expectedStatusCode: http.StatusCreated,
			saveError:          nil,
		},
		{
			// missing person_id
			mockRequestData: map[string]interface{}{
				"department_id": "department1",
				"workplace_id":  "workplace1",
				"timeslot_name": "timeslot1",
				"date":          "2021-01-01",
			},
			expectedStatusCode: http.StatusBadRequest,
			saveError:          nil,
		},
		{
			// missing department_id
			mockRequestData: map[string]interface{}{
				"person_id":     "person1",
				"workplace_id":  "workplace1",
				"timeslot_name": "timeslot1",
				"date":          "2021-01-01",
			},
			expectedStatusCode: http.StatusBadRequest,
			saveError:          nil,
		},
		{
			// missing Date
			mockRequestData: map[string]interface{}{
				"person_id":     "person1",
				"department_id": "department1",
				"workplace_id":  "workplace1",
				"timeslot_name": "timeslot1",
			},
			expectedStatusCode: http.StatusBadRequest,
			saveError:          nil,
		},
		{
			mockRequestData: map[string]interface{}{
				"person_id":     "person1",
				"department_id": "department1",
				"workplace_id":  "workplace1",
				"timeslot_name": "timeslot1",
				"date":          "2021-01-01",
			},
			// repository error
			expectedStatusCode: 500,
			saveError:          errors.New("repository error"),
		},
	}

	for _, testStep := range testSteps {
		t.Run("Test Assign Person To Workday ", func(t *testing.T) {
			workdayRepository.On("AssignPersonToWorkday").Return(testStep.saveValue, testStep.saveError)

			// get GIN context
			w := httptest.NewRecorder()
			c, err := mock.NewTestContextBuilder(w).
				WithBody(testStep.mockRequestData).
				WithMethod("POST").Build()
			if err != nil {
				t.Errorf("Error while building context: %s", err)
			}

			workdayService.AssignPersonToWorkday(c)
			response := w.Result()

			if response.StatusCode != testStep.expectedStatusCode {
				t.Errorf("Expected status code %d, got %d", testStep.expectedStatusCode, response.StatusCode)
			}
		})
	}
}

func TestUnassignPersonFromWorkday(t *testing.T) {
	workdayRepository := mock.NewWorkdayRepositoryMock()
	workdayService := WorkdayServiceImpl{
		WorkdayRepository: workdayRepository,
	}

	testSteps := []ServiceTestPOST{
		{
			// valid request
			mockRequestData: map[string]interface{}{
				"person_id":     "person1",
				"department_id": "department1",
				"workplace_id":  "workplace1",
				"timeslot_name": "timeslot1",
				"date":          "2021-01-01",
			},
			expectedStatusCode: http.StatusOK,
			saveError:          nil,
		},
		{
			// missing person_id
			mockRequestData: map[string]interface{}{
				"department_id": "department1",
				"workplace_id":  "workplace1",
				"timeslot_name": "timeslot1",
				"date":          "2021-01-01",
			},
			expectedStatusCode: http.StatusBadRequest,
			saveError:          nil,
		},
		{
			// missing department_id
			mockRequestData: map[string]interface{}{
				"person_id":     "person1",
				"workplace_id":  "workplace1",
				"timeslot_name": "timeslot1",
				"date":          "2021-01-01",
			},
			expectedStatusCode: http.StatusBadRequest,
			saveError:          nil,
		},
		{
			// missing Date
			mockRequestData: map[string]interface{}{
				"person_id":     "person1",
				"department_id": "department1",
				"workplace_id":  "workplace1",
				"timeslot_name": "timeslot1",
			},
			expectedStatusCode: http.StatusBadRequest,
			saveError:          nil,
		},
		{
			mockRequestData: map[string]interface{}{
				"person_id":     "person1",
				"department_id": "department1",
				"workplace_id":  "workplace1",
				"timeslot_name": "timeslot1",
				"date":          "2021-01-01",
			},
			// repository error
			expectedStatusCode: 500,
			saveError:          errors.New("repository error"),
		},
	}

	for _, testStep := range testSteps {
		t.Run("Test Unassign Person From Workday ", func(t *testing.T) {
			workdayRepository.On("UnassignPersonFromWorkday").Return(testStep.saveValue, testStep.saveError)

			// get GIN context
			w := httptest.NewRecorder()
			c, err := mock.NewTestContextBuilder(w).
				WithBody(testStep.mockRequestData).
				WithMethod("POST").Build()
			if err != nil {
				t.Errorf("Error while building context: %s", err)
			}

			workdayService.UnassignPersonFromWorkday(c)
			response := w.Result()

			if response.StatusCode != testStep.expectedStatusCode {
				t.Errorf("Expected status code %d, got %d", testStep.expectedStatusCode, response.StatusCode)
			}
		})
	}
}
