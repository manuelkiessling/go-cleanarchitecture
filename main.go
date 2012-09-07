package main

import (
	"application"
	"interfaces"
	"interfaces/repositories"
	"net/http"
)

func main() {
	orderManager := application.OrderManager{}
	orderManager.UserRepository = new(repositories.FakeUserRepo)
	orderManager.OrderRepository = new(repositories.FakeOrderRepo)
	orderManager.ItemRepository = new(repositories.FakeItemRepo)

	webservice := interfaces.Webservice{}
	webservice.OrderManager = orderManager

	http.HandleFunc("/orders", func(res http.ResponseWriter, req *http.Request) {
		webservice.HandleShowOrder(res, req)
	})
	http.ListenAndServe(":8080", nil)
}
