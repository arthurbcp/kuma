package style

import "github.com/charmbracelet/lipgloss"

var (
	TextTileColor = lipgloss.Color("#2C040A")
	BgColor       = lipgloss.Color("#EDAFB8")
	FgColor       = lipgloss.Color("#EDAFB8")
	ErrorColor    = lipgloss.Color("#D50000")
	SuccessColor  = lipgloss.Color("#008000")
)

var (
	TitleStyle        = lipgloss.NewStyle().Background(lipgloss.Color(BgColor)).Foreground(TextTileColor).Bold(true).Padding(0, 1, 0)
	ErrorStyle        = lipgloss.NewStyle().Background(lipgloss.Color(ErrorColor)).Foreground(lipgloss.Color("#FFF")).Bold(true).Padding(0, 1, 0)
	LogStyle          = lipgloss.NewStyle().Foreground(lipgloss.Color(FgColor))
	FocusedStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color(BgColor)).Bold(true)
	SelectedItemStyle = lipgloss.NewStyle().PaddingLeft(1).Foreground(FgColor).Bold(true)
)
