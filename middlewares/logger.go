package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/valyala/fasthttp"
)

func LoggerHttprouter(handler httprouter.Handle, name string) httprouter.Handle {
	return (func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		start := time.Now()

		handler(w, r, ps)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

func LoggerFastHttp(name string) func(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
		return func(ctx *fasthttp.RequestCtx) {
			start := time.Now()

			defer track(ctx, start, name)
			handler(ctx)
		}
	}
}

func track(ctx *fasthttp.RequestCtx, start time.Time, name string) {
	log.Printf(
		"%s\t%s\t%s\t%s",
		string(ctx.Method()),
		string(ctx.RequestURI()),
		name,
		time.Since(start),
	)
}
