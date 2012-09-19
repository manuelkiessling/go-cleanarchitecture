package interfaces

import (
	"usecases"
	"io"
	"net/http"
	"strconv"
)

type OrderInteractor interface {
	Items(userId, orderId int) ([]usecases.Item, error)
	Add(userId, orderId, itemId int) error
}

type WebserviceHandler struct {
	OrderInteractor OrderInteractor
}

func (handler WebserviceHandler) ShowOrder(res http.ResponseWriter, req *http.Request) {
	userId, _ := strconv.Atoi(req.FormValue("userId"))
	orderId, _ := strconv.Atoi(req.FormValue("orderId"))
	items, _ := handler.OrderInteractor.Items(userId, orderId)
	for i := range items {
		io.WriteString(res, "item id: "+strconv.Itoa(items[i].Id)+"\n")
		io.WriteString(res, "item name: "+items[i].Name+"\n")
		io.WriteString(res, "item value: "+strconv.FormatFloat(items[i].Value, 'f', 2, 64)+"\n")
	}
}
