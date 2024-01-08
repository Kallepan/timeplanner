package service

import (
	"api-gateway/app/domain/dao"
	"api-gateway/app/domain/dco"
	"api-gateway/app/domain/dto"
	"api-gateway/app/mock"
	"database/sql"
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
	dummyDescription := "dummy_description"
	var testSteps = []ServiceTestPUT{
		{
			mockRequestData: map[string]interface{}{
				"name":        "name_new",
				"description": "description_new",
			},
			mockValue:          nil,
			expectedValue:      nil,
			mockError:          gorm.ErrRecordNotFound,
			expectedStatusCode: 404,
			queryParams: map[string]string{
				"permissionID": "00000000-0000-0000-0000-000000000001",
			},
		},
		{
			mockRequestData: map[string]interface{}{
				"name":        "name_new",
				"description": "description_new",
			},
			mockValue: dao.Permission{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name:        "name_old",
				Description: sql.NullString{String: dummyDescription, Valid: true},
			},
			expectedValue: dao.Permission{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name:        "name_new",
				Description: sql.NullString{String: "description_new", Valid: true},
			},
			mockError:          nil,
			expectedStatusCode: 200,
			queryParams: map[string]string{
				"permissionID": "00000000-0000-0000-0000-000000000001",
			},
		},
		{
			// update with nil description
			mockRequestData: map[string]interface{}{
				"name":        "name_new",
				"description": nil,
			},
			mockValue: dao.Permission{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name:        "name_old",
				Description: sql.NullString{String: "description_old", Valid: true},
			},
			expectedValue: dao.Permission{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name:        "name_new",
				Description: sql.NullString{String: "", Valid: false},
			},
			mockError:          nil,
			expectedStatusCode: 200,
			queryParams: map[string]string{
				"permissionID": "00000000-0000-0000-0000-000000000001",
			},
		},
		{
			// update with no description
			mockRequestData: map[string]interface{}{
				"name": "name_new",
			},
			mockValue: dao.Permission{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name:        "name_old",
				Description: sql.NullString{String: "description_old", Valid: true},
			},
			expectedValue: dao.Permission{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name:        "name_new",
				Description: sql.NullString{String: "", Valid: false},
			},
			mockError:          nil,
			expectedStatusCode: 200,
			queryParams: map[string]string{
				"permissionID": "00000000-0000-0000-0000-000000000001",
			},
		},
	}

	// Run test
	for i, testStep := range testSteps {
		// Set mock data
		permissionRepositoryMock.On("FindPermissionById").Return(testStep.mockValue, testStep.mockError)
		permissionRepositoryMock.On("Save").Return(testStep.expectedValue, testStep.mockError)

		// get GIN context
		w := httptest.NewRecorder()
		c := mock.GetGinTestContextWithBody(w, "PUT", testStep.QueryParamsToGinParams(), testStep.mockRequestData)

		// Run function
		permissionService.UpdatePermission(c)

		// Check result
		response := w.Result()
		if response.StatusCode != testStep.expectedStatusCode {
			t.Errorf("Step: %d. Expected status code %d but got %d", i, testStep.expectedStatusCode, response.StatusCode)
		}

		if testStep.expectedValue == nil {
			continue
		}

		// Check mock data
		var responseBody dto.APIResponse[dco.PermissionResponse]
		if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
			t.Errorf("Step: %d. Error when unmarshalling response body: %s", i, err.Error())
		}

		// compare name
		if responseBody.Data.Name != testStep.expectedValue.(dao.Permission).Name {
			t.Errorf("Step: %d. Expected name %s but got %s", i, testStep.expectedValue.(dao.Permission).Name, responseBody.Data.Name)
		}
		// compare description
		if *responseBody.Data.Description != testStep.expectedValue.(dao.Permission).Description.String {
			t.Errorf("Step: %d. Expected description %s but got %s", i, testStep.expectedValue.(dao.Permission).Description.String, *responseBody.Data.Description)
		}
	}
}

func TestDeletePermission(t *testing.T) {
	// Create mock object
	permissionRepositoryMock := mock.NewPermissionRepositoryMock()
	permissionService := PermissionServiceImpl{
		PermissionRepository: &permissionRepositoryMock,
	}

	// Set mock data
	testSteps := []ServiceTestDELETE{
		{
			mockValue: dao.Permission{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name:        "name_1",
				Description: sql.NullString{String: "description_1", Valid: true},
			},
			mockError:          nil,
			expectedStatusCode: 200,
			queryParams: map[string]string{
				"permissionID": "00000000-0000-0000-0000-000000000001",
			},
		},
		{
			mockValue: dao.Permission{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name:        "name_1",
				Description: sql.NullString{String: "description_1", Valid: true},
			},
			mockError:          gorm.ErrRecordNotFound,
			expectedStatusCode: 404,
			queryParams: map[string]string{
				"permissionID": "00000000-0000-0000-0000-000000000001",
			},
		},
	}

	// Run test
	for _, testStep := range testSteps {
		// Set mock data
		permissionRepositoryMock.On("DeletePermissionById").Return(testStep.mockValue, testStep.mockError)

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
			mockRequestData: map[string]interface{}{
				"name":        "name_1",
				"description": "description_1",
			},
			expectedValue: dao.Permission{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name:        "name_1",
				Description: sql.NullString{String: "description_1", Valid: true},
			},
			mockError:          nil,
			expectedStatusCode: 201,
		},
		{
			mockRequestData:    map[string]interface{}{},
			expectedValue:      nil,
			mockError:          nil, // validation error does not need to be mocked
			expectedStatusCode: 400,
		},
		{
			mockRequestData: map[string]interface{}{
				"name": "name_1",
			},
			expectedValue: dao.Permission{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name:        "name_1",
				Description: sql.NullString{String: "", Valid: false},
			},
			mockError:          nil,
			expectedStatusCode: 201,
		},
	}

	// Run test
	for i, testStep := range testSteps {
		// Set mock data
		permissionRepositoryMock.On("Save").Return(testStep.expectedValue, testStep.mockError)

		// get GIN context
		w := httptest.NewRecorder()
		c := mock.GetGinTestContextWithBody(w, "POST", gin.Params{}, testStep.mockRequestData)

		// Run function
		permissionService.AddPermission(c)

		// Check result
		response := w.Result()
		if response.StatusCode != testStep.expectedStatusCode {
			t.Errorf("Step: %d. Expected status code %d but got %d", i, testStep.expectedStatusCode, response.StatusCode)
		}

		if testStep.expectedValue == nil {
			continue
		}

		// Check mock data
		var responseBody dto.APIResponse[dco.PermissionResponse]
		if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
			t.Errorf("Step: %d. Error when unmarshalling response body: %s", i, err.Error())
		}

		// compare name
		if responseBody.Data.Name != testStep.expectedValue.(dao.Permission).Name {
			t.Errorf("Step: %d. Expected name %s but got %s", i, testStep.expectedValue.(dao.Permission).Name, responseBody.Data.Name)
		}
		// compare description
		if *responseBody.Data.Description != testStep.expectedValue.(dao.Permission).Description.String {
			t.Errorf("Step: %d. Expected description %s but got %s", i, testStep.expectedValue.(dao.Permission).Description.String, *responseBody.Data.Description)
		}
	}
}

func TestGetAllPermissions(t *testing.T) {
	// Create mock object
	permissionRepositoryMock := mock.NewPermissionRepositoryMock()
	permissionService := PermissionServiceImpl{
		PermissionRepository: &permissionRepositoryMock,
	}

	dummyDescription := "dummy_description"
	// Set mock data
	testSteps := []ServiceTestGET{
		{
			mockValue: []dao.Permission{
				{
					BaseModel: dao.BaseModel{
						ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					},
					Name:        "name_1",
					Description: sql.NullString{String: dummyDescription, Valid: true},
				},
				{
					BaseModel: dao.BaseModel{
						ID: uuid.MustParse("00000000-0000-0000-0000-000000000002"),
					},
					Name:        "name_2",
					Description: sql.NullString{String: "", Valid: false},
				},
			},
			expectedResponse: []dao.Permission{
				{
					BaseModel: dao.BaseModel{
						ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					},
					Name:        "name_1",
					Description: sql.NullString{String: dummyDescription, Valid: true},
				},
				{
					BaseModel: dao.BaseModel{
						ID: uuid.MustParse("00000000-0000-0000-0000-000000000002"),
					},
					Name:        "name_2",
					Description: sql.NullString{String: "", Valid: false},
				},
			},
			mockError:          nil,
			expectedStatusCode: 200,
		},
		{
			mockValue:          []dao.Permission{},
			expectedResponse:   nil,
			mockError:          nil,
			expectedStatusCode: 200,
		},
	}

	// Run test
	for i, testStep := range testSteps {
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
			t.Errorf("Step: %d.Expected status code %d but got %d", i, testStep.expectedStatusCode, response.StatusCode)
		}

		if testStep.expectedResponse == nil {
			continue
		}

		// Check mock data
		var responseBody dto.APIResponse[[]dco.PermissionResponse]
		if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
			t.Errorf("Step: %d Error when unmarshalling response body: %s", i, err.Error())
		}

		for j, permission := range responseBody.Data {
			// compare name
			if permission.Name != testStep.expectedResponse.([]dao.Permission)[j].Name {
				t.Errorf("Step: %d. Expected name %s but got %s", i, testStep.expectedResponse.([]dao.Permission)[j].Name, permission.Name)
			}
			// compare description
			if *permission.Description != testStep.expectedResponse.([]dao.Permission)[j].Description.String {
				t.Errorf("Step: %d. Expected description %s but got %s", i, testStep.expectedResponse.([]dao.Permission)[j].Description.String, *permission.Description)
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
	dummyDescription := "dummy_description"
	testSteps := []ServiceTestGET{
		{

			mockValue: dao.Permission{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name:        "name_1",
				Description: sql.NullString{String: dummyDescription, Valid: true},
			},
			expectedResponse: dao.Permission{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Name:        "name_1",
				Description: sql.NullString{String: dummyDescription, Valid: true},
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
				Description: sql.NullString{String: "", Valid: false},
			},
			expectedResponse:   nil,
			mockError:          gorm.ErrRecordNotFound,
			expectedStatusCode: 404,
		},
	}

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

		if testStep.expectedResponse == nil {
			continue
		}

		// Check mock data
		var responseBody dto.APIResponse[dco.PermissionResponse]
		if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
			t.Errorf("Error when unmarshalling response body: %s", err.Error())

		}

		// compare name
		if responseBody.Data.Name != testStep.expectedResponse.(dao.Permission).Name {
			t.Errorf("Expected name %s but got %s", testStep.expectedResponse.(dao.Permission).Name, responseBody.Data.Name)
		}

		// compare description
		if *responseBody.Data.Description != testStep.expectedResponse.(dao.Permission).Description.String {
			t.Errorf("Expected description %s but got %s", testStep.expectedResponse.(dao.Permission).Description.String, *responseBody.Data.Description)
		}
	}
}
