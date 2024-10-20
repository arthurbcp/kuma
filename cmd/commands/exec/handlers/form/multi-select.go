package execFormHandlers

import (
	execBuilders "github.com/arthurbcp/kuma/v2/cmd/commands/exec/builders"
	"github.com/arthurbcp/kuma/v2/cmd/constants"
	"github.com/charmbracelet/huh"
)

func HandleMultiSelect(input map[string]interface{}, vars map[string]interface{}) (*huh.MultiSelect[string], string, *[]string, error) {
	var err error

	label, err := execBuilders.BuildStringValue("label", input, vars, false, constants.MultiSelectComponent)
	if err != nil {
		return nil, "", nil, err
	}
	description, err := execBuilders.BuildStringValue("description", input, vars, false, constants.MultiSelectComponent)
	if err != nil {
		return nil, "", nil, err
	}
	out, err := execBuilders.BuildStringValue("out", input, vars, true, constants.MultiSelectComponent)
	if err != nil {
		return nil, "", nil, err
	}
	limit, err := execBuilders.BuildIntValue("limit", input, vars, false, constants.MultiSelectComponent)
	if err != nil {
		return nil, "", nil, err
	}
	options := []huh.Option[string]{}
	if mapOptions, ok := input["options"].([]interface{}); ok {
		for _, option := range mapOptions {
			optionMap := option.(map[string]interface{})
			label, err := execBuilders.BuildStringValue("label", optionMap, vars, true, constants.MultiSelectOptionComponent)
			if err != nil {
				return nil, "", nil, err
			}
			value, err := execBuilders.BuildStringValue("value", optionMap, vars, false, constants.MultiSelectOptionComponent)
			if err != nil {
				return nil, "", nil, err
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

		return h, out, &outValue, nil
	}
	return nil, out, nil, nil
}
