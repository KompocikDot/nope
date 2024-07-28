package internal

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type NopeMode string

const (
	MODE_BROWSE = "BROWSE"
	MODE_EDIT   = "EDIT"
	MODE_INSERT = "INSERT"
)

type NopeModel struct {
	list  list.Model
	input textinput.Model

	mode NopeMode
}

func (m NopeModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m NopeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		if m.mode != MODE_BROWSE {
			return nonBrowseActions(m, msg)
		}

		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "d":
			m.list.RemoveItem(m.list.Index())
		case "e":
			m.mode = MODE_EDIT
			current := m.list.SelectedItem().(todo)

			m.input.SetValue(current.description)
			m.input.Focus()

		case "enter":
			i, ok := m.list.SelectedItem().(todo)
			index := m.list.Index()
			if !ok {
				return m, nil
			}

			i.ToggleComplete()
			m.list.SetItem(index, i)
			m.mode = MODE_BROWSE

		case "i":
			m.input.Focus()
			m.mode = MODE_INSERT
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func nonBrowseActions(m NopeModel, msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "esc":
		m.mode = MODE_BROWSE
		return m, nil
	case "enter":
		newValue := m.input.Value()

		if m.mode == MODE_EDIT {
			m.list.RemoveItem(m.list.Index())
			m.list.InsertItem(m.list.Index(), NewTodo(m.input.Value()))
		} else {
			m.list.InsertItem(len(m.list.Items()), NewTodo(m.input.Value()))
			if newValue == "" {
				m.input.Blur()
				break
			}
		}

		m.mode = MODE_BROWSE
		m.input.SetValue("")
	}

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m NopeModel) View() string {
	out := m.list.View() + "\n"

	if m.mode != MODE_BROWSE {
		return out + editAddModeTagStyle.Render(string(m.mode)) + " " + m.input.View()
	}
	return out + browseModeTagStyle.Render(string(m.mode))
}

func NewNopeModel() *NopeModel {
	li := list.New([]list.Item{
		todo{description: "add inputs"},
		todo{description: "delete items"},
	}, todoDelegate{}, 14, 20)

	li.SetShowStatusBar(false)
	li.SetShowTitle(false)

	input := textinput.New()

	return &NopeModel{
		list:  li,
		input: input,
		mode:  MODE_BROWSE,
	}
}
