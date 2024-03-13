package dco

/** Responses **/
type DepartmentInPersonResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type WorkplaceInPersonResponse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	DepartmentID string `json:"department_id"`
}
type WeekdayResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type PersonResponse struct {
	Base
	ID           string  `json:"id"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Email        string  `json:"email"`
	Active       *bool   `json:"active"`
	WorkingHours float64 `json:"working_hours"`

	Workplaces  []WorkplaceInPersonResponse  `json:"workplaces,omitempty"`
	Departments []DepartmentInPersonResponse `json:"departments,omitempty"`
	Weekdays    []WeekdayResponse            `json:"weekdays,omitempty"`
}

/** Requests **/
type PersonRequest struct {
	ID           string  `json:"id" binding:"required,alphanum,min=4,max=4"`
	FirstName    string  `json:"first_name" binding:"required"`
	LastName     string  `json:"last_name" binding:"required"`
	Email        string  `json:"email" binding:"required,email"`
	Active       *bool   `json:"active" binding:"required"`
	WorkingHours float64 `json:"working_hours"`
}
