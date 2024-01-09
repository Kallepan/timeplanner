package service

import (
	"encoding/json"
	"errors"
	"net/http/httptest"
	"planner-backend/app/domain/dao"
	"planner-backend/app/domain/dco"
	"planner-backend/app/domain/dto"
	"planner-backend/app/mock"
	"planner-backend/app/pkg"
	"testing"
)

func TestDeleteWorkplace(t *testing.T) {
	workplaceRepository := mock.NewWorkplaceRepositoryMock()
	workplaceService := workplaceServiceImpl{
		workplaceRepository: workplaceRepository,
	}

	testSteps := []ServiceTestDELETE{
		{
			mockValue: dao.Workplace{
				Name: "test",
			},
			queryParams: map[string]string{
				"departmentName": "test",
				"workplaceName":  "test",
			},
			mockError:          nil,
			expectedStatusCode: 200,
		},
		{
			mockValue: dao.Workplace{
				Name: "test",
			},
			queryParams: map[string]string{
				"departmentName": "test,",
			},
			mockError:          pkg.ErrNoRows,
			expectedStatusCode: 400,
		},
		{
			mockValue: dao.Workplace{
				Name: "test",
			},
			queryParams: map[string]string{
				"workplaceName": "test,",
			},
			mockError:          pkg.ErrNoRows,
			expectedStatusCode: 400,
		},
	}

	for i, testStep := range testSteps {
		workplaceRepository.On("Delete").Return(testStep.mockValue, testStep.mockError)

		// get GIN context
		w := httptest.NewRecorder()
		c := mock.GetGinTestContext(w, "DELETE", testStep.QueryParamsToGinParams())

		workplaceService.DeleteWorkplace(c)
		response := w.Result()

		if response.StatusCode != testStep.expectedStatusCode {
			t.Errorf("Test Step %d: Expected status code %d, got %d", i, testStep.expectedStatusCode, response.StatusCode)
		}
	}

}

func TestUpdateWorkplace(t *testing.T) {
	workplaceRepository := mock.NewWorkplaceRepositoryMock()
	workplaceService := workplaceServiceImpl{
		workplaceRepository: workplaceRepository,
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
			expectedStatusCode: 200,
			findError:          nil,
			saveError:          nil,
			queryParams: map[string]string{
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
			expectedStatusCode: 404,
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
			expectedStatusCode: 400,
			findError:          nil,
			saveError:          nil,
		},
	}

	for i, testStep := range testSteps {
		workplaceRepository.On("FindWorkplaceByName").Return(testStep.findValue, testStep.findError)
		workplaceRepository.On("Save").Return(testStep.saveValue, testStep.saveError)

		// get GIN context
		w := httptest.NewRecorder()
		c := mock.GetGinTestContextWithBody(w, "PUT", testStep.QueryParamsToGinParams(), testStep.mockRequestData)

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
	workplaceRepository := mock.NewWorkplaceRepositoryMock()
	workplaceService := workplaceServiceImpl{
		workplaceRepository: workplaceRepository,
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
			expectedStatusCode: 201,
			findError:          pkg.ErrNoRows,
			saveError:          nil,
			queryParams: map[string]string{
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
			expectedStatusCode: 400,
			findError:          pkg.ErrNoRows,
			saveError:          nil,
		},
		{
			mockRequestData:    map[string]interface{}{},
			findValue:          nil,
			saveValue:          nil,
			expectedStatusCode: 400,
			findError:          pkg.ErrNoRows,
			saveError:          nil,
			queryParams: map[string]string{
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
			queryParams: map[string]string{
				"departmentName": "test",
			},
		},
		{
			mockRequestData:    map[string]interface{}{},
			findValue:          nil,
			saveValue:          nil,
			expectedStatusCode: 400,
			findError:          nil,
			saveError:          pkg.ErrNoRows,
		},
	}

	for i, testStep := range testSteps {
		t.Run("Test Add Workplace", func(t *testing.T) {
			workplaceRepository.On("FindWorkplaceByName").Return(testStep.findValue, testStep.findError)
			workplaceRepository.On("Save").Return(testStep.saveValue, testStep.saveError)

			// get GIN context
			w := httptest.NewRecorder()
			c := mock.GetGinTestContextWithBody(w, "POST", testStep.QueryParamsToGinParams(), testStep.mockRequestData)

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
	workplaceRepository := mock.NewWorkplaceRepositoryMock()
	workplaceService := workplaceServiceImpl{
		workplaceRepository: workplaceRepository,
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
			expectedStatusCode: 200,
			mockError:          nil,
			queryParams: map[string]string{
				"departmentName": "test",
			},
		},
		{
			mockValue:          nil,
			expectedResponse:   nil,
			expectedStatusCode: 400,
			mockError:          nil,
		},
		{
			mockValue:          nil,
			expectedResponse:   nil,
			expectedStatusCode: 200,
			mockError:          pkg.ErrNoRows,
			queryParams: map[string]string{
				"departmentName": "test",
			},
		},
	}

	for i, testStep := range testSteps {
		t.Run("Test Get All Workplaces", func(t *testing.T) {
			workplaceRepository.On("FindAllWorkplaces").Return(testStep.mockValue, testStep.mockError)

			// get GIN context
			w := httptest.NewRecorder()
			c := mock.GetGinTestContext(w, "GET", testStep.QueryParamsToGinParams())

			workplaceService.GetAllWorkplaces(c)
			response := w.Result()

			if response.StatusCode != testStep.expectedStatusCode {
				t.Errorf("Test Step %d: Expected status code %d, got %d", i, testStep.expectedStatusCode, response.StatusCode)
			}

			if response.StatusCode != 200 {
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
	workplaceRepository := mock.NewWorkplaceRepositoryMock()
	workplaceService := workplaceServiceImpl{
		workplaceRepository: workplaceRepository,
	}

	testSteps := []ServiceTestGET{
		{
			mockValue: dao.Workplace{
				Name: "test",
			},
			expectedResponse: dao.Workplace{
				Name: "test",
			},
			expectedStatusCode: 200,
			mockError:          nil,
			queryParams: map[string]string{
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
			expectedStatusCode: 400,
			mockError:          nil,
			queryParams: map[string]string{
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
			expectedStatusCode: 400,
			mockError:          nil,
			queryParams: map[string]string{
				"workplaceName": "test",
			},
		},
		{
			mockValue:          nil,
			expectedResponse:   nil,
			expectedStatusCode: 400,
			mockError:          nil,
			queryParams:        map[string]string{},
		},
		{
			mockValue:          nil,
			expectedResponse:   nil,
			expectedStatusCode: 404,
			mockError:          pkg.ErrNoRows,
			queryParams: map[string]string{
				"workplaceName":  "test",
				"departmentName": "test",
			},
		},
	}

	for i, testStep := range testSteps {
		t.Run("Test Get Workplace By Name", func(t *testing.T) {
			workplaceRepository.On("FindWorkplaceByName").Return(testStep.mockValue, testStep.mockError)

			// get GIN context
			w := httptest.NewRecorder()
			c := mock.GetGinTestContext(w, "GET", testStep.QueryParamsToGinParams())

			workplaceService.GetWorkplaceByName(c)
			response := w.Result()

			if response.StatusCode != testStep.expectedStatusCode {
				t.Errorf("Test Step %d: Expected status code %d, got %d", i, testStep.expectedStatusCode, response.StatusCode)
			}

			if response.StatusCode != 200 {
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
