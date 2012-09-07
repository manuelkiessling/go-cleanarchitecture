package repositories

import (
	"usecases"
	"domain"
	"strconv"
)

type FakeUserRepo struct{}

func (r FakeUserRepo) Store(user usecases.User) error {
	return nil
}

func (r FakeUserRepo) FindById(id int) usecases.User {
	u := usecases.User{}
	u.Id = 111
	u.Customer = domain.Customer{Id: 555}
	u.IsAdmin = false
	return u
}

type FakeOrderRepo struct{}

func (r FakeOrderRepo) Store(order domain.Order) error {
	return nil
}

func (r FakeOrderRepo) FindById(id int) domain.Order {
	c := domain.Customer{}
	c.Id = 555
	c.Name = "Cus T. Omer"
	o := domain.Order{}
	o.Id = 87654
	o.Customer = c
	o.Items = make([]domain.Item, 5)
	for i := 0; i < 5; i++ {
		it := domain.Item{}
		it.Id = 40 + i
		it.Name = "Item Number #" + strconv.Itoa(40+i)
		it.Value = 10.45 + float64(i)
		it.Available = true
		o.Items[i] = it
	}
	return o
}

type FakeItemRepo struct{}

func (r FakeItemRepo) Store(item domain.Item) error {
	return nil
}

func (r FakeItemRepo) FindById(id int) domain.Item {
	return domain.Item{}
}
