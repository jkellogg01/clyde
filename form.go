package main

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	methodpicker "github.com/jkellogg01/clyde/tea-components/method-picker"
)

type Form struct {
	Method methodpicker.MethodPicker
	URL    textinput.Model
	Body   textarea.Model
}

func (m Form) Init() tea.Cmd {
	return nil
}

func (m Form) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Form) View() string {
	return "this is the form"
}
