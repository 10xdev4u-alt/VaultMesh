package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// PeerListView renders the list of connected peers.
func PeerListView(peers []string) string {
	content := "Connected Peers:\n\n"

	green := lipgloss.NewStyle().Foreground(lipgloss.Color("42"))

	for _, p := range peers {
		indicator := green.Render("●")
		content += fmt.Sprintf("%s %s\n", indicator, p)
	}

	return lipgloss.NewStyle().
		Padding(1, 2).
		Render(content)
}
