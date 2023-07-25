package machine

import (
	"github.com/mvrdgs/patterns/state/items"
	"github.com/mvrdgs/patterns/state/states"
)

type Machine struct {
	state states.Interface
}

func New(itemList map[int]items.Item) states.Interface {
	m := Machine{
		state: states.New(itemList),
	}

	return &m
}

func (m *Machine) InsertCoin(money int) {
	m.state.InsertCoin(money)
}

func (m *Machine) Select(itemID int) {
	m.state.Select(itemID)
}

func (m *Machine) Remove(index int) {
	m.state.Remove(index)
}

func (m *Machine) Cancel() {
	m.state.Cancel()
}

func (m *Machine) Next() {
	m.state.Next()
}
