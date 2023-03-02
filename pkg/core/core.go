package core

import (
	"context"
	"fmt"

	"confinio/pkg/engine/router"
	"confinio/pkg/settings"
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
		router router.Router
		config Configuration
	}

	// Configuration defines parameters used by the application
	Configuration struct {
		HTTPEngine struct {
			Name             string `json:"name"`
			ListenAddress    string `json:"listen_address"`
			ListenAddressTLS string `json:"tls_listen_address"`
			CertFile         string `json:"tls_cert_file"`
			KeyFile          string `json:"tls_key_file"`
		} `json:"http_engine"`
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

	k.createRouter()
	return nil
}

// Run executes top-level logic of the application components
func (k *Kernel) Run(failure chan<- error) {
	if err := k.router.Execute(); err != nil {
		failure <- fmt.Errorf(
			"unable to start HTTP engine: %w",
			err,
		)
	}
}
