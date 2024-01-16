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
				"departmentName": "test",
				"workplaceName":  "test",
			},
			mockError:          nil,
			expectedStatusCode: http.StatusOK,
		},
		{
			mockValue: dao.Workplace{
				Name: "test",
			},
			params: map[string]string{
				"departmentName": "test,",
			},
			mockError:          pkg.ErrNoRows,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			mockValue: dao.Workplace{
				Name: "test",
			},
			params: map[string]string{
				"workplaceName": "test,",
			},
			mockError:          pkg.ErrNoRows,
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for i, testStep := range testSteps {
		WorkplaceRepository.On("Delete").Return(testStep.mockValue, testStep.mockError)

		// get GIN context
		w := httptest.NewRecorder()
		c := mock.GetGinTestContext(w, "DELETE", testStep.ParamsToGinParams(), nil)

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
				"departmentName": "test",
				"workplaceName":  "oldName",
			},
		},
		{
			mockRequestData: map[string]interface{}{
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
		WorkplaceRepository.On("FindWorkplaceByName").Return(testStep.findValue, testStep.findError)
		WorkplaceRepository.On("Save").Return(testStep.saveValue, testStep.saveError)

		// get GIN context
		w := httptest.NewRecorder()
		c := mock.GetGinTestContext(w, "PUT", testStep.ParamsToGinParams(), testStep.mockRequestData)

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
				"departmentName": "test",
			},
		},
		{
			mockRequestData: map[string]interface{}{
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
				"departmentName": "test",
			},
		},
		{
			mockRequestData: map[string]interface{}{
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
				"departmentName": "test",
			},
		},
		{
			mockRequestData: map[string]interface{}{
				"name": "test",
			},
			findValue:          nil,
			saveValue:          nil,
			expectedStatusCode: 500,
			findError:          pkg.ErrNoRows,
			saveError:          pkg.ErrNoRows,
			params: map[string]string{
				"departmentName": "test",
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
			WorkplaceRepository.On("FindWorkplaceByName").Return(testStep.findValue, testStep.findError)
			WorkplaceRepository.On("Save").Return(testStep.saveValue, testStep.saveError)

			// get GIN context
			w := httptest.NewRecorder()
			c := mock.GetGinTestContext(w, "POST", testStep.ParamsToGinParams(), testStep.mockRequestData)

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
				"departmentName": "test",
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
				"departmentName": "test",
			},
		},
	}

	for i, testStep := range testSteps {
		t.Run("Test Get All Workplaces", func(t *testing.T) {
			WorkplaceRepository.On("FindAllWorkplaces").Return(testStep.mockValue, testStep.mockError)

			// get GIN context
			w := httptest.NewRecorder()
			c := mock.GetGinTestContext(w, "GET", testStep.ParamsToGinParams(), nil)

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

func TestGetWorkplaceByName(t *testing.T) {
	WorkplaceRepository := mock.NewWorkplaceRepositoryMock()
	workplaceService := WorkplaceServiceImpl{
		WorkplaceRepository: WorkplaceRepository,
	}

	testSteps := []ServiceTestGET{
		{
			mockValue: dao.Workplace{
				Name: "test",
			},
			expectedResponse: dao.Workplace{
				Name: "test",
			},
			expectedStatusCode: http.StatusOK,
			mockError:          nil,
			params: map[string]string{
				"workplaceName":  "test",
				"departmentName": "test",
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
				"departmentName": "test",
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
				"workplaceName": "test",
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
				"workplaceName":  "test",
				"departmentName": "test",
			},
		},
	}

	for i, testStep := range testSteps {
		t.Run("Test Get Workplace By Name", func(t *testing.T) {
			WorkplaceRepository.On("FindWorkplaceByName").Return(testStep.mockValue, testStep.mockError)

			// get GIN context
			w := httptest.NewRecorder()
			c := mock.GetGinTestContext(w, "GET", testStep.ParamsToGinParams(), nil)

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
