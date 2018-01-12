package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	c "go-boilerplate/config"
	"go-boilerplate/klogl"
)

var Session *sql.DB

func init() {
	var err error
	Session, err = sql.Open("mysql", fmt.Sprintf("%v:%v@%v(%v)/", c.AppConfig.MysqlUser, c.AppConfig.MysqlPassword, c.AppConfig.MysqlProtocol, c.AppConfig.MysqlHost))
	if err != nil {
		klogl.Log.Fatalf("Can't connect to MySQL Database: %v", err)
	}
	err = Session.Ping()
	if err != nil {
		klogl.Log.Fatalf("Can't ping to MySQL Database: %v", err)
	}

	klogl.Log.Infoln("Mysql has been initialized")
}

type Todo struct {
	Id   int
	Todo string
}

type Todos []Todo

func TodoListExample() Todos {
	rows, err := Session.Query("select id, todo from todos")
	var list Todos
	var t Todo
	for rows.Next() {
		err = rows.Scan(&t.Id, &t.Todo)
		if err != nil {
			fmt.Println(err.Error())
		}
		list = append(list, t)
	}
	return list
}
