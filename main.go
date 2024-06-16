package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	width  int
	height int
}

func New() *model {
	m := new(model)
	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
    return m, cmd
}

func (m model) View() string {
    return ""
}

func main() {
	f, err := tea.LogToFile("debug.log", "[DEBUG]")
	if err != nil {
		log.Fatalf("err: %s", err)
	}
	defer f.Close()

	p := tea.NewProgram(New(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatalf("err: %s", err)
	}
}
