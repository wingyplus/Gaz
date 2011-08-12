package gaz

import (	
	"fmt"
	"testing"
	"github.com/bmizerany/assert"
	mymy "github.com/ziutek/mymysql"
)



func TestQuery(t *testing.T) {
	db := new(Connection).DB("test")
	rows := db.Query("select * from User").([]*mymy.Row)
	
	for _, row := range rows {
		for _, col := range row.Data {
			fmt.Printf ("%v ", string(col.([]byte)))
		}
		fmt.Println()
	}
}

func TestInsert(t *testing.T) {
	dataset := new(Connection).DB("test").C("User")
	data := map[string]string {"name":"grean", "password":"1234", "email":"hello@grean.com"}
	
	err, ok := dataset.Insert(data)
	if !ok {
		panic(err)
	}
}

func TestGet(t *testing.T) {
	dataset := new(Connection).DB("test").C("User")
	row := dataset.Get("1").(*mymy.Row)
	assert.Equal(t, "barzaar", row.Str(1))
	assert.Equal(t, "hello@bazaar.com", row.Str(3))
}

func TestFindOne(t *testing.T) {
	dataset := new(Connection).DB("test").C("User")
	row := dataset.FindOne(Params{"name":"grean"}).(*mymy.Row)
	assert.Equal(t, "hello1@grean.com", row.Str(3))
}

func TestExtractField(t *testing.T) {
	
}