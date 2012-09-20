package infrastructure

type SqliteHandler struct {
}

func (db SqliteHandler) init() {
}

func (db SqliteHandler) Begin() {
}

func (db SqliteHandler) Prepare(statement string) Statement {
}

func (db SqliteHandler) Commit() {}

type SqliteStatement struct {
}

func (stmt SqliteStatement) Execute(values ...string) SqliteRow {
	return SqliteRow{}
}

type SqliteRow struct {}

func (r SqliteRow) getFieldValue(name string) string {
	return val
}
