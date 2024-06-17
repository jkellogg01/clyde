package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type QueryInput struct {
	value   string
	infield textinput.Model
}

func NewQueryInput() *QueryInput {
	m := new(QueryInput)
    m.infield = textinput.New()
    m.infield.Prompt = ""
    m.infield.Placeholder = "query..."
    m.infield.Focus()
	return m
}

func (m QueryInput) Init() tea.Cmd {
	return nil
}

func (m QueryInput) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
    m.infield, cmd = m.infield.Update(msg)
	return m, cmd
}

func (m QueryInput) View() string {
    return lipgloss.JoinHorizontal(
        0,
        m.infield.View(),
    )
}
