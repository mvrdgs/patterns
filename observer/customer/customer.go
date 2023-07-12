package customer

import "fmt"

type Customer interface {
	Update()
	GetID() int64
}

type customer struct {
	ID int64
}

func NewCustomer(id int64) Customer {
	return &customer{id}
}

func (c *customer) Update() {
	fmt.Printf("Customer %d notified\n", c.ID)
}

func (c *customer) GetID() int64 {
	return c.ID
}
