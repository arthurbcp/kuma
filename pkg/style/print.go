package style

import (
	"fmt"

	"github.com/arthurbcp/kuma/internal/debug"
	"github.com/gookit/color"
)

func TitlePrint(text string) {
	fmt.Println(TitleStyle.Render(text) + "\n")
}

func LogPrint(text string) {
	fmt.Println(LogStyle.Render(text) + "\n")
}

func CheckMarkPrint(text string) {
	fmt.Println(CheckStyle.Render("✔") + text)
}

func CrossMarkPrint(text string) {
	fmt.Println(CrossMarkStyle.Render("✖") + text)
}

func ErrorPrint(text string) {
	fmt.Println(ErrorStyle.Render(text) + "\n")
}

func DebugPrint(header, text string) {
	fmt.Println()
	if debug.Debug {
		color.New(color.FgBlack, color.BgYellow).Println(" - " + header + " - ")
		color.Gray.Println(text)
	}
}
