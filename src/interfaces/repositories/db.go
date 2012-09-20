package repositories

import (
	"usecases"
	"domain"
	"strconv"
)

type dbHandler interface {
	Begin()
	Prepare(query string) Statement
	Commit()
}

type Statement interface {
	Execute(values ...string) Row
}

type Row interface {
	getFieldValue(name string) string
}

type UserRepo struct {
	Db dbHandler
}

func (repo UserRepo) Store(user usecases.User) error {
	repo.Db.Begin()
	stmt := repo.Db.Prepare("INSERT INTO users (id, customer_id, is_admin) VALUES (?, ?, ?)")
	var isAdmin string
	if user.IsAdmin {
		isAdmin = "1"
	} else {
		isAdmin = "0"
	}
	_ = stmt.Execute(strconv.Itoa(user.Id), strconv.Itoa(user.Customer.Id), isAdmin)
	repo.Db.Commit()
	return nil
}

func (repo UserRepo) FindById(id int) usecases.User {
	repo.Db.Begin()
	stmt := repo.Db.Prepare("SELECT is_admin, customer_id FROM users WHERE id = ? LIMIT 1")
	row := stmt.Execute(string(id))
	u := usecases.User{}
	u.Id = id
	u.Customer = domain.Customer{Id: 555}
	isAdmin := row.getFieldValue("is_admin")
	u.IsAdmin = false
	if isAdmin == "1" {
		u.IsAdmin = true
	}
	return u
}
