package main

import (
	"github.com/mvrdgs/patterns/observer/customer"
	"github.com/mvrdgs/patterns/observer/product"
)

func main() {
	prod := product.New(1, "Shoes", 100)

	cust1 := customer.NewCustomer(1)
	cust2 := customer.NewCustomer(2)
	cust3 := customer.NewCustomer(3)
	cust4 := customer.NewCustomer(4)
	cust5 := customer.NewCustomer(5)

	prod.AttachCustomer(cust1, 90)
	prod.AttachCustomer(cust2, 80)
	prod.AttachCustomer(cust3, 70)
	prod.AttachCustomer(cust4, 60)
	prod.AttachCustomer(cust5, 50)

	prod.UpdatePrice(80)
	prod.DetachCustomer(cust1)
	prod.DetachCustomer(cust2)
	prod.UpdatePrice(60)

	prod.AttachCustomer(cust5, 110)
	prod.UpdatePrice(100)
}
