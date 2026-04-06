package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/lipgloss"
)

// UploadView renders the active uploads and their progress.
func UploadView(filename string, percent float64) string {
	p := progress.New(progress.WithDefaultGradient())
	bar := p.ViewAs(percent)

	content := fmt.Sprintf(
		"Uploading: %s\n\n%s %.0f%%",
		filename, bar, percent*100,
	)

	return lipgloss.NewStyle().
		Padding(1, 2).
		Render(content)
}
