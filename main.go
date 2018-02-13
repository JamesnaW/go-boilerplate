package main

import (
	"fmt"

	"github.com/valyala/fasthttp"

	c "go-boilerplate/config"
	"go-boilerplate/klogl"
	"go-boilerplate/routes"
)

func main() {
	router := route.NewFastRouter()
	port := fmt.Sprintf(":%v", c.AppConfig.Port)
	klogl.Log.Fatal(fasthttp.ListenAndServe(port, router.Handler))
}
