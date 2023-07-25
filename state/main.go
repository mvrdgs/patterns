package main

import (
	"github.com/mvrdgs/patterns/state/items"
	"github.com/mvrdgs/patterns/state/machine"
)

func main() {
	i := map[int]items.Item{
		1: {
			Name:  "Capuccino",
			Price: 200,
		},
		2: {
			Name:  "Chocolate quente",
			Price: 150,
		},
	}

	m := machine.New(i)

	m.InsertCoin(100)
	m.InsertCoin(100)
	m.Next()
	m.Select(1)
	m.Select(2)
	m.Next()
	m.Cancel()
	m.Cancel()
	m.InsertCoin(100)
	m.InsertCoin(50)
	m.InsertCoin(25)
	m.Next()
	m.Select(1)
	m.Select(2)
	m.Next()
	m.Next()
}
