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

func TestUpdateUser(t *testing.T) {
	/* Test Update User */
	// Create Mock Repo
	userRepoMock := mock.NewUserRepositoryMock()
	userService := UserServiceImpl{
		UserRepository: &userRepoMock,
	}

	testSteps := []ServiceTestPUT{
		{
			// Test Update Department without finding one
			mockRequestData: map[string]interface{}{
				"username":      "TEST",
				"email":         "test@example.com",
				"password":      "testpassword",
				"department_id": "00000000-0000-0000-0000-000000000003",
			},
			findValue:          nil,
			saveValue:          nil,
			findError:          gorm.ErrRecordNotFound,
			expectedStatusCode: 404,
			queryParams: map[string]string{
				"userID": "00000000-0000-0000-0000-000000000001",
			},
		},
		{
			// Test Update Department
			mockRequestData: map[string]interface{}{
				"username":      "TEST",
				"email":         "test@example.com",
				"password":      "testpassword",
				"department_id": "00000000-0000-0000-0000-000000000003",
			},
			findValue: dao.User{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Username:     "TEST",
				Email:        "test@example.com",
				DepartmentID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			},
			saveValue: dao.User{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Username: "TEST",
				Email:    "test@example.com",
				Department: dao.Department{
					BaseModel: dao.BaseModel{
						ID: uuid.MustParse("00000000-0000-0000-0000-000000000003"),
					},
				},
			},
			findError:          nil,
			saveError:          nil,
			expectedStatusCode: 200,
			queryParams: map[string]string{
				"userID": "00000000-0000-0000-0000-000000000001",
			},
		},
		{
			// update admin to true
			mockRequestData: map[string]interface{}{
				"username":      "TEST",
				"email":         "test@example.com",
				"password":      "testpassword",
				"department_id": "00000000-0000-0000-0000-000000000001",
				"is_admin":      true,
			},
			findValue: dao.User{
				Username:     "TEST",
				Email:        "test@example.com",
				DepartmentID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				IsAdmin:      false,
			},
			saveValue: dao.User{
				Username:     "TEST",
				Email:        "test@example.com",
				DepartmentID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				IsAdmin:      true,
			},
			findError:          nil,
			saveError:          nil,
			expectedStatusCode: 200,
			queryParams: map[string]string{
				"userID": "00000000-0000-0000-0000-000000000001",
			},
		},
		{
			// update keep admin false
			mockRequestData: map[string]interface{}{
				"username":      "TEST",
				"email":         "test@example.com",
				"password":      "testpassword",
				"department_id": "00000000-0000-0000-0000-000000000001",
				"is_admin":      false,
			},
			findValue: dao.User{
				Username:     "TEST",
				Email:        "test@example.com",
				DepartmentID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				IsAdmin:      false,
			},
			saveValue: dao.User{
				Username:     "TEST",
				Email:        "test@example.com",
				DepartmentID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				IsAdmin:      false,
			},
			findError:          nil,
			saveError:          nil,
			expectedStatusCode: 200,
			queryParams: map[string]string{
				"userID": "00000000-0000-0000-0000-000000000001",
			},
		},
		{
			mockRequestData: map[string]interface{}{
				"username":      "TEST",
				"email":         "test@example.com",
				"password":      "testpassword",
				"department_id": "00000000-0000-0000-0000-000000000001",
				"is_admin":      false,
			},
			findValue: dao.User{
				Username:     "TEST",
				Email:        "test@example.com",
				DepartmentID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				IsAdmin:      false,
			},
			saveValue: dao.User{
				Username:     "TEST",
				Email:        "test@example.com",
				DepartmentID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				IsAdmin:      false,
			},
			findError:          nil,
			saveError:          nil,
			expectedStatusCode: 200,
			queryParams: map[string]string{
				"userID": "00000000-0000-0000-0000-000000000001",
			},
		},
		{
			// Test Update User, remove admin
			mockRequestData: map[string]interface{}{
				"username":      "TEST",
				"email":         "test@example.com",
				"password":      "testpassword",
				"department_id": "00000000-0000-0000-0000-000000000001",
				"is_admin":      false,
			},
			findValue: dao.User{
				Username:     "TEST",
				Email:        "test@example.com",
				DepartmentID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				IsAdmin:      true,
			},
			saveValue: dao.User{
				Username:     "TEST",
				Email:        "test@example.com",
				DepartmentID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				IsAdmin:      false,
			},
			findError:          nil,
			saveError:          nil,
			expectedStatusCode: 200,
			queryParams: map[string]string{
				"userID": "00000000-0000-0000-0000-000000000001",
			},
		},
		{
			// Test Update User, without admin
			mockRequestData: map[string]interface{}{
				"username":      "TEST",
				"email":         "test@example.com",
				"password":      "testpassword",
				"department_id": "00000000-0000-0000-0000-000000000001",
			},
			findValue: dao.User{
				Username:     "TEST",
				Email:        "test@example.com",
				DepartmentID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			},
			saveValue: dao.User{
				Username:     "TEST",
				Email:        "test@example.com",
				DepartmentID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			},
			findError:          nil,
			saveError:          nil,
			expectedStatusCode: 200,
			queryParams: map[string]string{
				"userID": "00000000-0000-0000-0000-000000000001",
			},
		},
		{
			// Test Update User, invalid request email
			mockRequestData: map[string]interface{}{
				"username":      "TEST",
				"email":         "",
				"password":      "testpassword",
				"department_id": uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			},
			findValue:          dao.User{},
			saveValue:          dao.User{},
			expectedStatusCode: 400,
		},
	}

	for i, testStep := range testSteps {
		// Set mock data
		userRepoMock.On("FindUserById").Return(testStep.findValue, testStep.findError)
		userRepoMock.On("Save").Return(testStep.saveValue, testStep.saveError)

		// get GIN context
		w := httptest.NewRecorder()
		c := mock.GetGinTestContextWithBody(w, "PUT", testStep.QueryParamsToGinParams(), testStep.mockRequestData)

		// Call function
		userService.UpdateUser(c)

		response := w.Result()
		if response.StatusCode != testStep.expectedStatusCode {
			t.Errorf("Step: %d. Expected status code %d, but got %d", i, testStep.expectedStatusCode, response.StatusCode)
		}

		if testStep.saveValue == nil {
			continue
		}

		// Read response body
		var responseBody dto.APIResponse[dco.UserResponse]
		if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
			t.Errorf("Step: %d. Error when decoding response body: %s", i, err.Error())
		}

		// compare attributes
		if responseBody.Data.Username != testStep.saveValue.(dao.User).Username {
			t.Errorf("Step: %d. Expected username %s, but got %s", i, testStep.saveValue.(dao.User).Username, responseBody.Data.Username)
		}

		if responseBody.Data.Email != testStep.saveValue.(dao.User).Email {
			t.Errorf("Step: %d. Expected email %s, but got %s", i, testStep.saveValue.(dao.User).Email, responseBody.Data.Email)
		}

		if responseBody.Data.IsAdmin != testStep.saveValue.(dao.User).IsAdmin {
			t.Errorf("Step: %d. Expected is_admin %t, but got %t", i, testStep.saveValue.(dao.User).IsAdmin, responseBody.Data.IsAdmin)
		}
	}
}

func TestDeleteUser(t *testing.T) {
	/* Test Delete User */
	// Create Mock Repo
	userRepoMock := mock.NewUserRepositoryMock()
	userService := UserServiceImpl{
		UserRepository: &userRepoMock,
	}

	testSteps := []ServiceTestDELETE{
		{
			mockValue: dao.User{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Username: "testuser",
				Email:    "test@example.com",
			},
			mockError:          nil,
			expectedStatusCode: 200,
			queryParams: map[string]string{
				"userID": "00000000-0000-0000-0000-000000000001",
			},
		},
		{
			mockValue:          nil,
			mockError:          gorm.ErrRecordNotFound,
			expectedStatusCode: 404,
			queryParams: map[string]string{
				"userID": "00000000-0000-0000-0000-000000000001",
			},
		},
	}

	for _, testStep := range testSteps {
		// Prime mock
		userRepoMock.On("DeleteUser").Return(nil, testStep.mockError)

		// get GIN context
		w := httptest.NewRecorder()
		c := mock.GetGinTestContext(w, "DELETE", testStep.QueryParamsToGinParams())

		// Call function
		userService.DeleteUser(c)

		response := w.Result()
		if response.StatusCode != testStep.expectedStatusCode {
			t.Errorf("Expected status code %d, but got %d", testStep.expectedStatusCode, response.StatusCode)
		}
	}
}

func TestAddUser(t *testing.T) {
	/* Test Add User */
	// Create Mock Repo
	userRepoMock := mock.NewUserRepositoryMock()
	userService := UserServiceImpl{
		UserRepository: &userRepoMock,
	}

	testSteps := []ServiceTestPOST{
		{
			// Test Add User, admin true
			mockRequestData: map[string]interface{}{
				"username":      "TEST",
				"email":         "test@example.com",
				"password":      "testpassword",
				"department_id": "00000000-0000-0000-0000-000000000001",
				"is_admin":      true,
			},
			findValue: nil,
			saveValue: dao.User{
				Username:     "TEST",
				Email:        "test@example.com",
				DepartmentID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				IsAdmin:      true,
			},
			findError:          gorm.ErrRecordNotFound,
			saveError:          nil,
			expectedStatusCode: 201,
		},
		{
			// Test Add User, admin false implicit
			mockRequestData: map[string]interface{}{
				"username":      "TEST",
				"email":         "test@example.com",
				"password":      "testpassword",
				"department_id": "00000000-0000-0000-0000-000000000001",
			},
			findValue: nil,
			saveValue: dao.User{
				Username:     "TEST",
				Email:        "test@example.com",
				DepartmentID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				IsAdmin:      false,
			},
			findError:          gorm.ErrRecordNotFound,
			saveError:          nil,
			expectedStatusCode: 201,
		},
		{
			// Test Add User, admin false explicit
			mockRequestData: map[string]interface{}{
				"username":      "TEST",
				"email":         "test@example.com",
				"password":      "testpassword",
				"department_id": "00000000-0000-0000-0000-000000000001",
				"is_admin":      false,
			},
			findValue: nil,
			saveValue: dao.User{
				Username:     "TEST",
				Email:        "test@example.com",
				DepartmentID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				IsAdmin:      false,
			},
			findError:          gorm.ErrRecordNotFound,
			saveError:          nil,
			expectedStatusCode: 201,
		},
		{
			// Test Add User, invalid request email
			mockRequestData: map[string]interface{}{
				"username":      "TEST",
				"email":         "",
				"password":      "testpassword",
				"department_id": uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			},
			findValue:          dao.User{},
			saveValue:          nil,
			findError:          gorm.ErrRecordNotFound,
			saveError:          nil,
			expectedStatusCode: 400,
		},
		{
			// Test Add User, invalid request username
			mockRequestData: map[string]interface{}{
				"username":      "TEST123",
				"email":         "testmail@example.com",
				"password":      "testpassword",
				"department_id": uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				"is_admin":      false,
			},
			findValue:          nil,
			saveValue:          nil,
			findError:          gorm.ErrRecordNotFound,
			saveError:          nil,
			expectedStatusCode: 400,
		},
	}

	for i, testStep := range testSteps {
		// Prime mock
		userRepoMock.On("FindUserByUsername").Return(testStep.findValue, testStep.findError)
		userRepoMock.On("Save").Return(testStep.saveValue, testStep.saveError)

		// get GIN context
		w := httptest.NewRecorder()
		c := mock.GetGinTestContextWithBody(w, "POST", gin.Params{}, testStep.mockRequestData)

		// Call function
		userService.AddUser(c)

		response := w.Result()
		if response.StatusCode != testStep.expectedStatusCode {
			t.Errorf("Step: %d. Expected status code %d, but got %d", i, testStep.expectedStatusCode, response.StatusCode)
		}

		if testStep.saveValue == nil {
			continue
		}

		// Read response body
		var responseBody dto.APIResponse[dco.UserResponse]
		if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
			t.Errorf("Step: %d. Error when decoding response body: %s", i, err.Error())
		}

		// Compare response body
		if responseBody.Data.Username != testStep.saveValue.(dao.User).Username {
			t.Errorf("Step: %d. Expected username %s, but got %s", i, testStep.saveValue.(dao.User).Username, responseBody.Data.Username)
		}

		if responseBody.Data.Email != testStep.saveValue.(dao.User).Email {
			t.Errorf("Step: %d. Expected email %s, but got %s", i, testStep.saveValue.(dao.User).Email, responseBody.Data.Email)
		}

		if responseBody.Data.IsAdmin != testStep.saveValue.(dao.User).IsAdmin {
			t.Errorf("Step: %d. Expected is_admin %t, but got %t", i, testStep.saveValue.(dao.User).IsAdmin, responseBody.Data.IsAdmin)
		}

		if responseBody.Data.Department.ID != testStep.saveValue.(dao.User).DepartmentID {
			t.Errorf("Step: %d. Expected department_id %s, but got %s", i, testStep.saveValue.(dao.User).DepartmentID, responseBody.Data.Department.ID)
		}
	}
}

func TestGetAllUsers(t *testing.T) {
	/* Test All Users */
	// Create Mock Repo
	userRepoMock := mock.NewUserRepositoryMock()
	userService := UserServiceImpl{
		UserRepository: &userRepoMock,
	}

	testSteps := []ServiceTestGET{
		{
			// Get All Users, hide password
			mockValue: []dao.User{
				{
					BaseModel: dao.BaseModel{
						ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					},
					Username:     "testuser",
					Email:        "test@example.com",
					DepartmentID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					Password:     "asdasd",
					IsAdmin:      true,
				},
				{
					BaseModel: dao.BaseModel{
						ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					},
					Username:     "testuser",
					Email:        "test@example.com",
					DepartmentID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					Password:     "asdasd",
				},
				{
					BaseModel: dao.BaseModel{
						ID: uuid.MustParse("00000000-0000-0000-0000-000000000002"),
					},
					Username:     "testuser2",
					Email:        "test2@example.com",
					DepartmentID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
			},
			expectedResponse: []dao.User{
				{
					BaseModel: dao.BaseModel{
						ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					},
					Username:     "testuser",
					Email:        "test@example.com",
					DepartmentID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					IsAdmin:      true,
				},
				{
					BaseModel: dao.BaseModel{
						ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					},
					Username:     "testuser",
					Email:        "test@example.com",
					DepartmentID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					Password:     "",
					IsAdmin:      false,
				},
				{
					BaseModel: dao.BaseModel{
						ID: uuid.MustParse("00000000-0000-0000-0000-000000000002"),
					},
					Username:     "testuser2",
					Email:        "test2@example.com",
					DepartmentID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					IsAdmin:      false,
				},
			},
			mockError:          nil,
			expectedStatusCode: 200,
		},
	}

	for i, testStep := range testSteps {
		// Prime mock
		userRepoMock.On("FindAllUsers").Return(testStep.mockValue, testStep.mockError)

		// get GIN context
		w := httptest.NewRecorder()
		c := mock.GetGinTestContext(w, "GET", gin.Params{})

		// Call function
		userService.GetAllUsers(c)

		response := w.Result()
		if response.StatusCode != testStep.expectedStatusCode {
			t.Errorf("Step: %d. Expected status code %d, but got %d", i, testStep.expectedStatusCode, response.StatusCode)
		}
		if testStep.expectedResponse == nil {
			continue
		}

		// Read response body
		var responseBody dto.APIResponse[[]dco.UserResponse]
		if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
			t.Errorf("Step: %d. Error when decoding response body: %s", i, err.Error())
		}

		// Compare response body
		for j, user := range responseBody.Data {
			if user.Username != testStep.expectedResponse.([]dao.User)[j].Username {
				t.Errorf("Step: %d. Expected username %s, but got %s", i, testStep.expectedResponse.([]dao.User)[j].Username, user.Username)
			}

			if user.Email != testStep.expectedResponse.([]dao.User)[j].Email {
				t.Errorf("Step: %d. Expected email %s, but got %s", i, testStep.expectedResponse.([]dao.User)[j].Email, user.Email)
			}

			if user.IsAdmin != testStep.expectedResponse.([]dao.User)[j].IsAdmin {
				t.Errorf("Step: %d. Expected is_admin %t, but got %t", i, testStep.expectedResponse.([]dao.User)[j].IsAdmin, user.IsAdmin)
			}

			if user.Department.ID != testStep.expectedResponse.([]dao.User)[j].DepartmentID {
				t.Errorf("Step: %d. Expected department_id %s, but got %s", i, testStep.expectedResponse.([]dao.User)[j].DepartmentID, user.Department.ID)
			}
		}
	}
}

func TestGetUser(t *testing.T) {
	/* Test Get User */
	// Create Mock Repo
	userRepoMock := mock.NewUserRepositoryMock()
	userService := UserServiceImpl{
		UserRepository: &userRepoMock,
	}

	testSteps := []ServiceTestGET{
		{
			// Test Get User, hide password
			mockValue: dao.User{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Username:     "testuser",
				Email:        "test@example.com",
				DepartmentID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Password:     "asdasd",
				IsAdmin:      true,
			},
			expectedResponse: dao.User{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Username:     "testuser",
				Email:        "test@example.com",
				DepartmentID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Password:     "",
				IsAdmin:      true,
			},
			mockError:          nil,
			expectedStatusCode: 200,
		},
		{
			// Test Get User, admin false
			mockValue: dao.User{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Username:     "testuser",
				Email:        "test@example.com",
				DepartmentID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Password:     "asdasd",
			},
			expectedResponse: dao.User{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Username:     "testuser",
				Email:        "test@example.com",
				DepartmentID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Password:     "",
				IsAdmin:      false,
			},
			mockError:          nil,
			expectedStatusCode: 200,
		},
		{
			// Test Get User, admin true
			mockValue:          dao.User{},
			expectedResponse:   nil,
			mockError:          gorm.ErrRecordNotFound,
			expectedStatusCode: 404,
		},
		{
			// Test Get User, with permissions
			mockValue: dao.User{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Permissions: []dao.Permission{
					{
						BaseModel: dao.BaseModel{
							ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
						},
						Name:        "testpermission",
						Description: sql.NullString{String: "", Valid: false},
					},
				},
			},
			expectedResponse: dao.User{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Permissions: []dao.Permission{
					{
						BaseModel: dao.BaseModel{
							ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
						},
						Name:        "testpermission",
						Description: sql.NullString{String: "", Valid: false},
					},
				},
			},
			mockError:          nil,
			expectedStatusCode: 200,
		},
	}

	for i, testStep := range testSteps {
		// Prime mock
		userRepoMock.On("FindUserById").Return(testStep.mockValue, testStep.mockError)

		// get GIN context
		w := httptest.NewRecorder()
		c := mock.GetGinTestContext(w, "GET", gin.Params{
			{Key: "userID", Value: "00000000-0000-0000-0000-000000000001"},
		})

		// Call function
		userService.GetUserById(c)

		response := w.Result()
		if response.StatusCode != testStep.expectedStatusCode {
			t.Errorf("Step: %d. Expected status code %d, but got %d", i, testStep.expectedStatusCode, response.StatusCode)
		}

		if testStep.expectedResponse == nil {
			continue
		}

		// Read response body
		var responseBody dto.APIResponse[dco.UserResponse]
		if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
			t.Errorf("Step: %d. Error when decoding response body: %s", i, err.Error())
		}

		// Compare response body
		if responseBody.Data.Username != testStep.expectedResponse.(dao.User).Username {
			t.Errorf("Step: %d. Expected username %s, but got %s", i, testStep.expectedResponse.(dao.User).Username, responseBody.Data.Username)
		}

		if responseBody.Data.Email != testStep.expectedResponse.(dao.User).Email {
			t.Errorf("Step: %d. Expected email %s, but got %s", i, testStep.expectedResponse.(dao.User).Email, responseBody.Data.Email)
		}

		if responseBody.Data.IsAdmin != testStep.expectedResponse.(dao.User).IsAdmin {
			t.Errorf("Step: %d. Expected is_admin %t, but got %t", i, testStep.expectedResponse.(dao.User).IsAdmin, responseBody.Data.IsAdmin)
		}

		if responseBody.Data.Department.ID != testStep.expectedResponse.(dao.User).DepartmentID {
			t.Errorf("Step: %d. Expected department_id %s, but got %s", i, testStep.expectedResponse.(dao.User).DepartmentID, responseBody.Data.Department.ID)
		}
	}
}

func TestAddPermissionToUser(t *testing.T) {
	/* Test Add Permission */
	// Create Mock Repo
	userRepoMock := mock.NewUserRepositoryMock()
	userService := UserServiceImpl{
		UserRepository: &userRepoMock,
	}

	testSteps := []ServiceTestPOST{
		{
			// Test Add Permission
			mockRequestData: map[string]interface{}{
				"userID":       "TEST",
				"permissionID": "00000000-0000-0000-0000-000000000001",
			},

			saveValue:          nil,
			saveError:          nil,
			expectedStatusCode: 201,
		},
		{
			// Test Add Permission, not found
			mockRequestData: map[string]interface{}{
				"userID":       "TEST",
				"permissionID": "00000000-0000-0000-0000-000000000001",
			},
			saveValue:          nil,
			saveError:          gorm.ErrRecordNotFound,
			expectedStatusCode: 400, // Here we expect 400 because the user or permission do not exist
		},
	}

	for i, testStep := range testSteps {
		// Prime mock
		userRepoMock.On("AddPermissionToUser").Return(testStep.saveValue, testStep.saveError)

		// get GIN context
		w := httptest.NewRecorder()
		c := mock.GetGinTestContextWithBody(w, "POST", gin.Params{
			{Key: "userID", Value: "00000000-0000-0000-0000-000000000001"},
			{Key: "permissionID", Value: "00000000-0000-0000-0000-000000000001"},
		}, testStep.mockRequestData)

		// Call function
		userService.AddPermission(c)

		response := w.Result()
		if response.StatusCode != testStep.expectedStatusCode {
			t.Errorf("Step: %d. Expected status code %d, but got %d", i, testStep.expectedStatusCode, response.StatusCode)
		}
	}
}

func TestDeletePermissionFromUser(t *testing.T) {
	/* Test Delete Permission */
	// Create Mock Repo
	userRepoMock := mock.NewUserRepositoryMock()
	userService := UserServiceImpl{
		UserRepository: &userRepoMock,
	}

	testSteps := []ServiceTestGET{
		{
			// Test Delete Permission
			mockValue:          nil,
			expectedResponse:   nil,
			mockError:          nil,
			expectedStatusCode: 200,
		},
		{
			// Test Delete Permission, not found
			mockValue:          nil,
			expectedResponse:   nil,
			mockError:          gorm.ErrRecordNotFound,
			expectedStatusCode: 400,
		},
	}

	for i, testStep := range testSteps {
		// Prime mock
		userRepoMock.On("DeletePermissionFromUser").Return(nil, testStep.mockError)

		// get GIN context
		w := httptest.NewRecorder()
		c := mock.GetGinTestContext(w, "DELETE", gin.Params{
			{Key: "userID", Value: "00000000-0000-0000-0000-000000000001"},
			{Key: "permissionID", Value: "00000000-0000-0000-0000-000000000001"},
		})

		// Call function
		userService.DeletePermission(c)

		response := w.Result()
		if response.StatusCode != testStep.expectedStatusCode {
			t.Errorf("Step: %d. Expected status code %d, but got %d", i, testStep.expectedStatusCode, response.StatusCode)
		}
	}
}
