package execFormHandlers

import (
	"os"

	execBuilders "github.com/arthurbcp/kuma/cmd/commands/exec/builders"
	"github.com/arthurbcp/kuma/pkg/style"
	"github.com/charmbracelet/huh"
)

func HandleInput(input map[string]interface{}, vars map[string]interface{}) *huh.Input {
	var err error
	data := vars["data"].(map[string]interface{})

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
	placeholder, err := execBuilders.BuildStringValue("placeholder", input, vars, false)
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}
	var outValue string
	h := huh.NewInput().
		Title(label).
		Placeholder(placeholder).
		Value(&outValue)

	data[out] = out

	return h
}
