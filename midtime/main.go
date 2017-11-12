package midtime

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

type Middleware func(fasthttp.RequestHandler) fasthttp.RequestHandler

func Chain(mdList []Middleware, h fasthttp.RequestHandler) fasthttp.RequestHandler {
	if h == nil {
		h = func(ctx *fasthttp.RequestCtx) {
			fmt.Println("Error: no last handler")
		}
	}

	for i := range mdList {
		h = mdList[len(mdList)-1-i](h)
	}
	return h
}
