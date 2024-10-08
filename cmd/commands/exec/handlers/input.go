package execHandlers

import (
	"os"

	execBuilders "github.com/arthurbcp/kuma-cli/cmd/commands/exec/builders"
	"github.com/arthurbcp/kuma-cli/cmd/ui/textInput"
	"github.com/arthurbcp/kuma-cli/cmd/ui/utils/program"
	"github.com/arthurbcp/kuma-cli/pkg/style"
	tea "github.com/charmbracelet/bubbletea"
)

func HandleInput(input map[string]interface{}, vars map[string]interface{}) {
	var err error
	program := program.NewProgram()
	data := vars["data"].(map[string]interface{})

	label, err := execBuilders.BuildStringValue("label", input, vars)
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}

	out, err := execBuilders.BuildStringValue("out", input, vars)
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}

	other, err := execBuilders.BuildBoolValue("other", input, vars)
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}

	multi, err := execBuilders.BuildBoolValue("multi", input, vars)
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}

	if mapOptions, ok := input["options"].([]interface{}); ok {
		options, err := execBuilders.BuildOptions(program, mapOptions, other, multi, label, vars)
		if err != nil {
			style.ErrorPrint(err.Error())
			os.Exit(1)
		}
		data[out] = options
	} else {
		output := &textInput.Output{}
		p := tea.NewProgram(textInput.InitialTextInputModel(output, label, program))
		_, err := p.Run()

		program.ExitCLI(p)

		if err != nil {
			style.ErrorPrint("error running program: " + err.Error())
			os.Exit(1)
		}
		data[out] = output.Output
	}
}
