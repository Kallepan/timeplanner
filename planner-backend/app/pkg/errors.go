package pkg

import "errors"

var (
	// Error codes
	ErrNoRows                   = errors.New("no rows in result set")
	ErrValidation               = errors.New("validation error")
	ErrDidNotCreateRelationship = errors.New("did not create relationship")
)
