package gaz

import mymy "github.com/ziutek/mymysql"

type MySql struct {
	*mymy.MySQL
}

const (
	proto    = "tcp",
	laddr    = "",
	raddr    = "127.0.0.1:3306",
	user     = "root",
	pass     = "root",
	database = "test",
)

func(m *MySql) Query(query string) interface{} {
	m.MySQL = mymy.New("tcp", "", "127.0.0.1:3306", "root", "root", "test")
	if err := m.Connect() ; err != nil {
		panic("cannot connect")
	}
	defer m.MySQL.Close()
	
	rows, _, err := m.MySQL.Query(query)
	
	if err != nil {
		panic(err)
	}
	
	return rows
}

func(m *MySql) Insert(p interface{}) (interface{}, bool) {
    return true, true
}

func(m *MySql) Get(id string) interface{} {
    return true
}
