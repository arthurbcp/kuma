package execHandlers

import (
	"os"

	"github.com/arthurbcp/kuma-cli/cmd/shared"
	"github.com/arthurbcp/kuma-cli/internal/helpers"
	"github.com/arthurbcp/kuma-cli/pkg/filesystem"
	"github.com/arthurbcp/kuma-cli/pkg/style"
	"github.com/spf13/afero"
)

func HandleRun(name string, vars map[string]interface{}) {
	fs := filesystem.NewFileSystem(afero.NewOsFs())
	data, err := helpers.UnmarshalFile(shared.KumaRunsPath, fs)
	if err != nil {
		style.ErrorPrint("parsing file error: " + err.Error())
		os.Exit(1)
	}

	run, ok := data[name]
	if !ok {
		style.ErrorPrint("run not found: " + name)
		os.Exit(1)
	}

	for _, step := range run.([]interface{}) {
		step := step.(map[string]interface{})
		for key, value := range step {
			if key == "cmd" {
				HandleCommand(value.(string), vars)
			} else if key == "input" {
				HandleInput(value.(map[string]interface{}), vars)
			} else if key == "log" {
				HandleLog(value.(string), vars)
			} else if key == "run" {
				HandleRun(value.(string), vars)
			} else if key == "create" {
				HandleCreate(value.(map[string]interface{}), vars)
			} else if key == "load" {
				HandleLoad(value.(map[string]interface{}), vars)
			}
		}
	}
}
