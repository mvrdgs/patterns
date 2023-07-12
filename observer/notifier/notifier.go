package notifier

import (
	"fmt"

	"github.com/mvrdgs/patterns/observer/customer"
)

type Notifier interface {
	Attach(obs customer.Customer, wantedPrice float64)
	Detach(obs customer.Customer)
	Notify(price float64)
}

type Observer struct {
	customer    customer.Customer
	wantedPrice float64
}

type notifier struct {
	observers []Observer
}

func NewPriceNotifier() Notifier {
	return &notifier{}
}

func (n *notifier) Attach(obs customer.Customer, wantedPrice float64) {
	n.Detach(obs)

	observer := Observer{obs, wantedPrice}
	n.observers = append(n.observers, observer)
}

func (n *notifier) Detach(obs customer.Customer) {
	for i, o := range n.observers {
		if o.customer.GetID() == obs.GetID() {
			n.observers = append(n.observers[:i], n.observers[i+1:]...)
			return
		}
	}
}

func (n *notifier) Notify(price float64) {
	fmt.Print("Notifying customers...\n")
	for _, o := range n.observers {
		if price <= o.wantedPrice {
			o.customer.Update()
		}
	}
	fmt.Print("Done!\n")
}
