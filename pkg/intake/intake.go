// Package intake handles the reading in of the
// wav audio source
package intake

import (
	"io"
	"os"

	"go.uber.org/zap"
)

const (
	// RESOLUTION of sampling
	RESOLUTION = 1 // seconds
)

// Intake  describes our audio consumer
type Intake struct {
	c          *Config
	r          io.Reader
	resolution uint
	Logger     *zap.Logger
}

// New returns a new Intake
func New(c *Config) (*Intake, error) {
	err := c.Validate()
	if err != nil {
		return nil, err
	}

	reader, err := os.Open(c.Filename)
	if err != nil {
		return nil, err
	}

	i := &Intake{
		r:          reader,
		c:          c,
		Logger:     c.Logger,
		resolution: RESOLUTION,
	}

	return i, err
}
