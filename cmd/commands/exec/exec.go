// run.go
//
// Package run defines the 'run' subcommand for the Kuma CLI.
// It handles generating project scaffolds based on Go templates.
package exec

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"strings"

	"github.com/arthurbcp/kuma-cli/cmd/shared"
	"github.com/arthurbcp/kuma-cli/cmd/steps"
	"github.com/arthurbcp/kuma-cli/cmd/ui/multiSelectInput"
	"github.com/arthurbcp/kuma-cli/cmd/ui/selectInput"
	"github.com/arthurbcp/kuma-cli/cmd/ui/textInput"

	"github.com/arthurbcp/kuma-cli/internal/domain"
	"github.com/arthurbcp/kuma-cli/internal/handlers"
	"github.com/arthurbcp/kuma-cli/internal/helpers"
	"github.com/arthurbcp/kuma-cli/pkg/filesystem"
	"github.com/arthurbcp/kuma-cli/pkg/style"
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
				handleCommand(value.(string), vars)
			} else if key == "input" {
				handleInput(value.(map[string]interface{}), vars)
			} else if key == "log" {
				handleLog(value.(string), vars)
			} else if key == "run" {
				ExecRun(value.(string), vars)
			} else if key == "create" {
				handleCreate(value.(map[string]interface{}), vars)
			} else if key == "load" {
				handleLoad(value.(map[string]interface{}), vars)
			}
		}
	}
}

func handleInput(input map[string]interface{}, vars map[string]interface{}) {
	data := vars["data"].(map[string]interface{})
	label, ok := input["label"].(string)
	if !ok {
		style.ErrorPrint("label is required for input")
		os.Exit(1)
	}
	out, ok := input["out"].(string)
	if !ok {
		style.ErrorPrint("out is required for input")
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
				style.ErrorPrint("error running program: " + err.Error())
				os.Exit(1)
			}
			selectedChoices := make([]string, 0)
			for key, value := range output.Choices {
				if value {
					selectedChoices = append(selectedChoices, key)
				}
			}
			data[out] = selectedChoices
		} else {
			output := &selectInput.Selection{}
			p := tea.NewProgram(selectInput.InitialSelectInputModel(options, output, label, other, skippable, false))
			_, err := p.Run()
			if err != nil {
				style.ErrorPrint("error running program: " + err.Error())
				os.Exit(1)
			}
			data[out] = output.Choice
		}
	} else {
		output := &textInput.Output{}
		p := tea.NewProgram(textInput.InitialTextInputModel(output, label, false))
		_, err := p.Run()
		if err != nil {
			style.ErrorPrint("error running program: " + err.Error())
			os.Exit(1)
		}
		data[out] = output.Output
	}
}

func handleLog(log string, vars map[string]interface{}) {
	var err error
	helpers := helpers.NewHelpers()
	log, err = helpers.ReplaceVars(log, vars, helpers.GetFuncMap())
	if err != nil {
		style.ErrorPrint("parsing log error: " + err.Error())
		os.Exit(1)
	}
	style.LogPrint(log)
}

func handleCommand(cmdStr string, vars map[string]interface{}) {
	var err error
	helpers := helpers.NewHelpers()
	cmdStr, err = helpers.ReplaceVars(cmdStr, vars, helpers.GetFuncMap())
	if err != nil {
		style.ErrorPrint("parsing command error: " + err.Error())
		os.Exit(1)
	}
	style.LogPrint(fmt.Sprintf("running: %s", cmdStr))
	cmdArgs := strings.Split(cmdStr, " ")
	execCmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	// Set the command's standard output to the console
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr

	// Execute the command
	err = execCmd.Run()
	if err != nil {
		style.ErrorPrint("command error: " + err.Error())
		os.Exit(1)
	}
}

func handleCreate(data map[string]interface{}, vars map[string]interface{}) {
	fs := filesystem.NewFileSystem(afero.NewOsFs())
	helpers := helpers.NewHelpers()
	// Initialize a new Builder with the provided configurations.
	builder, err := domain.NewBuilder(fs, helpers, domain.NewConfig(".", shared.KumaTemplatesPath))
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}
	from, ok := data["from"].(string)
	if !ok {
		style.ErrorPrint("from is required")
		os.Exit(1)
	}
	err = builder.SetBuilderDataFromFile(from, vars)
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}

	// Execute the build process using the BuilderHandler.
	if err = handlers.NewBuilderHandler(builder).Build(); err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}
}

func handleLoad(load map[string]interface{}, vars map[string]interface{}) {
	var err error
	data := vars["data"].(map[string]interface{})
	helpers := helpers.NewHelpers()
	fs := filesystem.NewFileSystem(afero.NewOsFs())
	from, ok := load["from"].(string)
	if !ok {
		style.ErrorPrint("from is required")
		os.Exit(1)
	}
	from, err = helpers.ReplaceVars(from, vars, helpers.GetFuncMap())
	if err != nil {
		style.ErrorPrint("parsing from error: " + err.Error())
		os.Exit(1)
	}
	out, ok := load["out"].(string)
	if !ok {
		style.ErrorPrint("out is required")
		os.Exit(1)
	}
	var fileVars map[string]interface{}
	_, err = url.ParseRequestURI(from)
	if err != nil {
		fileVars, err = helpers.UnmarshalFile(from, fs)
		if err != nil {
			style.ErrorPrint("parsing file error: " + err.Error())
			os.Exit(1)
		}
	} else {
		style.TitlePrint("downloading variables file")
		varsContent, err := fs.ReadFileFromURL(from)
		if err != nil {
			style.ErrorPrint("reading file error: " + err.Error())
			os.Exit(1)
		}
		splitURL := strings.Split(from, "/")
		fileVars, err = helpers.UnmarshalByExt(splitURL[len(splitURL)-1], []byte(varsContent))
		if err != nil {
			style.ErrorPrint("parsing file error: " + err.Error())
			os.Exit(1)
		}
	}
	data[out] = fileVars
}

// init sets up flags for the 'run' subcommand and binds them to variables.
func init() {
	// Repository name
	ExecRunCmd.Flags().StringVarP(&Run, "run", "r", "", "run to use")
	ExecRunCmd.MarkFlagRequired("run")
}
