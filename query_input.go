package main

import (
	"errors"

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
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			segments...,
		),
		m.err,
	)
}
