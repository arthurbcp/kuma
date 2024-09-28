package helpers

import (
	"fmt"
	"strings"

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
