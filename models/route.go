package model

import (
	// "github.com/julienschmidt/httprouter"
	"github.com/valyala/fasthttp"
)

type Route struct {
	Name    string
	Method  string
	Pattern string
	// HandlerFunc httprouter.Handle
	HandlerFunc fasthttp.RequestHandler
}

type Routes []Route
