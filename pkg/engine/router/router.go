package router

import (
	"fmt"
	"log"
	"net"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/reuseport"
)

type (
	// Router describes behaviour of the HTTP runtime components
	Router interface {
		Execute() error
	}
)

type (
	// Runtime implements "Router" interface
	Runtime struct {
		c        *RuntimeConfiguration
		mux      *router.Router
		listener net.Listener
	}

	// RuntimeConfiguration defines parameters used by HTTP runtime components
	RuntimeConfiguration struct {
		ServerName       string
		ListenAddress    string
		ListenAddressTLS string
		CertFile         string
		KeyFile          string
	}
)

// NewRouter returns a reference to a new instance of "Runtime" type
func NewRouter(rc *RuntimeConfiguration) *Runtime {
	return &Runtime{
		c:   rc,
		mux: router.New(),
	}
}

// Execute configures and runs HTTP/S server-side component(s)
func (r *Runtime) Execute() error {
	var (
		fault = make(chan error)
	)

	if r.c.ListenAddressTLS != "" && r.c.CertFile != "" && r.c.KeyFile != "" {
		if err := r.start(
			true,
			r.c.ListenAddressTLS,
			"Starting HTTPS runtime components: %s",
			fault,
		); err != nil {
			return err
		}
	}

	if err := r.start(
		false,
		r.c.ListenAddress,
		"Starting HTTP runtime components: %s",
		fault,
	); err != nil {
		return err
	}

	return <-fault
}

func (r *Runtime) start(serveTLS bool, address, message string, fault chan<- error) error {
	s, listener, err := r.create(address)
	if err != nil {
		return err
	}

	go func() {
		log.Printf(message, address)
		if serveTLS {
			handlerFunc := r.handler(true)
			s.Handler = fasthttp.CompressHandler(handlerFunc)
			if err := s.ServeTLS(listener, r.c.CertFile, r.c.KeyFile); err != nil {
				fault <- err
			}
		} else {
			handlerFunc := r.handler(false)
			s.Handler = fasthttp.CompressHandler(handlerFunc)
			if r.listener != nil {
				listener = r.listener
			}
			if err := s.Serve(listener); err != nil {
				fault <- err
			}
		}
	}()

	return nil
}

func (r *Runtime) create(address string) (*fasthttp.Server, net.Listener, error) {
	listener, err := reuseport.Listen("tcp4", address)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to initialize a port listener: %w", err)
	}

	return &fasthttp.Server{
		Name:                 r.c.ServerName,
		NoDefaultContentType: true,
	}, listener, nil
}
