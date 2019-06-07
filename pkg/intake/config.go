package intake

import (
	"github.com/asciifaceman/drawdio/pkg/validate"
	"go.uber.org/zap"
	validator "gopkg.in/go-playground/validator.v9"
)

// Config describes configuration for Intake
type Config struct {
	Filename string `validate:"required"`
	Logger   *zap.Logger
}

// Validate validates the contents of Config
func (c *Config) Validate() error {
	v, err := validate.NewValidator(validator.New())
	if err != nil {
		return err
	}
	if err = v.Validate(c); err != nil {
		return err
	}
	return nil
}
