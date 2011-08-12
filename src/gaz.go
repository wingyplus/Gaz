package gaz

//import "strconv"
import mymy "github.com/ziutek/mymysql"

type Connection struct {
	*mymy.MySQL
}

type Database struct {
	Connection *Connection
	db         string
}

type DataSet struct {
	DB         *Database
	Name       string
}

/*
type MySql struct {
	Connection *Connection
	DB         DataStore
	Table      Database
}
*/

type Params map[string]interface{}

const (
	proto    = "tcp"
	laddr    = ""
	raddr    = "127.0.0.1:3306"
	user     = "root"
	pass     = "root"
)

/*
type Database struct {
	Session *Session
	Name    string
}

type Collection struct {
	DB       Database
	Name     string // "collection"
	FullName string // "db.collection"
}

func (session *Session) DB(name string) Database {
	return Database{session, name}
}

func (database Database) C(name string) Collection {
	return Collection{database, name, database.Name + "." + name}
}
*/

func(m *Connection) DB(db string) *Database {
	return &Database{m, db}
}

func(database *Database) C(table string) *DataSet {
	return &DataSet{database, table}
}

func(database *Database) new() {
	database.Connection.MySQL = mymy.New(proto, laddr, raddr, user, pass, database.db)
}

func(database *Database) close() {
	database.Connection.Close()
}

func(database *Database) Query(query string) interface{} {
	database.new()
	if err := database.Connection.Connect() ; err != nil {
		panic("cannot connect")
	}
	defer database.close()
	
	rows, _, err := database.Connection.MySQL.Query(query)
	
	if err != nil {
		panic(err)
	}
	
	return rows
}
/*
func(m *Connection) insert(p interface{}, table string) (interface{}, bool) {
	m.new()
	if err := m.Connect() ; err != nil {
		panic("cannot connect")
	}
	defer m.close()
	
	data := p.(map[string]string)
	rows, _, _ := m.MySQL.Query("SELECT * FROM " + table)
	
	query := "INSERT INTO " + table + "(id, email, name, password) VALUES (" + strconv.Itoa(len(rows)+1)
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

func(m *Connection) get(id string) interface{} {
	rows := m.Query("SELECT * FROM User WHERE id=" + id).([]*mymy.Row)
	
	return rows[0]
}

func(m *Connection) findOne(p Params) interface{} {
	query := "SELECT * FROM User WHERE "
	for key, value := range p {
		query += key + "='" + value.(string) + "'"
	}
	
	rows := m.Query(query).([]*mymy.Row)
	return rows[0]
}
*/
