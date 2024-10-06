// run.go
//
// Package run defines the 'run' subcommand for the Kuma CLI.
// It handles generating project scaffolds based on Go templates.
package exec

import (
	"os"

	execHandlers "github.com/arthurbcp/kuma-cli/cmd/commands/exec/handlers"
	"github.com/arthurbcp/kuma-cli/cmd/shared"

	"github.com/arthurbcp/kuma-cli/internal/helpers"
	"github.com/arthurbcp/kuma-cli/pkg/filesystem"
	"github.com/arthurbcp/kuma-cli/pkg/style"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var (
	Run string
)

// ExecRunCmd represents the 'run' subcommand.
var ExecRunCmd = &cobra.Command{
	Use:   "exec",
	Short: "Run a specific pipeline",
	Run: func(cmd *cobra.Command, args []string) {
		vars := map[string]interface{}{
			"data": map[string]interface{}{},
		}
		ExecRun(Run, vars)
	},
}

func ExecRun(name string, vars map[string]interface{}) {
	helpers := helpers.NewHelpers()
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
				execHandlers.HandleCommand(value.(string), vars)
			} else if key == "input" {
				execHandlers.HandleInput(value.(map[string]interface{}), vars)
			} else if key == "log" {
				execHandlers.HandleLog(value.(string), vars)
			} else if key == "run" {
				ExecRun(value.(string), vars)
			} else if key == "create" {
				execHandlers.HandleCreate(value.(map[string]interface{}), vars)
			} else if key == "load" {
				execHandlers.HandleLoad(value.(map[string]interface{}), vars)
			}
		}
	}
}

// init sets up flags for the 'run' subcommand and binds them to variables.
func init() {
	// Repository name
	ExecRunCmd.Flags().StringVarP(&Run, "run", "r", "", "run to use")
	ExecRunCmd.MarkFlagRequired("run")
}
