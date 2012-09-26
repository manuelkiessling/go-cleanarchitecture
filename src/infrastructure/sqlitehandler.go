package infrastructure

import (
	"fmt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"interfaces/repositories"
)

type SqliteHandler struct {
	Conn *sql.DB
}

func (handler *SqliteHandler) Execute(statement string) repositories.Row {
	fmt.Println(statement)
	rows, err := handler.Conn.Query(statement)
	if err != nil {
		fmt.Println(err)
		return new(SqliteRow)
	}
	rows.Next()
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

func NewSqliteHandler(dbfileName string) *SqliteHandler {
	conn, _ := sql.Open("sqlite3", dbfileName)
	sqliteHandler := new(SqliteHandler)
	sqliteHandler.Conn = conn
	return sqliteHandler
}
