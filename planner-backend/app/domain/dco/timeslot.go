package dco

/** Responses **/
type WeekdayResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	// We use string here because we only need the time like "08:00"
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type TimeslotResponse struct {
	Base
	Name           string            `json:"name"`
	Active         *bool             `json:"active"`
	DepartmentName string            `json:"department_name"`
	WorkplaceName  string            `json:"workplace_name"`
	Weekdays       []WeekdayResponse `json:"weekdays"`
}

/** Requests **/
type WeekdayRequest struct {
	ID        string  `json:"id" binding:"required"`
	StartTime *string `json:"start_time" binding:"omitempty"`
	EndTime   *string `json:"end_time" binding:"omitempty"`
}

type TimeslotRequest struct {
	Name   string `json:"name" binding:"required"`
	Active *bool  `json:"active" binding:"omitempty"`
}
