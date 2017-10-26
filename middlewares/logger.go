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

func LoggerFastHttp(handler fasthttp.RequestHandler, name string) fasthttp.RequestHandler {
	return (func(ctx *fasthttp.RequestCtx) {
		start := time.Now()

		handler(ctx)

		log.Printf(
			"%s\t%s\t%s\t%s",
			string(ctx.Method()),
			string(ctx.RequestURI()),
			name,
			time.Since(start),
		)
	})
}
