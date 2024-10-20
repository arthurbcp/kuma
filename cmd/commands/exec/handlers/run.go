package execHandlers

import (
	"fmt"
	"log"
	"os"

	execFormHandlers "github.com/arthurbcp/kuma/v2/cmd/commands/exec/handlers/form"
	"github.com/arthurbcp/kuma/v2/cmd/constants"
	"github.com/arthurbcp/kuma/v2/cmd/shared"
	"github.com/arthurbcp/kuma/v2/internal/domain"
	"github.com/arthurbcp/kuma/v2/internal/services"
	"github.com/arthurbcp/kuma/v2/pkg/filesystem"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/afero"
)

func HandleRun(name, moduleName string, vars map[string]interface{}) error {
	var err error
	var run = &domain.Run{}
	fs := filesystem.NewFileSystem(afero.NewOsFs())
	if moduleName != "" {
		moduleService := services.NewModuleService(shared.KumaFilesPath, fs)
		modules, err := moduleService.GetAll()
		if err != nil {
			return err
		}
		module := modules[moduleName]
		run, err = moduleService.GetRun(&module, name, shared.KumaFilesPath+"/"+moduleName+"/"+shared.KumaRunsPath)

		if err != nil {
			return err
		}
	} else {
		runService := services.NewRunService(shared.KumaRunsPath, fs)
		run, err = runService.Get(name)
		if err != nil {
			return err
		}
	}

	for _, step := range run.Steps {
		step := step.(map[string]interface{})
		for key, value := range step {
			switch key {
			case constants.CmdHandler:
				err := HandleCommand(value.(string), vars)
				if err != nil {
					return fmt.Errorf("[handler: %s] - %s", constants.CmdHandler, err.Error())
				}
			case constants.LogHandler:
				err := HandleLog(value.(string), vars)
				if err != nil {
					return fmt.Errorf("[handler: %s] - %s", constants.LogHandler, err.Error())
				}
			case constants.RunHandler:
				err := HandleRun(value.(string), moduleName, vars)
				if err != nil {
					return fmt.Errorf("[handler: %s] - %s", constants.RunHandler, err.Error())
				}
			case constants.CreateHandler:
				err := HandleCreate(moduleName, value.(map[string]interface{}), vars)
				if err != nil {
					return fmt.Errorf("[handler: %s] - %s", constants.CreateHandler, err.Error())
				}
			case constants.LoadHandler:
				err := HandleLoad(value.(map[string]interface{}), vars)
				if err != nil {
					return fmt.Errorf("[handler: %s] - %s", constants.LoadHandler, err.Error())
				}
			case constants.WhenHandler:
				err := HandleWhen(moduleName, value.(map[string]interface{}), vars)
				if err != nil {
					return fmt.Errorf("[handler: %s] - %s", constants.WhenHandler, err.Error())
				}
			case constants.ModifyHandler:
				err := HandleModify(moduleName, value.(map[string]interface{}), vars)
				if err != nil {
					return fmt.Errorf("[handler: %s] - %s", constants.ModifyHandler, err.Error())
				}
			case constants.FormHandler:
				err := execFormHandlers.HandleForm(value.(map[string]interface{}), vars)
				if err != nil {
					return fmt.Errorf("[handler: %s] - %s", constants.FormHandler, err.Error())
				}
			case constants.DefineHandler:
				err := HandleDefine(value.(map[string]interface{}), vars)
				if err != nil {
					return fmt.Errorf("[handler: %s] - %s", constants.DefineHandler, err.Error())
				}
			default:
				return fmt.Errorf("invalid handler type: %s", key)
			}
		}
	}
	return nil
}

func ExitCLI(tprogram *tea.Program) {
	if err := tprogram.ReleaseTerminal(); err != nil {
		log.Fatal(err)
	}
	os.Exit(1)
}
