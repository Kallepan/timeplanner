package dco

import (
	"strings"
	"time"
)

/* Requests */
type AssignPersonToWorkdayRequest struct {
	PersonID       string `json:"person_id" binding:"required"`
	DepartmentName string `json:"department_name" binding:"required"`
	WorkplaceName  string `json:"workplace_name" binding:"required"`
	TimeslotName   string `json:"timeslot_name" binding:"required"`
	Date           string `json:"date" binding:"required"`
}

func (r *AssignPersonToWorkdayRequest) Validate() error {
	// validate date in format: yyyy-mm-

	// Ensure all strings are uppercase
	r.PersonID = strings.ToUpper(r.PersonID)
	r.DepartmentName = strings.ToUpper(r.DepartmentName)
	r.WorkplaceName = strings.ToUpper(r.WorkplaceName)
	r.TimeslotName = strings.ToUpper(r.TimeslotName)

	_, err := time.Parse("2006-01-02", r.Date)
	if err != nil {
		return err
	}

	return nil
}

type UnassignPersonFromWorkdayRequest struct {
	PersonID       string `json:"person_id" binding:"required"`
	DepartmentName string `json:"department_name" binding:"required"`
	WorkplaceName  string `json:"workplace_name" binding:"required"`
	TimeslotName   string `json:"timeslot_name" binding:"required"`
	Date           string `json:"date" binding:"required"`
}

func (r *UnassignPersonFromWorkdayRequest) Validate() error {
	// validate date in format: yyyy-mm-

	// Ensure all strings are uppercase
	r.PersonID = strings.ToUpper(r.PersonID)
	r.DepartmentName = strings.ToUpper(r.DepartmentName)
	r.WorkplaceName = strings.ToUpper(r.WorkplaceName)
	r.TimeslotName = strings.ToUpper(r.TimeslotName)

	_, err := time.Parse("2006-01-02", r.Date)
	if err != nil {
		return err
	}

	return nil
}

/* Responses */
type WorkdayPersonResponse struct {
	ID           string `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	WorkingHours int64  `json:"working_hours"`
}
type WorkdayResponse struct {
	Department string `json:"department"`
	Workplace  string `json:"workplace"`
	Timeslot   string `json:"timeslot"`
	Date       string `json:"date"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
	Weekday    string `json:"weekday"`

	// Assigned Person can be nil
	Person *WorkdayPersonResponse `json:"person"`
}
