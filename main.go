package main

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"

	c "go-boilerplate/config"
	"go-boilerplate/routes"
)

func main() {
	router := route.NewFastRouter()
	port := fmt.Sprintf(":%v", c.Port)
	log.Fatal(fasthttp.ListenAndServe(port, router.Handler))
}
