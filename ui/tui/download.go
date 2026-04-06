package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// DownloadView renders the active downloads status.
func DownloadView(cid string, speed float64, eta string) string {
	speedStr := lipgloss.NewStyle().Foreground(AccentColor).Render(fmt.Sprintf("%.2f MB/s", speed))
	etaStr := lipgloss.NewStyle().Foreground(AccentColor).Render(eta)

	content := fmt.Sprintf(
		"Downloading: %s\nSpeed:       %s\nETA:         %s",
		cid, speedStr, etaStr,
	)

	return lipgloss.NewStyle().
		Padding(1, 2).
		Render(content)
}
