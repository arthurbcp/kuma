package execHandlers

import (
	"os"

	"github.com/arthurbcp/kuma-cli/cmd/steps"
	"github.com/arthurbcp/kuma-cli/cmd/ui/multiSelectInput"
	"github.com/arthurbcp/kuma-cli/cmd/ui/selectInput"
	"github.com/arthurbcp/kuma-cli/cmd/ui/textInput"
	"github.com/arthurbcp/kuma-cli/internal/helpers"
	"github.com/arthurbcp/kuma-cli/pkg/style"
	tea "github.com/charmbracelet/bubbletea"
)

func HandleInput(input map[string]interface{}, vars map[string]interface{}) {
	var err error
	helpers := helpers.NewHelpers()
	data := vars["data"].(map[string]interface{})

	label, ok := input["label"].(string)
	if !ok {
		style.ErrorPrint("label is required for input")
		os.Exit(1)
	}
	label, err = helpers.ReplaceVars(label, vars, helpers.GetFuncMap())
	if err != nil {
		style.ErrorPrint("parsing label error: " + err.Error())
		os.Exit(1)
	}

	out, ok := input["out"].(string)
	if !ok {
		style.ErrorPrint("out is required for input")
		os.Exit(1)
	}
	out, err = helpers.ReplaceVars(out, vars, helpers.GetFuncMap())
	if err != nil {
		style.ErrorPrint("parsing out error: " + err.Error())
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
		options, err := getOptions(mapOptions, vars)
		if err != nil {
			style.ErrorPrint("parsing options error: " + err.Error())
			os.Exit(1)
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

func getOptions(mapOptions []interface{}, vars map[string]interface{}) ([]steps.Item, error) {
	var err error
	helpers := helpers.NewHelpers()
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
