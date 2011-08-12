package gaz

import "strconv"
import mymy "github.com/ziutek/mymysql"

type Connection struct {
	*mymy.MySQL
}

type MySql struct {
	connection *Connection
	db         string
	table      string
}

type Params map[string]interface{}

const (
	proto    = "tcp"
	laddr    = ""
	raddr    = "127.0.0.1:3306"
	user     = "root"
	pass     = "root"
	db       = "test"
)

func(m *MySql) DB(db string) Database {}

func(m *MySql) C(table string) DataSet {}

func(m *Connection) new() {
	m.MySQL = mymy.New(proto, laddr, raddr, user, pass, db)
}

func(m *Connection) close() {
	m.Close()
}

func(m *Connection) Query(query string) interface{} {
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

func(m *Connection) Insert(p interface{}) (interface{}, bool) {
	m.new()
	if err := m.Connect() ; err != nil {
		panic("cannot connect")
	}
	defer m.close()
	
	data := p.(map[string]string)
	rows, _, _ := m.MySQL.Query("SELECT * FROM User")
	
	query := "INSERT INTO User(id, email, name, password) VALUES (" + strconv.Itoa(len(rows)+1)
	for _, value := range data {
		query += "," + "'" + value + "'"
	}
	query += ")"
	
	_, _, err := m.MySQL.Query(query)
	if(err != nil) {
		return err, false
	}
	
	return nil, true
}

func(m *Connection) Get(id string) interface{} {
	rows := m.Query("SELECT * FROM User WHERE id=" + id).([]*mymy.Row)
	
	return rows[0]
}

func(m *Connection) FindOne(p Params) interface{} {
	query := "SELECT * FROM User WHERE "
	for key, value := range p {
		query += key + "='" + value.(string) + "'"
	}
	
	rows := m.Query(query).([]*mymy.Row)
	return rows[0]
}

