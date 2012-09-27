package infrastructure

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"interfaces"
)

type SqliteHandler struct {
	Conn *sql.DB
}

func (handler *SqliteHandler) Execute(statement string) {
	handler.Conn.Exec(statement)
}

func (handler *SqliteHandler) Query(statement string) interfaces.Row {
	//fmt.Println(statement)
	rows, err := handler.Conn.Query(statement)
	if err != nil {
		fmt.Println(err)
		return new(SqliteRow)
	}
	row := new(SqliteRow)
	row.Rows = rows
	return row
}

type SqliteRow struct {
	Rows *sql.Rows
}

func (r SqliteRow) Scan(dest ...interface{}) {
	r.Rows.Scan(dest...)
}

func (r SqliteRow) Next() bool {
	return r.Rows.Next()
}

func NewSqliteHandler(dbfileName string) *SqliteHandler {
	conn, _ := sql.Open("sqlite3", dbfileName)
	sqliteHandler := new(SqliteHandler)
	sqliteHandler.Conn = conn
	return sqliteHandler
}
