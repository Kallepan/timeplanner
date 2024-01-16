package router

import (
	"net/http"
	"net/http/httptest"
	"planner-backend/app/mock"
	"planner-backend/config"
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
		WorkplaceCtrl:  &mock.WorkplaceControllerMock{},
		TimeslotCtrl:   &mock.TimeslotControllerMock{},
		WeekdayCtrl:    &mock.WeekdayControllerMock{},
		PersonCtrl:     &mock.PersonControllerMock{},
		PersonRelCtrl:  &mock.PersonRelControllerMock{},
		WorkdayCtrl:    &mock.WorkdayControllerMock{},
	}

	t.Run("Test System Routes", func(t *testing.T) {
		var testSteps = []RouterTest{
			{httpMethod: "GET", url: "/api/v1/planner/ping", result: "{\"message\":\"pong\"}"},
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
