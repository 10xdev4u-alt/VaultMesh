package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// FileBrowserView renders the list of available files.
func FileBrowserView(files []string) string {
	content := "File Browser (Search: /):\n\n"

	for _, f := range files {
		content += fmt.Sprintf("📄 %s\n", f)
	}

	return lipgloss.NewStyle().
		Padding(1, 2).
		Render(content)
}
