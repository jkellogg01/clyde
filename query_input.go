package main

import (
	"errors"
	"log"
	"net/url"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type QueryInput struct {
	value   string
	infield textinput.Model
	err     error
}

func NewQueryInput() *QueryInput {
	m := new(QueryInput)
	m.infield = textinput.New()
	m.infield.Prompt = ""
	m.infield.Placeholder = "https://www.google.com/search?q=something"
	m.infield.Focus()
	return m
}

var (
	ErrIncompleteQuery     = errors.New("this query is incomplete!")
	ErrUnsupportedProtocol = errors.New("this protocol is not supported!")
)

func (m QueryInput) validate() error {
	log.Printf("parsing input: %s", m.value)
	parsed, err := url.Parse(m.value)
	if err != nil {
		return err
	}
    log.Printf("found scheme '%s' and hostname '%s'", parsed.Scheme, parsed.Hostname())
	if parsed.Scheme == "" || parsed.Hostname() == "" {
		return ErrIncompleteQuery
	}
	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return ErrUnsupportedProtocol
	}
    log.Printf("the url '%s' is valid!", parsed.String())
	return nil
}

func (m QueryInput) Init() tea.Cmd {
	return nil
}

type validateMsg string

func (m QueryInput) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.infield.Width = msg.Width - 3
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			// right now this is just so that the list of tokens gets logged to the console
			parsed, err := url.Parse(m.value)
			if err != nil {
				log.Printf("ERROR: %s", err)
			}
			log.Printf("Parsed URL: %s", parsed.String())
		default:
			m.infield, _ = m.infield.Update(msg)
			return m, tea.Tick(time.Second, func(t time.Time) tea.Msg {
				return validateMsg(m.value)
			})
		}
	case validateMsg:
		if string(msg) == m.value {
			m.err = m.validate()
		}
	}
	m.value = m.infield.Value()
	return m, cmd
}

func (m QueryInput) View() string {
	var inputErr string
	if m.err != nil {
		inputErr = m.err.Error()
	}
	return lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Render(lipgloss.JoinHorizontal(
			lipgloss.Left,
			m.infield.View(),
		)),
		lipgloss.NewStyle().Foreground(lipgloss.ANSIColor(1)).Render(inputErr),
	)
}
