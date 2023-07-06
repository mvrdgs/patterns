package caretaker

import (
	"sync"

	"github.com/mvrdgs/patterns/memento/memento"
)

var (
	once sync.Once
	ct   Caretaker
)

type Caretaker interface {
	Store(m memento.Memento)
	Undo()
	Redo()
}

type caretaker struct {
	index    int
	mementos []memento.Memento
}

func (c *caretaker) Store(m memento.Memento) {
	if c.index < len(c.mementos)-1 {
		c.mementos = append(c.mementos[:c.index+1], m)
	} else {
		c.mementos = append(c.mementos, m)
	}

	c.index = len(c.mementos) - 1
}

func (c *caretaker) Undo() {
	if c.index > 0 {
		c.index--
	}

	m := c.mementos[c.index]
	m.Restore()
}

func (c *caretaker) Redo() {
	if c.index < len(c.mementos)-1 {
		c.index++
	}

	m := c.mementos[c.index]
	m.Restore()
}

func Get() Caretaker {
	once.Do(func() {
		ct = &caretaker{}
	})

	return ct
}
