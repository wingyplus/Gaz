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
//	fmt.Printf ("%T\n", rows)
	
	for _, row := range rows {
		for _, col := range row.Data {
			fmt.Printf ("%v ", string(col.([]byte)))
		}
		fmt.Println()
	}
}
