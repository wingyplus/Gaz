package gaz

import (
	"testing"
	"github.com/bmizerany/assert"
	mymy "github.com/ziutek/mymysql"
)



func TestConnections(t *testing.T) {
	db := mymy.New("tcp", "", "127.0.0.1:3306", "root", "root", "test")
	conn := &MySql{db}
	defer conn.db.Close()
	assert.Equal(t, nil, conn.db.Connect())
}