// run.go
//
// Package run defines the 'run' subcommand for the Kuma CLI.
// It handles generating project scaffolds based on Go templates.
package exec

import (
	"os"

	execHandlers "github.com/arthurbcp/kuma/cmd/commands/exec/handlers"
	"github.com/arthurbcp/kuma/cmd/shared"
	"github.com/arthurbcp/kuma/cmd/ui/selectInput"
	"github.com/arthurbcp/kuma/cmd/ui/utils/program"
	"github.com/arthurbcp/kuma/cmd/ui/utils/steps"
	"github.com/arthurbcp/kuma/internal/services"
	"github.com/arthurbcp/kuma/pkg/filesystem"
	"github.com/arthurbcp/kuma/pkg/style"
	tea "github.com/charmbracelet/bubbletea"

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
		Execute()
	},
}

func Execute() {
	if Run == "" {
		Run = handleTea()
	}
	vars := map[string]interface{}{
		"data": map[string]interface{}{},
	}
	execHandlers.HandleRun(Run, vars)
}

func handleTea() string {

	program := program.NewProgram()
	runService := services.NewRunService(shared.KumaRunsPath, filesystem.NewFileSystem(afero.NewOsFs()))
	runs, err := runService.GetAll()
	if err != nil {
		style.ErrorPrint("getting runs error: " + err.Error())
		os.Exit(1)
	}
	var options = make([]steps.Item, 0)
	for key, run := range runs {
		options = append(options, steps.NewItem(
			key,
			key,
			run.Description,
			[]string{},
		))
	}

	output := &selectInput.Selection{}
	p := tea.NewProgram(selectInput.InitialSelectInputModel(options, output, "Select a run", false, program))
	_, err = p.Run()

	program.ExitCLI(p)

	if err != nil {
		style.ErrorPrint("error running program: " + err.Error())
		os.Exit(1)
	}
	return output.Choice
}

// init sets up flags for the 'run' subcommand and binds them to variables.
func init() {
	// Repository name
	ExecRunCmd.Flags().StringVarP(&Run, "run", "r", "", "run to use")
}
