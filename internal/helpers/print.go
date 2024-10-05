package helpers

import (
	"fmt"

	"github.com/arthurbcp/kuma-cli/internal/debug"
	"github.com/arthurbcp/kuma-cli/pkg/style"
	"github.com/charmbracelet/lipgloss"
	"github.com/gookit/color"
)

var (
	CheckStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color(style.SuccessColor)).Bold(true).Padding(0, 1, 0)
	CrossMarkStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(style.ErrorColor)).Bold(true).Padding(0, 1, 0)
)

func (h *Helpers) TitlePrint(text string) {
	fmt.Println(style.TitleStyle.Render(text) + "\n")
}

func (h *Helpers) LogPrint(text string) {
	fmt.Println(style.LogStyle.Render(text) + "\n")
}

func (h *Helpers) CheckMarkPrint(text string) {
	fmt.Println(CheckStyle.Render("✔") + text)
}

func (h *Helpers) CrossMarkPrint(text string) {
	fmt.Println(CrossMarkStyle.Render("✖") + text)
}

func (h *Helpers) ErrorPrint(text string) {
	fmt.Println(style.ErrorStyle.Render(text) + "\n")
}

func (h *Helpers) DebugPrint(header, text string) {
	fmt.Println()
	if debug.Debug {
		color.New(color.FgBlack, color.BgYellow).Println(" - " + header + " - ")
		color.Gray.Println(text)
	}
}
