package service

import (
	"api-gateway/app/domain/dao"
	"api-gateway/app/domain/dco"
	"api-gateway/app/domain/dto"
	"api-gateway/app/middleware"
	"api-gateway/app/mock"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func TestLoginSimple(t *testing.T) {
	// define test struct
	type authLoginTest struct {
		mockRequestData    map[string]interface{}
		expectedStatusCode int
		expectedValue      dao.User
		cookieExpected     bool
		mockError          error
	}

	// mock
	mockUserRepository := mock.NewUserRepositoryMock()

	authService := AuthServiceImpl{
		UserRepository: &mockUserRepository,
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("test"), bcrypt.DefaultCost)
	testSteps := []authLoginTest{
		{
			mockRequestData: map[string]interface{}{
				"username": "test",
				"password": "test",
			},
			expectedValue: dao.User{
				Username: "test",
				Password: string(hashedPassword),
			},
			expectedStatusCode: http.StatusUnauthorized, // Assuming 400 is the status code for bad request
			cookieExpected:     false,
			mockError:          gorm.ErrRecordNotFound,
		},
		{
			mockRequestData: map[string]interface{}{
				"username": "test",
				"password": "test",
			},
			expectedValue: dao.User{
				Username: "test",
				Password: string(hashedPassword),
			},
			expectedStatusCode: 500, // Assuming 500 is the status code for server errors
			cookieExpected:     false,
			mockError:          errors.New("some error"),
		},
		{
			mockRequestData: map[string]interface{}{
				"username": "test",
				"password": "test",
			},
			expectedValue: dao.User{
				Username: "test",
				Password: string(hashedPassword),
			},
			expectedStatusCode: http.StatusOK,
			cookieExpected:     true,
			mockError:          nil,
		},
		{
			mockRequestData: map[string]interface{}{
				"username": "test",
				"password": "test",
			},
			expectedValue:      dao.User{},
			expectedStatusCode: http.StatusUnauthorized,
			cookieExpected:     false,
			mockError:          nil,
		},
	}

	for i, testStep := range testSteps {
		t.Run(fmt.Sprintf("Test Step: %d", i), func(t *testing.T) {
			// Set mock data
			mockUserRepository.On("FindUserByUsername").Return(testStep.expectedValue, testStep.mockError)

			w := httptest.NewRecorder()
			ctx := mock.GetGinTestContext(w, "POST", gin.Params{}, testStep.mockRequestData)

			authService.Login(ctx)

			if w.Code != testStep.expectedStatusCode {
				t.Errorf("Expected status code %d but got %d", testStep.expectedStatusCode, w.Code)
			}

			// check the httpOnly cookie
			cookie := w.Header().Get("Set-Cookie")

			if testStep.cookieExpected {
				if cookie == "" {
					t.Errorf("Expected cookie to be set but got empty")
				}
			} else {
				if cookie != "" {
					t.Errorf("Expected cookie to be empty but got %s", cookie)
				}
			}
		})
	}
}

func TestMe(t *testing.T) {
	// define test struct
	type authMeTest struct {
		data               map[string]string
		query              map[string]string
		expectedStatusCode int
		expectedValue      dao.User
		setCookie          bool
		mockError          error
	}

	// mock
	mockUserRepository := mock.NewUserRepositoryMock()
	authService := AuthServiceImpl{
		UserRepository: &mockUserRepository,
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("test"), bcrypt.DefaultCost)
	testSteps := []authMeTest{
		{
			// check if admin can access 'his' department
			data: map[string]string{
				"username": "test",
			},
			query: map[string]string{
				"department": "test",
			},
			expectedValue: dao.User{
				IsAdmin: true,
				Department: dao.Department{
					Name: "test",
				},
			},
			setCookie:          true,
			expectedStatusCode: http.StatusOK,
		},
		{
			// check if admin can access all department
			data: map[string]string{
				"username": "test",
			},
			query: map[string]string{
				"department": "test2",
			},
			expectedValue: dao.User{
				IsAdmin: true,
				Department: dao.Department{
					Name: "test",
				},
			},
			setCookie:          true,
			expectedStatusCode: http.StatusOK,
		},
		{
			// check if user can access correct department
			data: map[string]string{
				"username": "test",
			},
			query: map[string]string{
				"department": "test",
			},
			expectedValue: dao.User{
				Username: "test",
				IsAdmin:  false,
				Department: dao.Department{
					Name: "test",
				},
			},
			setCookie:          true,
			expectedStatusCode: http.StatusOK,
		},
		{
			// check if user can access different department
			data: map[string]string{
				"username": "test",
			},
			query: map[string]string{
				"department": "test2",
			},
			expectedValue: dao.User{
				IsAdmin: false,
				Department: dao.Department{
					Name: "test",
				},
			},
			setCookie:          true,
			expectedStatusCode: http.StatusUnauthorized,
		},
		{
			data: map[string]string{
				"username": "test",
			},
			expectedValue: dao.User{
				Username: "test",
				Password: string(hashedPassword),
				Department: dao.Department{
					BaseModel: dao.BaseModel{
						ID: uuid.MustParse("00000000-0000-0000-0000-000000000000"),
					},
					Name: "testDepartment",
				},
			},
			setCookie:          true,
			expectedStatusCode: http.StatusOK,
		},
		{
			data: map[string]string{
				"username": "test",
			},
			expectedValue: dao.User{
				BaseModel: dao.BaseModel{
					ID: uuid.MustParse("00000000-0000-0000-0000-000000000000"),
				},
				Username: "test",
				Password: string(hashedPassword),
				Department: dao.Department{
					BaseModel: dao.BaseModel{
						ID: uuid.MustParse("00000000-0000-0000-0000-000000000000"),
					},
					Name: "test",
				},
			},
			setCookie:          true,
			expectedStatusCode: 401,
			mockError:          gorm.ErrRecordNotFound,
		},
		{
			data: map[string]string{
				"username": "test",
			},
			expectedValue:      dao.User{},
			expectedStatusCode: 401,
			setCookie:          false,
		},
		{
			data: map[string]string{
				"username": "test",
			},
			expectedValue:      dao.User{},
			expectedStatusCode: 500,
			mockError:          errors.New("some error"),
			setCookie:          true,
		},
	}

	for i, testStep := range testSteps {
		t.Run(fmt.Sprintf("Test step %d", i), func(t *testing.T) {

			// Set mock data
			mockUserRepository.On("FindUserByUsername").Return(testStep.expectedValue, testStep.mockError)

			w := httptest.NewRecorder()
			ctx, err := mock.NewTestContextBuilder().WithMethod("GET").WithBody(testStep.data).WithQueries(testStep.query).WithResponseRecorder(w).Build()
			if err != nil {
				t.Error("Error happened: when build test context", "error", err)
			}

			// generate mock token
			if testStep.setCookie {
				token, err := mock.GenerateMockToken(testStep.expectedValue)
				if err != nil {
					t.Error("Error happened: when generate mock token", "error", err)
				}
				claim, _ := middleware.DecodeToken(token)
				ctx.Set("retrievedToken", claim)
			}

			// send request
			authService.Me(ctx)

			// check status code
			if w.Code != testStep.expectedStatusCode {
				t.Errorf("Step %d. Expected status code %d but got %d", i, testStep.expectedStatusCode, w.Code)
			}

			// check if user is returned in the response
			if testStep.expectedStatusCode == http.StatusOK {
				var responseBody dto.APIResponse[dco.UserResponse]
				if err := json.Unmarshal(w.Body.Bytes(), &responseBody); err != nil {
					t.Error("Error happened: when unmarshal response body", "error", err)
				}
				if responseBody.Data.Username != testStep.expectedValue.Username {
					t.Errorf("Expected username %s but got %s", testStep.expectedValue.Username, responseBody.Data.Username)
				}

			}
		})
	}

}

func TestLogoutSimple(t *testing.T) {
	// define test struct
	type authLogoutTest struct {
		expectedStatusCode int
		cookieExpected     bool
	}

	authService := AuthServiceImpl{}

	testSteps := []authLogoutTest{
		// for now the logout function will always return 200
		{
			expectedStatusCode: http.StatusOK,
			cookieExpected:     true,
		},
	}

	for i, testStep := range testSteps {
		t.Run(fmt.Sprintf("Test step %d", i), func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx := mock.GetGinTestContext(w, "POST", gin.Params{}, nil)

			authService.Logout(ctx)

			if w.Code != testStep.expectedStatusCode {
				t.Errorf("Expected status code %d but got %d", testStep.expectedStatusCode, w.Code)
			}

			// check the httpOnly cookie
			cookie := w.Header().Get("Set-Cookie")

			if testStep.cookieExpected {
				if cookie == "" {
					t.Errorf("Expected cookie to be set but got empty")
				}
			} else {
				if cookie != "" {
					t.Errorf("Expected cookie to be empty but got %s", cookie)
				}
			}
		})
	}
}

func TestCheckAdmin(t *testing.T) {
	// define test struct
	type authCheckAdminTest struct {
		expectedStatusCode int
		isAdmin            bool
		setCookie          bool
	}

	authService := AuthServiceImpl{}

	testSteps := []authCheckAdminTest{
		{
			expectedStatusCode: http.StatusOK,
			setCookie:          true,
			isAdmin:            true,
		},
		{
			expectedStatusCode: http.StatusUnauthorized,
			setCookie:          false,
		},
		{
			expectedStatusCode: http.StatusUnauthorized,
			setCookie:          true,
			isAdmin:            false,
		},
	}

	for i, testStep := range testSteps {
		t.Run(fmt.Sprintf("Test step %d", i), func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx := mock.GetGinTestContext(w, "GET", gin.Params{}, nil)

			// generate mock token
			if testStep.setCookie {
				token, err := mock.GenerateMockToken(dao.User{IsAdmin: testStep.isAdmin})
				if err != nil {
					t.Error("Error happened: when generate mock token", "error", err)
				}
				claim, _ := middleware.DecodeToken(token)
				ctx.Set("retrievedToken", claim)
			}

			authService.CheckAdmin(ctx)

			if w.Code != testStep.expectedStatusCode {
				t.Errorf("Expected status code %d but got %d", testStep.expectedStatusCode, w.Code)
			}

			// check response body
			if testStep.expectedStatusCode == http.StatusUnauthorized {
				var responseBody dto.APIResponse[bool]
				if err := json.Unmarshal(w.Body.Bytes(), &responseBody); err != nil {
					t.Error("Error happened: when unmarshal response body", "error", err)
				}
				if responseBody.Data != false {
					t.Errorf("Expected false but got %t", responseBody.Data)
				}
			}

			if testStep.expectedStatusCode == http.StatusOK {
				var responseBody dto.APIResponse[bool]
				if err := json.Unmarshal(w.Body.Bytes(), &responseBody); err != nil {
					t.Error("Error happened: when unmarshal response body", "error", err)
				}
				if responseBody.Data != true {
					t.Errorf("Expected true but got %t", responseBody.Data)
				}
			}
		})
	}
}
