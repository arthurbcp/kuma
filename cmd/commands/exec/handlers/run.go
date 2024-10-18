package execHandlers

import (
	"log"
	"os"

	execFormHandlers "github.com/arthurbcp/kuma/v2/cmd/commands/exec/handlers/form"
	"github.com/arthurbcp/kuma/v2/cmd/shared"
	"github.com/arthurbcp/kuma/v2/internal/domain"
	"github.com/arthurbcp/kuma/v2/internal/services"
	"github.com/arthurbcp/kuma/v2/pkg/filesystem"
	"github.com/arthurbcp/kuma/v2/pkg/style"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/afero"
)

func HandleRun(name, moduleName string, vars map[string]interface{}) {
	var err error
	var run = &domain.Run{}
	fs := filesystem.NewFileSystem(afero.NewOsFs())
	if moduleName != "" {
		moduleService := services.NewModuleService(shared.KumaFilesPath, fs)
		modules, err := moduleService.GetAll()
		if err != nil {
			style.ErrorPrint(err.Error())
			os.Exit(1)
		}
		module := modules[moduleName]
		run, err = moduleService.GetRun(&module, name, shared.KumaFilesPath+"/"+moduleName+"/"+shared.KumaRunsPath)

		if err != nil {
			style.ErrorPrint(err.Error())
			os.Exit(1)
		}
	} else {
		runService := services.NewRunService(shared.KumaRunsPath, fs)
		run, err = runService.Get(name)
		if err != nil {
			style.ErrorPrint(err.Error())
			os.Exit(1)
		}
	}

	for _, step := range run.Steps {
		step := step.(map[string]interface{})
		for key, value := range step {
			switch key {
			case "cmd":
				HandleCommand(value.(string), vars)
			case "log":
				HandleLog(value.(string), vars)
			case "run":
				HandleRun(value.(string), moduleName, vars)
			case "create":
				HandleCreate(moduleName, value.(map[string]interface{}), vars)
			case "load":
				HandleLoad(value.(map[string]interface{}), vars)
			case "when":
				HandleWhen(moduleName, value.(map[string]interface{}), vars)
			case "modify":
				HandleModify(value.(map[string]interface{}), vars)
			case "form":
				execFormHandlers.HandleForm(value.(map[string]interface{}), vars)
			default:
				style.ErrorPrint("invalid step type: " + key)
				os.Exit(1)
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
