package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
	doc := ""

	tabs := []string{"Dashboard", "Uploads", "Peers"}
	var renderedTabs []string
	for i, t := range tabs {
		if i == m.activeTab {
			renderedTabs = append(renderedTabs, ActiveTabStyle.Render(t))
		} else {
			renderedTabs = append(renderedTabs, TabStyle.Render(t))
		}
	}
	doc += lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...) + "\n\n"

	switch m.activeTab {
	case 0:
		doc += DashboardView(12, 1050000000, 500, 1200)
	case 1:
		doc += UploadView("research_paper.pdf", 0.65)
	case 2:
		doc += "Peer List View (Coming Soon)"
	}

	return doc + "\n\n (Q to quit)"
}

func NewApp() *tea.Program {
	return tea.NewProgram(model{})
}
