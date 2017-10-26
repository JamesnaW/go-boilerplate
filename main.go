package main

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"

	c "github.com/JamesnaW/boilerplate/config"
	"github.com/JamesnaW/boilerplate/routes"
	// "github.com/JamesnaW/boilerplate/cassandra"
	"github.com/JamesnaW/boilerplate/mysql"
)

func main() {
	// CassandraSession := cassandra.Session
	// defer CassandraSession.Close()

	MysqlSession := mysql.Session
	defer MysqlSession.Close()

	router := route.NewFastRouter()
	port := fmt.Sprintf(":%v", c.Port)
	log.Fatal(fasthttp.ListenAndServe(port, router.Handler))
}
