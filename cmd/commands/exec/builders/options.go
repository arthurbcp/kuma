package execBuilders

import (
	"os"

	"github.com/arthurbcp/kuma/cmd/ui/multiSelectInput"
	"github.com/arthurbcp/kuma/cmd/ui/selectInput"
	"github.com/arthurbcp/kuma/cmd/ui/utils/program"
	"github.com/arthurbcp/kuma/cmd/ui/utils/steps"
	"github.com/arthurbcp/kuma/internal/helpers"
	"github.com/arthurbcp/kuma/pkg/style"
	tea "github.com/charmbracelet/bubbletea"
)

func BuildOptions(program *program.Program, mapOptions []interface{}, other bool, multi bool, label string, vars map[string]interface{}) (interface{}, error) {
	var err error

	options, err := getOptions(mapOptions, vars)
	if err != nil {
		return nil, err
	}
	if multi {
		choices := make(map[string]bool)
		for _, option := range options {
			choices[option.Value] = false
		}
		output := &multiSelectInput.Selection{
			Choices: choices,
		}
		p := tea.NewProgram(multiSelectInput.InitialMultiSelectInputModel(options, output, label, program))
		_, err := p.Run()

		program.ExitCLI(p)

		if err != nil {
			return nil, err
		}
		selectedChoices := make([]string, 0)
		for key, value := range output.Choices {
			if value {
				selectedChoices = append(selectedChoices, key)
			}
		}

		return selectedChoices, nil
	}
	output := &selectInput.Selection{}
	p := tea.NewProgram(selectInput.InitialSelectInputModel(options, output, label, other, program))
	_, err = p.Run()

	program.ExitCLI(p)

	if err != nil {
		style.ErrorPrint("error running program: " + err.Error())
		os.Exit(1)
	}
	return output.Choice, nil
}

func getOptions(mapOptions []interface{}, vars map[string]interface{}) ([]steps.Item, error) {
	var err error

	options := make([]steps.Item, len(mapOptions))
	for i, option := range mapOptions {

		label := option.(map[string]interface{})["label"].(string)
		label, err = helpers.ReplaceVars(label, vars, helpers.GetFuncMap())
		if err != nil {
			style.ErrorPrint("parsing out error: " + err.Error())
			os.Exit(1)
		}

		value := option.(map[string]interface{})["value"].(string)
		value, err = helpers.ReplaceVars(value, vars, helpers.GetFuncMap())
		if err != nil {
			style.ErrorPrint("parsing out error: " + err.Error())
			os.Exit(1)
		}

		options[i] = steps.Item{
			Label: label,
			Value: value,
		}
	}
	return options, nil
}
