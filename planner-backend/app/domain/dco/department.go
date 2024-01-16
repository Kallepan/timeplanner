package dco

type DepartmentResponse struct {
	Base
	Name string `json:"name"`
}

type DepartmentRequest struct {
	Name string `json:"name" binding:"required"`
}
