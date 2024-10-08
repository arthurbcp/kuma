package style

import "github.com/charmbracelet/lipgloss"

var (
	TextTileColor = lipgloss.Color("#2C040A")
	BgColor       = lipgloss.Color("#EDAFB8")
	FgColor       = lipgloss.Color("#EDAFB8")
	ErrorColor    = lipgloss.Color("#D50000")
	WhiteColor    = lipgloss.Color("#FFFFFF")
	SuccessColor  = lipgloss.Color("#008000")
	TagsColor     = lipgloss.Color("#11011a")
)

var (
	TitleStyle        = lipgloss.NewStyle().Background(BgColor).Foreground(TextTileColor).Bold(true).Padding(0, 1, 0)
	ErrorStyle        = lipgloss.NewStyle().Background(ErrorColor).Foreground(WhiteColor).Bold(true).Padding(0, 1, 0)
	LogStyle          = lipgloss.NewStyle().Foreground(FgColor)
	FocusedStyle      = lipgloss.NewStyle().Foreground(BgColor).Bold(true)
	SelectedItemStyle = lipgloss.NewStyle().PaddingLeft(1).Foreground(FgColor).Bold(true)
	CheckStyle        = lipgloss.NewStyle().Foreground(SuccessColor).Bold(true).Padding(0, 1, 0)
	CrossMarkStyle    = lipgloss.NewStyle().Foreground(ErrorColor).Bold(true).Padding(0, 1, 0)
	TagsStyle         = lipgloss.NewStyle().Foreground(TagsColor)
	DescriptionStyle  = lipgloss.NewStyle().Foreground(WhiteColor)
)
