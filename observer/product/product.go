package product

import (
	"github.com/mvrdgs/patterns/observer/customer"
	"github.com/mvrdgs/patterns/observer/notifier"
)

type Interface interface {
	GetID() int64
	GetName() string
	GetPrice() float64
	UpdatePrice(price float64)
	AttachCustomer(c customer.Customer, price float64)
	DetachCustomer(c customer.Customer)
}

type product struct {
	notifier notifier.Notifier
	id       int64
	name     string
	price    float64
}

func New(id int64, name string, price float64) Interface {
	return &product{
		notifier: notifier.NewPriceNotifier(),
		id:       id,
		name:     name,
		price:    price,
	}
}

func (p *product) GetID() int64 {
	return p.id
}

func (p *product) GetName() string {
	return p.name
}

func (p *product) GetPrice() float64 {
	return p.price
}

func (p *product) UpdatePrice(price float64) {
	p.price = price
	p.notifier.Notify(price)
}

func (p *product) AttachCustomer(c customer.Customer, price float64) {
	p.notifier.Attach(c, price)
}

func (p *product) DetachCustomer(c customer.Customer) {
	p.notifier.Detach(c)
}
