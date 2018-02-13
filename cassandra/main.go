package cassandra

import (
	c "go-boilerplate/config"
	"go-boilerplate/klogl"

	"github.com/gocql/gocql"
)

var Session *gocql.Session

func init() {
	var err error

	cluster := gocql.NewCluster(c.AppConfig.CassandraHost)
	// cluster.Keyspace = "keyspace"
	// cluster.Keyspace = "users"
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}

	klogl.Log.Infoln("Cassandra has been initialized")
}

type Todos []interface{}

func TodoListExample() Todos {
	rows := Session.Query("select * from todos").Iter()
	var list Todos
	m := map[string]interface{}{}
	for rows.MapScan(m) {
		list = append(list, m)
		m = map[string]interface{}{}
	}
	return list
}
