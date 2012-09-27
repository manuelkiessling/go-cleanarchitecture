package infrastructure

import (
	"fmt"
	"testing"
)

func Test_SqliteHandler(t *testing.T) {
	h := NewSqliteHandler("/var/tmp/sqlitehandler_test.db")
	h.Execute("DROP TABLE foo")
	h.Execute("CREATE TABLE foo (id integer, name varchar(42))")
	h.Execute("INSERT INTO foo (id, name) VALUES (23, 'johndoe')")
	row := h.Query("SELECT id, name FROM foo LIMIT 1")
	var id int
	var name string
	row.Next()
	row.Scan(&id, &name)
	if id != 23 {
		fmt.Println(id)
		t.Error()
	}
}
