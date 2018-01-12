package middleware

import (
	"go-boilerplate/klogl"
	"time"

	"github.com/valyala/fasthttp"
)

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
	klogl.Log.Infof(
		"%s\t%s\t%s\t%s",
		string(ctx.Method()),
		string(ctx.RequestURI()),
		name,
		time.Since(start),
	)
}
