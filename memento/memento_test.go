package main

import (
	"testing"

	"github.com/mvrdgs/patterns/memento/caretaker"
	"github.com/mvrdgs/patterns/memento/text"
	"github.com/stretchr/testify/assert"
)

func TestMemento(t *testing.T) {
	textEditor := text.NewTextEditor()
	ct := caretaker.Get()

	contents := []string{"Hello", "Hello,", "Hello, World", "Hello, World!!!"}

	for _, content := range contents {
		textEditor.SetContent(content)
		assert.Equal(t, content, textEditor.GetContent())
	}

	for i := len(contents) - 1; i > 0; i-- {
		ct.Undo()
		assert.Equal(t, contents[i-1], textEditor.GetContent())
	}

	ct.Undo()
	assert.Equal(t, "Hello", textEditor.GetContent())

	for i := 0; i < len(contents)-1; i++ {
		ct.Redo()
		assert.Equal(t, contents[i+1], textEditor.GetContent())
	}

	ct.Redo()
	assert.Equal(t, "Hello, World!!!", textEditor.GetContent())

	ct.Undo()
	assert.Equal(t, "Hello, World", textEditor.GetContent())
	ct.Undo()
	assert.Equal(t, "Hello,", textEditor.GetContent())

	textEditor.SetContent("Hello, people")
	assert.Equal(t, "Hello, people", textEditor.GetContent())

	ct.Undo()
	assert.Equal(t, "Hello,", textEditor.GetContent())

	ct.Redo()
	assert.Equal(t, "Hello, people", textEditor.GetContent())
}
