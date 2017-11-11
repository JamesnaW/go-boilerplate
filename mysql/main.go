package mysql

import (
	"database/sql"
	"fmt"
	c "github.com/timeff/go-boilerplate/config"

	_ "github.com/go-sql-driver/mysql"
)

var Session *sql.DB

func init() {
	var err error
	Session, err = sql.Open("mysql", fmt.Sprintf("%v:%v@%v(%v)/", c.MysqlUser, c.MysqlPassword, c.MysqlProtocol, c.MysqlHost))
	if err != nil {
		panic(err)
	}
	err = Session.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Mysql init done")
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
