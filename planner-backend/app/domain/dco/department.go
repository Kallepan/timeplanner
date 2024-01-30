package dco

type DepartmentResponse struct {
	Base
	Name string `json:"name"`
	ID   string `json:"id"`
}

type DepartmentRequest struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
