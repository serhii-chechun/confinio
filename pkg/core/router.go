package core

import (
	"confinio/pkg/router"
)

func (k *Kernel) createServers() {
	var (
		servers = make([]router.Router, len(k.config.Servers))
	)

	for i, s := range k.config.Servers {
		servers[i] = router.NewRouter(
			&router.RuntimeConfiguration{
				ServerName:       s.HTTPEngine.ServerName,
				ListenAddress:    s.HTTPEngine.ListenAddress,
				ListenAddressTLS: s.HTTPEngine.ListenAddressTLS,
				CertFile:         s.HTTPEngine.CertFile,
				KeyFile:          s.HTTPEngine.KeyFile,
			},
		)
	}

	if k.servers == nil {
		k.servers = servers
	}
}
