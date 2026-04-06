package tui

import "github.com/charmbracelet/lipgloss"

var (
	PrimaryColor = lipgloss.Color("#7C3AED") // Deep Purple
	AccentColor  = lipgloss.Color("#06B6D4") // Cyan

	TabStyle = lipgloss.NewStyle().
			Padding(0, 2).
			Foreground(lipgloss.Color("240"))

	ActiveTabStyle = TabStyle.Copy().
			Foreground(PrimaryColor).
			Bold(true).
			Border(lipgloss.NormalBorder(), false, false, true, false).
			BorderForeground(PrimaryColor)
)
