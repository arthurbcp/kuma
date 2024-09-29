package helpers

import (
	"fmt"
	"strings"

	"github.com/arthurbcp/kuma-cli/internal/debug"
	"github.com/gookit/color"
)

func HeaderPrint(text string) {
	fmt.Println()
	color.New(color.FgWhite, color.BgMagenta).Println(" - " + strings.ToUpper(text) + " - ")
}

func CheckMarkPrint(text string) {
	color.Gray.Println("  ✅ " + text)
}

func CrossMarkPrint(text string) {
	color.Gray.Println("  ❌ " + text)
}

func ErrorPrint(text string) {
	color.Red.Println(text)
}

func DebugPrint(header, text string) {
	fmt.Println()
	if debug.Debug {
		color.New(color.FgBlack, color.BgYellow).Println(" - " + header + " - ")
		color.Gray.Println(text)
	}
}
