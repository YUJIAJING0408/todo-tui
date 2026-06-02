package app

import "charm.land/lipgloss/v2"

var (
	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("205")).
			MarginBottom(1)

	TodoTextStyle = lipgloss.NewStyle()

	DoneTextStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Strikethrough(true)

	CursorLineStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("236"))

	HelpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))

	MiddleStyle = lipgloss.NewStyle().Width(3).Align(lipgloss.Center)
)
