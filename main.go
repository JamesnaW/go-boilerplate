package main

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"

	c "github.com/timeff/go-boilerplate/config"
	"github.com/timeff/go-boilerplate/routes"
	// "github.com/timeff/go-boilerplate/cassandra"
	// "github.com/timeff/go-boilerplate/mysql"
)

func main() {
	// CassandraSession := cassandra.Session
	// defer CassandraSession.Close()

	// MysqlSession := mysql.Session
	// defer MysqlSession.Close()

	router := route.NewFastRouter()
	port := fmt.Sprintf(":%v", c.Port)
	log.Fatal(fasthttp.ListenAndServe(port, router.Handler))
}
