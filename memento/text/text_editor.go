package text

import (
	"github.com/mvrdgs/patterns/memento/caretaker"
	"github.com/mvrdgs/patterns/memento/memento"
)

type mementoImpl struct {
	content string
	editor  *Editor
}

func (m *mementoImpl) Restore() {
	m.editor.content = m.content
}

func NewMemento(c string, e *Editor) memento.Memento {
	return &mementoImpl{
		content: c,
		editor:  e,
	}
}

type Editor struct {
	content string
}

func NewTextEditor() *Editor {
	return &Editor{}
}

func (e *Editor) SetContent(content string) {
	e.content = content
	m := NewMemento(e.content, e)
	caretaker.Get().Store(m)
}

func (e *Editor) GetContent() string {
	return e.content
}
