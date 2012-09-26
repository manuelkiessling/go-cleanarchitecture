package repositories

import (
	"usecases"
	"domain"
	"fmt"
)

type dbHandler interface {
	Execute(query string) Row
}

type Row interface {
	Scan(dest ...interface{})
}

type DbUserRepo struct {
	DbHandler dbHandler
}

func (repo DbUserRepo) Store(user usecases.User) error {
	var isAdmin string
	if user.IsAdmin {
		isAdmin = "yes"
	} else {
		isAdmin = "no"
	}
	repo.DbHandler.Execute(fmt.Sprintf("INSERT INTO users (id, customer_id, is_admin) VALUES ('%d', '%d', '%v')", user.Id, user.Customer.Id, isAdmin))
	return nil
}

func (repo DbUserRepo) FindById(id int) usecases.User {
	row := repo.DbHandler.Execute(fmt.Sprintf("SELECT is_admin, customer_id FROM users WHERE id = '%d' LIMIT 1", id))
	u := usecases.User{}
	u.Id = id
	var isAdmin string
	var customerId int
	row.Scan(&isAdmin, &customerId)
	u.IsAdmin = false
	if isAdmin == "yes" {
		u.IsAdmin = true
	}
	u.Customer = domain.Customer{Id: customerId}
	return u
}
