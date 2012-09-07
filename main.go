package main

import (
	"usecases"
	"interfaces"
	"interfaces/repositories"
	"net/http"
)

func main() {
	orderInteractor := usecases.OrderInteractor{}
	orderInteractor.UserRepository = new(repositories.FakeUserRepo)
	orderInteractor.OrderRepository = new(repositories.FakeOrderRepo)
	orderInteractor.ItemRepository = new(repositories.FakeItemRepo)

	handler := interfaces.WebserviceHandler{}
	handler.OrderInteractor = orderInteractor

	http.HandleFunc("/orders", func(res http.ResponseWriter, req *http.Request) {
		handler.ShowOrder(res, req)
	})
	http.ListenAndServe(":8080", nil)
}
