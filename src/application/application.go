package application

import (
	"domain"
	"errors"
	"fmt"
)

type UserRepository interface {
	Store(user User) error
	FindById(id int) User
}

type User struct {
	Id         int
	IsAdmin    bool
	CustomerId int // it's an id and not a pointer to a domain.Customer because we are referring to something on a different layer!
}

type Item struct {
	Id    int
	Name  string
	Value float64
}

type Logger interface {
	Log(message string) error // Adding "[Application]" in front of every application logging event is up to the implementing layer, while the content of the event is up to this layer
	// Severity levels could be added by this layer, but their meaning must be related to this layer
}

type OrderManager struct {
	UserRepository  UserRepository
	OrderRepository domain.OrderRepository
	ItemRepository  domain.ItemRepository
	Logger          Logger
}

func (orderManager *OrderManager) GetItemsForOrder(userId int, orderId int) ([]Item, error) {
	user := orderManager.UserRepository.FindById(userId)
	order := orderManager.OrderRepository.FindById(orderId)
	if user.CustomerId != order.Customer.Id {
		orderManager.Logger.Log(fmt.Sprintf("User #%i (customer #%i) is not allowed to see items in order #%i (which belongs to customer #%i)", user.Id, user.CustomerId, order.Id, order.Customer.Id))
		return nil, errors.New("Not allowed")
	}
	items := make([]Item, len(order.Items))
	for i := range order.Items {
		items[i].Id = order.Items[i].Id
		items[i].Name = order.Items[i].Name
		items[i].Value = order.Items[i].Value
	}
	return items, nil
}

func (orderManager *OrderManager) AddItemToOrder(userId int, orderId int, itemId int) error {
	user := orderManager.UserRepository.FindById(userId)
	order := orderManager.OrderRepository.FindById(orderId)
	if user.CustomerId != order.Customer.Id {
		orderManager.Logger.Log(fmt.Sprintf("User #%i (customer #%i) is not allowed to add items to order #%i (which belongs to customer #%i)", user.Id, user.CustomerId, order.Id, order.Customer.Id))
		return errors.New("Not allowed")
	}
	item := orderManager.ItemRepository.FindById(itemId)
	order.AddItem(item)
	orderManager.OrderRepository.Store(order)
	orderManager.Logger.Log(fmt.Sprintf("User added item '%s' (#%i) to order #%i", item.Name, item.Id, order.Id))
	return nil
}

type AdminOrderManager struct {
	OrderManager
}

func (adminOrderManager *AdminOrderManager) AddItemToOrder(userId int, orderId int, itemId int) error { // Even though the exact same things end up in the database, this is a different action!
	user := adminOrderManager.UserRepository.FindById(userId)
	order := adminOrderManager.OrderRepository.FindById(orderId)
	if !user.IsAdmin {
		adminOrderManager.Logger.Log(fmt.Sprintf("User #%i (customer #%i) is not allowed to add items to order #%i (which belongs to customer #%i)", user.Id, user.CustomerId, order.Id, order.Customer.Id))
		return errors.New("Not allowed")
	}
	item := adminOrderManager.ItemRepository.FindById(itemId)
	order.AddItem(item)
	adminOrderManager.OrderRepository.Store(order)
	adminOrderManager.Logger.Log(fmt.Sprintf("Admin added item '%s' (#%i) to order #%i", item.Name, item.Id, order.Id))
	return nil
}
