package gaz

import (	
	"fmt"
	"testing"
//	"github.com/bmizerany/assert"
	mymy "github.com/ziutek/mymysql"
)



func TestQuery(t *testing.T) {
	conn := new(MySql)
	rows := conn.Query("select * from User").([]*mymy.Row)
	
	for _, row := range rows {
		for _, col := range row.Data {
			fmt.Printf ("%v ", string(col.([]byte)))
		}
		fmt.Println()
	}
}

func TestInsert(t *testing.T) {
	conn := new(MySql)
	data := map[string]string {"name":"grean", "password":"1234", "email":"hello@grean.com"}
	
	err, ok := conn.Insert(data)
	if !ok {
		panic(err)
	}
}

