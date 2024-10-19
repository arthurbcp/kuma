package execFormHandlers

import (
	"os"

	execBuilders "github.com/arthurbcp/kuma/v2/cmd/commands/exec/builders"
	"github.com/arthurbcp/kuma/v2/cmd/constants"
	"github.com/arthurbcp/kuma/v2/pkg/style"
	"github.com/charmbracelet/huh"
)

func HandleMultiSelect(input map[string]interface{}, vars map[string]interface{}) (*huh.MultiSelect[string], string, *[]string) {
	var err error

	label, err := execBuilders.BuildStringValue("label", input, vars, false, constants.MultiSelectComponent)
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}
	description, err := execBuilders.BuildStringValue("description", input, vars, false, constants.MultiSelectComponent)
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}
	out, err := execBuilders.BuildStringValue("out", input, vars, true, constants.MultiSelectComponent)
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}
	limit, err := execBuilders.BuildIntValue("limit", input, vars, false, constants.MultiSelectComponent)
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}
	options := []huh.Option[string]{}
	if mapOptions, ok := input["options"].([]interface{}); ok {
		for _, option := range mapOptions {
			optionMap := option.(map[string]interface{})
			label, err := execBuilders.BuildStringValue("label", optionMap, vars, true, constants.MultiSelectOptionComponent)
			if err != nil {
				style.ErrorPrint(err.Error())
				os.Exit(1)
			}
			value, err := execBuilders.BuildStringValue("value", optionMap, vars, false, constants.MultiSelectOptionComponent)
			if err != nil {
				style.ErrorPrint(err.Error())
				os.Exit(1)
			}
			if value == "" {
				value = label
			}
			options = append(options, huh.NewOption[string](label, value))
		}

		var outValue []string
		h := huh.NewMultiSelect[string]().
			Title(label).
			Description(description).
			Options(options...).
			Value(&outValue)

		if limit > 0 {
			h.Limit(limit)
		}

		return h, out, &outValue
	}
	return nil, out, nil
}
