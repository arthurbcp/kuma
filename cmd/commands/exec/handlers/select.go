package execHandlers

import (
	"os"

	execBuilders "github.com/arthurbcp/kuma/cmd/commands/exec/builders"
	"github.com/arthurbcp/kuma/cmd/ui/utils/program"
	"github.com/arthurbcp/kuma/pkg/style"
)

func HandleSelect(input map[string]interface{}, vars map[string]interface{}) {
	var err error
	program := program.NewProgram()
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

	other, err := execBuilders.BuildBoolValue("other", input, vars, true)
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}

	multi, err := execBuilders.BuildBoolValue("multi", input, vars, true)
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
	}
}
