package gaz

import "strings"
import mymy "github.com/ziutek/mymysql"

type Connection struct {
	*mymy.MySQL
}

type Database struct {
	connection *Connection
	db_name    string
}

type DataSet struct {
	db         Database
	table_name string
}

type Params map[string]interface{}

const (
	proto    = "tcp"
	laddr    = ""
	raddr    = "127.0.0.1:3306"
	user     = "root"
	pass     = "root"
)

func(conn *Connection) DB(db string) Database {
	return Database{conn, db}
}

func(database Database) C(table string) DataSet {
	return DataSet{database, table}
}

func(database *Database) new() {
	database.connection.MySQL = mymy.New(proto, laddr, raddr, user, pass, database.db_name)
}

func(database *Database) close() {
	database.connection.Close()
}

func(database *Database) Query(query string) interface{} {
	database.new()
	if err := database.connection.Connect() ; err != nil {
		panic("cannot connect")
	}
	defer database.close()
	
	rows, _, err := database.connection.MySQL.Query(query)
	
	if err != nil {
		panic(err)
	}
	
	return rows
}

// extractField is return Field and Data_Types of Table
func(dataset *DataSet) extractField() map[string]string {
	rows := dataset.db.Query("DESC " + dataset.table_name).([]*mymy.Row)
	field := make(map[string]string)
	for _, row := range rows {
		field[row.Str(0)] = row.Str(1)
	}
	return field
}

func(dataset *DataSet) Insert(p interface{}) (interface{}, bool) {
	dataset.db.new()
	
	field := dataset.extractField()
	data := p.(map[string]string)
	
	var sub_query, data_query string
	for key, _ := range field {
		if key == "id" {
			continue
		}
		sub_query += key + " "
		data_query += "'" + data[key] + "' " 
	}
	query := "INSERT INTO " + dataset.table_name + "(" + strings.Replace(strings.Replace(sub_query, " ", ",", len(field)-2), " ", "", -1) + ") VALUES (" + strings.Replace(strings.Replace(data_query, " ", ",", len(field)-2), " ", "", -1) + ")"

	if err := dataset.db.connection.Connect() ; err != nil {
		panic("cannot connect")
	}
	defer dataset.db.close()
	_, _, err := dataset.db.connection.MySQL.Query(query)
	if(err != nil) {
		return err, false
	}
	
	return nil, true
}

func(dataset *DataSet) Get(id string) interface{} {
	rows := dataset.db.Query("SELECT * FROM " + dataset.table_name + " WHERE id=" + id).([]*mymy.Row)
	
	return rows[0]
}

func(dataset *DataSet) FindOne(p Params) interface{} {
	query := "SELECT * FROM " + dataset.table_name + " WHERE "
	for key, value := range p {
		query += key + "='" + value.(string) + "'"
	}
	
	rows := dataset.db.Query(query).([]*mymy.Row)
	return rows[0]
}

