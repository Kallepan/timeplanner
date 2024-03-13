package dco

type WorkplaceResponse struct {
	Base
	Name         string `json:"name"`
	ID           string `json:"id"`
	DepartmentID string `json:"department_id"`
}

type WorkplaceRequest struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
