package style

import "github.com/charmbracelet/lipgloss"

var (
	BgColor = lipgloss.Color("#EDAFB8")
	FgColor = lipgloss.Color("#EDAFB8")
)

var (
	TitleStyle        = lipgloss.NewStyle().Background(lipgloss.Color(BgColor)).Foreground(lipgloss.Color("#2C040A")).Bold(true).Padding(0, 1, 0)
	ErrorStyle        = lipgloss.NewStyle().Padding(0, 1, 0)
	LogStyle          = lipgloss.NewStyle().Foreground(lipgloss.Color(FgColor))
	FocusedStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color(BgColor)).Bold(true)
	SelectedItemStyle = lipgloss.NewStyle().PaddingLeft(1).Foreground(FgColor).Bold(true)
)
