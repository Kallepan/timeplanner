package service

import (
	"auth-backend/app/domain/dao"
	"auth-backend/app/domain/dco"
	"auth-backend/app/domain/dto"
	"auth-backend/app/middleware"
	"auth-backend/app/mock"
	"encoding/json"
	"errors"
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
			expectedStatusCode: 400, // Assuming 400 is the status code for bad request
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
			expectedStatusCode: 200,
			cookieExpected:     true,
			mockError:          nil,
		},
		{
			mockRequestData: map[string]interface{}{
				"username": "test",
				"password": "test",
			},
			expectedValue:      dao.User{},
			expectedStatusCode: 400,
			cookieExpected:     false,
			mockError:          nil,
		},
	}

	for _, testStep := range testSteps {
		// Set mock data
		mockUserRepository.On("FindUserByUsername").Return(testStep.expectedValue, testStep.mockError)

		w := httptest.NewRecorder()
		ctx := mock.GetGinTestContextWithBody(w, "POST", gin.Params{}, testStep.mockRequestData)

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
	}
}

func TestMe(t *testing.T) {
	// define test struct
	type authMeTest struct {
		queryParams        map[string]string
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
			queryParams: map[string]string{
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
			expectedStatusCode: 200,
		},
		{
			queryParams: map[string]string{
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
			queryParams: map[string]string{
				"username": "test",
			},
			expectedValue:      dao.User{},
			expectedStatusCode: 401,
			setCookie:          false,
		},
		{
			queryParams: map[string]string{
				"username": "test",
			},
			expectedValue:      dao.User{},
			expectedStatusCode: 500,
			mockError:          errors.New("some error"),
			setCookie:          true,
		},
	}

	for i, testStep := range testSteps {
		// Set mock data
		mockUserRepository.On("FindUserByUsername").Return(testStep.expectedValue, testStep.mockError)

		w := httptest.NewRecorder()
		ctx := mock.GetGinTestContextWithBody(w, "GET", gin.Params{}, nil)

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
		if testStep.expectedStatusCode == 200 {
			var responseBody dto.APIResponse[dco.UserResponse]
			if err := json.Unmarshal(w.Body.Bytes(), &responseBody); err != nil {
				t.Error("Error happened: when unmarshal response body", "error", err)
			}
			if responseBody.Data.Username != testStep.expectedValue.Username {
				t.Errorf("Expected username %s but got %s", testStep.expectedValue.Username, responseBody.Data.Username)
			}

		}
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
			expectedStatusCode: 200,
			cookieExpected:     true,
		},
	}

	for _, testStep := range testSteps {
		w := httptest.NewRecorder()
		ctx := mock.GetGinTestContextWithBody(w, "POST", gin.Params{}, nil)

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
	}
}
