package gaz

import mymy "github.com/ziutek/mymysql"

type MySql struct {
	db *mymy.MySQL
}

func(m *MySql) Query(query string) interface{} {
	rows, _, err := MySql.db.Query(query)
	
	if err != nil {
		panic(err)
	}
	
	return rows
}
