package dco

import (
	"encoding/json"

	"github.com/google/uuid"
)

/** Responses **/
type DepartmentResponse struct {
	BaseModel

	Name string `json:"name"`
}

type PermissionResponse struct {
	BaseModel

	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type UserResponse struct {
	BaseModel

	Username    string               `json:"username"`
	Email       string               `json:"email"`
	IsAdmin     bool                 `json:"is_admin"`
	Department  DepartmentResponse   `json:"department"`
	Permissions []PermissionResponse `json:"permissions"`
}

func (res UserResponse) MarshalJSON() ([]byte, error) {
	/* MarshalJSON is a function to marshal user response
	 * @param res is UserResponse
	 * @return ([]byte, error)
	 */
	type Alias UserResponse
	x := Alias(res)

	// handle boolean
	if res.IsAdmin {
		res.IsAdmin = true
	} else {
		res.IsAdmin = false
	}

	return json.Marshal(x)
}

/** Requests **/
type DepartmentRequest struct {
	Name string `json:"name" binding:"required"`
}

type PermissionRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description" binding:"omitempty"`
}

type UserRequest struct {
	Username string `json:"username" binding:"required,alpha,len=4,excludesall=!@#$%^&*()_+-="`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	IsAdmin  bool   `json:"is_admin"`

	DepartmentID uuid.UUID `json:"department_id" binding:"required"`
}

func (req UserRequest) MarshalJSON() ([]byte, error) {
	/* MarshalJSON is a function to marshal user request
	 * @param req is UserRequest
	 * @return ([]byte, error)
	 */
	type Alias UserRequest
	x := Alias(req)

	// handle boolean
	if req.IsAdmin {
		req.IsAdmin = true
	} else {
		req.IsAdmin = false
	}

	return json.Marshal(x)
}
