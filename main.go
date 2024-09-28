package main

import (
	// "github.com/xV0lk/myvpn/cmd"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/xV0lk/myvpn/tui"
)

func main() {
	// cmd.Execute()
	p := tea.NewProgram(tui.InitialListModel())
	if _, error := p.Run(); error != nil {
		fmt.Printf("There was an error running the program: %v", error)
		os.Exit(1)
	}

}
