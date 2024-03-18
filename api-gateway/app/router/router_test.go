package router

import (
	"api-gateway/app/domain/dao"
	"api-gateway/app/mock"
	"api-gateway/config"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var token string
var authErrorString = "{\"response_key\":\"Unauthorized\",\"response_message\":\"Unauthorized\",\"data\":null}"

func TestMain(m *testing.M) {
	// Generate a mock token
	user := dao.User{Username: "test"}
	t, err := mock.GenerateMockToken(user)
	if err != nil {
		fmt.Printf("Error generating token: %v", err)
		os.Exit(1)
	}

	token = t

	// Run the tests and exit
	os.Exit(m.Run())
	token = ""
}

type RouterTest struct {
	httpMethod       string // GET, POST, PUT, DELETE
	url              string // /api/v1/user
	expectedResponse string // {"message": "GetAll"}
	shouldLogin      bool
}

func TestRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)
	init := &config.Injector{
		DB:             nil,
		SystemCtrl:     &mock.SystemControllerMock{},
		DepartmentCtrl: &mock.DepartmentControllerMock{},
		UserCtrl:       &mock.UserControllerMock{},
		PermissionCtrl: &mock.PermissionControllerMock{},
	}

	t.Run("Test Department Routes", func(t *testing.T) {
		var testSteps = []RouterTest{
			{httpMethod: "GET", url: "/api/v1/department", expectedResponse: "{\"message\":\"GetAll\"}", shouldLogin: true},
			{httpMethod: "GET", url: "/api/v1/department/1", expectedResponse: "{\"message\":\"Get\"}", shouldLogin: true},
			{httpMethod: "POST", url: "/api/v1/department", expectedResponse: "{\"message\":\"Create\"}", shouldLogin: true},
			{httpMethod: "PUT", url: "/api/v1/department/1", expectedResponse: "{\"message\":\"Update\"}", shouldLogin: true},
			{httpMethod: "DELETE", url: "/api/v1/department/1", expectedResponse: "{\"message\":\"Delete\"}", shouldLogin: true},
			{httpMethod: "DELETE", url: "/api/v1/department/1", expectedResponse: authErrorString, shouldLogin: false},
			{httpMethod: "GET", url: "/api/v1/department", expectedResponse: authErrorString, shouldLogin: false},
			{httpMethod: "GET", url: "/api/v1/department/1", expectedResponse: authErrorString, shouldLogin: false},
			{httpMethod: "POST", url: "/api/v1/department", expectedResponse: authErrorString, shouldLogin: false},
			{httpMethod: "PUT", url: "/api/v1/department/1", expectedResponse: authErrorString, shouldLogin: false},
			{httpMethod: "DELETE", url: "/api/v1/department/1", expectedResponse: authErrorString, shouldLogin: false},
		}

		for i, testStep := range testSteps {
			router := Init(init)
			w := httptest.NewRecorder()

			req, _ := http.NewRequest(testStep.httpMethod, testStep.url, nil)
			if testStep.shouldLogin {
				req.AddCookie(&http.Cookie{
					Name:  "Authorization",
					Value: token,
				})
			}

			router.ServeHTTP(w, req)

			if w.Body.String() != testStep.expectedResponse {
				t.Errorf("Expected body to be %v, got %v", testStep.expectedResponse, w.Body.String())
			}
			t.Logf("Test %v passed", i)
		}
	})

	t.Run("Test User Routes", func(t *testing.T) {
		var testSteps = []RouterTest{
			{httpMethod: "GET", url: "/api/v1/user", expectedResponse: "{\"message\":\"GetAll\"}", shouldLogin: true},
			{httpMethod: "GET", url: "/api/v1/user/1", expectedResponse: "{\"message\":\"Get\"}", shouldLogin: true},
			{httpMethod: "POST", url: "/api/v1/user", expectedResponse: "{\"message\":\"Create\"}", shouldLogin: true},
			{httpMethod: "PUT", url: "/api/v1/user/1", expectedResponse: "{\"message\":\"Update\"}", shouldLogin: true},
			{httpMethod: "DELETE", url: "/api/v1/user/1", expectedResponse: "{\"message\":\"Delete\"}", shouldLogin: true},
			{httpMethod: "POST", url: "/api/v1/user/1/permission/1", expectedResponse: "{\"message\":\"AddPermission\"}", shouldLogin: true},
			{httpMethod: "DELETE", url: "/api/v1/user/1/permission/1", expectedResponse: "{\"message\":\"DeletePermission\"}", shouldLogin: true},
			{httpMethod: "GET", url: "/api/v1/user", expectedResponse: authErrorString, shouldLogin: false},
			{httpMethod: "GET", url: "/api/v1/user/detail", expectedResponse: authErrorString, shouldLogin: false},
			{httpMethod: "POST", url: "/api/v1/user", expectedResponse: authErrorString, shouldLogin: false},
			{httpMethod: "PUT", url: "/api/v1/user/1", expectedResponse: authErrorString, shouldLogin: false},
			{httpMethod: "DELETE", url: "/api/v1/user/1", expectedResponse: authErrorString, shouldLogin: false},
			{httpMethod: "POST", url: "/api/v1/user/1/permission/1", expectedResponse: authErrorString, shouldLogin: false},
			{httpMethod: "DELETE", url: "/api/v1/user/1/permission/1", expectedResponse: authErrorString, shouldLogin: false},
		}

		for i, testStep := range testSteps {
			router := Init(init)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(testStep.httpMethod, testStep.url, nil)
			if testStep.shouldLogin {
				req.AddCookie(&http.Cookie{
					Name:  "Authorization",
					Value: token,
				})
			}

			router.ServeHTTP(w, req)

			if w.Body.String() != testStep.expectedResponse {
				t.Errorf("Expected body to be %v, got %v", testStep.expectedResponse, w.Body.String())
			}
			t.Logf("Test %v passed", i)
		}
	})

	t.Run("Test Permission Routes", func(t *testing.T) {
		var testSteps = []RouterTest{
			{httpMethod: "GET", url: "/api/v1/permission", expectedResponse: "{\"message\":\"GetAll\"}", shouldLogin: true},
			{httpMethod: "GET", url: "/api/v1/permission/1", expectedResponse: "{\"message\":\"Get\"}", shouldLogin: true},
			{httpMethod: "POST", url: "/api/v1/permission", expectedResponse: "{\"message\":\"Create\"}", shouldLogin: true},
			{httpMethod: "PUT", url: "/api/v1/permission/1", expectedResponse: "{\"message\":\"Update\"}", shouldLogin: true},
			{httpMethod: "DELETE", url: "/api/v1/permission/1", expectedResponse: "{\"message\":\"Delete\"}", shouldLogin: true},
			{httpMethod: "GET", url: "/api/v1/permission", expectedResponse: authErrorString, shouldLogin: false},
			{httpMethod: "GET", url: "/api/v1/permission/1", expectedResponse: authErrorString, shouldLogin: false},
			{httpMethod: "POST", url: "/api/v1/permission", expectedResponse: authErrorString, shouldLogin: false},
			{httpMethod: "PUT", url: "/api/v1/permission/1", expectedResponse: authErrorString, shouldLogin: false},
			{httpMethod: "DELETE", url: "/api/v1/permission/1", expectedResponse: authErrorString, shouldLogin: false},
		}

		for i, testStep := range testSteps {
			router := Init(init)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(testStep.httpMethod, testStep.url, nil)
			if testStep.shouldLogin {
				req.AddCookie(&http.Cookie{
					Name:  "Authorization",
					Value: token,
				})
			}

			router.ServeHTTP(w, req)

			if w.Body.String() != testStep.expectedResponse {
				t.Errorf("Expected body to be %v, got %v", testStep.expectedResponse, w.Body.String())
			}
			t.Logf("Test %v passed", i)
		}
	})
}

func TestPing(t *testing.T) {
	gin.SetMode(gin.TestMode)
	init := &config.Injector{
		DB:             nil,
		SystemCtrl:     &mock.SystemControllerMock{},
		DepartmentCtrl: &mock.DepartmentControllerMock{},
		UserCtrl:       &mock.UserControllerMock{},
		PermissionCtrl: &mock.PermissionControllerMock{},
	}

	router := Init(init)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/ping", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200 got %v", w.Code)
	}

	if w.Body.String() != "{\"message\":\"pong\"}" {
		t.Errorf("Expected body to be {\"message\":\"pong\"}, got %v", w.Body.String())
	}

}

type ResponseRecorder struct {
	*httptest.ResponseRecorder
	closeNotify chan bool
}

func NewResponseRecorder() *ResponseRecorder {
	return &ResponseRecorder{
		httptest.NewRecorder(),
		make(chan bool, 1),
	}
}

func (r *ResponseRecorder) CloseNotify() <-chan bool {
	return r.closeNotify
}
func TestPlannerProxy(t *testing.T) {
	// crate mock server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "{\"message\":\"mock server response\"}")
	}))
	defer mockServer.Close()

	// set proxy url
	os.Setenv("PLANNER_BACKEND_TARGET", mockServer.URL)

	// prepare routes
	gin.SetMode(gin.TestMode)
	init := &config.Injector{
		DB:             nil,
		SystemCtrl:     &mock.SystemControllerMock{},
		DepartmentCtrl: &mock.DepartmentControllerMock{},
		UserCtrl:       &mock.UserControllerMock{},
		PermissionCtrl: &mock.PermissionControllerMock{},
	}
	router := Init(init)

	// proxy to planner
	w := NewResponseRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/planner/test", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200 got %v", w.Code)
	}

	if w.Body.String() != "{\"message\":\"mock server response\"}\n" {
		t.Errorf("Expected body to be {\"message\":\"mock server response\"}, got %v", w.Body.String())
	}
}
