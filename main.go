package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	input  tea.Model
	width  int
	height int
}

func New() *model {
	m := new(model)
	m.input = NewQueryInput()
	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.input.View()
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
