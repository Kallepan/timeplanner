package router

import (
	"api-gateway/app/mock"
	"api-gateway/config"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

type RouterTest struct {
	httpMethod string // GET, POST, PUT, DELETE
	url        string // /api/v1/user
	result     string // {"message": "GetAll"}
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
			{httpMethod: "GET", url: "/api/v1/department", result: "{\"message\":\"GetAll\"}"},
			{httpMethod: "GET", url: "/api/v1/department/1", result: "{\"message\":\"Get\"}"},
			{httpMethod: "POST", url: "/api/v1/department", result: "{\"message\":\"Create\"}"},
			{httpMethod: "PUT", url: "/api/v1/department/1", result: "{\"message\":\"Update\"}"},
			{httpMethod: "DELETE", url: "/api/v1/department/1", result: "{\"message\":\"Delete\"}"},
		}

		for i, testStep := range testSteps {
			router := Init(init)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(testStep.httpMethod, testStep.url, nil)
			router.ServeHTTP(w, req)

			if w.Code != http.StatusOK {
				t.Errorf("Expected status code 200 got %v", w.Code)
			}

			if w.Body.String() != testStep.result {
				t.Errorf("Expected body to be %v, got %v", testStep.result, w.Body.String())
			}
			t.Logf("Test %v passed", i)
		}
	})

	t.Run("Test User Routes", func(t *testing.T) {
		var testSteps = []RouterTest{
			{httpMethod: "GET", url: "/api/v1/user", result: "{\"message\":\"GetAll\"}"},
			{httpMethod: "GET", url: "/api/v1/user/detail", result: "{\"message\":\"Get\"}"},
			{httpMethod: "POST", url: "/api/v1/user", result: "{\"message\":\"Create\"}"},
			{httpMethod: "PUT", url: "/api/v1/user/1", result: "{\"message\":\"Update\"}"},
			{httpMethod: "DELETE", url: "/api/v1/user/1", result: "{\"message\":\"Delete\"}"},
			{httpMethod: "POST", url: "/api/v1/user/1/permission/1", result: "{\"message\":\"AddPermission\"}"},
			{httpMethod: "DELETE", url: "/api/v1/user/1/permission/1", result: "{\"message\":\"DeletePermission\"}"},
		}

		for i, testStep := range testSteps {
			router := Init(init)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(testStep.httpMethod, testStep.url, nil)
			router.ServeHTTP(w, req)

			if w.Code != http.StatusOK {
				t.Errorf("Expected status code 200 got %v", w.Code)
			}

			if w.Body.String() != testStep.result {
				t.Errorf("Expected body to be %v, got %v", testStep.result, w.Body.String())
			}
			t.Logf("Test %v passed", i)
		}
	})

	t.Run("Test Permission Routes", func(t *testing.T) {
		var testSteps = []RouterTest{
			{httpMethod: "GET", url: "/api/v1/permission", result: "{\"message\":\"GetAll\"}"},
			{httpMethod: "GET", url: "/api/v1/permission/1", result: "{\"message\":\"Get\"}"},
			{httpMethod: "POST", url: "/api/v1/permission", result: "{\"message\":\"Create\"}"},
			{httpMethod: "PUT", url: "/api/v1/permission/1", result: "{\"message\":\"Update\"}"},
			{httpMethod: "DELETE", url: "/api/v1/permission/1", result: "{\"message\":\"Delete\"}"},
		}

		for i, testStep := range testSteps {
			router := Init(init)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(testStep.httpMethod, testStep.url, nil)
			router.ServeHTTP(w, req)

			if w.Code != http.StatusOK {
				t.Errorf("Expected status code 200 got %v", w.Code)
			}

			if w.Body.String() != testStep.result {
				t.Errorf("Expected body to be %v, got %v", testStep.result, w.Body.String())
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
