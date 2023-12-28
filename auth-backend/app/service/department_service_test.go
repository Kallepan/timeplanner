package service

import (
	"auth-backend/app/domain/dao"
	"auth-backend/app/domain/dto"
	"auth-backend/app/mock"
	"database/sql"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TestUpdateDepartment(t *testing.T) {
	// Create mock object
	mockDepartmentRepository := mock.NewDepartmentRepositoryMock()

	// Create object to be tested
	departmentService := DepartmentServiceImpl{
		DepartmentRepository: &mockDepartmentRepository,
	}

	var testSteps = []ServiceTestPOST{
		{
			data: map[string]interface{}{
				"name": "Department 3",
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
				Name: "Department 3",
			},
			mockError: nil,
			expectedStatusCode: 200,
		},
		{
			data: map[string]interface{}{
				"name": "Department 1",
			},			
			mockValue: dao.Department{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name: "Department 1",
			},
			expectedValue: nil,
			mockError: sql.ErrNoRows,
			expectedStatusCode: 404,
		},
	}

	for _, testStep := range testSteps {
		// Set mock data
		mockDepartmentRepository.On("FindDepartmentById").Return(testStep.mockValue, testStep.mockError)
		mockDepartmentRepository.On("Save").Return(testStep.expectedValue, testStep.mockError)

		// get GIN context
		w := httptest.NewRecorder()
		ctx := mock.GetGinTestContextWithBody(w, "PUT", gin.Params{
			{Key: "departmentID", Value: testStep.mockValue.(dao.Department).ID.String()},
		}, testStep.data)
	
		// Test function
		departmentService.UpdateDepartment(ctx)

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

func TestDeleteDepartment(t *testing.T) {
	// Create mock object
	mockDepartmentRepository := mock.NewDepartmentRepositoryMock()

	// Create object to be tested
	departmentService := DepartmentServiceImpl{
		DepartmentRepository: &mockDepartmentRepository,
	}

	var testSteps = []ServiceTestGET{
		{
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
				Name: "Department 1",
			},
			mockError: nil,
			expectedStatusCode: 200,
		},
		{
			mockValue: dao.Department{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name: "Department 1",
			},
			expectedValue: nil,
			mockError: sql.ErrNoRows,
			expectedStatusCode: 200, // if data is not found, it will return 200 as well
		},
	}

	for _, testStep := range testSteps {
		// Set mock data
		mockDepartmentRepository.On("Delete").Return(testStep.expectedValue, testStep.mockError)

		// get GIN context
		w := httptest.NewRecorder()
		ctx := mock.GetGinTestContext(w, "DELETE", gin.Params{
			{Key: "departmentID", Value: testStep.mockValue.(dao.Department).ID.String()},
		})

		// Test function
		departmentService.DeleteDepartment(ctx)

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
	}
}

func TestAddDepartment(t *testing.T) {
	// Create mock object
	mockDepartmentRepository := mock.NewDepartmentRepositoryMock()

	// Create object to be tested
	departmentService := DepartmentServiceImpl{
		DepartmentRepository: &mockDepartmentRepository,
	}

	var testSteps = []ServiceTestPOST{
		{
			data: map[string]interface{}{
				"name": "Department 1",
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
				Name: "Department 1",
			},
			mockError: nil,
			expectedStatusCode: 201,
		},
	}

	for _, testStep := range testSteps {
		// Set mock data
		mockDepartmentRepository.On("Save").Return(testStep.mockValue, testStep.mockError)

		// get GIN context
		w := httptest.NewRecorder()
		ctx := mock.GetGinTestContextWithBody(w, "POST", gin.Params{}, testStep.data)
	
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

	var testSteps = []ServiceTestGET{
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
			expectedValue: []dao.Department{
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
			mockError: nil,
			expectedStatusCode: 200,
		},
		{
			mockValue: []dao.Department{},
			expectedValue: nil,
			mockError: nil,
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
		if testStep.expectedValue == nil {
			continue
		}

		// Assert response body
		var responseBody dto.APIResponse[[]dao.Department]
		if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
			t.Errorf("Error when decoding response body. Error: %s", err)
		}
		if responseBody.Data[0].ID != testStep.expectedValue.([]dao.Department)[0].ID {
			t.Errorf("Expected ID %s, but got %s", testStep.expectedValue.([]dao.Department)[0].ID.String(), responseBody.Data[0].ID.String())
		}
		if responseBody.Data[0].Name != testStep.expectedValue.([]dao.Department)[0].Name {
			t.Errorf("Expected Name %s, but got %s", testStep.expectedValue.([]dao.Department)[0].Name, responseBody.Data[0].Name)
		}
		if responseBody.Data[1].ID != testStep.expectedValue.([]dao.Department)[1].ID {
			t.Errorf("Expected ID %s, but got %s", testStep.expectedValue.([]dao.Department)[1].ID.String(), responseBody.Data[1].ID.String())
		}
		if responseBody.Data[1].Name != testStep.expectedValue.([]dao.Department)[1].Name {
			t.Errorf("Expected Name %s, but got %s", testStep.expectedValue.([]dao.Department)[1].Name, responseBody.Data[1].Name)
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

	var testSteps = []ServiceTestGET{
		{
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
				Name: "Department 1",
			},
			mockError: nil,
			expectedStatusCode: 200,
		},
		{
			mockValue: dao.Department{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name: "Department 1",
			},
			expectedValue: nil,
			mockError: sql.ErrNoRows,
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