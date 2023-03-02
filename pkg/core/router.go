package core

import (
	"confinio/pkg/engine/router"
)

func (k *Kernel) createRouter() {
	var (
		c = k.config
	)

	engine := router.NewRouter(
		&router.RuntimeConfiguration{
			EngineName:       c.HTTPEngine.Name,
			ListenAddress:    c.HTTPEngine.ListenAddress,
			ListenAddressTLS: c.HTTPEngine.ListenAddressTLS,
			CertFile:         c.HTTPEngine.CertFile,
			KeyFile:          c.HTTPEngine.KeyFile,
		},
	)
	if k.router == nil {
		k.router = engine
	}
}
