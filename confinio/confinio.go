package confinio

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"confinio/pkg/core"
)

const (
	_name    = "confinio"
	_version = "v0.0.1"
)

var (
	// _confinio defines the state properties of the "Main" function
	_confinio = struct {
		ctx  context.Context
		core core.Core
	}{
		ctx:  context.Background(),
		core: core.NewKernel(),
	}
)

// Main implements entry point of the application
func Main() {
	var (
		configFilename = flag.String("c", "", "configuration filename")
		showVersion    = flag.Bool("v", false, "current version")
	)

	flag.Parse()
	if *showVersion {
		println(
			_name,
			_version,
		)
		return
	}

	failure := make(chan error)
	if err := _confinio.core.Prepare(
		_confinio.ctx,
		*configFilename,
	); err != nil {
		println(err.Error())
		os.Exit(1)
	}
	go _confinio.core.Run(failure)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case s := <-quit:
		log.Printf("shutting down, received OS signal %s", s)
	case <-_confinio.ctx.Done():
		log.Println("shutting down, global context has terminated")
	case err := <-failure:
		log.Printf("shutting down due to a runtime failure: %s", err)
	}
}
