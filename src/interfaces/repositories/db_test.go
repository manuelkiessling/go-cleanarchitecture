package interfaces_test

import (
	"interfaces"
	"testing"
	"usecases"
	"domain"
	"infrastructure"
)

func Test_UserRepository_Store(t *testing.T) {
	dbHandler := infrastructure.NewSqliteHandler("/var/tmp/test.sqlite")
	dbHandler.Execute("DROP TABLE users")
	dbHandler.Execute("CREATE TABLE users (id INTEGER, customer_id INTEGER, is_admin VARCHAR(3))")

	r := new(interfaces.DbUserRepo)
	r.DbHandler = dbHandler

	c := domain.Customer{}
	c.Id = 555
	u := usecases.User{}
	u.Id = 6
	u.IsAdmin = true
	u.Customer = c
	r.Store(u)

	u = r.FindById(6)
	if u.Id != 6 {
		t.Error()
	}
	if u.Customer.Id != 555 {
		t.Error()
	}
	if u.IsAdmin != true {
		t.Error()
	}
}
