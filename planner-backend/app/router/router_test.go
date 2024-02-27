package router

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"planner-backend/app/mock"
	"planner-backend/config"
	"testing"

	gatewayDAO "api-gateway/app/domain/dao"
	gatewayMock "api-gateway/app/mock"

	"github.com/gin-gonic/gin"
)

var token string
var authErrorString = "{\"response_key\":\"Unauthorized\",\"response_message\":\"Unauthorized\",\"data\":null}"

func TestMain(m *testing.M) {
	// Generate a mock token
	user := gatewayDAO.User{Username: "test"}
	t, err := gatewayMock.GenerateMockToken(user)
	if err != nil {
		slog.Error("Error generating token: %v", err)
		os.Exit(1)
	}

	token = t

	os.Exit(m.Run())
	token = ""
}

type RouterTest struct {
	httpMethod       string // GET, POST, PUT, DELETE
	url              string // /api/v1/user
	expectedResponse string
	shouldLogin      bool
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
		AbsenceCtrl:    &mock.AbsenceControllerMock{},
	}

	t.Run("Test System Routes", func(t *testing.T) {
		var testSteps = []RouterTest{
			{httpMethod: "GET", url: "/api/v1/planner/ping", expectedResponse: "{\"message\":\"pong\"}", shouldLogin: false},
		}

		for i, testStep := range testSteps {
			router := Init(init)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest(testStep.httpMethod, testStep.url, nil)

			router.ServeHTTP(w, req)
			if w.Code != http.StatusOK {
				t.Errorf("Expected status code http.StatusOK got %v", w.Code)
			}

			if w.Body.String() != testStep.expectedResponse {
				t.Errorf("Expected body to be %v, got %v", testStep.expectedResponse, w.Body.String())
			}

			t.Logf("Test %v passed", i)
		}
	})

	t.Run("Test Department and Subroutes", func(t *testing.T) {
		var testSteps = []RouterTest{
			// not logged in
			{httpMethod: "GET", url: "/api/v1/planner/department/", expectedResponse: "{\"message\":\"GetAll\"}", shouldLogin: false},
			{httpMethod: "GET", url: "/api/v1/planner/department/1", expectedResponse: "{\"message\":\"Get\"}", shouldLogin: false},
			{httpMethod: "GET", url: "/api/v1/planner/department/1/workplace/", expectedResponse: "{\"message\":\"GetAll\"}", shouldLogin: false},
			{httpMethod: "GET", url: "/api/v1/planner/department/1/workplace/1", expectedResponse: "{\"message\":\"Get\"}", shouldLogin: false},
			{httpMethod: "GET", url: "/api/v1/planner/department/1/workplace/1/timeslot/", expectedResponse: "{\"message\":\"GetAll\"}", shouldLogin: false},
			{httpMethod: "GET", url: "/api/v1/planner/department/1/workplace/1/timeslot/1", expectedResponse: "{\"message\":\"Get\"}", shouldLogin: false},
			{httpMethod: "GET", url: "/api/v1/planner/department/1/absency/", expectedResponse: "{\"message\":\"GetAll\"}", shouldLogin: false},
			// logged in
			{httpMethod: "GET", url: "/api/v1/planner/department/", expectedResponse: "{\"message\":\"GetAll\"}", shouldLogin: true},
			{httpMethod: "GET", url: "/api/v1/planner/department/1", expectedResponse: "{\"message\":\"Get\"}", shouldLogin: true},
			{httpMethod: "GET", url: "/api/v1/planner/department/1/workplace/", expectedResponse: "{\"message\":\"GetAll\"}", shouldLogin: true},
			{httpMethod: "GET", url: "/api/v1/planner/department/1/workplace/1", expectedResponse: "{\"message\":\"Get\"}", shouldLogin: true},
			{httpMethod: "GET", url: "/api/v1/planner/department/1/workplace/1/timeslot/", expectedResponse: "{\"message\":\"GetAll\"}", shouldLogin: true},
			{httpMethod: "GET", url: "/api/v1/planner/department/1/workplace/1/timeslot/1", expectedResponse: "{\"message\":\"Get\"}", shouldLogin: true},
			{httpMethod: "GET", url: "/api/v1/planner/department/1/absency/", expectedResponse: "{\"message\":\"GetAll\"}", shouldLogin: true},

			{httpMethod: "POST", url: "/api/v1/planner/department/", expectedResponse: authErrorString, shouldLogin: false},
			{httpMethod: "PUT", url: "/api/v1/planner/department/1", expectedResponse: authErrorString, shouldLogin: false},
			{httpMethod: "DELETE", url: "/api/v1/planner/department/1", expectedResponse: authErrorString, shouldLogin: false},
			{httpMethod: "POST", url: "/api/v1/planner/department/1/workplace/", expectedResponse: authErrorString, shouldLogin: false},
			{httpMethod: "PUT", url: "/api/v1/planner/department/1/workplace/1", expectedResponse: authErrorString, shouldLogin: false},
			{httpMethod: "DELETE", url: "/api/v1/planner/department/1/workplace/1", expectedResponse: authErrorString, shouldLogin: false},
			{httpMethod: "POST", url: "/api/v1/planner/department/1/workplace/1/timeslot/", expectedResponse: authErrorString, shouldLogin: false},
			{httpMethod: "PUT", url: "/api/v1/planner/department/1/workplace/1/timeslot/1", expectedResponse: authErrorString, shouldLogin: false},
			{httpMethod: "DELETE", url: "/api/v1/planner/department/1/workplace/1/timeslot/1", expectedResponse: authErrorString, shouldLogin: false},
			{httpMethod: "POST", url: "/api/v1/planner/department/1/workplace/1/timeslot/1/weekday/", expectedResponse: authErrorString, shouldLogin: false},
			{httpMethod: "DELETE", url: "/api/v1/planner/department/1/workplace/1/timeslot/1/weekday/", expectedResponse: authErrorString, shouldLogin: false},

			// persons logged in
			{httpMethod: "GET", url: "/api/v1/planner/person/", expectedResponse: "{\"message\":\"GetAll\"}", shouldLogin: true},
			{httpMethod: "GET", url: "/api/v1/planner/person/", expectedResponse: "{\"message\":\"GetAll\"}", shouldLogin: false},
			{httpMethod: "GET", url: "/api/v1/planner/person/1", expectedResponse: "{\"message\":\"Get\"}", shouldLogin: true},
			{httpMethod: "GET", url: "/api/v1/planner/person/1", expectedResponse: "{\"message\":\"Get\"}", shouldLogin: false},

			{httpMethod: "POST", url: "/api/v1/planner/person/", expectedResponse: authErrorString, shouldLogin: false},
			{httpMethod: "POST", url: "/api/v1/planner/person/", expectedResponse: "{\"message\":\"Create\"}", shouldLogin: true},
			{httpMethod: "PUT", url: "/api/v1/planner/person/1", expectedResponse: authErrorString, shouldLogin: false},
			{httpMethod: "PUT", url: "/api/v1/planner/person/1", expectedResponse: "{\"message\":\"Update\"}", shouldLogin: true},
			{httpMethod: "DELETE", url: "/api/v1/planner/person/1", expectedResponse: authErrorString, shouldLogin: false},
			{httpMethod: "DELETE", url: "/api/v1/planner/person/1", expectedResponse: "{\"message\":\"Delete\"}", shouldLogin: true},

			{httpMethod: "GET", url: "/api/v1/planner/person/1/absency", expectedResponse: "{\"message\":\"FindAbsencyForPerson\"}", shouldLogin: false},
			{httpMethod: "GET", url: "/api/v1/planner/person/1/absency", expectedResponse: "{\"message\":\"FindAbsencyForPerson\"}", shouldLogin: true},
			{httpMethod: "DELETE", url: "/api/v1/planner/person/1/absency/1", expectedResponse: authErrorString, shouldLogin: false},
			{httpMethod: "DELETE", url: "/api/v1/planner/person/1/absency/1", expectedResponse: "{\"message\":\"RemoveAbsency\"}", shouldLogin: true},
			{httpMethod: "POST", url: "/api/v1/planner/person/1/absency", expectedResponse: authErrorString, shouldLogin: false},
			{httpMethod: "POST", url: "/api/v1/planner/person/1/absency", expectedResponse: "{\"message\":\"AddAbsency\"}", shouldLogin: true},

			{httpMethod: "GET", url: "/api/v1/planner/workday/", expectedResponse: "{\"message\":\"GetWorkdaysForDepartmentAndDate\"}", shouldLogin: false},
			{httpMethod: "GET", url: "/api/v1/planner/workday/", expectedResponse: "{\"message\":\"GetWorkdaysForDepartmentAndDate\"}", shouldLogin: true},

			{httpMethod: "PUT", url: "/api/v1/planner/workday/", expectedResponse: authErrorString, shouldLogin: false},
			{httpMethod: "PUT", url: "/api/v1/planner/workday/", expectedResponse: "{\"message\":\"UpdateWorkday\"}", shouldLogin: true},
			{httpMethod: "POST", url: "/api/v1/planner/workday/assign", expectedResponse: authErrorString, shouldLogin: false},
			{httpMethod: "POST", url: "/api/v1/planner/workday/assign", expectedResponse: "{\"message\":\"AssignPersonToWorkday\"}", shouldLogin: true},
			{httpMethod: "DELETE", url: "/api/v1/planner/workday/assign", expectedResponse: authErrorString, shouldLogin: false},
			{httpMethod: "DELETE", url: "/api/v1/planner/workday/assign", expectedResponse: "{\"message\":\"UnassignPersonFromWorkday\"}", shouldLogin: true},
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
				t.Errorf("URL: %s Expected body to be %v, got %v", testStep.url, testStep.expectedResponse, w.Body.String())
			}
			t.Logf("Test %v passed", i)
		}
	})

}
