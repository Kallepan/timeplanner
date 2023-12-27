package dao

import (
	"github.com/google/uuid"
)

type Permission struct {
	BaseModel
	// This is a simple permissions model
	// It is used to define what a user can do
	Name string `gorm:"type:varchar(255);column:name;not null" json:"name"`
	Description string `gorm:"type:varchar(255);column:description;not null" json:"description"`
}

type Department struct {
	// This is a simple department model
	BaseModel

	Name string `gorm:"type:varchar(255);column:name;not null" json:"name"`
}

type User struct {
	// This is a simple user model
	BaseModel
	Username string `gorm:"type:varchar(255);column:username;not null;index:idx_username,unique" json:"username"`
	// ->:false read-only field
	Password string `gorm:"type:varchar(255);column:password;->:false;<-:create" json:"-"`
	Email string `gorm:"type:varchar(255);column:email;not null" json:"email"`

	// Each User belongs to a department
	DepartmentID uuid.UUID `gorm:"type:uuid;column:department_id;not null" json:"department_id"`
	Department Department `gorm:"foreignKey:DepartmentID;references:ID" json:"department"`

	// Each User has multiple permissions
	Permissions []Permission `gorm:"many2many:user_permissions;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"permissions"`
}
