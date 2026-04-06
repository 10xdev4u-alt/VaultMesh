package tui

import "github.com/charmbracelet/lipgloss"

// Theme holds the color definitions for the UI.
type Theme struct {
	Primary lipgloss.TerminalColor
	Accent  lipgloss.TerminalColor
	Text    lipgloss.TerminalColor
}

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

	// DarkTheme is the default deep purple theme.
	DarkTheme = Theme{
		Primary: lipgloss.Color("#7C3AED"),
		Accent:  lipgloss.Color("#06B6D4"),
		Text:    lipgloss.Color("#FFFFFF"),
	}

	// LightTheme is an alternative high-contrast theme.
	LightTheme = Theme{
		Primary: lipgloss.Color("#4C1D95"),
		Accent:  lipgloss.Color("#0891B2"),
		Text:    lipgloss.Color("#000000"),
	}
)

// CurrentTheme can be toggled by the user.
var CurrentTheme = DarkTheme
