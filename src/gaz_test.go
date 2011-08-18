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
	data := map[string]interface{} {"email":"hello@grean.com", "name":"grean", "password":"1234"}
	
	err, ok := dataset.Insert(data)
	assert.Equal(t, true, ok)
	assert.Equal(t, nil, err)
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
	dataset := new(Connection).DB("test").C("User")
	field := dataset.extractField()
	assert.Equal(t, "int(11)", field["id"])
	assert.Equal(t, "varchar(40)", field["name"])
	assert.Equal(t, "varchar(10)", field["password"])
	assert.Equal(t, "varchar(100)", field["email"])
}

func TestMappingData(t *testing.T) {
	test_data := map[string]interface{} {"id":1, "username": "hylo", "score":1.23}
	result := maptype(test_data)
	assert.Equal(t, "'hylo'", result["username"])
	assert.Equal(t, "1", result["id"])
	assert.Equal(t, "1.23", result["score"])
}