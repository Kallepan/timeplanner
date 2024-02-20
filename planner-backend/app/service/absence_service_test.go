package service

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"planner-backend/app/domain/dao"
	"planner-backend/app/mock"
	"testing"
	"time"
)

func TestGetAllAbsencies(t *testing.T) {
	absenceMockRepo := mock.NewAbsenceRepositoryMock()
	absenceService := AbsenceServiceImpl{
		AbsenceRepository: absenceMockRepo,
	}

	testSteps := []ServiceTestGET{
		{
			params: map[string]string{
				"departmentID": "1",
			},
			queries: map[string]string{
				"date": "2021-01-01",
			},
			expectedStatusCode: http.StatusOK,
			mockValue: []dao.Absence{
				{
					PersonID:  "1",
					Date:      "2021-01-01",
					Reason:    "Sick",
					CreatedAt: time.Now(),
				},
			},
			mockError: nil,
		},
		{
			params: map[string]string{},
			queries: map[string]string{
				"date": "2021-01-01",
			},
			expectedStatusCode: http.StatusBadRequest,
			mockValue:          []dao.Absence{},
			mockError:          nil,
		},
		{
			params: map[string]string{
				"departmentID": "1",
			},
			queries:            map[string]string{},
			expectedStatusCode: http.StatusBadRequest,
			mockValue:          []dao.Absence{},
			mockError:          nil,
		},
		{
			params: map[string]string{
				"departmentID": "1",
			},
			queries: map[string]string{
				"date": "2021-01",
			},
			expectedStatusCode: http.StatusBadRequest,
			mockValue:          []dao.Absence{},
			mockError:          nil,
		},
	}

	for i, testStep := range testSteps {
		t.Run(fmt.Sprintf("Running TestGetAllAbsencies. Step: %d", i), func(t *testing.T) {
			w := httptest.NewRecorder()
			c, err := mock.NewTestContextBuilder(w).WithMethod("GET").WithQueries(testStep.queries).WithMapParams(testStep.params).Build()
			if err != nil {
				t.Errorf("Error while creating test context: %v", err)
			}

			absenceService.GetAllAbsencies(c)

			if w.Code != testStep.expectedStatusCode {
				t.Errorf("Expected status code: %d, got: %d", testStep.expectedStatusCode, w.Code)
			}
		})
	}
}
