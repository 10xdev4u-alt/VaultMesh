package tui

import (
	"github.com/charmbracelet/lipgloss"
)

// SettingsView renders the configuration panel.
func SettingsView() string {
	content := "Node Settings:\n\n"
	content += "Port:           8080\n"
	content += "Data Shards:    3\n"
	content += "Parity Shards:  2\n"
	content += "Auto-Heal:      Enabled\n"

	return lipgloss.NewStyle().
		Padding(1, 2).
		Render(content)
}
