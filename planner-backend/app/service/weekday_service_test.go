package service

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"planner-backend/app/domain/dao"
	"planner-backend/app/domain/dco"
	"planner-backend/app/domain/dto"
	"planner-backend/app/mock"
	"planner-backend/app/pkg"
	"testing"
)

func TestBulkUpdateWeekdaysForTimeslot(t *testing.T) {
	WeekdayRepository := mock.NewWeekdayRepositoryMock()
	TimeslotRepository := mock.NewTimeslotRepositoryMock()
	weekdayService := WeekdayServiceImpl{
		WeekdayRepository:  WeekdayRepository,
		TimeslotRepository: TimeslotRepository,
	}

	testSteps := []ServiceTestPOST{
		{
			mockRequestData: map[string]interface{}{
				"weekdays": []map[string]interface{}{
					{
						"id":         1,
						"start_time": "08:00",
						"end_time":   "09:00",
					},
				},
			},
			findValue: dao.Timeslot{
				Name: "test",
			},
			saveValue: []dao.OnWeekday{
				{
					ID: 1,
				},
			},
			expectedStatusCode: http.StatusCreated,
			params: map[string]string{
				"departmentID": "test",
				"workplaceID":  "test",
				"timeslotID":   "test",
			},
		},
		{
			mockRequestData:    map[string]interface{}{},
			findValue:          nil,
			saveValue:          nil,
			expectedStatusCode: http.StatusBadRequest,
			findError:          nil,
			saveError:          pkg.ErrNoRows,
		},
		{
			mockRequestData: map[string]interface{}{
				"weekdays": []map[string]interface{}{
					{
						"id":         1,
						"start_time": "08:00",
						"end_time":   "09:00",
					},
				},
			},
			findValue:          nil,
			saveValue:          nil,
			expectedStatusCode: http.StatusBadRequest,
			findError:          pkg.ErrNoRows,
			saveError:          nil,
			params: map[string]string{
				"departmentID": "test",
				"workplaceID":  "test",
				"timeslotID":   "test",
			},
		},
		// Add more test steps here...
	}

	for i, testStep := range testSteps {
		t.Run("Test Bulk Update Weekdays For Timeslot", func(t *testing.T) {
			TimeslotRepository.On("FindTimeslotByID").Return(testStep.findValue, testStep.findError)
			WeekdayRepository.On("DeleteAllWeekdaysFromTimeslot").Return(nil, testStep.additionalError)
			WeekdayRepository.On("AddWeekdaysToTimeslot").Return(testStep.saveValue, testStep.saveError)

			// get GIN context
			w := httptest.NewRecorder()
			c, err := mock.NewTestContextBuilder(w).WithMethod("POST").WithMapParams(testStep.params).WithBody(testStep.mockRequestData).Build()
			if err != nil {
				t.Errorf("Test Step %d: Error when getting GIN context", i)
			}

			weekdayService.BulkUpdateWeekdaysForTimeslot(c)
			response := w.Result()

			if response.StatusCode != testStep.expectedStatusCode {
				t.Errorf("Test Step %d: Expected status code %d, got %d", i, testStep.expectedStatusCode, response.StatusCode)
			}

			if response.StatusCode != http.StatusCreated {
				return
			}

			var responseBody dto.APIResponse[[]dco.WeekdayResponse]
			if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
				t.Errorf("Test Step %d: Error when decoding response body", i)
			}

			for _, weekday := range responseBody.Data {
				if weekday.ID != testStep.saveValue.([]dao.OnWeekday)[0].ID {
					t.Errorf("Test Step %d: Expected response body %v, got %v", i, testStep.saveValue, responseBody.Data)
				}
			}
		})
	}
}

func TestUpdateWeekdayForTimeslot(t *testing.T) {
	WeekdayRepository := mock.NewWeekdayRepositoryMock()
	TimeslotRepository := mock.NewTimeslotRepositoryMock()
	weekdayService := WeekdayServiceImpl{
		WeekdayRepository:  WeekdayRepository,
		TimeslotRepository: TimeslotRepository,
	}

	testSteps := []ServiceTestPUT{
		{
			// valid
			mockRequestData: map[string]interface{}{
				"id":         1,
				"start_time": "08:00",
				"end_time":   "09:00",
			},
			findValue: dao.Timeslot{
				Name: "test",
			},
			saveValue: []dao.OnWeekday{
				{
					ID: 1,
				},
			},
			expectedStatusCode: http.StatusOK,
			findError:          nil,
			saveError:          nil,
			params: map[string]string{
				"departmentID": "test",
				"workplaceID":  "test",
				"timeslotID":   "test",
			},
		},
		{
			mockRequestData: map[string]interface{}{
				"id":         1,
				"start_time": "08:00",
				"end_time":   "09:00",
			},
			findValue:          nil,
			saveValue:          nil,
			expectedStatusCode: http.StatusBadRequest,
			findError:          pkg.ErrNoRows,
			saveError:          nil,
			params: map[string]string{
				"departmentID": "test",
				"workplaceID":  "test",
				"timeslotID":   "test",
			},
		},
		{
			mockRequestData: map[string]interface{}{
				"id":         1,
				"start_time": "16:00",
				"end_time":   "09:00",
			},
			findValue: nil,
			saveValue: []dao.OnWeekday{
				{
					ID: 1,
				},
			},
			expectedStatusCode: http.StatusBadRequest,
			findError:          pkg.ErrNoRows,
			saveError:          nil,
		},
		{
			mockRequestData: map[string]interface{}{
				"id":         1,
				"start_time": "08:00",
				"end_time":   "09:00",
			},
			findValue: nil,
			saveValue: []dao.OnWeekday{
				{
					ID: 1,
				},
			},
			expectedStatusCode: http.StatusBadRequest,
			findError:          pkg.ErrNoRows,
			saveError:          nil,
		},
		{
			mockRequestData:    map[string]interface{}{},
			findValue:          nil,
			saveValue:          nil,
			expectedStatusCode: http.StatusBadRequest,
			findError:          pkg.ErrNoRows,
			saveError:          nil,
			params: map[string]string{
				"departmentID": "test",
				"workplaceID":  "test",
				"timeslotID":   "test",
			},
		},
		{
			mockRequestData:    map[string]interface{}{},
			findValue:          nil,
			saveValue:          nil,
			expectedStatusCode: http.StatusBadRequest,
			findError:          nil,
			saveError:          pkg.ErrNoRows,
		},
	}

	for i, testStep := range testSteps {
		t.Run("Test Update Weekday For Timeslot", func(t *testing.T) {
			TimeslotRepository.On("FindTimeslotByID").Return(testStep.findValue, testStep.findError)
			WeekdayRepository.On("UpdateWeekdayForTimeslot").Return(testStep.saveValue, testStep.saveError)

			// get GIN context
			w := httptest.NewRecorder()
			c, err := mock.NewTestContextBuilder(w).WithMethod("PUT").WithMapParams(testStep.params).WithBody(testStep.mockRequestData).Build()
			if err != nil {
				t.Errorf("Test Step %d: Error when getting GIN context", i)
			}

			weekdayService.UpdateWeekdayForTimeslot(c)
			response := w.Result()

			if response.StatusCode != testStep.expectedStatusCode {
				t.Errorf("Test Step %d: Expected status code %d, got %d", i, testStep.expectedStatusCode, response.StatusCode)
			}

			if response.StatusCode != http.StatusOK {
				return
			}

			var responseBody dto.APIResponse[[]dco.OnWeekdayResponse]
			if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
				t.Errorf("Test Step %d: Error when decoding response body", i)
			}

			for _, weekday := range responseBody.Data {
				if weekday.ID != testStep.saveValue.([]dao.OnWeekday)[0].ID {
					t.Errorf("Test Step %d: Expected response body %v, got %v", i, testStep.saveValue, responseBody.Data)
				}
			}
		})
	}
}

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
				"id":         1,
				"start_time": "08:00",
				"end_time":   "09:00",
			},
			findValue: dao.Timeslot{
				Name: "test",
			},
			saveValue: []dao.OnWeekday{
				{
					ID: 1,
				},
			},
			expectedStatusCode: http.StatusCreated,
			findError:          nil,
			saveError:          nil,
			params: map[string]string{
				"departmentID": "test",
				"workplaceID":  "test",
				"timeslotID":   "test",
			},
		},
		{
			mockRequestData: map[string]interface{}{
				"id":         1,
				"start_time": "08:00",
				"end_time":   "09:00",
			},
			findValue:          nil,
			saveValue:          nil,
			expectedStatusCode: http.StatusBadRequest,
			findError:          pkg.ErrNoRows,
			saveError:          nil,
			params: map[string]string{
				"departmentID": "test",
				"workplaceID":  "test",
				"timeslotID":   "test",
			},
		},
		{
			mockRequestData: map[string]interface{}{
				"id":         1,
				"start_time": "08:00",
				"end_time":   "09:00",
			},
			findValue: nil,
			saveValue: []dao.OnWeekday{
				{
					ID: 1,
				},
			},
			expectedStatusCode: http.StatusBadRequest,
			findError:          pkg.ErrNoRows,
			saveError:          nil,
		},
		{
			mockRequestData:    map[string]interface{}{},
			findValue:          nil,
			saveValue:          nil,
			expectedStatusCode: http.StatusBadRequest,
			findError:          pkg.ErrNoRows,
			saveError:          nil,
			params: map[string]string{
				"departmentID": "test",
				"workplaceID":  "test",
				"timeslotID":   "test",
			},
		},
		{
			mockRequestData:    map[string]interface{}{},
			findValue:          nil,
			saveValue:          nil,
			expectedStatusCode: http.StatusBadRequest,
			findError:          nil,
			saveError:          pkg.ErrNoRows,
		},
	}

	for i, testStep := range testSteps {
		t.Run("Test Add Weekday To Timeslot", func(t *testing.T) {
			TimeslotRepository.On("FindTimeslotByID").Return(testStep.findValue, testStep.findError)
			WeekdayRepository.On("AddWeekdayToTimeslot").Return(testStep.saveValue, testStep.saveError)

			// get GIN context
			w := httptest.NewRecorder()
			c, err := mock.NewTestContextBuilder(w).WithMethod("POST").WithMapParams(testStep.params).WithBody(testStep.mockRequestData).Build()
			if err != nil {
				t.Errorf("Test Step %d: Error when getting GIN context", i)
			}

			weekdayService.AddWeekdayToTimeslot(c)
			response := w.Result()

			if response.StatusCode != testStep.expectedStatusCode {
				t.Errorf("Test Step %d: Expected status code %d, got %d", i, testStep.expectedStatusCode, response.StatusCode)
			}

			if response.StatusCode != http.StatusCreated {
				return
			}

			var responseBody dto.APIResponse[[]dco.OnWeekdayResponse]
			if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
				t.Errorf("Test Step %d: Error when decoding response body", i)
			}

			for _, weekday := range responseBody.Data {
				if weekday.ID != testStep.saveValue.([]dao.OnWeekday)[0].ID {
					t.Errorf("Test Step %d: Expected response body %v, got %v", i, testStep.saveValue, responseBody.Data)
				}
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
				"id": 1,
			},
			expectedStatusCode: http.StatusOK,
			findValue: dao.OnWeekday{
				ID: 1,
			},
			findError: nil,
			params: map[string]string{
				"departmentID": "test",
				"workplaceID":  "test",
				"timeslotID":   "test",
				"weekdayName":  "test",
			},
		},
		{ // Test 2
			mockRequestData:    map[string]interface{}{},
			expectedStatusCode: http.StatusBadRequest,
			findValue:          nil,
			findError:          pkg.ErrNoRows,
			params: map[string]string{
				"departmentID": "test",
				"workplaceID":  "test",
				"timeslotID":   "test",
				"weekdayName":  "test",
			},
		},
		{ // Test 3
			expectedStatusCode: http.StatusBadRequest,
			findValue:          nil,
			findError:          pkg.ErrNoRows,
			params: map[string]string{
				"departmentID": "test",
				"workplaceID":  "test",
				"timeslotID":   "test",
				"weekdayName":  "test",
			},
		},
	}

	for i, testStep := range testSteps {
		t.Run("Test Delete Weekday From Timeslot", func(t *testing.T) {
			WeekdayRepository.On("DeleteWeekdayFromTimeslot").Return(testStep.findValue, testStep.findError)

			// get GIN context
			w := httptest.NewRecorder()
			c, err := mock.NewTestContextBuilder(w).WithMethod("DELETE").WithMapParams(testStep.params).WithBody(testStep.mockRequestData).Build()
			if err != nil {
				t.Errorf("Test Step %d: Error when getting GIN context", i)
			}

			weekdayService.DeleteWeekdayFromTimeslot(c)
			response := w.Result()

			if response.StatusCode != testStep.expectedStatusCode {
				t.Errorf("Test Step %d: Expected status code %d, got %d", i, testStep.expectedStatusCode, response.StatusCode)
			}
		})
	}
}
