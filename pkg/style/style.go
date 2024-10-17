package style

import "github.com/charmbracelet/lipgloss"

var (
	BgColor          = lipgloss.Color("#EDAFB8")
	PrimaryTextColor = lipgloss.Color("#EDAFB8")
	ErrorColor       = lipgloss.Color("#D50000")
	WhiteColor       = lipgloss.Color("#FFFFFF")
	SuccessColor     = lipgloss.Color("#008000")
)

var (
	TitleStyle        = lipgloss.NewStyle().Foreground(PrimaryTextColor).Bold(true).Padding(0, 1, 0)
	ErrorStyle        = lipgloss.NewStyle().Background(ErrorColor).Foreground(WhiteColor).Bold(true).Padding(0, 1, 0)
	LogStyle          = lipgloss.NewStyle().Foreground(PrimaryTextColor)
	FocusedStyle      = lipgloss.NewStyle().Foreground(BgColor).Bold(true)
	SelectedItemStyle = lipgloss.NewStyle().PaddingLeft(1).Foreground(PrimaryTextColor).Bold(true)
	CheckStyle        = lipgloss.NewStyle().Foreground(SuccessColor).Bold(true).Padding(0, 1, 0)
	CrossMarkStyle    = lipgloss.NewStyle().Foreground(ErrorColor).Bold(true).Padding(0, 1, 0)
	TagsStyle         = lipgloss.NewStyle().Foreground(PrimaryTextColor)
	DescriptionStyle  = lipgloss.NewStyle().Foreground(WhiteColor)
)
