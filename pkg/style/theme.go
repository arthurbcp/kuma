package style

import (
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

var (
	NormalBg  = lipgloss.AdaptiveColor{Light: "235", Dark: "252"}
	Primary   = lipgloss.Color("#EDAFB8")
	Secondary = lipgloss.Color("#89CFF0")
	Cream     = lipgloss.AdaptiveColor{Light: "#FFFDF5", Dark: "#FFFDF5"}
	Third     = lipgloss.Color("#F780E2")
	Error     = lipgloss.AdaptiveColor{Light: "#FF4672", Dark: "#ED567A"}
	Success   = lipgloss.AdaptiveColor{Light: "#02BA84", Dark: "#02BF87"}
)

var (
	TitleWithBgStyle  = lipgloss.NewStyle().Background(Primary).Foreground(Cream).Bold(true).Padding(0, 1, 0)
	TitleStyle        = lipgloss.NewStyle().Foreground(Primary).Bold(true)
	ErrorStyle        = lipgloss.NewStyle().Background(Error).Foreground(Cream).Bold(true).Padding(0, 1, 0)
	LogStyle          = lipgloss.NewStyle().Foreground(Cream)
	FocusedStyle      = lipgloss.NewStyle().Foreground(Primary).Bold(true)
	SelectedItemStyle = lipgloss.NewStyle().Foreground(Third).Bold(true)
	CheckStyle        = lipgloss.NewStyle().Foreground(Success).Bold(true).Padding(0, 1, 0)
	CrossMarkStyle    = lipgloss.NewStyle().Foreground(Error).Bold(true).Padding(0, 1, 0)
	TagsStyle         = lipgloss.NewStyle().Foreground(Secondary)
	DescriptionStyle  = lipgloss.NewStyle().Foreground(Cream)
)

func KumaTheme() *huh.Theme {
	t := huh.ThemeBase()

	t.Focused.Base = t.Focused.Base.BorderForeground(lipgloss.Color("238"))
	t.Focused.Title = t.Focused.Title.Foreground(Primary).Bold(true)
	t.Focused.NoteTitle = t.Focused.NoteTitle.Foreground(Primary).Bold(true).MarginBottom(1)
	t.Focused.Directory = t.Focused.Directory.Foreground(Primary)
	t.Focused.Description = t.Focused.Description.Foreground(lipgloss.AdaptiveColor{Light: "", Dark: "243"})
	t.Focused.ErrorIndicator = t.Focused.ErrorIndicator.Foreground(Error)
	t.Focused.ErrorMessage = t.Focused.ErrorMessage.Foreground(Error)
	t.Focused.SelectSelector = t.Focused.SelectSelector.Foreground(Third)
	t.Focused.NextIndicator = t.Focused.NextIndicator.Foreground(Third)
	t.Focused.PrevIndicator = t.Focused.PrevIndicator.Foreground(Third)
	t.Focused.Option = t.Focused.Option.Foreground(NormalBg)
	t.Focused.MultiSelectSelector = t.Focused.MultiSelectSelector.Foreground(Third)
	t.Focused.SelectedOption = t.Focused.SelectedOption.Foreground(Secondary)
	t.Focused.SelectedPrefix = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#02CF92", Dark: "#02A877"}).SetString("✓ ")
	t.Focused.UnselectedPrefix = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "", Dark: "243"}).SetString("• ")
	t.Focused.UnselectedOption = t.Focused.UnselectedOption.Foreground(NormalBg)
	t.Focused.FocusedButton = t.Focused.FocusedButton.Foreground(Cream).Background(Third)
	t.Focused.Next = t.Focused.FocusedButton
	t.Focused.BlurredButton = t.Focused.BlurredButton.Foreground(NormalBg).Background(lipgloss.AdaptiveColor{Light: "252", Dark: "237"})

	t.Focused.TextInput.Cursor = t.Focused.TextInput.Cursor.Foreground(Secondary)
	t.Focused.TextInput.Placeholder = t.Focused.TextInput.Placeholder.Foreground(lipgloss.AdaptiveColor{Light: "248", Dark: "238"})
	t.Focused.TextInput.Prompt = t.Focused.TextInput.Prompt.Foreground(Third)

	t.Blurred = t.Focused
	t.Blurred.Base = t.Focused.Base.BorderStyle(lipgloss.HiddenBorder())
	t.Blurred.NextIndicator = lipgloss.NewStyle()
	t.Blurred.PrevIndicator = lipgloss.NewStyle()

	return t
}
