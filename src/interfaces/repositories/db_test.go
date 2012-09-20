package repositories

import (
	"testing"
	"fmt"
	"usecases"
	"domain"
)

type testBag struct {
	LastStatement string
	LastValues []string
}

var bag testBag

type testDbHandler struct {
}

func (db testDbHandler) init() {
	
}

func (db testDbHandler) Begin() {}

func (db testDbHandler) Prepare(statement string) Statement {
	prepared := testStatement{}
	prepared.Statement = statement
	return prepared
}

func (db testDbHandler) Commit() {}

type testStatement struct {
	Statement string
}

func (stmt testStatement) Execute(values ...string) Row {
	fmt.Println("executing " + stmt.Statement + "with")
	bag.LastStatement = stmt.Statement
	for i, value := range(values) {
		bag.LastValues[i] = value
		fmt.Printf("%d - %v\n", i, value)
	}
	return testRow{}
}

type testRow struct {}

func (r testRow) getFieldValue(name string) string {
	val := "4"
	return val
}

func Test_UserRepository_FindById(t *testing.T) {
	bag.LastValues = make([]string, 10)
	r := UserRepo{}
	r.Db = testDbHandler{}
	u := usecases.User{}
	c := domain.Customer{}
	c.Id = 999
	c.Name = "foo"
	u.Id = 6
	u.IsAdmin = true
	u.Customer = c
	r.Store(u)
	if bag.LastStatement != "INSERT INTO users (id, customer_id, is_admin) VALUES (?, ?, ?)" {
		t.Error()
	}
	if bag.LastValues[0] != "6" {
		t.Error()
	}
}
