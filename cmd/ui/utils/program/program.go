package program

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type Program struct {
	Exit bool
}

func NewProgram() *Program {
	return &Program{
		Exit: false,
	}
}

func (p *Program) ExitCLI(tprogram *tea.Program) {
	if p.Exit {
		if err := tprogram.ReleaseTerminal(); err != nil {
			log.Fatal(err)
		}
		os.Exit(1)
	}
}
