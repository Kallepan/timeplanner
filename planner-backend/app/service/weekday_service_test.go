package service

import (
	"encoding/json"
	"net/http/httptest"
	"planner-backend/app/domain/dao"
	"planner-backend/app/domain/dco"
	"planner-backend/app/domain/dto"
	"planner-backend/app/mock"
	"planner-backend/app/pkg"
	"testing"
)

func TestAddWeekdayToTimeslot(t *testing.T) {
	WeekdayRepository := mock.NewWeekdayRepositoryMock()
	TimeslotRepository := mock.NewTimeslotRepositoryMock()
	weekdayService := WeekdayServiceImpl{
		WeekdayRepository:  WeekdayRepository,
		TimeslotRepository: TimeslotRepository,
	}

	testSteps := []ServiceTestPOST{
		{
			mockRequestData: map[string]interface{}{
				"name":       "test",
				"start_time": "08:00",
				"end_time":   "09:00",
			},
			findValue: dao.Timeslot{
				Name: "test",
			},
			saveValue: []dao.OnWeekday{
				{
					ID: "MON",
				},
			},
			expectedStatusCode: 200,
			findError:          nil,
			saveError:          nil,
			queryParams: map[string]string{
				"departmentName": "test",
				"workplaceName":  "test",
				"timeslotName":   "test",
			},
		},
		{
			mockRequestData: map[string]interface{}{
				"name":       "test",
				"start_time": "08:00",
				"end_time":   "09:00",
			},
			findValue:          nil,
			saveValue:          nil,
			expectedStatusCode: 400,
			findError:          pkg.ErrNoRows,
			saveError:          nil,
			queryParams: map[string]string{
				"departmentName": "test",
				"workplaceName":  "test",
				"timeslotName":   "test",
			},
		},
		{
			mockRequestData: map[string]interface{}{
				"name":       "test",
				"start_time": "08:00",
				"end_time":   "09:00",
			},
			findValue: nil,
			saveValue: []dao.OnWeekday{
				{
					ID: "MON",
				},
			},
			expectedStatusCode: 400,
			findError:          pkg.ErrNoRows,
			saveError:          nil,
		},
		{
			mockRequestData:    map[string]interface{}{},
			findValue:          nil,
			saveValue:          nil,
			expectedStatusCode: 400,
			findError:          pkg.ErrNoRows,
			saveError:          nil,
			queryParams: map[string]string{
				"departmentName": "test",
				"workplaceName":  "test",
				"timeslotName":   "test",
			},
		},
		{
			mockRequestData:    map[string]interface{}{},
			findValue:          nil,
			saveValue:          nil,
			expectedStatusCode: 400,
			findError:          nil,
			saveError:          pkg.ErrNoRows,
		},
	}

	for i, testStep := range testSteps {
		t.Run("Test Add Weekday To Timeslot", func(t *testing.T) {
			TimeslotRepository.On("FindTimeslotByName").Return(testStep.findValue, testStep.findError)
			WeekdayRepository.On("AddWeekdayToTimeslot").Return(testStep.saveValue, testStep.saveError)

			// get GIN context
			w := httptest.NewRecorder()
			c := mock.GetGinTestContextWithBody(w, "POST", testStep.QueryParamsToGinParams(), testStep.mockRequestData)

			weekdayService.AddWeekdayToTimeslot(c)
			response := w.Result()

			if response.StatusCode != testStep.expectedStatusCode {
				t.Errorf("Test Step %d: Expected status code %d, got %d", i, testStep.expectedStatusCode, response.StatusCode)
			}

			if response.StatusCode != 201 {
				return
			}

			var responseBody dto.APIResponse[dco.WeekdayResponse]
			if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
				t.Errorf("Test Step %d: Error when decoding response body", i)
			}

			if responseBody.Data.ID != testStep.saveValue.(dao.OnWeekday).ID {
				t.Errorf("Test Step %d: Expected response body %v, got %v", i, testStep.saveValue, responseBody.Data)
			}
		})
	}
}

func TestDeleteWeekdayFromTimeslot(t *testing.T) {
	WeekdayRepository := mock.NewWeekdayRepositoryMock()
	TimeslotRepository := mock.NewTimeslotRepositoryMock()
	weekdayService := WeekdayServiceImpl{
		WeekdayRepository:  WeekdayRepository,
		TimeslotRepository: TimeslotRepository,
	}

	testSteps := []ServiceTestPOST{
		{ // Test 1
			mockRequestData: map[string]interface{}{
				"name": "test",
			},
			expectedStatusCode: 200,
			findValue: dao.OnWeekday{
				ID: "test",
			},
			findError: nil,
			queryParams: map[string]string{
				"departmentName": "test",
				"workplaceName":  "test",
				"timeslotName":   "test",
				"weekdayName":    "test",
			},
		},
		{ // Test 2
			mockRequestData:    map[string]interface{}{},
			expectedStatusCode: 400,
			findValue:          nil,
			findError:          pkg.ErrNoRows,
			queryParams: map[string]string{
				"departmentName": "test",
				"workplaceName":  "test",
				"timeslotName":   "test",
				"weekdayName":    "test",
			},
		},
		{ // Test 3
			expectedStatusCode: 400,
			findValue:          nil,
			findError:          pkg.ErrNoRows,
			queryParams: map[string]string{
				"departmentName": "test",
				"workplaceName":  "test",
				"timeslotName":   "test",
				"weekdayName":    "test",
			},
		},
	}

	for i, testStep := range testSteps {
		t.Run("Test Delete Weekday From Timeslot", func(t *testing.T) {
			WeekdayRepository.On("DeleteWeekdayFromTimeslot").Return(testStep.findValue, testStep.findError)

			// get GIN context
			w := httptest.NewRecorder()
			c := mock.GetGinTestContextWithBody(w, "DELETE", testStep.QueryParamsToGinParams(), testStep.mockRequestData)

			weekdayService.DeleteWeekdayFromTimeslot(c)
			response := w.Result()

			if response.StatusCode != testStep.expectedStatusCode {
				t.Errorf("Test Step %d: Expected status code %d, got %d", i, testStep.expectedStatusCode, response.StatusCode)
			}
		})
	}
}
