package internal

import "github.com/charmbracelet/lipgloss"

var (
	modeTagStyle = lipgloss.NewStyle().Bold(true)

	editAddModeTagStyle = lipgloss.NewStyle().PaddingLeft(1).PaddingRight(1).
				Background(lipgloss.Color("#9ECE6A")).
				Foreground(lipgloss.Color("#15161E")).
				Inherit(modeTagStyle)

	browseModeTagStyle = lipgloss.NewStyle().PaddingLeft(1).PaddingRight(1).
				Background(lipgloss.Color("#7AA2F7")).
				Foreground(lipgloss.Color("#15161E")).
				Inherit(modeTagStyle)

	itemStyle                  = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle          = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("120"))
	completedSelectedItemStyle = lipgloss.NewStyle().Strikethrough(true).PaddingLeft(2).Foreground(lipgloss.Color("120"))
	completedItemStyle         = lipgloss.NewStyle().Strikethrough(true).PaddingLeft(4)
)
