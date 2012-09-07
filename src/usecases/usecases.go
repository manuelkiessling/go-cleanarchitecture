package usecases

import (
	"domain"
	"fmt"
)

type UserRepository interface {
	Store(user User) error
	FindById(id int) User
}

type User struct {
	Id         int
	IsAdmin    bool
	Customer   domain.Customer
}

type Item struct {
        Id    int
        Name  string
        Value float64
}

type Logger interface {
	Log(message string) error
}

type OrderInteractor struct {
	UserRepository  UserRepository
	OrderRepository domain.OrderRepository
	ItemRepository  domain.ItemRepository
	Logger          Logger
}

func (interactor *OrderInteractor) Items(userId, orderId int) ([]Item, error) {
	var items []Item
	user := interactor.UserRepository.FindById(userId)
	order := interactor.OrderRepository.FindById(orderId)
	if user.Customer.Id != order.Customer.Id {
		message := "User #%i (customer #%i) "
		message += "is not allowed to see items "
		message += "in order #%i (of customer #%i)"
		err := fmt.Errorf(message,
		                  user.Id,
		                  user.Customer.Id,
		                  order.Id,
		                  order.Customer.Id)
		interactor.Logger.Log(err.Error())
		items = make([]Item, 0)
		return items, err
	}
	items = make([]Item, len(order.Items))
	for i, item := range order.Items {
		items[i] = Item{item.Id, item.Name, item.Value}
	}
	return items, nil
}

func (interactor *OrderInteractor) Add(userId, orderId, itemId int) error {
	var message string
	user := interactor.UserRepository.FindById(userId)
	order := interactor.OrderRepository.FindById(orderId)
	if user.Customer.Id != order.Customer.Id {
	message  = "User #%i (customer #%i) "
	message += "is not allowed to add items "
	message += "to order #%i (of customer #%i)"
		err := fmt.Errorf(message,
		                  user.Id,
		                  user.Customer.Id,
		                  order.Id,
		                  order.Customer.Id)
		interactor.Logger.Log(err.Error())
		return err
	}
	item := interactor.ItemRepository.FindById(itemId)
	if domainErr := order.Add(item); domainErr != nil {
		message  = "Could not add item #%i "
		message += "to order #%i (of customer #%i) "
		message += "as user #%i because a business "
		message += "rule was violated: '%s'"
		err := fmt.Errorf(message,
		                  item.Id,
		                  order.Id,
		                  order.Customer.Id,
		                  user.Id,
		                  domainErr.Error())
		interactor.Logger.Log(err.Error())
		return err
	}
	interactor.OrderRepository.Store(order)
	interactor.Logger.Log(fmt.Sprintf(
	                    "User added item '%s' (#%i) to order #%i",
                      item.Name, item.Id, order.Id))
	return nil
}

type AdminOrderInteractor struct {
	OrderInteractor
}

func (interactor *AdminOrderInteractor) Add(userId, orderId, itemId int) error {
	var message string
	user := interactor.UserRepository.FindById(userId)
	order := interactor.OrderRepository.FindById(orderId)
	if !user.IsAdmin {
		message  = "User #%i (customer #%i) "
		message += "is not allowed to add items "
		message += "to order #%i (of customer #%i), "
		message += "because he is not an administrator"
		err := fmt.Errorf(message,
		                  user.Id,
		                  user.Customer.Id,
		                  order.Id,
		                  order.Customer.Id)
		interactor.Logger.Log(err.Error())
		return err
	}
	item := interactor.ItemRepository.FindById(itemId)
	if domainErr := order.Add(item); domainErr != nil {
		message  = "Could not add item #%i "
		message += "to order #%i (of customer #%i) "
		message += "as user #%i because a business "
		message += "rule was violated: '%s'"
		err := fmt.Errorf(message,
		                  item.Id,
		                  order.Id,
		                  order.Customer.Id,
		                  user.Id,
                      domainErr.Error())
		interactor.Logger.Log(err.Error())
		return err
	}
	interactor.OrderRepository.Store(order)
	interactor.Logger.Log(fmt.Sprintf(
		"Admin added item '%s' (#%i) to order #%i",
		item.Name, item.Id, order.Id))
	return nil
}
