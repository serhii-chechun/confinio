package core

import (
	"porta/pkg/engine"
)

func (k *Kernel) createEngines() {
	var (
		e = make([]engine.Engine, len(k.config.Servers))
	)

	for i, s := range k.config.Servers {
		e[i] = engine.NewRuntime(
			&engine.RuntimeConfiguration{
				ServerName:       s.HTTPEngine.ServerName,
				ListenAddress:    s.HTTPEngine.ListenAddress,
				ListenAddressTLS: s.HTTPEngine.ListenAddressTLS,
				CertFile:         s.HTTPEngine.CertFile,
				KeyFile:          s.HTTPEngine.KeyFile,
			},
		)
	}

	if k.engines == nil {
		k.engines = e
	}
}
