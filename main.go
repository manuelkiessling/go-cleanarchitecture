package main

import (
	"usecases"
	"interfaces"
	"infrastructure"
	"net/http"
)

func main() {
	dbHandler := infrastructure.NewSqliteHandler("/var/tmp/production.sqlite")

	orderInteractor := new(usecases.OrderInteractor)
	orderInteractor.UserRepository = new(interfaces.DbUserRepo)
	orderInteractor.ItemRepository = new(interfaces.DbItemRepo)
	orderInteractor.OrderRepository = new(interfaces.DbOrderRepo)

	orderInteractor.UserRepository.DbHandler = dbHandler
	orderInteractor.ItemRepository.DbHandler = dbHandler
	orderInteractor.OrderRepository.DbHandler = dbHandler

	handler := interfaces.WebserviceHandler{}
	handler.OrderInteractor = orderInteractor

	http.HandleFunc("/orders", func(res http.ResponseWriter, req *http.Request) {
		handler.ShowOrder(res, req)
	})
	http.ListenAndServe(":8080", nil)
}
