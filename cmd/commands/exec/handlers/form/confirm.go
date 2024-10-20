package execFormHandlers

import (
	"os"

	execBuilders "github.com/arthurbcp/kuma/v2/cmd/commands/exec/builders"
	"github.com/arthurbcp/kuma/v2/cmd/constants"
	"github.com/arthurbcp/kuma/v2/pkg/style"
	"github.com/charmbracelet/huh"
)

func HandleConfirm(input map[string]interface{}, vars map[string]interface{}) (*huh.Confirm, string, *bool) {
	var err error
	data := vars["data"].(map[string]interface{})

	label, err := execBuilders.BuildStringValue("label", input, vars, false, constants.ConfirmComponent)
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}
	description, err := execBuilders.BuildStringValue("description", input, vars, false, constants.ConfirmComponent)
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}
	affirmative, err := execBuilders.BuildStringValue("affirmative", input, vars, false, constants.ConfirmComponent)
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}
	negative, err := execBuilders.BuildStringValue("negative", input, vars, false, constants.ConfirmComponent)
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}
	out, err := execBuilders.BuildStringValue("out", input, vars, true, constants.ConfirmComponent)
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}

	if affirmative == "" {
		affirmative = "Yes"
	}
	if negative == "" {
		negative = "No"
	}

	var outValue bool
	h := huh.NewConfirm().
		Title(label).
		Description(description).
		Affirmative(affirmative).
		Negative(negative).
		Value(&outValue)

	data[out] = out

	return h, out, &outValue
}
