package app

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"porta/pkg/core"
)

const (
	_name    = "confinio-porta"
	_version = "v0.0.3"
)

var (
	_main = struct {
		ctx  context.Context
		core core.Core
	}{
		ctx:  context.Background(),
		core: core.NewKernel(),
	}
)

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
	if err := _main.core.Prepare(
		_main.ctx,
		*configFilename,
	); err != nil {
		println(err.Error())
		os.Exit(1)
	}
	go _main.core.Run(failure)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case s := <-quit:
		log.Printf("Shutting down: received OS signal %s", s)
	case <-_main.ctx.Done():
		log.Println("Shutting down: global context has terminated")
	case err := <-failure:
		log.Printf("Shutting down: runtime failure: %s", err)
	}
}
