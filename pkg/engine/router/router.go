package router

type (
	// Router describes behaviour of the HTTP server components
	Router interface {
		RunEngine() error
	}
)

type (
	// Engine implements "Router" interface
	Engine struct{}
)

// NewEngine returns a reference to a new instance of "Engine" type
func NewEngine() *Engine {
	return &Engine{}
}

// RunEngine configures and runs HTTP/S server(s)
func (e *Engine) RunEngine() error {
	println("confinio: router engine is running...")
	return nil
}
