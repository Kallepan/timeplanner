package router

import (
	"api-gateway/app/mock"
	"api-gateway/config"
	"net/http"
	"net/http/httptest"
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

			if w.Code != 200 {
				t.Errorf("Expected status code 200, got %v", w.Code)
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
			{httpMethod: "GET", url: "/api/v1/user/1", result: "{\"message\":\"Get\"}"},
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

			if w.Code != 200 {
				t.Errorf("Expected status code 200, got %v", w.Code)
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

			if w.Code != 200 {
				t.Errorf("Expected status code 200, got %v", w.Code)
			}

			if w.Body.String() != testStep.result {
				t.Errorf("Expected body to be %v, got %v", testStep.result, w.Body.String())
			}
			t.Logf("Test %v passed", i)
		}
	})
}
