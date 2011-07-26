package gaz

import (
	"testing"
	"github.com/bmizerany/assert"
	mymy "github.com/ziutek/mymysql"
)



func TestAsserts(t *testing.T) {
	db := mymy.New("tcp", "", "127.0.0.1:3306", "root", "root", "test")
	defer db.Close()
	conn := &MySql{db}

	assert.Equal(t, nil, conn.Connect())
}