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

	"github.com/gin-gonic/gin"
)

func TestDeleteDepartment(t *testing.T) {
	departmentMockRepo := mock.NewDepartmentRepositoryMock()
	departmentService := DepartmentServiceImpl{
		DepartmentRepository: departmentMockRepo,
	}

	testSteps := []ServiceTestDELETE{
		{
			mockValue: dao.Department{
				Name: "test",
			},
			queryParams: map[string]string{
				"departmentName": "test",
			},
			mockError:          nil,
			expectedStatusCode: 200,
		},
		{
			mockValue: dao.Department{
				Name: "test",
			},
			queryParams: map[string]string{
				"departmentName": "test",
			},
			mockError:          pkg.ErrNoRows,
			expectedStatusCode: 200,
		},
	}

	for i, testStep := range testSteps {
		departmentMockRepo.On("Delete").Return(testStep.mockValue, testStep.mockError)

		// get GIN context
		w := httptest.NewRecorder()
		c := mock.GetGinTestContext(w, "DELETE", gin.Params{})

		departmentService.DeleteDepartment(c)
		response := w.Result()

		if response.StatusCode != testStep.expectedStatusCode {
			t.Errorf("Test Step %d: Expected status code %d, got %d", i, testStep.expectedStatusCode, response.StatusCode)
		}
	}

}

func TestUpdateDepartment(t *testing.T) {
	departmentMockRepo := mock.NewDepartmentRepositoryMock()
	departmentService := DepartmentServiceImpl{
		DepartmentRepository: departmentMockRepo,
	}

	testSteps := []ServiceTestPUT{
		{
			mockRequestData: map[string]interface{}{
				"name": "newName",
			},
			findValue: dao.Department{
				Name: "oldName",
			},
			saveValue: dao.Department{
				Name: "newName",
			},
			expectedStatusCode: 200,
			findError:          nil,
			saveError:          nil,
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
			findValue: dao.Department{
				Name: "oldName",
			},
			saveValue: dao.Department{
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
		departmentMockRepo.On("FindDepartmentByName").Return(testStep.findValue, testStep.findError)
		departmentMockRepo.On("Save").Return(testStep.saveValue, testStep.saveError)

		// get GIN context
		w := httptest.NewRecorder()
		c := mock.GetGinTestContextWithBody(w, "PUT", gin.Params{}, testStep.mockRequestData)

		departmentService.UpdateDepartment(c)
		response := w.Result()

		if response.StatusCode != testStep.expectedStatusCode {
			t.Errorf("Test Step %d: Expected status code %d, got %d", i, testStep.expectedStatusCode, response.StatusCode)
		}

		if response.StatusCode != 201 {
			return
		}

		var responseBody dto.APIResponse[dco.DepartmentResponse]
		if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
			t.Errorf("Test Step %d: Error when decoding response body", i)
		}

		if responseBody.Data.Name != testStep.saveValue.(dao.Department).Name {
			t.Errorf("Test Step %d: Expected response body %v, got %v", i, testStep.saveValue, responseBody.Data)
		}
	}
}

func TestAddDepartment(t *testing.T) {
	departmentMockRepo := mock.NewDepartmentRepositoryMock()
	departmentService := DepartmentServiceImpl{
		DepartmentRepository: departmentMockRepo,
	}

	testSteps := []ServiceTestPOST{
		{
			mockRequestData: map[string]interface{}{
				"name": "test",
			},
			findValue: nil,
			saveValue: dao.Department{
				Name: "test",
			},
			expectedStatusCode: 201,
			findError:          pkg.ErrNoRows,
			saveError:          nil,
		},
		{
			// conflict
			mockRequestData: map[string]interface{}{
				"name": "test",
			},
			findValue: dao.Department{
				Name: "test",
			},
			saveValue:          nil,
			expectedStatusCode: 409,
			findError:          nil,
			saveError:          nil,
		},
		{
			mockRequestData:    map[string]interface{}{},
			findValue:          nil,
			saveValue:          nil,
			expectedStatusCode: 400,
			findError:          pkg.ErrNoRows,
			saveError:          nil,
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
		t.Run("Test Add Department", func(t *testing.T) {
			departmentMockRepo.On("FindDepartmentByName").Return(testStep.findValue, testStep.findError)
			departmentMockRepo.On("Save").Return(testStep.saveValue, testStep.saveError)

			// get GIN context
			w := httptest.NewRecorder()
			c := mock.GetGinTestContextWithBody(w, "POST", gin.Params{}, testStep.mockRequestData)

			departmentService.AddDepartment(c)
			response := w.Result()

			if response.StatusCode != testStep.expectedStatusCode {
				t.Errorf("Test Step %d: Expected status code %d, got %d", i, testStep.expectedStatusCode, response.StatusCode)
			}

			if response.StatusCode != 201 {
				return
			}

			var responseBody dto.APIResponse[dco.DepartmentResponse]
			if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
				t.Errorf("Test Step %d: Error when decoding response body", i)
			}

			if responseBody.Data.Name != testStep.saveValue.(dao.Department).Name {
				t.Errorf("Test Step %d: Expected response body %v, got %v", i, testStep.saveValue, responseBody.Data)
			}
		})
	}
}

func TestGetAllDepartments(t *testing.T) {
	departmentMockRepo := mock.NewDepartmentRepositoryMock()
	departmentService := DepartmentServiceImpl{
		DepartmentRepository: departmentMockRepo,
	}

	testSteps := []ServiceTestGET{
		{
			mockValue: []dao.Department{
				{
					Name: "test",
				},
			},
			expectedResponse: []dao.Department{
				{
					Name: "test",
				},
			},
			expectedStatusCode: 200,
			mockError:          nil,
		},
		{
			mockValue:          nil,
			expectedResponse:   nil,
			expectedStatusCode: 200,
			mockError:          nil,
		},
		{
			mockValue:          nil,
			expectedResponse:   nil,
			expectedStatusCode: 500,
			mockError:          pkg.ErrNoRows,
		},
	}

	for i, testStep := range testSteps {
		t.Run("Test Get All Departments", func(t *testing.T) {
			departmentMockRepo.On("FindAllDepartments").Return(testStep.mockValue, testStep.mockError)

			// get GIN context
			w := httptest.NewRecorder()
			c := mock.GetGinTestContext(w, "GET", gin.Params{})

			departmentService.GetAllDepartments(c)
			response := w.Result()

			if response.StatusCode != testStep.expectedStatusCode {
				t.Errorf("Test Step %d: Expected status code %d, got %d", i, testStep.expectedStatusCode, response.StatusCode)
			}

			if response.StatusCode != 200 {
				return
			}

			var responseBody dto.APIResponse[[]dco.DepartmentResponse]
			if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
				t.Errorf("Test Step %d: Error when decoding response body", i)
			}

			for i, department := range responseBody.Data {
				if department.Name != testStep.expectedResponse.([]dao.Department)[i].Name {
					t.Errorf("Test Step %d: Expected response body %v, got %v", i, testStep.expectedResponse, responseBody.Data)
				}
			}
		})
	}
}

func TestGetDepartmentByName(t *testing.T) {
	departmentMockRepo := mock.NewDepartmentRepositoryMock()
	departmentService := DepartmentServiceImpl{
		DepartmentRepository: departmentMockRepo,
	}

	testSteps := []ServiceTestGET{
		{
			mockValue: dao.Department{
				Name: "test",
			},
			expectedResponse: dao.Department{
				Name: "test",
			},
			expectedStatusCode: 200,
			mockError:          nil,
		},
		{
			mockValue:          nil,
			expectedResponse:   nil,
			expectedStatusCode: 404,
			mockError:          pkg.ErrNoRows,
		},
	}

	for i, testStep := range testSteps {
		t.Run("Test Get Department By Name", func(t *testing.T) {
			departmentMockRepo.On("FindDepartmentByName").Return(testStep.mockValue, testStep.mockError)

			// get GIN context
			w := httptest.NewRecorder()
			c := mock.GetGinTestContext(w, "GET", gin.Params{
				{Key: "departmentName", Value: "test"},
			})

			departmentService.GetDepartmentByName(c)
			response := w.Result()

			if response.StatusCode != testStep.expectedStatusCode {
				t.Errorf("Test Step %d: Expected status code %d, got %d", i, testStep.expectedStatusCode, response.StatusCode)
			}

			if response.StatusCode != 200 {
				return
			}

			var responseBody dto.APIResponse[dco.DepartmentResponse]
			if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
				t.Errorf("Test Step %d: Error when decoding response body", i)
			}

			if responseBody.Data.Name != testStep.expectedResponse.(dao.Department).Name {
				t.Errorf("Test Step %d: Expected response body %v, got %v", i, testStep.expectedResponse, responseBody.Data)
			}
		})
	}
}
