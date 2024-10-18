package execFormHandlers

import (
	"os"

	execBuilders "github.com/arthurbcp/kuma/cmd/commands/exec/builders"
	"github.com/arthurbcp/kuma/pkg/style"
	"github.com/charmbracelet/huh"
)

func HandleSelect(input map[string]interface{}, vars map[string]interface{}) (*huh.Select[string], string, *string) {
	var err error

	label, err := execBuilders.BuildStringValue("label", input, vars, false)
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}

	out, err := execBuilders.BuildStringValue("out", input, vars, true)
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}

	options := []huh.Option[string]{}
	if mapOptions, ok := input["options"].([]interface{}); ok {
		for _, option := range mapOptions {
			optionMap := option.(map[string]interface{})
			label, err := execBuilders.BuildStringValue("label", optionMap, vars, true)
			if err != nil {
				style.ErrorPrint(err.Error())
				os.Exit(1)
			}
			value, err := execBuilders.BuildStringValue("value", optionMap, vars, false)
			if err != nil {
				style.ErrorPrint(err.Error())
				os.Exit(1)
			}
			if value == "" {
				value = label
			}
			options = append(options, huh.NewOption[string](label, value))
		}

		var outValue string
		h := huh.NewSelect[string]().
			Title(label).
			Options(options...).
			Value(&outValue)

		return h, out, &outValue
	}
	return nil, out, nil
}
