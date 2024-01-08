package service

import (
	"api-gateway/app/domain/dao"
	"api-gateway/app/domain/dto"
	"api-gateway/app/mock"
	"encoding/json"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func TestUpdateDepartment(t *testing.T) {
	// Create mock object
	mockDepartmentRepository := mock.NewDepartmentRepositoryMock()

	// Create object to be tested
	departmentService := DepartmentServiceImpl{
		DepartmentRepository: &mockDepartmentRepository,
	}

	testSteps := []ServiceTestPUT{
		{
			mockRequestData: map[string]interface{}{
				"name": "Department 2",
			},
			mockValue: dao.Department{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name: "Department 1",
			},
			expectedValue: dao.Department{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name: "Department 2",
			},
			mockError:          nil,
			expectedStatusCode: 200,
			queryParams: map[string]string{
				"departmentID": "00000000-0000-0000-0000-000000000001",
			},
		},
		{
			mockRequestData: map[string]interface{}{
				"name": "Department 1",
			},
			mockValue:          nil,
			expectedValue:      nil,
			mockError:          gorm.ErrRecordNotFound,
			expectedStatusCode: 404,
			queryParams: map[string]string{
				"departmentID": "00000000-0000-0000-0000-000000000001",
			},
		},
	}

	for i, testStep := range testSteps {
		// Set mock data
		mockDepartmentRepository.On("FindDepartmentById").Return(testStep.mockValue, testStep.mockError)
		mockDepartmentRepository.On("Save").Return(testStep.expectedValue, testStep.mockError)

		// get GIN context
		w := httptest.NewRecorder()
		ctx := mock.GetGinTestContextWithBody(w, "PUT", testStep.QueryParamsToGinParams(), testStep.mockRequestData)

		// Test function
		departmentService.UpdateDepartment(ctx)

		// Assert
		// Get response from GIN context
		response := w.Result()

		// Assert status code
		if response.StatusCode != testStep.expectedStatusCode {
			t.Errorf("Step: %d, Expected status code %d, but got %d", i, testStep.expectedStatusCode, response.StatusCode)
		}

		// continue the loop if expectedData is nil
		if testStep.expectedValue == nil {
			continue
		}

		// Assert response body
		var responseBody dto.APIResponse[dao.Department]
		if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
			t.Errorf("Error when decoding response body. Error: %s", err)
		}
		if responseBody.Data.ID != testStep.expectedValue.(dao.Department).ID {
			t.Errorf("Expected ID %s, but got %s", testStep.expectedValue.(dao.Department).ID.String(), responseBody.Data.ID.String())
		}
		if responseBody.Data.Name != testStep.expectedValue.(dao.Department).Name {
			t.Errorf("Expected Name %s, but got %s", testStep.expectedValue.(dao.Department).Name, responseBody.Data.Name)
		}
	}
}

func TestDeleteDepartment(t *testing.T) {
	// Create mock object
	mockDepartmentRepository := mock.NewDepartmentRepositoryMock()

	// Create object to be tested
	departmentService := DepartmentServiceImpl{
		DepartmentRepository: &mockDepartmentRepository,
	}

	testSteps := []ServiceTestDELETE{
		{
			mockValue: dao.Department{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name: "Department 1",
			},
			mockError:          nil,
			expectedStatusCode: 200,
			queryParams: map[string]string{
				"departmentID": "00000000-0000-0000-0000-000000000001",
			},
		},
		{
			mockValue:          nil,
			mockError:          gorm.ErrRecordNotFound,
			expectedStatusCode: 404,
			queryParams: map[string]string{
				"departmentID": "00000000-0000-0000-0000-000000000001",
			},
		},
	}

	for i, testStep := range testSteps {
		// Set mock data
		mockDepartmentRepository.On("DeleteDepartmentById").Return(testStep.mockValue, testStep.mockError)

		// get GIN context
		w := httptest.NewRecorder()
		ctx := mock.GetGinTestContext(w, "DELETE", testStep.QueryParamsToGinParams())

		// Test function
		departmentService.DeleteDepartment(ctx)

		// Assert
		// Get response from GIN context
		response := w.Result()

		// Assert status code
		if response.StatusCode != testStep.expectedStatusCode {
			t.Errorf("Step: %d. Expected status code %d, but got %d", i, testStep.expectedStatusCode, response.StatusCode)
		}
	}
}

func TestAddDepartment(t *testing.T) {
	// Create mock object
	mockDepartmentRepository := mock.NewDepartmentRepositoryMock()

	// Create object to be tested
	departmentService := DepartmentServiceImpl{
		DepartmentRepository: &mockDepartmentRepository,
	}

	testSteps := []ServiceTestPOST{
		{
			mockRequestData: map[string]interface{}{
				"name": "Department 1",
			},
			expectedValue: dao.Department{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name: "Department 1",
			},
			mockError:          nil,
			expectedStatusCode: 201,
		},
		{
			mockRequestData:    map[string]interface{}{},
			expectedValue:      nil,
			mockError:          nil,
			expectedStatusCode: 400,
		},
	}

	for _, testStep := range testSteps {
		// Set mock data
		mockDepartmentRepository.On("Save").Return(testStep.expectedValue, testStep.mockError)

		// get GIN context
		w := httptest.NewRecorder()
		ctx := mock.GetGinTestContextWithBody(w, "POST", gin.Params{}, testStep.mockRequestData)

		// Test function
		departmentService.AddDepartment(ctx)

		// Assert
		// Get response from GIN context
		response := w.Result()

		// Assert status code
		if response.StatusCode != testStep.expectedStatusCode {
			t.Errorf("Expected status code %d, but got %d", testStep.expectedStatusCode, response.StatusCode)
		}

		// continue the loop if expectedData is nil
		if testStep.expectedValue == nil {
			continue
		}

		// Assert response body
		var responseBody dto.APIResponse[dao.Department]
		if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
			t.Errorf("Error when decoding response body. Error: %s", err)
		}
		if responseBody.Data.ID != testStep.expectedValue.(dao.Department).ID {
			t.Errorf("Expected ID %s, but got %s", testStep.expectedValue.(dao.Department).ID.String(), responseBody.Data.ID.String())
		}
		if responseBody.Data.Name != testStep.expectedValue.(dao.Department).Name {
			t.Errorf("Expected Name %s, but got %s", testStep.expectedValue.(dao.Department).Name, responseBody.Data.Name)
		}
	}
}

func TestGetAllDepartments(t *testing.T) {
	// Create mock object
	mockDepartmentRepository := mock.NewDepartmentRepositoryMock()

	// Create object to be tested
	departmentService := DepartmentServiceImpl{
		DepartmentRepository: &mockDepartmentRepository,
	}

	testSteps := []ServiceTestGET{
		{
			mockValue: []dao.Department{
				{
					BaseModel: dao.BaseModel{
						ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					},
					Name: "Department 1",
				},
				{
					BaseModel: dao.BaseModel{
						ID: uuid.MustParse("00000000-0000-0000-0000-000000000002"),
					},
					Name: "Department 2",
				},
			},
			expectedResponse: []dao.Department{
				{
					BaseModel: dao.BaseModel{
						ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					},
					Name: "Department 1",
				},
				{
					BaseModel: dao.BaseModel{
						ID: uuid.MustParse("00000000-0000-0000-0000-000000000002"),
					},
					Name: "Department 2",
				},
			},
			mockError:          nil,
			expectedStatusCode: 200,
		},
		{
			mockValue:          []dao.Department{},
			expectedResponse:   nil,
			mockError:          nil,
			expectedStatusCode: 200,
		},
	}

	for _, testStep := range testSteps {
		// Set mock data
		mockDepartmentRepository.On("FindAllDepartments").Return(testStep.mockValue, testStep.mockError)

		// get GIN context
		w := httptest.NewRecorder()
		ctx := mock.GetGinTestContext(w, "GET", gin.Params{})

		// Test function
		departmentService.GetAllDepartments(ctx)

		// Assert
		// Get response from GIN context
		response := w.Result()

		// Assert status code
		if response.StatusCode != testStep.expectedStatusCode {
			t.Errorf("Expected status code %d, but got %d", testStep.expectedStatusCode, response.StatusCode)
		}

		// continue the loop if expectedData is nil
		if testStep.expectedResponse == nil {
			continue
		}

		// Assert response body
		var responseBody dto.APIResponse[[]dao.Department]
		if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
			t.Errorf("Error when decoding response body. Error: %s", err)
		}

		// compare body
		if !reflect.DeepEqual(responseBody.Data, testStep.expectedResponse) {
			t.Errorf("Expected response body %v, but got %v", testStep.expectedResponse, responseBody.Data)
		}
	}
}

func TestGetDepartmentById(t *testing.T) {
	// Create mock object
	mockDepartmentRepository := mock.NewDepartmentRepositoryMock()

	// Create object to be tested
	departmentService := DepartmentServiceImpl{
		DepartmentRepository: &mockDepartmentRepository,
	}

	testSteps := []ServiceTestGET{
		{
			mockValue: dao.Department{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name: "Department 1",
			},
			expectedResponse: dao.Department{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name: "Department 1",
			},
			mockError:          nil,
			expectedStatusCode: 200,
		},
		{
			mockValue: dao.Department{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name: "Department 1",
			},
			expectedResponse:   nil,
			mockError:          gorm.ErrRecordNotFound,
			expectedStatusCode: 404,
		},
	}

	for _, testStep := range testSteps {
		// Set mock data
		mockDepartmentRepository.On("FindDepartmentById").Return(testStep.mockValue, testStep.mockError)

		// get GIN context
		w := httptest.NewRecorder()
		ctx := mock.GetGinTestContext(w, "GET", gin.Params{
			{Key: "departmentID", Value: testStep.mockValue.(dao.Department).ID.String()},
		})

		// Test function
		departmentService.GetDepartmentById(ctx)

		// Assert
		// Get response from GIN context
		response := w.Result()

		// Assert status code
		if response.StatusCode != testStep.expectedStatusCode {
			t.Errorf("Expected status code %d, but got %d", testStep.expectedStatusCode, response.StatusCode)
		}

		// continue the loop if expectedData is nil
		if testStep.expectedResponse == nil {
			continue
		}

		// Assert response body
		var responseBody dto.APIResponse[dao.Department]
		if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
			t.Errorf("Error when decoding response body. Error: %s", err)
		}

		// compare body
		if !reflect.DeepEqual(responseBody.Data, testStep.expectedResponse) {
			t.Errorf("Expected response body %v, but got %v", testStep.expectedResponse, responseBody.Data)
		}
	}
}
