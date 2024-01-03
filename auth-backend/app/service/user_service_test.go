package service

import (
	"auth-backend/app/constant"
	"auth-backend/app/domain/dao"
	"auth-backend/app/domain/dto"
	"auth-backend/app/mock"
	"auth-backend/app/pkg"
	"encoding/json"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func TestGetUser(t *testing.T) {
	/* Test Get User */
	// Create Mock Repo
	userRepoMock := mock.NewUserRepositoryMock()
	userService := UserServiceImpl{
		UserRepository: &userRepoMock,
	}

	testSteps := []ServiceTestGET{
		{
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
			expectedValue: dao.User{
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
			description:        "Test Get User, hide password",
		},
		{
			mockValue: dao.User{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Username:     "testuser",
				Email:        "test@example.com",
				DepartmentID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Password:     "asdasd",
			},
			expectedValue: dao.User{
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
			description:        "Test Get User, admin false",
		},
		{
			mockValue:          dao.User{},
			expectedValue:      nil,
			mockError:          gorm.ErrRecordNotFound,
			expectedStatusCode: 404,
			description:        "Test Get User, not found",
		},
		{
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
						Description: nil,
					},
				},
			},
			expectedValue: dao.User{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				},
				Permissions: []dao.Permission{
					{
						BaseModel: dao.BaseModel{
							ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
						},
						Name:        "testpermission",
						Description: nil,
					},
				},
			},
			mockError:          nil,
			expectedStatusCode: 200,
			description:        "Test Get User, with permissions",
		},
	}

	for _, testStep := range testSteps {
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
			t.Errorf("Expected status code %d, but got %d", testStep.expectedStatusCode, response.StatusCode)
		}

		if testStep.expectedValue == nil {
			continue
		}

		// Read response body
		var responseBody dto.APIResponse[dao.User]
		if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
			t.Errorf("Error when decoding response body: %s", err.Error())
		}

		// Compare response body
		if !reflect.DeepEqual(responseBody.Data, testStep.expectedValue) {
			t.Errorf("Expected response body %v, but got %v", testStep.expectedValue, responseBody.Data)
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
			expectedValue: []dao.User{
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
			description:        "Test Get All Users",
		},
	}

	for _, testStep := range testSteps {
		// Prime mock
		userRepoMock.On("FindAllUsers").Return(testStep.mockValue, testStep.mockError)

		// get GIN context
		w := httptest.NewRecorder()
		c := mock.GetGinTestContext(w, "GET", gin.Params{})

		// Call function
		userService.GetAllUsers(c)

		response := w.Result()
		if response.StatusCode != testStep.expectedStatusCode {
			t.Errorf("Expected status code %d, but got %d", testStep.expectedStatusCode, response.StatusCode)
		}
		if testStep.expectedValue == nil {
			continue
		}

		// Read response body
		var responseBody dto.APIResponse[[]dao.User]
		if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
			t.Errorf("Error when decoding response body: %s", err.Error())
		}

		// Compare response body
		if !reflect.DeepEqual(responseBody.Data, testStep.expectedValue) {
			t.Errorf("Expected response body %v, but got %v", testStep.expectedValue, responseBody.Data)
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
			data: map[string]interface{}{
				"username":      "TEST",
				"email":         "test@example.com",
				"password":      "testpassword",
				"department_id": "00000000-0000-0000-0000-000000000001",
				"is_admin":      true,
			},
			expectedValue: dao.User{
				Username:     "TEST",
				Email:        "test@example.com",
				DepartmentID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				IsAdmin:      true,
			},
			mockError:          nil,
			expectedStatusCode: 201,
			description:        "Test Add User, admin true",
		},
		{
			data: map[string]interface{}{
				"username":      "TEST",
				"email":         "test@example.com",
				"password":      "testpassword",
				"department_id": "00000000-0000-0000-0000-000000000001",
			},
			expectedValue: dao.User{
				Username:     "TEST",
				Email:        "test@example.com",
				DepartmentID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				IsAdmin:      false,
			},
			mockError:          nil,
			expectedStatusCode: 201,
			description:        "Test Add User, admin false implicit",
		},
		{
			data: map[string]interface{}{
				"username":      "TEST",
				"email":         "test@example.com",
				"password":      "testpassword",
				"department_id": "00000000-0000-0000-0000-000000000001",
				"is_admin":      false,
			},
			expectedValue: dao.User{
				Username:     "TEST",
				Email:        "test@example.com",
				DepartmentID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				IsAdmin:      false,
			},
			mockError:          nil,
			expectedStatusCode: 201,
			description:        "Test Add User, admin false explicit",
		},
		{
			data: map[string]interface{}{
				"username":      "TEST",
				"email":         "",
				"password":      "testpassword",
				"department_id": uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			},
			expectedValue:      dao.User{},
			mockError:          pkg.NewException(constant.InvalidRequest),
			expectedStatusCode: 400,
			description:        "Test Add User, invalid request email",
		},
		{
			data: map[string]interface{}{
				"username":      "TEST123",
				"email":         "testmail@example.com",
				"password":      "testpassword",
				"department_id": uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				"is_admin":      false,
			},
			expectedValue:      dao.User{},
			mockError:          pkg.NewException(constant.InvalidRequest),
			expectedStatusCode: 400,
			description:        "Test Add User, invalid request username",
		},
	}

	for i, testStep := range testSteps {
		// Prime mock
		userRepoMock.On("Save").Return(testStep.expectedValue, testStep.mockError)

		// get GIN context
		w := httptest.NewRecorder()
		c := mock.GetGinTestContextWithBody(w, "POST", gin.Params{}, testStep.data)

		// Call function
		userService.AddUser(c)

		response := w.Result()
		if response.StatusCode != testStep.expectedStatusCode {
			t.Errorf("Step: %d. Expected status code %d, but got %d", i, testStep.expectedStatusCode, response.StatusCode)
		}

		if testStep.expectedValue == nil {
			continue
		}

		// Read response body
		var responseBody dto.APIResponse[dao.User]
		if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
			t.Errorf("Step: %d. Error when decoding response body: %s", i, err.Error())
		}

		// Compare response body
		if !reflect.DeepEqual(responseBody.Data, testStep.expectedValue) {
			t.Errorf("Step: %d. Expected response body %v, but got %v", i, testStep.expectedValue, responseBody.Data)
		}

	}
}

func TestUpdateUser(t *testing.T) {
	/* Test Update User */
	// Create Mock Repo
	userRepoMock := mock.NewUserRepositoryMock()
	userService := UserServiceImpl{
		UserRepository: &userRepoMock,
	}

	testSteps := []ServiceTestPUT{
		{
			data: map[string]interface{}{
				"username":      "TEST",
				"email":         "test@example.com",
				"password":      "testpassword",
				"department_id": "00000000-0000-0000-0000-000000000001",
				"is_admin":      true,
			},
			mockValue: dao.User{
				Username:     "TEST",
				Email:        "test@example.com",
				DepartmentID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				IsAdmin:      false,
			},
			expectedValue: dao.User{
				Username:     "TEST",
				Email:        "test@example.com",
				DepartmentID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				IsAdmin:      true,
			},
			mockError:          nil,
			expectedStatusCode: 200,
			description:        "Test Update User, make admin true",
		},
		{
			data: map[string]interface{}{
				"username":      "TEST",
				"email":         "test@example.com",
				"password":      "testpassword",
				"department_id": "00000000-0000-0000-0000-000000000001",
				"is_admin":      false,
			},
			mockValue: dao.User{
				Username:     "TEST",
				Email:        "test@example.com",
				DepartmentID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				IsAdmin:      false,
			},
			expectedValue: dao.User{
				Username:     "TEST",
				Email:        "test@example.com",
				DepartmentID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				IsAdmin:      false,
			},
			mockError:          nil,
			expectedStatusCode: 200,
			description:        "Test Update User, keep admin false",
		},
		{
			data: map[string]interface{}{
				"username":      "TEST",
				"email":         "test@example.com",
				"password":      "testpassword",
				"department_id": "00000000-0000-0000-0000-000000000001",
				"is_admin":      false,
			},
			mockValue: dao.User{
				Username:     "TEST",
				Email:        "test@example.com",
				DepartmentID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				IsAdmin:      false,
			},
			expectedValue: dao.User{
				Username:     "TEST",
				Email:        "test@example.com",
				DepartmentID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				IsAdmin:      false,
			},
			mockError:          nil,
			expectedStatusCode: 200,
			description:        "Test Update User, keep admin false",
		},
		{
			data: map[string]interface{}{
				"username":      "TEST",
				"email":         "test@example.com",
				"password":      "testpassword",
				"department_id": "00000000-0000-0000-0000-000000000001",
				"is_admin":      false,
			},
			mockValue: dao.User{
				Username:     "TEST",
				Email:        "test@example.com",
				DepartmentID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				IsAdmin:      true,
			},
			expectedValue: dao.User{
				Username:     "TEST",
				Email:        "test@example.com",
				DepartmentID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				IsAdmin:      false,
			},
			mockError:          nil,
			expectedStatusCode: 200,
			description:        "Test Update User, remove admin",
		},
		{
			data: map[string]interface{}{
				"username":      "TEST",
				"email":         "test@example.com",
				"password":      "testpassword",
				"department_id": "00000000-0000-0000-0000-000000000001",
			},
			mockValue: dao.User{
				Username:     "TEST",
				Email:        "test@example.com",
				DepartmentID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			},
			expectedValue: dao.User{
				Username:     "TEST",
				Email:        "test@example.com",
				DepartmentID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			},
			mockError:          nil,
			expectedStatusCode: 200,
			description:        "Test Update User",
		},
		{
			data: map[string]interface{}{
				"username":      "TEST",
				"email":         "",
				"password":      "testpassword",
				"department_id": uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			},
			expectedValue:      dao.User{},
			mockError:          pkg.NewException(constant.InvalidRequest),
			expectedStatusCode: 400,
			description:        "Test Update User, invalid request email",
		},
	}

	for _, testStep := range testSteps {
		// Set mock data
		userRepoMock.On("FindUserById").Return(testStep.mockValue, testStep.mockError)
		userRepoMock.On("Save").Return(testStep.expectedValue, testStep.mockError)

		// get GIN context
		w := httptest.NewRecorder()
		c := mock.GetGinTestContextWithBody(w, "PUT", gin.Params{
			{Key: "userID", Value: "00000000-0000-0000-0000-000000000001"},
		}, testStep.data)

		// Call function
		userService.UpdateUser(c)

		response := w.Result()
		if response.StatusCode != testStep.expectedStatusCode {
			t.Errorf("Expected status code %d, but got %d", testStep.expectedStatusCode, response.StatusCode)
		}

		if testStep.expectedValue == nil {
			continue
		}

		// Read response body
		var responseBody dto.APIResponse[dao.User]
		if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
			t.Errorf("Error when decoding response body: %s", err.Error())
		}
		// Compare response body
		if !reflect.DeepEqual(responseBody.Data, testStep.expectedValue) {
			t.Errorf("Expected response body %v, but got %v", testStep.expectedValue, responseBody.Data)
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

	testSteps := []ServiceTestGET{
		{
			mockValue:          nil,
			expectedValue:      nil,
			mockError:          nil,
			expectedStatusCode: 200,
			description:        "Test Delete User",
		},
		{
			mockValue:          nil,
			expectedValue:      nil,
			mockError:          gorm.ErrRecordNotFound,
			expectedStatusCode: 404,
			description:        "Test Delete User, not found",
		},
	}

	for _, testStep := range testSteps {
		// Prime mock
		userRepoMock.On("DeleteUser").Return(nil, testStep.mockError)

		// get GIN context
		w := httptest.NewRecorder()
		c := mock.GetGinTestContext(w, "DELETE", gin.Params{
			{Key: "userID", Value: "00000000-0000-0000-0000-000000000001"},
		})

		// Call function
		userService.DeleteUser(c)

		response := w.Result()
		if response.StatusCode != testStep.expectedStatusCode {
			t.Errorf("Expected status code %d, but got %d", testStep.expectedStatusCode, response.StatusCode)
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
			data: map[string]interface{}{
				"userID":       "TEST",
				"permissionID": "00000000-0000-0000-0000-000000000001",
			},
			expectedValue:      pkg.Null(),
			mockError:          nil,
			expectedStatusCode: 201,
			description:        "Test Add Permission",
		},
		{
			data: map[string]interface{}{
				"userID":       "TEST",
				"permissionID": "00000000-0000-0000-0000-000000000001",
			},
			expectedValue:      pkg.Null(),
			mockError:          gorm.ErrRecordNotFound,
			expectedStatusCode: 400, // Here we expect 400 because the user or permission do not exist
			description:        "Test Add Permission, not found",
		},
	}

	for _, testStep := range testSteps {
		// Prime mock
		userRepoMock.On("AddPermissionToUser").Return(testStep.expectedValue, testStep.mockError)

		// get GIN context
		w := httptest.NewRecorder()
		c := mock.GetGinTestContextWithBody(w, "POST", gin.Params{
			{Key: "userID", Value: "00000000-0000-0000-0000-000000000001"},
			{Key: "permissionID", Value: "00000000-0000-0000-0000-000000000001"},
		}, testStep.data)

		// Call function
		userService.AddPermission(c)

		response := w.Result()
		if response.StatusCode != testStep.expectedStatusCode {
			t.Errorf("Expected status code %d, but got %d", testStep.expectedStatusCode, response.StatusCode)
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
			mockValue:          nil,
			expectedValue:      nil,
			mockError:          nil,
			expectedStatusCode: 200,
			description:        "Test Delete Permission",
		},
		{
			mockValue:          nil,
			expectedValue:      nil,
			mockError:          gorm.ErrRecordNotFound,
			expectedStatusCode: 400,
			description:        "Test Delete Permission, not found",
		},
	}

	for _, testStep := range testSteps {
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
			t.Errorf("Expected status code %d, but got %d", testStep.expectedStatusCode, response.StatusCode)
		}
	}
}
