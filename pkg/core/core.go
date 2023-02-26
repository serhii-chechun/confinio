package core

import (
	"context"
	"fmt"

	"confinio/pkg/engine/router"
)

type (
	// Core describes behaviour of the central component
	Core interface {
		Prepare(ctx context.Context, configFile string) error
		Run(failure chan<- error)
	}
)

type (
	// Kernel implements "Core" interface and defines internal state of the central components
	Kernel struct {
		router router.Router
	}
)

// NewKernel returns a reference to a new instance of "Kernel" type
func NewKernel() *Kernel {
	return &Kernel{}
}

// Prepare applies application-wide initialization procedures
func (k *Kernel) Prepare(ctx context.Context, configFile string) error {
	k.router = router.NewEngine()
	return nil
}

// Run executes top-level logic of the application components
func (k *Kernel) Run(failure chan<- error) {
	if err := k.router.RunEngine(); err != nil {
		failure <- fmt.Errorf(
			"unable to start router engine: %w",
			err,
		)
	}
}
