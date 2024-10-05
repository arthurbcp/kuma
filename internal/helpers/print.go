package helpers

import (
	"fmt"

	"github.com/arthurbcp/kuma-cli/internal/debug"
	"github.com/arthurbcp/kuma-cli/pkg/style"
	"github.com/gookit/color"
)

func (h *Helpers) TitlePrint(text string) {
	fmt.Println(style.TitleStyle.Render(text))
}

func (h *Helpers) LogPrint(text string) {
	fmt.Println(style.LogStyle.Render(text))
}

func (h *Helpers) CheckMarkPrint(text string) {
	color.Gray.Println("  ✅ " + text)
}

func (h *Helpers) CrossMarkPrint(text string) {
	color.Gray.Println("  ❌ " + text)
}

func (h *Helpers) ErrorPrint(text string) {
	fmt.Println(style.ErrorStyle.Render(text))
}

func (h *Helpers) DebugPrint(header, text string) {
	fmt.Println()
	if debug.Debug {
		color.New(color.FgBlack, color.BgYellow).Println(" - " + header + " - ")
		color.Gray.Println(text)
	}
}
