package execFormHandlers

import (
	execBuilders "github.com/arthurbcp/kuma/v2/cmd/commands/exec/builders"
	"github.com/arthurbcp/kuma/v2/cmd/constants"
	"github.com/charmbracelet/huh"
)

func HandleText(input map[string]interface{}, vars map[string]interface{}) (*huh.Text, string, *string, error) {
	var err error
	data := vars["data"].(map[string]interface{})

	label, err := execBuilders.BuildStringValue("label", input, vars, false, constants.TextComponent)
	if err != nil {
		return nil, "", nil, err
	}
	description, err := execBuilders.BuildStringValue("description", input, vars, false, constants.TextComponent)
	if err != nil {
		return nil, "", nil, err
	}
	out, err := execBuilders.BuildStringValue("out", input, vars, true, constants.TextComponent)
	if err != nil {
		return nil, "", nil, err
	}
	placeholder, err := execBuilders.BuildStringValue("placeholder", input, vars, false, constants.TextComponent)
	if err != nil {
		return nil, "", nil, err
	}
	var outValue string
	h := huh.NewText().
		Title(label).
		Description(description).
		Placeholder(placeholder).
		Value(&outValue)

	data[out] = out

	return h, out, &outValue, nil
}
