package states

import "github.com/mvrdgs/patterns/state/items"

type Interface interface {
	InsertCoin(int)
	Select(int)
	Remove(int)
	Cancel()
	Next()
}

type state struct {
	currentState Interface

	InsertedMoney int
	ItemList      map[int]items.Item
	SelectedItems []items.Item
}

func New(itemList map[int]items.Item) Interface {
	s := &state{
		ItemList:      itemList,
		SelectedItems: []items.Item{},
	}
	s.currentState = NewReady(s)
	return s
}

func (s *state) SetState(state Interface) {
	s.currentState = state
}

func (s *state) InsertCoin(money int) {
	s.currentState.InsertCoin(money)
}

func (s *state) Select(itemID int) {
	s.currentState.Select(itemID)
}

func (s *state) Remove(index int) {
	s.currentState.Remove(index)
}

func (s *state) Cancel() {
	s.currentState.Cancel()
}

func (s *state) Next() {
	s.currentState.Next()
}
