package execFormHandlers

import (
	"os"

	execBuilders "github.com/arthurbcp/kuma/v2/cmd/commands/exec/builders"
	"github.com/arthurbcp/kuma/v2/cmd/constants"
	"github.com/arthurbcp/kuma/v2/pkg/style"
	"github.com/charmbracelet/huh"
)

func HandleText(input map[string]interface{}, vars map[string]interface{}) (*huh.Text, string, *string) {
	var err error
	data := vars["data"].(map[string]interface{})

	label, err := execBuilders.BuildStringValue("label", input, vars, false, constants.TextComponent)
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}
	description, err := execBuilders.BuildStringValue("description", input, vars, false, constants.TextComponent)
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}
	out, err := execBuilders.BuildStringValue("out", input, vars, true, constants.TextComponent)
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}
	placeholder, err := execBuilders.BuildStringValue("placeholder", input, vars, false, constants.TextComponent)
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}
	var outValue string
	h := huh.NewText().
		Title(label).
		Description(description).
		Placeholder(placeholder).
		Value(&outValue)

	data[out] = out

	return h, out, &outValue
}
