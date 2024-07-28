package internal

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type todo struct {
	Completed   bool   `json:"completed"`
	Description string `json:"description"`
}

func (t todo) FilterValue() string { return t.Description }
func (t todo) String() string      { return t.Description }
func (t *todo) ToggleComplete()    { t.Completed = !t.Completed }

func NewTodo(description string) todo { return todo{Description: description} }

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
		if t.Completed {
			return completedItemStyle.Render(strings.Join(s, " "))
		}
		return itemStyle.Render(strings.Join(s, " "))

	}

	if index == m.Index() {
		fn = func(s ...string) string {
			if t.Completed {
				return completedSelectedItemStyle.Render("> " + strings.Join(s, " "))
			}

			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(t.String()))
}
