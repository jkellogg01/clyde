package main

import (
	"errors"
	"log"
	"net/url"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type QueryInput struct {
	value   string
	infield textinput.Model
	err     string
}

func NewQueryInput() *QueryInput {
	m := new(QueryInput)
	m.infield = textinput.New()
	m.infield.Prompt = ""
	m.infield.Placeholder = "query..."
	m.infield.Focus()
	return m
}

var (
	ErrIncompleteQuery     = errors.New("this query is incomplete!")
	ErrUnsupportedProtocol = errors.New("this protocol is not supported!")
)

func (m QueryInput) segments() ([]string, error) {
	result := make([]string, 0)
	result = append(result, m.infield.View())
	if len(m.value) < 5 {
		return result, ErrIncompleteQuery
	}
	return result, nil
}

func (m QueryInput) Init() tea.Cmd {
	return nil
}

func (m QueryInput) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.infield.Width = msg.Width - 3
	case tea.KeyMsg:
		if msg.String() == "enter" {
			// right now this is just so that the list of tokens gets logged to the console
            parsed, err := url.Parse(m.value)
            if err != nil {
                log.Printf("ERROR: %s", err)
            }
            log.Printf("Parsed URL: %s", parsed.String())
		}
	}
	m.infield, cmd = m.infield.Update(msg)
	m.value = m.infield.Value()
	return m, cmd
}

func (m QueryInput) View() string {
	segments, err := m.segments()
	if err != nil {
		m.err = err.Error()
	}
	return lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Render(lipgloss.JoinHorizontal(
			lipgloss.Left,
			segments...,
		)),
		lipgloss.NewStyle().Foreground(lipgloss.ANSIColor(1)).Render(m.err),
	)
}
