package gaz

import "fmt"
import mymy "github.com/ziutek/mymysql"

type MySql struct {
	*mymy.MySQL
}

const (
	proto    = "tcp"
	laddr    = ""
	raddr    = "127.0.0.1:3306"
	user     = "root"
	pass     = "root"
	db       = "test"
)

func(m *MySql) new() {
	m.MySQL = mymy.New(proto, laddr, raddr, user, pass, db)
}

func(m *MySql) close() {
	m.Close()
}

func(m *MySql) Query(query string) interface{} {
	m.new()
	if err := m.Connect() ; err != nil {
		panic("cannot connect")
	}
	defer m.close()
	
	rows, _, err := m.MySQL.Query(query)
	
	if err != nil {
		panic(err)
	}
	
	return rows
}

func(m *MySql) Insert(p interface{}) (interface{}, bool) {
	m.new()
	if err := m.Connect() ; err != nil {
		panic("cannot connect")
	}
	defer m.close()
    return true, true
}

func(m *MySql) Get(id string) interface{} {
    return true
}
