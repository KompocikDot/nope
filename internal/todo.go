package internal

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type todo struct {
	completed   bool
	description string
}

func (t todo) FilterValue() string { return t.description }
func (t todo) String() string      { return t.description }
func (t *todo) ToggleComplete()    { t.completed = !t.completed }

func NewTodo(description string) todo { return todo{description: description} }

type todoDelegate struct{}

func (d todoDelegate) Height() int                             { return 1 }
func (d todoDelegate) Spacing() int                            { return 0 }
func (d todoDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }

func (d todoDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	t, ok := listItem.(todo)
	if !ok {
		return
	}

	fn := func(s ...string) string {
		if t.completed {
			return completedItemStyle.Render(strings.Join(s, " "))
		}
		return itemStyle.Render(strings.Join(s, " "))

	}

	if index == m.Index() {
		fn = func(s ...string) string {
			if t.completed {
				return completedSelectedItemStyle.Render("> " + strings.Join(s, " "))
			}

			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(t.String()))
}
