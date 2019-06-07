// Package validate provides a central
// way to validate individual fields and different
// implementations of config
package validate

import validator "gopkg.in/go-playground/validator.v9"

// Validator is a custom validator
type Validator struct {
	validator *validator.Validate
}

// NewValidator constructs a custom validator with all custom validators registered.
func NewValidator(v *validator.Validate) (*Validator, error) {
	cv := &Validator{
		validator: v,
	}
	err := cv.RegisterCustom()
	return cv, err
}

// RegisterCustom registers all custom validators for use.
func (cv *Validator) RegisterCustom() error {
	return nil
}

// Validate is the main validation step for structs
func (cv *Validator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
