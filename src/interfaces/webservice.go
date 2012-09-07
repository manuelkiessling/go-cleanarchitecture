package interfaces

import (
	"application"
	"io"
	"net/http"
	"strconv"
)

// This Layer knows about HTTP requests and responses, but not about listening to tcp etc.

type Webservice struct {
	OrderManager application.OrderManager
}

func (webservice Webservice) HandleShowOrder(res http.ResponseWriter, req *http.Request) {
	userId, _ := strconv.Atoi(req.FormValue("userId")) // We could handle HTTP auth issues here (Session, API token check etc), but whether or not a user may see an order totally is an issue of the application layer!
                                                     // litmus test: we would have to re-implement the auth check if we wrote a console script as the frontend for the app
	orderId, _ := strconv.Atoi(req.FormValue("orderId"))
	items, _ := webservice.OrderManager.GetItemsForOrder(userId, orderId)
	for i := range items {
		io.WriteString(res, "item id: "+strconv.Itoa(items[i].Id)+"\n")
		io.WriteString(res, "item name: "+items[i].Name+"\n")
		io.WriteString(res, "item value: "+strconv.FormatFloat(items[i].Value, 'f', 2, 64)+"\n")
	}
}
