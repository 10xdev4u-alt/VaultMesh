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
			m.activeTab = (m.activeTab + 1) % 5
		}
	}
	return m, nil
}

func (m model) View() string {
	doc := ""

	tabs := []string{"Dashboard", "Uploads", "Downloads", "Peers", "Network"}
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
		doc += DownloadView("bafybeigdyrzt5sfp7udm7hu76uh7m", 12.5, "2m 15s")
	case 3:
		doc += PeerListView([]string{"QmNnooDu7bfjPFoTBsPWCcqS2S2s7aPvwVfN2p7rQdEaJs", "QmQCU2Ecws3N79txbcocFQ977XLeqM6K1Y78T9fG6t4q8G", "QmbLHAnMo96F8tA6yHArD9Nn7yS85tshx5G7nQfG7xN9qD"})
	case 4:
		doc += NetworkView()
	}

	return doc + "\n\n (Q to quit)"
}

func NewApp() *tea.Program {
	return tea.NewProgram(model{})
}
