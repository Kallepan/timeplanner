package dco

import "planner-backend/app/pkg"

/** Responses **/
type OnWeekdayResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`

	// We use string here because we only need the time like "08:00"
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type TimeslotResponse struct {
	Base
	ID           string              `json:"id"`
	Name         string              `json:"name"`
	DepartmentID string              `json:"department_id"`
	WorkplaceID  string              `json:"workplace_id"`
	Weekdays     []OnWeekdayResponse `json:"weekdays"`
}

/** Requests **/
type WeekdayRequest struct {
	ID        int64   `json:"id" binding:"required"`
	StartTime *string `json:"start_time" binding:"omitempty"`
	EndTime   *string `json:"end_time" binding:"omitempty"`
}

// validate Weekday ID should be one of the following:
// 1, 2, 3, 4, 5, 6, 7
func (w *WeekdayRequest) Validate() error {
	/* Validate the weekday request */
	weekdayID := w.ID
	if weekdayID != 1 && weekdayID != 2 && weekdayID != 3 && weekdayID != 4 && weekdayID != 5 && weekdayID != 6 && weekdayID != 7 {
		return pkg.ErrValidation
	}

	return nil
}

// this is used for bulk updating the weekdays
type WeekdaysRequest struct {
	Weekdays []WeekdayRequest `json:"weekdays" binding:"required"`
}

func (w *WeekdaysRequest) Validate() error {
	for _, weekday := range w.Weekdays {
		if err := weekday.Validate(); err != nil {
			return err
		}
	}

	return nil
}

type TimeslotRequest struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
