package dco

import (
	"planner-backend/app/pkg"
	"time"
)

/** Responses **/
type AbsenceResponse struct {
	PersonID  string    `json:"person_id"`
	Reason    string    `json:"reason,omitempty"`
	Date      string    `json:"date"`
	CreatedAt time.Time `json:"created_at"`
}

/** Requests **/
type AbsenceRequest struct {
	Date   string  `json:"date" binding:"required"`
	Reason *string `json:"reason" binding:"omitempty"`
}

type RelAddDepartmentRequest struct {
	DepartmentID string `json:"department_id" binding:"required"`
}

type RelAddWorkplaceRequest struct {
	WorkplaceID  string `json:"workplace_id" binding:"required"`
	DepartmentID string `json:"department_id" binding:"required"`
}

type RelRemoveWorkplaceRequest struct {
	WorkplaceID  string `json:"workplace_id" binding:"required"`
	DepartmentID string `json:"department_id" binding:"required"`
}

type RelAddWeekdayRequest struct {
	WeekdayID string `json:"weekday_id" binding:"required"`
}

func (r *RelAddWeekdayRequest) Validate() error {
	/* Validate the weekday request */
	weekdayID := r.WeekdayID
	if weekdayID != "MON" && weekdayID != "TUE" && weekdayID != "WED" && weekdayID != "THU" && weekdayID != "FRI" && weekdayID != "SAT" && weekdayID != "SUN" {
		return pkg.ErrValidation
	}

	return nil
}
