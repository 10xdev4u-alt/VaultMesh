package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	activeTab int
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "tab":
			m.activeTab = (m.activeTab + 1) % 3
		}
	}
	return m, nil
}

func (m model) View() string {
	return "VaultMesh TUI (Press Tab to switch, Q to quit)"
}

func NewApp() *tea.Program {
	return tea.NewProgram(model{})
}
