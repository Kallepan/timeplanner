package dco

type WorkplaceResponse struct {
	Base
	Name string `json:"name"`
}

type WorkplaceRequest struct {
	Name string `json:"name" binding:"required"`
}
