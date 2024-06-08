package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	width  int
	height int
	active applicationStatus
}

type applicationStatus int

const (
	form applicationStatus = iota
	results
)

func New() *model {
	m := new(model)
	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if m.active == form {
				m.active = results
			}
		case "esc":
			if m.active == results {
				m.active = form
			}
		case "ctrl+c", "q":
			return m, tea.Quit
        default:
            log.Print(msg.String())
		}
	}
	return m, cmd
}

func (m model) View() string {
	var active string
	switch m.active {
	case form:
		active = "form mode"
	case results:
		active = "results mode"
	default:
		panic("invalid model state")
	}
	return lipgloss.Place(
		m.width,
		m.height,
		0.5,
		0.5,
		active,
	)
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
