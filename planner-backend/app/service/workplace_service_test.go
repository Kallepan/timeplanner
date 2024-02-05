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

func TestDeleteWorkplace(t *testing.T) {
	WorkplaceRepository := mock.NewWorkplaceRepositoryMock()
	workplaceService := WorkplaceServiceImpl{
		WorkplaceRepository: WorkplaceRepository,
	}

	testSteps := []ServiceTestDELETE{
		{
			mockValue: dao.Workplace{
				Name: "test",
			},
			params: map[string]string{
				"departmentID": "test",
				"workplaceID":  "test",
			},
			mockError:          nil,
			expectedStatusCode: http.StatusOK,
		},
		{
			mockValue: dao.Workplace{
				Name: "test",
			},
			params: map[string]string{
				"departmentID": "test,",
			},
			mockError:          pkg.ErrNoRows,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			mockValue: dao.Workplace{
				Name: "test",
			},
			params: map[string]string{
				"workplaceID": "test,",
			},
			mockError:          pkg.ErrNoRows,
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for i, testStep := range testSteps {
		WorkplaceRepository.On("Delete").Return(testStep.mockValue, testStep.mockError)

		// get GIN context
		w := httptest.NewRecorder()
		c, err := mock.NewTestContextBuilder(w).WithMethod("DELETE").WithMapParams(testStep.params).Build()
		if err != nil {
			t.Errorf("Test Step %d: Error when building context", i)
		}

		workplaceService.DeleteWorkplace(c)
		response := w.Result()

		if response.StatusCode != testStep.expectedStatusCode {
			t.Errorf("Test Step %d: Expected status code %d, got %d", i, testStep.expectedStatusCode, response.StatusCode)
		}
	}

}

func TestUpdateWorkplace(t *testing.T) {
	WorkplaceRepository := mock.NewWorkplaceRepositoryMock()
	workplaceService := WorkplaceServiceImpl{
		WorkplaceRepository: WorkplaceRepository,
	}

	testSteps := []ServiceTestPUT{
		{
			mockRequestData: map[string]interface{}{
				"id":   "test_id",
				"name": "newName",
			},
			findValue: dao.Workplace{
				Name: "oldName",
			},
			saveValue: dao.Workplace{
				Name: "newName",
			},
			expectedStatusCode: http.StatusOK,
			findError:          nil,
			saveError:          nil,
			params: map[string]string{
				"departmentID": "test",
				"workplaceID":  "oldName",
			},
		},
		{
			mockRequestData: map[string]interface{}{
				"id":   "test_id",
				"name": "newName",
			},
			findValue:          nil,
			saveValue:          nil,
			expectedStatusCode: http.StatusNotFound,
			findError:          pkg.ErrNoRows,
			saveError:          nil,
		},
		{
			mockRequestData: map[string]interface{}{
				"id":   "test_id",
				"name": "newName",
			},
			findValue: dao.Workplace{
				Name: "oldName",
			},
			saveValue: dao.Workplace{
				Name: "newName",
			},
			expectedStatusCode: 500,
			findError:          nil,
			saveError:          errors.New("Save error"),
		},
		{
			mockRequestData:    map[string]interface{}{},
			findValue:          nil,
			saveValue:          nil,
			expectedStatusCode: http.StatusBadRequest,
			findError:          nil,
			saveError:          nil,
		},
	}

	for i, testStep := range testSteps {
		WorkplaceRepository.On("FindWorkplaceByID").Return(testStep.findValue, testStep.findError)
		WorkplaceRepository.On("Save").Return(testStep.saveValue, testStep.saveError)

		// get GIN context
		w := httptest.NewRecorder()
		c, err := mock.NewTestContextBuilder(w).WithMethod("PUT").WithBody(testStep.mockRequestData).WithMapParams(testStep.params).Build()
		if err != nil {
			t.Errorf("Test Step %d: Error when building context", i)
		}

		workplaceService.UpdateWorkplace(c)
		response := w.Result()

		if response.StatusCode != testStep.expectedStatusCode {
			t.Errorf("Test Step %d: Expected status code %d, got %d", i, testStep.expectedStatusCode, response.StatusCode)
		}

		if response.StatusCode != 201 {
			return
		}

		var responseBody dto.APIResponse[dco.WorkplaceResponse]
		if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
			t.Errorf("Test Step %d: Error when decoding response body", i)
		}

		if responseBody.Data.Name != testStep.saveValue.(dao.Workplace).Name {
			t.Errorf("Test Step %d: Expected response body %v, got %v", i, testStep.saveValue, responseBody.Data)
		}
	}
}

func TestAddWorkplace(t *testing.T) {
	WorkplaceRepository := mock.NewWorkplaceRepositoryMock()
	workplaceService := WorkplaceServiceImpl{
		WorkplaceRepository: WorkplaceRepository,
	}

	testSteps := []ServiceTestPOST{
		{
			mockRequestData: map[string]interface{}{
				"id":   "test_id",
				"name": "test",
			},
			findValue: nil,
			saveValue: dao.Workplace{
				Name: "test",
			},
			expectedStatusCode: http.StatusCreated,
			findError:          pkg.ErrNoRows,
			saveError:          nil,
			params: map[string]string{
				"departmentID": "test",
			},
		},
		{
			mockRequestData: map[string]interface{}{
				"id":   "test_id",
				"name": "test",
			},
			findValue: dao.Workplace{
				Name: "test",
			},
			saveValue:          nil,
			expectedStatusCode: http.StatusConflict,
			findError:          nil,
			saveError:          nil,
			params: map[string]string{
				"departmentID": "test",
			},
		},
		{
			mockRequestData: map[string]interface{}{
				"id":   "test_id",
				"name": "test",
			},
			findValue: nil,
			saveValue: dao.Workplace{
				Name: "test",
			},
			expectedStatusCode: http.StatusBadRequest,
			findError:          pkg.ErrNoRows,
			saveError:          nil,
		},
		{
			mockRequestData:    map[string]interface{}{},
			findValue:          nil,
			saveValue:          nil,
			expectedStatusCode: http.StatusBadRequest,
			findError:          pkg.ErrNoRows,
			saveError:          nil,
			params: map[string]string{
				"departmentID": "test",
			},
		},
		{
			mockRequestData: map[string]interface{}{
				"id":   "test_id",
				"name": "test",
			},
			findValue:          dao.Workplace{},
			saveValue:          nil,
			expectedStatusCode: 500,
			findError:          pkg.ErrNoRows,
			saveError:          errors.New("Save error"),
			params: map[string]string{
				"departmentID": "test",
			},
		},
		{
			mockRequestData:    map[string]interface{}{},
			findValue:          nil,
			saveValue:          nil,
			expectedStatusCode: http.StatusBadRequest,
			findError:          nil,
			saveError:          pkg.ErrNoRows,
		},
	}

	for i, testStep := range testSteps {
		t.Run("Test Add Workplace", func(t *testing.T) {
			WorkplaceRepository.On("FindWorkplaceByID").Return(testStep.findValue, testStep.findError)
			WorkplaceRepository.On("Save").Return(testStep.saveValue, testStep.saveError)

			// get GIN context
			w := httptest.NewRecorder()
			c, err := mock.NewTestContextBuilder(w).WithMethod("POST").WithBody(testStep.mockRequestData).WithMapParams(testStep.params).Build()
			if err != nil {
				t.Errorf("Test Step %d: Error when building context", i)
			}

			workplaceService.AddWorkplace(c)
			response := w.Result()

			if response.StatusCode != testStep.expectedStatusCode {
				t.Errorf("Test Step %d: Expected status code %d, got %d", i, testStep.expectedStatusCode, response.StatusCode)
			}

			if response.StatusCode != 201 {
				return
			}

			var responseBody dto.APIResponse[dco.WorkplaceResponse]
			if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
				t.Errorf("Test Step %d: Error when decoding response body", i)
			}

			if responseBody.Data.Name != testStep.saveValue.(dao.Workplace).Name {
				t.Errorf("Test Step %d: Expected response body %v, got %v", i, testStep.saveValue, responseBody.Data)
			}
		})
	}
}

func TestGetAllWorkplaces(t *testing.T) {
	WorkplaceRepository := mock.NewWorkplaceRepositoryMock()
	workplaceService := WorkplaceServiceImpl{
		WorkplaceRepository: WorkplaceRepository,
	}

	testSteps := []ServiceTestGET{
		{
			mockValue: []dao.Workplace{
				{
					Name: "test",
				},
			},
			expectedResponse: []dco.WorkplaceResponse{
				{
					Name: "test",
				},
			},
			expectedStatusCode: http.StatusOK,
			mockError:          nil,
			params: map[string]string{
				"departmentID": "test",
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
			expectedStatusCode: http.StatusOK,
			mockError:          pkg.ErrNoRows,
			params: map[string]string{
				"departmentID": "test",
			},
		},
	}

	for i, testStep := range testSteps {
		t.Run("Test Get All Workplaces", func(t *testing.T) {
			WorkplaceRepository.On("FindAllWorkplaces").Return(testStep.mockValue, testStep.mockError)

			// get GIN context
			w := httptest.NewRecorder()
			c, err := mock.NewTestContextBuilder(w).WithMethod("GET").WithMapParams(testStep.params).Build()
			if err != nil {
				t.Errorf("Test Step %d: Error when building context", i)
			}

			workplaceService.GetAllWorkplaces(c)
			response := w.Result()

			if response.StatusCode != testStep.expectedStatusCode {
				t.Errorf("Test Step %d: Expected status code %d, got %d", i, testStep.expectedStatusCode, response.StatusCode)
			}

			if response.StatusCode != http.StatusOK {
				return
			}

			var responseBody dto.APIResponse[[]dco.WorkplaceResponse]
			if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
				t.Errorf("Test Step %d: Error when decoding response body", i)
			}

			for i, department := range responseBody.Data {
				if department.Name != testStep.expectedResponse.([]dco.WorkplaceResponse)[i].Name {
					t.Errorf("Test Step %d: Expected response body %v, got %v", i, testStep.expectedResponse, responseBody.Data)
				}
			}
		})
	}
}

func TestGetWorkplaceByID(t *testing.T) {
	WorkplaceRepository := mock.NewWorkplaceRepositoryMock()
	workplaceService := WorkplaceServiceImpl{
		WorkplaceRepository: WorkplaceRepository,
	}

	testSteps := []ServiceTestGET{
		{
			mockValue: dao.Workplace{
				ID:   "test",
				Name: "test",
			},
			expectedResponse: dao.Workplace{
				ID:   "test",
				Name: "test",
			},
			expectedStatusCode: http.StatusOK,
			mockError:          nil,
			params: map[string]string{
				"workplaceID":  "test",
				"departmentID": "test",
			},
		},
		{
			mockValue: dao.Workplace{
				ID:   "test",
				Name: "test",
			},
			expectedResponse: dao.Workplace{
				ID:   "test",
				Name: "test",
			},
			expectedStatusCode: http.StatusBadRequest,
			mockError:          nil,
			params: map[string]string{
				"departmentID": "test",
			},
		},
		{
			mockValue: dao.Workplace{
				Name: "test",
			},
			expectedResponse: dao.Workplace{
				Name: "test",
			},
			expectedStatusCode: http.StatusBadRequest,
			mockError:          nil,
			params: map[string]string{
				"workplaceID": "test",
			},
		},
		{
			mockValue:          nil,
			expectedResponse:   nil,
			expectedStatusCode: http.StatusBadRequest,
			mockError:          nil,
			params:             map[string]string{},
		},
		{
			mockValue:          nil,
			expectedResponse:   nil,
			expectedStatusCode: http.StatusNotFound,
			mockError:          pkg.ErrNoRows,
			params: map[string]string{
				"workplaceID":  "test",
				"departmentID": "test",
			},
		},
	}

	for i, testStep := range testSteps {
		t.Run("Test Get Workplace By Name", func(t *testing.T) {
			WorkplaceRepository.On("FindWorkplaceByID").Return(testStep.mockValue, testStep.mockError)

			// get GIN context
			w := httptest.NewRecorder()
			c, err := mock.NewTestContextBuilder(w).WithMethod("GET").WithMapParams(testStep.params).Build()
			if err != nil {
				t.Errorf("Test Step %d: Error when building context", i)
			}

			workplaceService.GetWorkplaceByName(c)
			response := w.Result()

			if response.StatusCode != testStep.expectedStatusCode {
				t.Errorf("Test Step %d: Expected status code %d, got %d", i, testStep.expectedStatusCode, response.StatusCode)
			}

			if response.StatusCode != http.StatusOK {
				return
			}

			var responseBody dto.APIResponse[dco.WorkplaceResponse]
			if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
				t.Errorf("Test Step %d: Error when decoding response body", i)
			}

			if responseBody.Data.Name != testStep.expectedResponse.(dao.Workplace).Name {
				t.Errorf("Test Step %d: Expected response body %v, got %v", i, testStep.expectedResponse, responseBody.Data)
			}
		})
	}
}
