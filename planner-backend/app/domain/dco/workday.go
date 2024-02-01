package dco

import (
	"errors"
	"time"
)

/* Requests */
type UpdateWorkdayRequest struct {
	StartTime string `json:"start_time" binding:"required"`
	EndTime   string `json:"end_time" binding:"required"`
	Comment   string `json:"comment" binding:"required"`
	Active    *bool  `json:"active" binding:"required"` // must be pointer due to the implementation in go
}

func (r *UpdateWorkdayRequest) Validate() error {
	// validate time in format: hh:mm:ss
	s, err := time.Parse("15:04:05", r.StartTime)
	if err != nil {
		return err
	}

	// validate time in format: hh:mm:ss
	e, err := time.Parse("15:04:05", r.EndTime)
	if err != nil {
		return err
	}

	// check if start time is before end time
	if s.After(e) {
		return errors.New("start time must be before end time")
	}

	return nil
}

type AssignPersonToWorkdayRequest struct {
	PersonID     string `json:"person_id" binding:"required"`
	DepartmentID string `json:"department_id" binding:"required"`
	WorkplaceID  string `json:"workplace_id" binding:"required"`
	TimeslotID   string `json:"timeslot_id" binding:"required"`
	Date         string `json:"date" binding:"required"`
}

func (r *AssignPersonToWorkdayRequest) Validate() error {
	// validate date in format: yyyy-mm-dd
	_, err := time.Parse("2006-01-02", r.Date)
	if err != nil {
		return err
	}

	return nil
}

type UnassignPersonFromWorkdayRequest struct {
	PersonID     string `json:"person_id" binding:"required"`
	DepartmentID string `json:"department_id" binding:"required"`
	WorkplaceID  string `json:"workplace_id" binding:"required"`
	TimeslotID   string `json:"timeslot_id" binding:"required"`
	Date         string `json:"date" binding:"required"`
}

func (r *UnassignPersonFromWorkdayRequest) Validate() error {
	// validate date in format: yyyy-mm-dd
	_, err := time.Parse("2006-01-02", r.Date)
	if err != nil {
		return err
	}

	return nil
}

/* Responses */
type WorkdayResponse struct {
	Department        DepartmentResponse `json:"department"`
	Workplace         WorkplaceResponse  `json:"workplace"`
	Timeslot          TimeslotResponse   `json:"timeslot"`
	Date              string             `json:"date"`
	StartTime         string             `json:"start_time"`
	DurationInMinutes int64              `json:"duration_in_minutes"`
	EndTime           string             `json:"end_time"`
	Weekday           string             `json:"weekday"`
	Comment           string             `json:"comment"`

	// Assigned Person can be nil
	Person *PersonResponse `json:"person"`
}
