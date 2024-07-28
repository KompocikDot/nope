package main

import (
	"fmt"
	"os"

	"github.com/KompocikDot/nope/internal"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	if _, err := tea.NewProgram(*internal.NewNopeModel()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
