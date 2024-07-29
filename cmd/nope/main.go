package main

import (
	"fmt"
	"os"

	"github.com/KompocikDot/nope/internal"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	savedTodos := internal.ReadTodos()
	model := internal.NewNopeModel(savedTodos)

	if _, err := tea.NewProgram(*model).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
