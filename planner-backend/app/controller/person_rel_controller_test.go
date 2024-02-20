package controller

import (
	"fmt"
	"net/http/httptest"
	"planner-backend/app/mock"
	"testing"

	"github.com/gin-gonic/gin"
)

type MockPersonRelService struct {
	Called map[string]bool
}

func (m *MockPersonRelService) AddAbsencyToPerson(ctx *gin.Context) {
	m.Called["AddAbsencyToPerson"] = true
}

func (m *MockPersonRelService) RemoveAbsencyFromPerson(ctx *gin.Context) {
	m.Called["RemoveAbsencyFromPerson"] = true
}

func (m *MockPersonRelService) FindAbsencyForPerson(ctx *gin.Context) {
	m.Called["FindAbsencyForPerson"] = true
}

func (m *MockPersonRelService) FindAbsencyForPersonInRange(ctx *gin.Context) {
	m.Called["FindAbsencyForPersonInRange"] = true
}

func (m *MockPersonRelService) AddDepartmentToPerson(ctx *gin.Context) {
	m.Called["AddDepartmentToPerson"] = true
}

func (m *MockPersonRelService) RemoveDepartmentFromPerson(ctx *gin.Context) {
	m.Called["RemoveDepartmentFromPerson"] = true
}

func (m *MockPersonRelService) AddWorkplaceToPerson(ctx *gin.Context) {
	m.Called["AddWorkplaceToPerson"] = true
}

func (m *MockPersonRelService) RemoveWorkplaceFromPerson(ctx *gin.Context) {
	m.Called["RemoveWorkplaceFromPerson"] = true
}

func (m *MockPersonRelService) AddWeekdayToPerson(ctx *gin.Context) {
	m.Called["AddWeekdayToPerson"] = true
}

func (m *MockPersonRelService) RemoveWeekdayFromPerson(ctx *gin.Context) {
	m.Called["RemoveWeekdayFromPerson"] = true
}

func NewMockPersonRelService() *MockPersonRelService {
	return &MockPersonRelService{Called: make(map[string]bool)}
}

func TestFindAbsencyForPerson(t *testing.T) {
	tests := []struct {
		params        map[string]string
		queries       map[string]string
		expecteCalled map[string]bool
	}{
		{
			params: map[string]string{
				"person_id": "1",
			},
			queries: map[string]string{
				"date": "2021-01-01",
			},
			expecteCalled: map[string]bool{"FindAbsencyForPerson": true},
		},
		{
			params: map[string]string{
				"person_id": "1",
			},
			queries:       map[string]string{},
			expecteCalled: map[string]bool{}, // no call
		},
		{
			params: map[string]string{
				"person_id": "1",
			},
			queries: map[string]string{
				"start_date": "2021-01-01",
				"end_date":   "2021-01-02",
			},
			expecteCalled: map[string]bool{"FindAbsencyForPersonInRange": true},
		},
	}

	for i, testStep := range tests {
		t.Run(fmt.Sprintf("Test Step %d", i), func(t *testing.T) {
			mockPersonRelService := NewMockPersonRelService()
			controller := PersonRelControllerImpl{PersonRelService: mockPersonRelService}

			w := httptest.NewRecorder()
			c, err := mock.NewTestContextBuilder(w).WithMapParams(testStep.params).WithQueries(testStep.queries).Build()
			if err != nil {
				t.Errorf("Error building test context: %v", err)
			}

			controller.FindAbsencyForPerson(c)

			for k, v := range testStep.expecteCalled {
				if mockPersonRelService.Called[k] != v {
					t.Errorf("Expected %s to be called: %v", k, v)
				}
			}

			for k, v := range mockPersonRelService.Called {
				if _, ok := testStep.expecteCalled[k]; !ok {
					t.Errorf("Unexpected call to %s: %v", k, v)
				}
			}
		})
	}
}
