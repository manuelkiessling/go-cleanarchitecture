package interfaces_test

import (
	"domain"
	_ "fmt"
	"infrastructure"
	"interfaces"
	"testing"
	"usecases"
)

func Test_UserRepository(t *testing.T) {
	dbHandler := infrastructure.NewSqliteHandler("/var/tmp/test1.sqlite")
	dbHandler.Execute("DROP TABLE users")
	dbHandler.Execute("DROP TABLE customers")
	dbHandler.Execute("CREATE TABLE users (id INTEGER, customer_id INTEGER, is_admin VARCHAR(3))")
	dbHandler.Execute("CREATE TABLE customers (id INTEGER, name VARCHAR(42))")

	handlers := make(map[string]interfaces.DbHandler)
	handlers["DbUserRepo"] = dbHandler
	handlers["DbCustomerRepo"] = dbHandler

	ur := interfaces.NewDbUserRepo(handlers)

	c := domain.Customer{}
	c.Id = 555
	c.Name = "John Doe"
	u := usecases.User{}
	u.Id = 6
	u.IsAdmin = true
	u.Customer = c
	ur.Store(u)

	u = ur.FindById(6)
	if u.Id != 6 {
		t.Error()
	}
	if u.IsAdmin != true {
		t.Error()
	}
	if u.Customer.Id != 555 {
		t.Error()
	}
	if u.Customer.Name != "John Doe" {
		t.Error()
	}
}

func Test_OrderRepository(t *testing.T) {
	dbHandler := infrastructure.NewSqliteHandler("/var/tmp/test2.sqlite")
	dbHandler.Execute("DROP TABLE users")
	dbHandler.Execute("DROP TABLE customers")
	dbHandler.Execute("DROP TABLE orders")
	dbHandler.Execute("DROP TABLE items")
	dbHandler.Execute("DROP TABLE items2orders")
	dbHandler.Execute("CREATE TABLE users (id INTEGER, customer_id INTEGER, is_admin VARCHAR(3))")
	dbHandler.Execute("CREATE TABLE customers (id INTEGER, name VARCHAR(42))")
	dbHandler.Execute("CREATE TABLE orders (id INTEGER, customer_id INTEGER)")
	dbHandler.Execute("CREATE TABLE items (id INTEGER, name VARCHAR(42), value FLOAT, available VARCHAR(3))")
	dbHandler.Execute("CREATE TABLE items2orders (item_id INTEGER, order_id INTEGER)")

	handlers := make(map[string]interfaces.DbHandler)
	handlers["DbUserRepo"] = dbHandler
	handlers["DbCustomerRepo"] = dbHandler
	handlers["DbOrderRepo"] = dbHandler
	handlers["DbItemRepo"] = dbHandler

	ur := interfaces.NewDbUserRepo(handlers)
	or := interfaces.NewDbOrderRepo(handlers)
	ir := interfaces.NewDbItemRepo(handlers)

	c := domain.Customer{}
	c.Id = 555
	c.Name = "John Doe"
	u := usecases.User{}
	u.Id = 6
	u.IsAdmin = true
	u.Customer = c
	ur.Store(u)

	i1 := domain.Item{}
	i1.Id = 101
	i1.Name = "Car"
	i1.Value = 125.45
	i1.Available = true
	ir.Store(i1)

	i2 := domain.Item{}
	i2.Id = 102
	i2.Name = "Table"
	i2.Value = 432.56
	i2.Available = false
	ir.Store(i2)

	i3 := domain.Item{}
	i3.Id = 103
	i3.Name = "Chair"
	i3.Value = 52.31
	i3.Available = true
	ir.Store(i3)

	o := domain.Order{}
	o.Id = 39
	o.Customer = c
	o.Add(i1)
	o.Add(i2) //fails because it's not available
	o.Add(i3)
	or.Store(o)

	o = or.FindById(39)

	if o.Customer.Name != "John Doe" {
		t.Error()
	}

	items := o.Items
	if items[1].Id != 103 {
		t.Error()
	}
}
