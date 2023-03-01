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
	// Router describes behaviour of the HTTP server components
	Router interface {
		RunEngine() error
	}
)

type (
	// Engine implements "Router" interface
	Engine struct {
		sc       *ServerConfiguration
		mux      *router.Router
		listener net.Listener
	}

	// ServerConfiguration defines parameters used by HTTP server
	ServerConfiguration struct {
		ServerName       string
		ListenAddress    string
		ListenAddressTLS string
		CertFile         string
		KeyFile          string
	}
)

// NewEngine returns a reference to a new instance of "Engine" type
func NewEngine(sc *ServerConfiguration) *Engine {
	return &Engine{
		sc:  sc,
		mux: router.New(),
	}
}

// RunEngine configures and runs HTTP/S server(s)
func (e *Engine) RunEngine() error {
	var (
		fault = make(chan error)
	)

	if e.sc.ListenAddressTLS != "" && e.sc.CertFile != "" && e.sc.KeyFile != "" {
		if err := e.spawnServer(true, e.sc.ListenAddressTLS, "Starting HTTPS engine: %s", fault); err != nil {
			return err
		}
	}

	if err := e.spawnServer(false, e.sc.ListenAddress, "Starting HTTP engine: %s", fault); err != nil {
		return err
	}

	return <-fault
}

func (e *Engine) spawnServer(serveTLS bool, address, message string, fault chan<- error) error {
	s, listener, err := e.createServer(address)
	if err != nil {
		return err
	}

	go func() {
		log.Printf(message, address)
		if serveTLS {
			handlerFunc := e.handler(true)
			s.Handler = fasthttp.CompressHandler(handlerFunc)
			if err := s.ServeTLS(listener, e.sc.CertFile, e.sc.KeyFile); err != nil {
				fault <- err
			}
		} else {
			handlerFunc := e.handler(false)
			s.Handler = fasthttp.CompressHandler(handlerFunc)
			if e.listener != nil {
				listener = e.listener
			}
			if err := s.Serve(listener); err != nil {
				fault <- err
			}
		}
	}()

	return nil
}

func (e *Engine) createServer(address string) (*fasthttp.Server, net.Listener, error) {
	listener, err := reuseport.Listen("tcp4", address)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to initialize a port listener: %w", err)
	}

	return &fasthttp.Server{
		Name:                 e.sc.ServerName,
		NoDefaultContentType: true,
	}, listener, nil
}
