package execRun

import (
	"os"

	execHandlers "github.com/arthurbcp/kuma/v2/cmd/commands/exec/handlers"
	"github.com/arthurbcp/kuma/v2/cmd/shared"
	"github.com/arthurbcp/kuma/v2/cmd/ui/selectInput"
	"github.com/arthurbcp/kuma/v2/cmd/ui/utils/program"
	"github.com/arthurbcp/kuma/v2/cmd/ui/utils/steps"
	"github.com/arthurbcp/kuma/v2/internal/services"
	"github.com/arthurbcp/kuma/v2/pkg/filesystem"
	"github.com/arthurbcp/kuma/v2/pkg/style"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var ExecCmd = &cobra.Command{
	Use:   "run",
	Short: "Execute a specific run without a module",
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

func Execute() {
	if shared.Run == "" {
		shared.Run = handleTea()
	}
	vars := map[string]interface{}{
		"data": map[string]interface{}{},
	}
	err := execHandlers.HandleRun(shared.Run, "", vars)
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}
}

func handleTea() string {
	var err error
	program := program.NewProgram()

	fs := filesystem.NewFileSystem(afero.NewOsFs())
	runService := services.NewRunService(shared.KumaRunsPath, fs)
	runs, err := runService.GetAll(true)
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

func init() {
	ExecCmd.Flags().StringVarP(&shared.Run, "run", "r", "", "run to use")
}
