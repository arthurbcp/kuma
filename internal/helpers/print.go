package helpers

import (
	"fmt"

	"github.com/arthurbcp/kuma-cli/internal/debug"
	"github.com/gookit/color"
)

func (h *Helpers) HeaderPrint(text string) {
	fmt.Println()
	color.New(color.FgWhite, color.BgMagenta).Println(" - " + text + " - ")
}

func (h *Helpers) CheckMarkPrint(text string) {
	color.Gray.Println("  ✅ " + text)
}

func (h *Helpers) CrossMarkPrint(text string) {
	color.Gray.Println("  ❌ " + text)
}

func (h *Helpers) ErrorPrint(text string) {
	color.Red.Println(text)
}

func (h *Helpers) DebugPrint(header, text string) {
	fmt.Println()
	if debug.Debug {
		color.New(color.FgBlack, color.BgYellow).Println(" - " + header + " - ")
		color.Gray.Println(text)
	}
}
