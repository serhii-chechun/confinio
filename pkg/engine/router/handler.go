package router

import (
	"log"

	"github.com/valyala/fasthttp"
)

func (e *Engine) handler(serveTLS bool) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		defer func() {
			if err := recover(); err != nil {
				ctx.SetStatusCode(500)
				log.Printf("engine-handler: unable to perform a request: %s: %s", ctx.URI(), err)
			}
		}()

		// handle a request
		e.mux.Handler(ctx)
	}
}
