package interfaces

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/manuelkiessling/go-cleanarchitecture/src/usecases"
)

type OrderInteractor interface {
	Items(userId, orderId int) ([]usecases.Item, error)
	Add(userId, orderId, itemId int) error
}

type WebserviceHandler struct {
	OrderInteractor OrderInteractor
}

func (handler WebserviceHandler) ShowOrder(res http.ResponseWriter, req *http.Request) {
	userId, err := strconv.Atoi(req.FormValue("userId"))
	if err != nil {
		panic(err)
	}
	orderId, err := strconv.Atoi(req.FormValue("orderId"))
	if err != nil {
		panic(err)
	}
	items, err := handler.OrderInteractor.Items(userId, orderId)
	if err != nil {
		panic(err)
	}
	for _, item := range items {
		io.WriteString(res, fmt.Sprintf("item id: %d\n", item.Id))
		io.WriteString(res, fmt.Sprintf("item name: %v\n", item.Name))
		io.WriteString(res, fmt.Sprintf("item value: %f\n", item.Value))
	}
}
