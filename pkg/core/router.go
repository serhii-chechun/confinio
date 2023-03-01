package core

import (
	"confinio/pkg/engine/router"
)

func (k *Kernel) createRouter() {
	routerEngine := router.NewEngine(
		&router.ServerConfiguration{
			ServerName:       k.config.WebServer.ServerName,
			ListenAddress:    k.config.WebServer.ListenAddress,
			ListenAddressTLS: k.config.WebServer.ListenAddressTLS,
			CertFile:         k.config.WebServer.CertFile,
			KeyFile:          k.config.WebServer.KeyFile,
		},
	)
	if k.router == nil {
		k.router = routerEngine
	}
}
