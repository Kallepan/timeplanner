package dco

import (
	"time"
)

/* Requests */
type AssignPersonToWorkdayRequest struct {
	PersonID     string `json:"person_id" binding:"required"`
	DepartmentID string `json:"department_id" binding:"required"`
	WorkplaceID  string `json:"workplace_id" binding:"required"`
	TimeslotName string `json:"timeslot_name" binding:"required"`
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
	TimeslotName string `json:"timeslot_name" binding:"required"`
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
type WorkdayPersonResponse struct {
	ID           string  `json:"id"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Email        string  `json:"email"`
	WorkingHours float64 `json:"working_hours"`
}
type WorkdayResponse struct {
	Department DepartmentResponse `json:"department"`
	Workplace  WorkplaceResponse  `json:"workplace"`
	Timeslot   TimeslotResponse   `json:"timeslot"`
	Date       string             `json:"date"`
	StartTime  string             `json:"start_time"`
	EndTime    string             `json:"end_time"`
	Weekday    string             `json:"weekday"`

	// Assigned Person can be nil
	Person *WorkdayPersonResponse `json:"person"`
}
