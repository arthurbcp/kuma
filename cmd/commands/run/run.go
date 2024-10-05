// run.go
//
// Package run defines the 'run' subcommand for the Kuma CLI.
// It handles generating project scaffolds based on Go templates.
package run

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/arthurbcp/kuma-cli/cmd/shared"
	"github.com/arthurbcp/kuma-cli/cmd/steps"
	"github.com/arthurbcp/kuma-cli/cmd/ui/multiSelectInput"
	"github.com/arthurbcp/kuma-cli/cmd/ui/selectInput"
	"github.com/arthurbcp/kuma-cli/cmd/ui/textInput"

	"github.com/arthurbcp/kuma-cli/internal/helpers"
	"github.com/arthurbcp/kuma-cli/pkg/filesystem"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var (
	Run       string
	Variables = make(map[string]interface{})
)

// RunCmd represents the 'run' subcommand.
var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a specific pipeline",
	Run: func(cmd *cobra.Command, args []string) {
		ExecRun(Run)
	},
}

func ExecRun(name string) {
	helpers := helpers.NewHelpers()
	fs := filesystem.NewFileSystem(afero.NewOsFs())
	data, err := helpers.UnmarshalFile(shared.KumaRunsPath, fs)
	if err != nil {
		helpers.ErrorPrint("parsing file error: " + err.Error())
		os.Exit(1)
	}
	run, ok := data[name]
	if !ok {
		helpers.ErrorPrint("run not found: " + name)
		os.Exit(1)
	}
	for _, step := range run.([]interface{}) {
		step := step.(map[string]interface{})
		for key, value := range step {
			if key == "cmd" {
				handleCommand(value.(string))
			} else if key == "input" {
				handleInput(value.(map[string]interface{}))
			} else if key == "log" {
				handleLog(value.(string))
			}
		}
	}
}

func handleInput(input map[string]interface{}) {
	helpers := helpers.NewHelpers()
	label, ok := input["label"].(string)
	if !ok {
		helpers.ErrorPrint("label is required for input")
		os.Exit(1)
	}
	out, ok := input["out"].(string)
	if !ok {
		helpers.ErrorPrint("out is required for input")
		os.Exit(1)
	}
	other := false
	if o, ok := input["other"]; ok {
		other = o.(bool)
	}

	skippable := false
	if s, ok := input["skippable"]; ok {
		skippable = s.(bool)
	}

	multi := false
	if m, ok := input["multi"]; ok {
		multi = m.(bool)
	}

	if mapOptions, ok := input["options"].([]interface{}); ok {
		options := make([]steps.Item, len(mapOptions))
		for i, option := range mapOptions {
			options[i] = steps.Item{
				Label: option.(map[string]interface{})["label"].(string),
				Value: option.(map[string]interface{})["value"].(string),
			}
		}
		if multi {
			choices := make(map[string]bool)
			for _, option := range options {
				choices[option.Value] = false
			}
			output := &multiSelectInput.Selection{
				Choices: choices,
			}
			p := tea.NewProgram(multiSelectInput.InitialMultiSelectInputModel(options, output, label, skippable, false))
			_, err := p.Run()
			if err != nil {
				helpers.ErrorPrint("error running program: " + err.Error())
				os.Exit(1)
			}
			selectedChoices := make([]string, 0)
			for key, value := range output.Choices {
				if value {
					selectedChoices = append(selectedChoices, key)
				}
			}
			Variables[out] = selectedChoices
		} else {
			output := &selectInput.Selection{}
			p := tea.NewProgram(selectInput.InitialSelectInputModel(options, output, label, other, skippable, false))
			_, err := p.Run()
			if err != nil {
				helpers.ErrorPrint("error running program: " + err.Error())
				os.Exit(1)
			}
			Variables[out] = output.Choice
		}
	} else {
		output := &textInput.Output{}
		p := tea.NewProgram(textInput.InitialTextInputModel(output, label, false))
		_, err := p.Run()
		if err != nil {
			helpers.ErrorPrint("error running program: " + err.Error())
			os.Exit(1)
		}
		Variables[out] = output.Output
	}
}

func handleLog(log string) {
	var err error
	helpers := helpers.NewHelpers()
	log, err = helpers.ReplaceVars(log, Variables, helpers.GetFuncMap())
	if err != nil {
		helpers.ErrorPrint("parsing log error: " + err.Error())
		os.Exit(1)
	}
	helpers.LogPrint(log)
}

func handleCommand(cmdStr string) {
	var err error
	helpers := helpers.NewHelpers()
	cmdStr, err = helpers.ReplaceVars(cmdStr, Variables, helpers.GetFuncMap())
	if err != nil {
		helpers.ErrorPrint("parsing command error: " + err.Error())
		os.Exit(1)
	}
	helpers.LogPrint(fmt.Sprintf("running: %s", cmdStr))
	cmdArgs := strings.Split(cmdStr, " ")
	execCmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	// Set the command's standard output to the console
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr

	// Execute the command
	err = execCmd.Run()
	if err != nil {
		helpers.ErrorPrint("command error: " + err.Error())
		os.Exit(1)
	}
}

// init sets up flags for the 'run' subcommand and binds them to variables.
func init() {
	// Repository name
	RunCmd.Flags().StringVarP(&Run, "name", "n", "", "run to use")
	RunCmd.MarkFlagRequired("name")
}
