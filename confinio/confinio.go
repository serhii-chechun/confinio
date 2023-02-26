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

type (
	// main defines the state properties of the "Main" function
	main struct {
		ctx  context.Context
		core core.Core
	}
)

const (
	_name    = "confinio"
	_version = "v0.0.1"
)

var (
	_main = main{
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
		println(_name, _version)
		return
	}

	failure := make(chan error)
	if err := _main.core.Prepare(_main.ctx, *configFilename); err != nil {
		println(err.Error())
		os.Exit(1)
	}
	go _main.core.Run(failure)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case s := <-quit:
		log.Printf("shutting down gracefully, received OS signal %s", s)
	case <-_main.ctx.Done():
		log.Println("shutting down, global context has terminated")
	case err := <-failure:
		log.Printf("shutting down due to a runtime failure: %s", err)
	}
}
