package service

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"planner-backend/app/domain/dao"
	"planner-backend/app/domain/dco"
	"planner-backend/app/domain/dto"
	"planner-backend/app/mock"
	"planner-backend/app/pkg"
	"testing"
)

func TestDeleteTimeslot(t *testing.T) {
	TimeslotRepository := mock.NewTimeslotRepositoryMock()
	timeslotService := TimeslotServiceImpl{
		TimeslotRepository: TimeslotRepository,
	}

	testSteps := []ServiceTestDELETE{
		{
			mockValue: dao.Timeslot{
				Name: "test",
			},
			params: map[string]string{
				"departmentName": "test",
				"workplaceName":  "test",
				"timeslotName":   "test",
			},
			mockError:          nil,
			expectedStatusCode: http.StatusOK,
		},
		{
			mockValue: dao.Timeslot{
				Name: "test",
			},
			params: map[string]string{
				"departmentName": "test",
				"workplaceName":  "test",
			},
			mockError:          pkg.ErrNoRows,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			mockValue: dao.Timeslot{
				Name: "test",
			},
			params: map[string]string{
				"departmentName": "test",
				"timeslotName":   "test",
			},
			mockError:          pkg.ErrNoRows,
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for i, testStep := range testSteps {
		TimeslotRepository.On("Delete").Return(testStep.mockValue, testStep.mockError)

		// get GIN context
		w := httptest.NewRecorder()
		c := mock.GetGinTestContext(w, "DELETE", testStep.ParamsToGinParams(), nil)

		timeslotService.DeleteTimeslot(c)
		response := w.Result()

		if response.StatusCode != testStep.expectedStatusCode {
			t.Errorf("Test Step %d: Expected status code %d, got %d", i, testStep.expectedStatusCode, response.StatusCode)
		}
	}
}

func TestUpdateTimeslot(t *testing.T) {
	TimeslotRepository := mock.NewTimeslotRepositoryMock()
	timeslotService := TimeslotServiceImpl{
		TimeslotRepository: TimeslotRepository,
	}

	testSteps := []ServiceTestPUT{
		{
			mockRequestData: map[string]interface{}{
				"name":   "newName",
				"active": true,
			},
			findValue: dao.Timeslot{
				Name: "oldName",
			},
			saveValue: dao.Timeslot{
				Name: "newName",
			},
			expectedStatusCode: http.StatusOK,
			findError:          nil,
			saveError:          nil,
			params: map[string]string{
				"departmentName": "test",
				"workplaceName":  "test",
				"timeslotName":   "oldName",
			},
		},
		{
			mockRequestData: map[string]interface{}{
				"name":   "newName",
				"active": false,
			},
			findValue: dao.Timeslot{
				Name: "oldName",
			},
			saveValue: dao.Timeslot{
				Name: "newName",
			},
			expectedStatusCode: http.StatusOK,
			findError:          nil,
			saveError:          nil,
			params: map[string]string{
				"departmentName": "test",
				"workplaceName":  "test",
				"timeslotName":   "oldName",
			},
		},
		{
			mockRequestData: map[string]interface{}{
				"name": "newName",
			},
			findValue: dao.Timeslot{
				Name: "oldName",
			},
			saveValue: dao.Timeslot{
				Name: "newName",
			},
			expectedStatusCode: http.StatusBadRequest,
			findError:          nil,
			saveError:          nil,
			params: map[string]string{
				"departmentName": "test",
				"workplaceName":  "test",
				"timeslotName":   "oldName",
			},
		},
		{
			mockRequestData: map[string]interface{}{
				"name": "newName",
			},
			findValue:          nil,
			saveValue:          nil,
			expectedStatusCode: http.StatusNotFound,
			findError:          pkg.ErrNoRows,
			saveError:          nil,
			params: map[string]string{
				"departmentName": "test",
				"workplaceName":  "test",
				"timeslotName":   "oldName",
			},
		},
		{
			mockRequestData: map[string]interface{}{
				"name": "newName",
			},
			findValue: dao.Timeslot{
				Name: "oldName",
			},
			saveValue: dao.Timeslot{
				Name: "newName",
			},
			expectedStatusCode: 500,
			findError:          nil,
			saveError:          errors.New("Save error"),
		},
		{
			mockRequestData:    map[string]interface{}{},
			findValue:          nil,
			saveValue:          nil,
			expectedStatusCode: http.StatusBadRequest,
			findError:          nil,
			saveError:          nil,
		},
	}

	for i, testStep := range testSteps {
		TimeslotRepository.On("FindTimeslotByName").Return(testStep.findValue, testStep.findError)
		TimeslotRepository.On("Save").Return(testStep.saveValue, testStep.saveError)

		// get GIN context
		w := httptest.NewRecorder()
		c := mock.GetGinTestContext(w, "PUT", testStep.ParamsToGinParams(), testStep.mockRequestData)

		timeslotService.UpdateTimeslot(c)
		response := w.Result()

		if response.StatusCode != testStep.expectedStatusCode {
			t.Errorf("Test Step %d: Expected status code %d, got %d", i, testStep.expectedStatusCode, response.StatusCode)
		}

		if response.StatusCode != http.StatusOK {
			return
		}

		var responseBody dto.APIResponse[dco.TimeslotResponse]
		if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
			t.Errorf("Test Step %d: Error when decoding response body", i)
		}

		if responseBody.Data.Name != testStep.saveValue.(dao.Timeslot).Name {
			t.Errorf("Test Step %d: Expected response body %v, got %v", i, testStep.saveValue, responseBody.Data)
		}
	}
}

func TestAddTimeslot(t *testing.T) {
	TimeslotRepository := mock.NewTimeslotRepositoryMock()
	timeslotService := TimeslotServiceImpl{
		TimeslotRepository: TimeslotRepository,
	}

	testSteps := []ServiceTestPOST{
		{
			mockRequestData: map[string]interface{}{
				"name":   "test",
				"active": true,
			},
			findValue: nil,
			saveValue: dao.Timeslot{
				Name: "test",
			},
			expectedStatusCode: http.StatusCreated,
			findError:          pkg.ErrNoRows,
			saveError:          nil,
			params: map[string]string{
				"departmentName": "test",
				"workplaceName":  "test",
			},
		},
		{
			mockRequestData: map[string]interface{}{
				"name":   "test",
				"active": false,
			},
			findValue: nil,
			saveValue: dao.Timeslot{
				Name: "test",
			},
			expectedStatusCode: http.StatusCreated,
			findError:          pkg.ErrNoRows,
			saveError:          nil,
			params: map[string]string{
				"departmentName": "test",
				"workplaceName":  "test",
			},
		},
		{
			mockRequestData: map[string]interface{}{
				"name": "test",
			},
			findValue: nil,
			saveValue: dao.Timeslot{
				Name: "test",
			},
			expectedStatusCode: http.StatusBadRequest,
			findError:          pkg.ErrNoRows,
			saveError:          nil,
			params: map[string]string{
				"departmentName": "test",
				"workplaceName":  "test",
			},
		},
		{
			mockRequestData: map[string]interface{}{
				"name":   "test",
				"active": true,
			},
			findValue: dao.Timeslot{
				Name: "test",
			},
			saveValue:          nil,
			expectedStatusCode: http.StatusConflict,
			findError:          nil,
			saveError:          nil,
			params: map[string]string{
				"departmentName": "test",
				"workplaceName":  "test",
			},
		},
		{
			mockRequestData: map[string]interface{}{
				"name":   "test",
				"active": true,
			},
			findValue: nil,
			saveValue: dao.Timeslot{
				Name: "test",
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
				"departmentName": "test",
				"workplaceName":  "test",
			},
		},
		{
			mockRequestData: map[string]interface{}{
				"name":   "test",
				"active": true,
			},
			findValue:          nil,
			saveValue:          nil,
			expectedStatusCode: 500,
			findError:          pkg.ErrNoRows,
			saveError:          pkg.ErrNoRows,
			params: map[string]string{
				"departmentName": "test",
				"workplaceName":  "test",
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
		t.Run("Test Add Timeslot", func(t *testing.T) {
			TimeslotRepository.On("FindTimeslotByName").Return(testStep.findValue, testStep.findError)
			TimeslotRepository.On("Save").Return(testStep.saveValue, testStep.saveError)

			// get GIN context
			w := httptest.NewRecorder()
			c := mock.GetGinTestContext(w, "POST", testStep.ParamsToGinParams(), testStep.mockRequestData)

			timeslotService.AddTimeslot(c)
			response := w.Result()

			if response.StatusCode != testStep.expectedStatusCode {
				t.Errorf("Test Step %d: Expected status code %d, got %d", i, testStep.expectedStatusCode, response.StatusCode)
			}

			if response.StatusCode != 201 {
				return
			}

			var responseBody dto.APIResponse[dco.TimeslotResponse]
			if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
				t.Errorf("Test Step %d: Error when decoding response body", i)
			}

			if responseBody.Data.Name != testStep.saveValue.(dao.Timeslot).Name {
				t.Errorf("Test Step %d: Expected response body %v, got %v", i, testStep.saveValue, responseBody.Data)
			}
		})
	}
}

func TestGetAllTimeslots(t *testing.T) {
	TimeslotRepository := mock.NewTimeslotRepositoryMock()
	timeslotService := TimeslotServiceImpl{
		TimeslotRepository: TimeslotRepository,
	}

	testSteps := []ServiceTestGET{
		{
			mockValue: []dao.Timeslot{
				{
					Name: "test",
				},
			},
			expectedResponse: []dco.TimeslotResponse{
				{
					Name: "test",
				},
			},
			expectedStatusCode: http.StatusOK,
			mockError:          nil,
			params: map[string]string{
				"departmentName": "test",
				"workplaceName":  "test",
			},
		},
		{
			mockValue:          nil,
			expectedResponse:   nil,
			expectedStatusCode: http.StatusBadRequest,
			mockError:          nil,
		},
		{
			mockValue:          nil,
			expectedResponse:   nil,
			expectedStatusCode: http.StatusOK,
			mockError:          pkg.ErrNoRows,
			params: map[string]string{
				"departmentName": "test",
				"workplaceName":  "test",
			},
		},
	}

	for i, testStep := range testSteps {
		t.Run("Test Get All Timeslots", func(t *testing.T) {
			TimeslotRepository.On("FindAllTimeslots").Return(testStep.mockValue, testStep.mockError)

			// get GIN context
			w := httptest.NewRecorder()
			c := mock.GetGinTestContext(w, "GET", testStep.ParamsToGinParams(), nil)

			timeslotService.GetAllTimeslots(c)
			response := w.Result()

			if response.StatusCode != testStep.expectedStatusCode {
				t.Errorf("Test Step %d: Expected status code %d, got %d", i, testStep.expectedStatusCode, response.StatusCode)
			}

			if response.StatusCode != http.StatusOK {
				return
			}

			var responseBody dto.APIResponse[[]dco.TimeslotResponse]
			if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
				t.Errorf("Test Step %d: Error when decoding response body", i)
			}

			for i, timeslot := range responseBody.Data {
				if timeslot.Name != testStep.expectedResponse.([]dco.TimeslotResponse)[i].Name {
					t.Errorf("Test Step %d: Expected response body %v, got %v", i, testStep.expectedResponse, responseBody.Data)
				}
			}
		})
	}
}

func TestGetTimeslotByName(t *testing.T) {
	TimeslotRepository := mock.NewTimeslotRepositoryMock()
	timeslotService := TimeslotServiceImpl{
		TimeslotRepository: TimeslotRepository,
	}

	testSteps := []ServiceTestGET{
		{
			mockValue: dao.Timeslot{
				Name: "test",
			},
			expectedResponse: dao.Timeslot{
				Name: "test",
			},
			expectedStatusCode: http.StatusOK,
			mockError:          nil,
			params: map[string]string{
				"timeslotName":   "test",
				"departmentName": "test",
				"workplaceName":  "test",
			},
		},
		{
			mockValue: dao.Timeslot{
				Name: "test",
			},
			expectedResponse: dao.Timeslot{
				Name: "test",
			},
			expectedStatusCode: http.StatusBadRequest,
			mockError:          nil,
			params: map[string]string{
				"departmentName": "test",
				"workplaceName":  "test",
			},
		},
		{
			mockValue: dao.Timeslot{
				Name: "test",
			},
			expectedResponse: dao.Timeslot{
				Name: "test",
			},
			expectedStatusCode: http.StatusBadRequest,
			mockError:          nil,
			params: map[string]string{
				"timeslotName":   "test",
				"departmentName": "test",
			},
		},
		{
			mockValue:          nil,
			expectedResponse:   nil,
			expectedStatusCode: http.StatusBadRequest,
			mockError:          nil,
			params:             map[string]string{},
		},
		{
			mockValue:          nil,
			expectedResponse:   nil,
			expectedStatusCode: http.StatusNotFound,
			mockError:          pkg.ErrNoRows,
			params: map[string]string{
				"timeslotName":   "test",
				"departmentName": "test",
				"workplaceName":  "test",
			},
		},
	}

	for i, testStep := range testSteps {
		t.Run("Test Get Timeslot By Name", func(t *testing.T) {
			TimeslotRepository.On("FindTimeslotByName").Return(testStep.mockValue, testStep.mockError)

			// get GIN context
			w := httptest.NewRecorder()
			c := mock.GetGinTestContext(w, "GET", testStep.ParamsToGinParams(), nil)

			timeslotService.GetTimeslotByName(c)
			response := w.Result()

			if response.StatusCode != testStep.expectedStatusCode {
				t.Errorf("Test Step %d: Expected status code %d, got %d", i, testStep.expectedStatusCode, response.StatusCode)
			}

			if response.StatusCode != http.StatusOK {
				return
			}

			var responseBody dto.APIResponse[dco.TimeslotResponse]
			if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
				t.Errorf("Test Step %d: Error when decoding response body", i)
			}

			if responseBody.Data.Name != testStep.expectedResponse.(dao.Timeslot).Name {
				t.Errorf("Test Step %d: Expected response body %v, got %v", i, testStep.expectedResponse, responseBody.Data)
			}
		})
	}
}
