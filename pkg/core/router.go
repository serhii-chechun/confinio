package core

import (
	"confinio/pkg/engine/router"
)

func (k *Kernel) createRouter() {
	engine := router.NewRouter(
		&router.RuntimeConfiguration{
			ServerName:       k.config.HTTPEngine.ServerName,
			ListenAddress:    k.config.HTTPEngine.ListenAddress,
			ListenAddressTLS: k.config.HTTPEngine.ListenAddressTLS,
			CertFile:         k.config.HTTPEngine.CertFile,
			KeyFile:          k.config.HTTPEngine.KeyFile,
		},
	)
	if k.router == nil {
		k.router = engine
	}
}
