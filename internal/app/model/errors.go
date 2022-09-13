package model

import "errors"

var (
	ErrValidation           = errors.New("model validation failed")
	ErrMissingRequiredField = errors.New("missing required field")
)
