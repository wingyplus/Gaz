package gaz

import "strings"
import mymy "github.com/ziutek/mymysql"
import "fmt"

type Connection struct {
	*mymy.MySQL
}

type Database struct {
	Connection *Connection
	Db_name    string
}

type DataSet struct {
	Db         Database
	Table_name string
}

type Params map[string]interface{}

const (
	proto    = "tcp"
	laddr    = ""
	raddr    = "127.0.0.1:3306"
	user     = "root"
	pass     = "root"
)

func(conn *Connection) DB(Db string) Database {
	return Database{conn, Db}
}

func(database Database) C(table string) DataSet {
	return DataSet{database, table}
}

func(database *Database) new() {
	database.Connection.MySQL = mymy.New(proto, laddr, raddr, user, pass, database.Db_name)
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

// extractField is return Field and Data_Types of Table
func(dataset *DataSet) extractField() map[string]string {
	rows := dataset.Db.Query("DESC " + dataset.Table_name).([]*mymy.Row)
	field := make(map[string]string)
	for _, row := range rows {
		field[row.Str(0)] = row.Str(1)
	}
	return field
}

func(dataset *DataSet) Insert(p interface{}) (interface{}, bool) {
	field := dataset.extractField()
	data := maptype(p.(map[string]interface{}))
	
	var sub_query, data_query string
	for key, _ := range field {
		if key == "id" {
			continue
		}
		sub_query += key + " "
		data_query += data[key] + " " 
	}
	query := "INSERT INTO " + dataset.Table_name + "(" + strings.Replace(strings.Replace(sub_query, " ", ",", len(field)-2), " ", "", -1) + ") VALUES (" + strings.Replace(strings.Replace(data_query, " ", ",", len(field)-2), " ", "", -1) + ")"
	fmt.Println(query)
	
	dataset.Db.new()
	if err := dataset.Db.Connection.Connect() ; err != nil {
		panic("cannot connect")
	}
	defer dataset.Db.close()
	_, _, err := dataset.Db.Connection.MySQL.Query(query)
	if(err != nil) {
		return err, false
	}
	
	return nil, true
}

func(dataset *DataSet) Get(id string) interface{} {
	rows := dataset.Db.Query("SELECT * FROM " + dataset.Table_name + " WHERE id=" + id).([]*mymy.Row)
	
	return rows[0]
}

func(dataset *DataSet) FindOne(p Params) interface{} {
	query := "SELECT * FROM " + dataset.Table_name + " WHERE "
	for key, value := range p {
		query += key + "='" + value.(string) + "'"
	}
	
	rows := dataset.Db.Query(query).([]*mymy.Row)
	return rows[0]
}

