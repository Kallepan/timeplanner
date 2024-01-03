package dao

import (
	"database/sql"

	"github.com/google/uuid"
)

type Permission struct {
	// This is a simple permissions model. It is used to define what a user can do
	BaseModel

	Name        string         `gorm:"type:varchar(255);column:name;unique;not null"`
	Description sql.NullString `gorm:"type:varchar(255);column:description;default:null"`
}

type Department struct {
	// This is a simple department model
	BaseModel

	Name string `gorm:"type:varchar(255);column:name;unique;not null"`
}

type User struct {
	// This is a simple user model
	BaseModel

	Username string `gorm:"type:varchar(255);column:username;not null;index:idx_username,unique"`
	// ->:false read-only field
	Password string `gorm:"type:varchar(255);column:password;->:false;<-:create"`
	Email    string `gorm:"type:varchar(255);column:email;not null"`

	// boolean field
	IsAdmin bool `gorm:"type:boolean;column:is_admin;not null;default:false" json:"is_admin"`

	// Each User belongs to a department
	DepartmentID uuid.UUID  `gorm:"type:uuid;column:department_id;not null"`
	Department   Department `gorm:"foreignKey:DepartmentID;references:ID"`

	// Each User has multiple permissions
	Permissions []Permission `gorm:"many2many:user_permissions;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
