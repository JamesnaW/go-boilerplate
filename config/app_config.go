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
type globalConfig struct {
	Port          int
	Name          string
	CassandraHost string
	MysqlHost     string
	MysqlUser     string
	MysqlPassword string
	MysqlProtocol string
	Environment   string
}

var (
	AppConfig globalConfig
)

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	name := os.Getenv("NAME")
	cassandraHost := os.Getenv("CASSANDRAHOST")
	mysqlHost := os.Getenv("MYSQLHOST")
	mysqlUser := os.Getenv("MYSQLUSER")
	mysqlPassword := os.Getenv("MYSQLPASSWORD")
	mysqlProtocol := os.Getenv("MYSQLPROTOCOL")
	environment := os.Getenv("ENVIRONMENT")

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
	if environment == "" {
		environment = "local"
	}

	if environment == "local" {
		flag.IntVar(&AppConfig.Port, "p", port, "Specify the port to listen to.")
		flag.StringVar(&AppConfig.Name, "n", cassandraHost, "Specify the project name.")
		flag.StringVar(&AppConfig.CassandraHost, "c", cassandraHost, "Specify the cassandra host.")
		flag.StringVar(&AppConfig.MysqlHost, "mh", mysqlHost, "Specify the mysql host.")
		flag.StringVar(&AppConfig.MysqlUser, "mu", mysqlUser, "Specify the mysql user.")
		flag.StringVar(&AppConfig.MysqlPassword, "mw", mysqlPassword, "Specify the mysql password.")
		flag.StringVar(&AppConfig.MysqlProtocol, "mp", mysqlProtocol, "Specify the mysql protocol.")
		flag.Parse()
	}

	fmt.Println("Run", AppConfig.Name, "project on port", AppConfig.Port)
}
