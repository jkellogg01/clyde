package methodpicker

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MethodPicker struct {
	list *methodList
}

func NewMethodPicker() *MethodPicker {
	return &MethodPicker{
		list: NewMethodList(),
	}
}

func (m MethodPicker) Init() tea.Cmd {
	return nil
}

func (m MethodPicker) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "space":
			m.list.Advance()
		}
	}
	return m, nil
}

func (m MethodPicker) View() string {
	return lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Render(m.list.View())
}
