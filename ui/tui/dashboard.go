package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// DashboardView renders the main status dashboard.
func DashboardView(peers int, storage int64, bwIn, bwOut int64) string {
	peersStr := lipgloss.NewStyle().Foreground(AccentColor).Render(fmt.Sprintf("%d", peers))
	storageStr := lipgloss.NewStyle().Foreground(AccentColor).Render(fmt.Sprintf("%.2f GB", float64(storage)/1e9))

	stats := fmt.Sprintf(
		"Connected Peers: %s\nStorage Used:    %s\nBandwidth:       IN: %d bytes | OUT: %d bytes",
		peersStr, storageStr, bwIn, bwOut,
	)

	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(PrimaryColor).
		Padding(1, 2).
		Render(stats)
}
