package service

import (
	"api-gateway/app/domain/dao"
	"api-gateway/app/domain/dto"
	"api-gateway/app/mock"
	"encoding/json"
	"fmt"
	"net/http"
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
			expectedStatusCode: http.StatusOK,
			findValue: dao.Department{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name: "Department 1",
			},
			saveValue: dao.Department{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name: "Department 2",
			},
			findError: nil,
			saveError: nil,
			params: map[string]string{
				"departmentID": "00000000-0000-0000-0000-000000000001",
			},
		},
		{
			mockRequestData: map[string]interface{}{
				"name": "Department 1",
			},
			findValue:          nil,
			saveValue:          nil,
			findError:          gorm.ErrRecordNotFound,
			saveError:          nil,
			expectedStatusCode: http.StatusNotFound,
			params: map[string]string{
				"departmentID": "00000000-0000-0000-0000-000000000001",
			},
		},
	}

	for i, testStep := range testSteps {
		t.Run(fmt.Sprintf("Test step %d", i), func(t *testing.T) {
			// Set mock data
			mockDepartmentRepository.On("FindDepartmentById").Return(testStep.findValue, testStep.findError)
			mockDepartmentRepository.On("Save").Return(testStep.saveValue, testStep.saveError)

			// get GIN context
			w := httptest.NewRecorder()
			ctx := mock.GetGinTestContext(w, "PUT", testStep.ParamsToGinParams(), testStep.mockRequestData)

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
			if testStep.saveValue == nil {
				return
			}

			// Assert response body
			var responseBody dto.APIResponse[dao.Department]
			if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
				t.Errorf("Error when decoding response body. Error: %s", err)
			}
			if responseBody.Data.ID != testStep.saveValue.(dao.Department).ID {
				t.Errorf("Expected ID %s, but got %s", testStep.saveValue.(dao.Department).ID.String(), responseBody.Data.ID.String())
			}
			if responseBody.Data.Name != testStep.saveValue.(dao.Department).Name {
				t.Errorf("Expected Name %s, but got %s", testStep.saveValue.(dao.Department).Name, responseBody.Data.Name)
			}
		})
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
			expectedStatusCode: http.StatusOK,
			params: map[string]string{
				"departmentID": "00000000-0000-0000-0000-000000000001",
			},
		},
		{
			mockValue:          nil,
			mockError:          gorm.ErrRecordNotFound,
			expectedStatusCode: http.StatusNotFound,
			params: map[string]string{
				"departmentID": "00000000-0000-0000-0000-000000000001",
			},
		},
	}

	for i, testStep := range testSteps {
		t.Run(fmt.Sprintf("Test step %d", i), func(t *testing.T) {
			// Set mock data
			mockDepartmentRepository.On("DeleteDepartmentById").Return(testStep.mockValue, testStep.mockError)

			// get GIN context
			w := httptest.NewRecorder()
			ctx := mock.GetGinTestContext(w, "DELETE", testStep.ParamsToGinParams(), nil)

			// Test function
			departmentService.DeleteDepartment(ctx)

			// Assert
			// Get response from GIN context
			response := w.Result()

			// Assert status code
			if response.StatusCode != testStep.expectedStatusCode {
				t.Errorf("Step: %d. Expected status code %d, but got %d", i, testStep.expectedStatusCode, response.StatusCode)
			}
		})
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
			findValue: nil,
			saveValue: dao.Department{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name: "Department 1",
			},
			findError:          gorm.ErrRecordNotFound,
			saveError:          nil,
			expectedStatusCode: http.StatusCreated,
		},
		{
			mockRequestData:    map[string]interface{}{},
			findValue:          nil,
			saveValue:          nil,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			mockRequestData: map[string]interface{}{
				"name": "Department 1",
			},
			findValue: dao.Department{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name: "Department 1",
			},
			saveValue:          nil,
			findError:          nil,
			saveError:          nil,
			expectedStatusCode: http.StatusConflict,
		},
	}

	for i, testStep := range testSteps {
		t.Run(fmt.Sprintf("Test step %d", i), func(t *testing.T) {
			// Set mock data
			mockDepartmentRepository.On("FindDepartmentByName").Return(testStep.findValue, testStep.findError)
			mockDepartmentRepository.On("Save").Return(testStep.saveValue, testStep.saveError)

			// get GIN context
			w := httptest.NewRecorder()
			ctx := mock.GetGinTestContext(w, "POST", gin.Params{}, testStep.mockRequestData)

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
			if testStep.saveValue == nil {
				return
			}

			// Assert response body
			var responseBody dto.APIResponse[dao.Department]
			if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
				t.Errorf("Error when decoding response body. Error: %s", err)
			}
			if responseBody.Data.ID != testStep.saveValue.(dao.Department).ID {
				t.Errorf("Expected ID %s, but got %s", testStep.saveValue.(dao.Department).ID.String(), responseBody.Data.ID.String())
			}
			if responseBody.Data.Name != testStep.saveValue.(dao.Department).Name {
				t.Errorf("Expected Name %s, but got %s", testStep.saveValue.(dao.Department).Name, responseBody.Data.Name)
			}
		})
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
			expectedStatusCode: http.StatusOK,
		},
		{
			mockValue:          []dao.Department{},
			expectedResponse:   nil,
			mockError:          nil,
			expectedStatusCode: http.StatusOK,
		},
	}

	for i, testStep := range testSteps {
		t.Run(fmt.Sprintf("Test step %d", i), func(t *testing.T) {
			// Set mock data
			mockDepartmentRepository.On("FindAllDepartments").Return(testStep.mockValue, testStep.mockError)

			// get GIN context
			w := httptest.NewRecorder()
			ctx := mock.GetGinTestContext(w, "GET", gin.Params{}, nil)

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
				return
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
		})
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
			expectedStatusCode: http.StatusOK,
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
			expectedStatusCode: http.StatusNotFound,
		},
	}

	for i, testStep := range testSteps {
		t.Run(fmt.Sprintf("Test step %d", i), func(t *testing.T) {
			// Set mock data
			mockDepartmentRepository.On("FindDepartmentById").Return(testStep.mockValue, testStep.mockError)

			// get GIN context
			w := httptest.NewRecorder()
			ctx := mock.GetGinTestContext(w, "GET", gin.Params{
				{Key: "departmentID", Value: testStep.mockValue.(dao.Department).ID.String()},
			}, nil)

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
				return
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
		})
	}
}
