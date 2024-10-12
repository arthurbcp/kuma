package execHandlers

import (
	"log"
	"os"

	"github.com/arthurbcp/kuma/cmd/shared"
	"github.com/arthurbcp/kuma/internal/services"
	"github.com/arthurbcp/kuma/pkg/filesystem"
	"github.com/arthurbcp/kuma/pkg/style"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/afero"
)

func HandleRun(name, module string, vars map[string]interface{}) {
	fs := filesystem.NewFileSystem(afero.NewOsFs())
	runService := services.NewRunService(shared.KumaFilesPath+"/"+module+"/"+shared.KumaRunsPath, fs)
	run, err := runService.Get(name)
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}

	for _, step := range run.Steps {
		step := step.(map[string]interface{})
		for key, value := range step {
			if key == "cmd" {
				HandleCommand(value.(string), vars)
			} else if key == "input" {
				HandleInput(value.(map[string]interface{}), vars)
			} else if key == "log" {
				HandleLog(value.(string), vars)
			} else if key == "run" {
				HandleRun(value.(string), module, vars)
			} else if key == "create" {
				HandleCreate(value.(map[string]interface{}), vars)
			} else if key == "load" {
				HandleLoad(value.(map[string]interface{}), vars)
			}
		}
	}
}

func ExitCLI(tprogram *tea.Program) {
	if err := tprogram.ReleaseTerminal(); err != nil {
		log.Fatal(err)
	}
	os.Exit(1)
}
