package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// NetworkView renders an ASCII representation of the network topology.
func NetworkView() string {
	node := lipgloss.NewStyle().
		Border(lipgloss.CircleBorder()).
		Padding(0, 1).
		Render("YOU")

	peer := lipgloss.NewStyle().
		Border(lipgloss.CircleBorder()).
		Foreground(AccentColor).
		Render("P")

	content := fmt.Sprintf(
		"      %s\n       | \n %s --- %s --- %s\n       | \n      %s",
		peer, peer, node, peer, peer,
	)

	return lipgloss.NewStyle().
		Padding(1, 2).
		Render("Network Topology:\n\n" + content)
}
