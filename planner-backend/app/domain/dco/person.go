package dco

import (
	"strings"
)

/** Responses **/
type PersonResponse struct {
	Base
	ID           string `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Active       *bool  `json:"active"`
	WorkingHours int64  `json:"working_hours"`

	Workplaces  []string `json:"workplaces,omitempty"`
	Departments []string `json:"departments,omitempty"`
	Weekdays    []string `json:"weekdays,omitempty"`
}

/** Requests **/
type PersonRequest struct {
	ID           string `json:"id" binding:"required,alphanum,min=4,max=4"`
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	Active       *bool  `json:"active" binding:"required"`
	WorkingHours int64  `json:"working_hours" binding:"required"`
}

func (p *PersonRequest) Validate() error {
	p.ID = strings.ToUpper(p.ID)
	p.ID = strings.TrimSpace(p.ID)

	return nil
}
