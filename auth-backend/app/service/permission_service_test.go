package service

import (
	"auth-backend/app/domain/dao"
	"auth-backend/app/domain/dto"
	"auth-backend/app/mock"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func TestUpdatePermission(t *testing.T) {
	// Create mock object
	permissionRepositoryMock := mock.NewPermissionRepositoryMock()
	permissionService := PermissionServiceImpl{
		PermissionRepository: &permissionRepositoryMock,
	}

	// Set mock data
	var testSteps = []ServiceTestPUT{
		{
			data: map[string]interface{}{
				"name":        "name_new",
				"description": "description_new",
			},
			mockValue: dao.Permission{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name:        "name_old",
				Description: new(string),
			},
			expectedValue: dao.Permission{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name:        "name_new",
				Description: new(string),
			},
			mockError:          nil,
			expectedStatusCode: 200,
		},
		{
			// update with nil description
			data: map[string]interface{}{
				"name":        "name_new",
				"description": nil,
			},
			mockValue: dao.Permission{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name:        "name_old",
				Description: new(string),
			},
			expectedValue: dao.Permission{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name:        "name_new",
				Description: nil,
			},
			mockError:          nil,
			expectedStatusCode: 200,
		},
		{
			// update with no description
			data: map[string]interface{}{
				"name": "name_new",
			},
			mockValue: dao.Permission{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name:        "name_old",
				Description: new(string),
			},
			expectedValue: dao.Permission{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name:        "name_new",
				Description: nil,
			},
			mockError:          nil,
			expectedStatusCode: 200,
		},
	}

	// define description as it is a pointer
	*testSteps[0].mockValue.(dao.Permission).Description = "description_old"
	*testSteps[0].expectedValue.(dao.Permission).Description = "description_new"

	// Run test
	for _, testStep := range testSteps {
		// Set mock data
		permissionRepositoryMock.On("FindPermissionById").Return(testStep.mockValue, testStep.mockError)
		permissionRepositoryMock.On("Save").Return(testStep.expectedValue, testStep.mockError)

		// get GIN context
		w := httptest.NewRecorder()
		c := mock.GetGinTestContextWithBody(w, "PUT", gin.Params{
			{Key: "permissionID", Value: testStep.mockValue.(dao.Permission).ID.String()},
		}, testStep.data)

		// Run function
		permissionService.UpdatePermission(c)

		// Check result
		response := w.Result()
		if response.StatusCode != testStep.expectedStatusCode {
			t.Errorf("Expected status code %d but got %d", testStep.expectedStatusCode, response.StatusCode)
		}

		if testStep.expectedValue == nil {
			continue
		}

		// Check mock data
		var responseBody dto.APIResponse[dao.Permission]
		if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
			t.Errorf("Error when unmarshalling response body: %s", err.Error())
		}
		if responseBody.Data.Name != testStep.expectedValue.(dao.Permission).Name {
			t.Errorf("Expected name %s but got %s", testStep.expectedValue.(dao.Permission).Name, responseBody.Data.Name)
		}
		expectedDescription := "nil"
		if testStep.expectedValue.(dao.Permission).Description != nil {
			expectedDescription = *testStep.expectedValue.(dao.Permission).Description
		}

		responseDescription := "nil"
		if responseBody.Data.Description != nil {
			responseDescription = *responseBody.Data.Description
		}

		if responseDescription != expectedDescription {
			t.Errorf("Expected description %s but got %s", expectedDescription, responseDescription)
		}
	}
}

func TestGetAllPermissions(t *testing.T) {
	// Create mock object
	permissionRepositoryMock := mock.NewPermissionRepositoryMock()
	permissionService := PermissionServiceImpl{
		PermissionRepository: &permissionRepositoryMock,
	}

	// Set mock data
	testSteps := []ServiceTestGET{
		{
			mockValue: []dao.Permission{
				{
					BaseModel: dao.BaseModel{
						ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					},
					Name:        "name_1",
					Description: new(string),
				},
				{
					BaseModel: dao.BaseModel{
						ID: uuid.MustParse("00000000-0000-0000-0000-000000000002"),
					},
					Name:        "name_2",
					Description: new(string),
				},
			},
			expectedValue: []dao.Permission{
				{
					BaseModel: dao.BaseModel{
						ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					},
					Name:        "name_1",
					Description: new(string),
				},
				{
					BaseModel: dao.BaseModel{
						ID: uuid.MustParse("00000000-0000-0000-0000-000000000002"),
					},
					Name:        "name_2",
					Description: new(string),
				},
			},
			mockError:          nil,
			expectedStatusCode: 200,
		},
		{
			mockValue:          []dao.Permission{},
			expectedValue:      nil,
			mockError:          nil,
			expectedStatusCode: 200,
		},
	}

	// define description as it is a pointer
	*testSteps[0].mockValue.([]dao.Permission)[0].Description = "description_1"
	*testSteps[0].mockValue.([]dao.Permission)[1].Description = "description_2"
	*testSteps[0].expectedValue.([]dao.Permission)[0].Description = "description_1"
	*testSteps[0].expectedValue.([]dao.Permission)[1].Description = "description_2"

	// Run test
	for _, testStep := range testSteps {
		// Set mock data
		permissionRepositoryMock.On("FindAllPermissions").Return(testStep.mockValue, testStep.mockError)

		// get GIN context
		w := httptest.NewRecorder()
		c := mock.GetGinTestContext(w, "GET", gin.Params{})

		// Run function
		permissionService.GetAllPermissions(c)

		// Check result
		response := w.Result()
		if response.StatusCode != testStep.expectedStatusCode {
			t.Errorf("Expected status code %d but got %d", testStep.expectedStatusCode, response.StatusCode)
		}

		if testStep.expectedValue == nil {
			continue
		}

		// Check mock data
		var responseBody dto.APIResponse[[]dao.Permission]
		if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
			t.Errorf("Error when unmarshalling response body: %s", err.Error())
		}

		for i, permission := range responseBody.Data {
			if permission.Name != testStep.expectedValue.([]dao.Permission)[i].Name {
				t.Errorf("Expected name %s but got %s", testStep.expectedValue.([]dao.Permission)[i].Name, permission.Name)
			}
			expectedDescription := "nil"
			if testStep.expectedValue.([]dao.Permission)[i].Description != nil {
				expectedDescription = *testStep.expectedValue.([]dao.Permission)[i].Description
			}

			responseDescription := "nil"
			if permission.Description != nil {
				responseDescription = *permission.Description
			}

			if responseDescription != expectedDescription {
				t.Errorf("Expected description %s but got %s", expectedDescription, responseDescription)
			}
		}
	}
}

func TestGetPermissionById(t *testing.T) {
	// Create mock object
	permissionRepositoryMock := mock.NewPermissionRepositoryMock()
	permissionService := PermissionServiceImpl{
		PermissionRepository: &permissionRepositoryMock,
	}

	// Set mock data
	testSteps := []ServiceTestGET{
		{

			mockValue: dao.Permission{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name:        "name_1",
				Description: new(string),
			},
			expectedValue: dao.Permission{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name:        "name_1",
				Description: new(string),
			},
			mockError:          nil,
			expectedStatusCode: 200,
		},
		{
			mockValue: dao.Permission{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name:        "name_1",
				Description: new(string),
			},
			expectedValue:      nil,
			mockError:          gorm.ErrRecordNotFound,
			expectedStatusCode: 404,
		},
	}

	// define description as it is a pointer
	*testSteps[0].mockValue.(dao.Permission).Description = "description_1"
	*testSteps[0].expectedValue.(dao.Permission).Description = "description_1"
	*testSteps[1].mockValue.(dao.Permission).Description = "description_1"

	// Run test
	for _, testStep := range testSteps {
		// Set mock data
		permissionRepositoryMock.On("FindPermissionById").Return(testStep.mockValue, testStep.mockError)

		// get GIN context
		w := httptest.NewRecorder()
		c := mock.GetGinTestContext(w, "GET", gin.Params{
			{Key: "permissionID", Value: testStep.mockValue.(dao.Permission).ID.String()},
		})

		// Run function
		permissionService.GetPermissionById(c)

		// Check result
		response := w.Result()
		if response.StatusCode != testStep.expectedStatusCode {
			t.Errorf("Expected status code %d but got %d", testStep.expectedStatusCode, response.StatusCode)
		}

		if testStep.expectedValue == nil {
			continue
		}

		// Check mock data
		var responseBody dto.APIResponse[dao.Permission]
		if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
			t.Errorf("Error when unmarshalling response body: %s", err.Error())

		}
		if responseBody.Data.Name != testStep.expectedValue.(dao.Permission).Name {
			t.Errorf("Expected name %s but got %s", testStep.expectedValue.(dao.Permission).Name, responseBody.Data.Name)
		}

		expectedDescription := "nil"
		if testStep.expectedValue.(dao.Permission).Description != nil {
			expectedDescription = *testStep.expectedValue.(dao.Permission).Description
		}

		responseDescription := "nil"
		if responseBody.Data.Description != nil {
			responseDescription = *responseBody.Data.Description
		}

		if responseDescription != expectedDescription {
			t.Errorf("Expected description %s but got %s", expectedDescription, responseDescription)
		}
	}
}

func TestDeletePermissionById(t *testing.T) {
	// Create mock object
	permissionRepositoryMock := mock.NewPermissionRepositoryMock()
	permissionService := PermissionServiceImpl{
		PermissionRepository: &permissionRepositoryMock,
	}

	// Set mock data
	testSteps := []ServiceTestGET{
		{
			mockValue: dao.Permission{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name:        "name_1",
				Description: new(string),
			},
			expectedValue: dao.Permission{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name:        "name_1",
				Description: new(string),
			},
			mockError:          nil,
			expectedStatusCode: 200,
		},
		{
			mockValue: dao.Permission{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name:        "name_1",
				Description: new(string),
			},
			expectedValue:      nil,
			mockError:          gorm.ErrRecordNotFound,
			expectedStatusCode: 404,
		},
	}

	// define description as it is a pointer
	*testSteps[0].mockValue.(dao.Permission).Description = "description_1"
	*testSteps[0].expectedValue.(dao.Permission).Description = "description_1"
	*testSteps[1].mockValue.(dao.Permission).Description = "description_1"

	// Run test
	for _, testStep := range testSteps {
		// Set mock data
		permissionRepositoryMock.On("FindPermissionById").Return(testStep.mockValue, testStep.mockError)
		permissionRepositoryMock.On("DeletePermissionById").Return(testStep.expectedValue, testStep.mockError)

		// get GIN context
		w := httptest.NewRecorder()
		c := mock.GetGinTestContext(w, "DELETE", gin.Params{
			{Key: "permissionID", Value: testStep.mockValue.(dao.Permission).ID.String()},
		})

		// Run function
		permissionService.DeletePermission(c)

		// Check result
		response := w.Result()
		if response.StatusCode != testStep.expectedStatusCode {
			t.Errorf("Expected status code %d but got %d", testStep.expectedStatusCode, response.StatusCode)
		}

		if testStep.expectedValue == nil {
			continue
		}

	}
}

func TestAddPermission(t *testing.T) {
	// Create mock object
	permissionRepositoryMock := mock.NewPermissionRepositoryMock()
	permissionService := PermissionServiceImpl{
		PermissionRepository: &permissionRepositoryMock,
	}

	// Set mock data
	testSteps := []ServiceTestPOST{
		{
			data: map[string]interface{}{
				"name":        "name_1",
				"description": "description_1",
			},
			mockValue: dao.Permission{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
			},
			expectedValue: dao.Permission{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
			},
			mockError:          nil,
			expectedStatusCode: 201,
		},
		{
			data: map[string]interface{}{},
			mockValue: dao.Permission{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name:        "name_1",
				Description: new(string),
			},
			expectedValue:      nil,
			mockError:          gorm.ErrRecordNotFound,
			expectedStatusCode: 400,
		},
		{
			data: map[string]interface{}{
				"name": "name_1",
			},
			mockValue: dao.Permission{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name:        "name_1",
				Description: new(string),
			},
			expectedValue: dao.Permission{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name:        "name_1",
				Description: new(string),
			},
			mockError:          nil,
			expectedStatusCode: 201,
		},
	}

	// Run test
	for _, testStep := range testSteps {
		// Set mock data
		permissionRepositoryMock.On("Save").Return(testStep.mockValue, testStep.mockError)

		// get GIN context
		w := httptest.NewRecorder()
		c := mock.GetGinTestContextWithBody(w, "POST", gin.Params{}, testStep.data)

		// Run function
		permissionService.AddPermission(c)

		// Check result
		response := w.Result()
		if response.StatusCode != testStep.expectedStatusCode {
			t.Errorf("Expected status code %d but got %d", testStep.expectedStatusCode, response.StatusCode)
		}

		if testStep.expectedValue == nil {
			continue
		}

		// Check mock data
		var responseBody dto.APIResponse[dao.Permission]
		if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
			t.Errorf("Error when unmarshalling response body: %s", err.Error())
		}
		if responseBody.Data.Name != testStep.expectedValue.(dao.Permission).Name {
			t.Errorf("Expected name %s but got %s", testStep.expectedValue.(dao.Permission).Name, responseBody.Data.Name)
		}

		expectedDescription := "nil"
		if testStep.expectedValue.(dao.Permission).Description != nil {
			expectedDescription = *testStep.expectedValue.(dao.Permission).Description
		}

		responseDescription := "nil"
		if responseBody.Data.Description != nil {
			responseDescription = *responseBody.Data.Description
		}

		if responseDescription != expectedDescription {
			t.Errorf("Expected description %s but got %s", expectedDescription, responseDescription)
		}
	}
}
