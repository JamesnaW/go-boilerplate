package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

///// Config with Flag /////
var (
	Port          int
	Name          string
	CassandraHost string
	MysqlHost     string
	MysqlUser     string
	MysqlPassword string
	MysqlProtocol string
)
var (
	port          int
	name          string
	cassandraHost string
	mysqlHost     string
	mysqlUser     string
	mysqlPassword string
	mysqlProtocol string
)

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port, _ = strconv.Atoi(os.Getenv("PORT"))
	name = os.Getenv("PROJECT")
	cassandraHost = os.Getenv("CASSANDRAHOST")
	mysqlHost = os.Getenv("MYSQLHOST")
	mysqlUser = os.Getenv("MYSQLUSER")
	mysqlPassword = os.Getenv("MYSQLPASSWORD")
	mysqlProtocol = os.Getenv("MYSQLPROTOCOL")
	if port == 0 {
		port = 5000
	}
	if name == "" {
		name = "go-Boilerplate"
	}
	if cassandraHost == "" {
		cassandraHost = "127.0.0.1"
	}
	if mysqlHost == "" {
		mysqlHost = "127.0.0.1:3306"
	}
	if mysqlUser == "" {
		mysqlUser = "root"
	}
	if mysqlPassword == "" {
		mysqlPassword = "root"
	}
	if mysqlProtocol == "" {
		mysqlProtocol = "tcp"
	}
	flag.IntVar(&Port, "p", port, "Specify the port to listen to.")
	flag.StringVar(&Name, "n", name, "Specify the project name.")
	flag.StringVar(&CassandraHost, "c", cassandraHost, "Specify the cassandra host.")
	flag.StringVar(&MysqlHost, "mh", mysqlHost, "Specify the mysql host.")
	flag.StringVar(&MysqlUser, "mu", mysqlUser, "Specify the mysql user.")
	flag.StringVar(&MysqlPassword, "mw", mysqlPassword, "Specify the mysql password.")
	flag.StringVar(&MysqlProtocol, "mp", mysqlProtocol, "Specify the mysql protocol.")
	flag.Parse()

	fmt.Println("Run", Name, "project on port", Port)
}
