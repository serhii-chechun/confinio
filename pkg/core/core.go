package core

import (
	"context"
	"fmt"

	"confinio/pkg/router"

	"github.com/pkg-wire/settings"
)

type (
	// Core describes behaviour of the central component
	Core interface {
		Prepare(ctx context.Context, configFile string) error
		Run(failure chan<- error)
	}
)

type (
	// Kernel implements "Core" interface and
	// defines internal state of the central components
	Kernel struct {
		servers []router.Router
		config  Configuration
	}
)

// NewKernel returns a reference to a new instance of "Kernel" type
func NewKernel() *Kernel {
	return &Kernel{}
}

// Prepare applies application-wide initialization procedures
func (k *Kernel) Prepare(ctx context.Context, configFile string) error {
	if err := settings.NewConfiguration().
		FromFile(configFile, settings.FileFormatJSON).
		Populate(&k.config); err != nil {
		return fmt.Errorf(
			"configuration processing issue: %w",
			err,
		)
	}

	k.createServers()
	return nil
}

// Run executes top-level logic of the application components
func (k *Kernel) Run(failure chan<- error) {
	for i := range k.servers {
		go k.spawnServer(i, failure)
	}
}

func (k *Kernel) spawnServer(idx int, failure chan<- error) {
	if err := k.servers[idx].Execute(); err != nil {
		failure <- fmt.Errorf(
			"unable to start HTTP engine: %w",
			err,
		)
	}
}
