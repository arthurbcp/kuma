package execHandlers

import (
	"log"
	"os"

	execFormHandlers "github.com/arthurbcp/kuma/v2/cmd/commands/exec/handlers/form"
	"github.com/arthurbcp/kuma/v2/cmd/constants"
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
			case constants.CmdHandler:
				HandleCommand(value.(string), vars)
			case constants.LogHandler:
				HandleLog(value.(string), vars)
			case constants.RunHandler:
				HandleRun(value.(string), moduleName, vars)
			case constants.CreateHandler:
				HandleCreate(moduleName, value.(map[string]interface{}), vars)
			case constants.LoadHandler:
				HandleLoad(value.(map[string]interface{}), vars)
			case constants.WhenHandler:
				HandleWhen(moduleName, value.(map[string]interface{}), vars)
			case constants.ModifyHandler:
				HandleModify(value.(map[string]interface{}), vars)
			case constants.FormHandler:
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
