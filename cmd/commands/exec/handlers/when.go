package execHandlers

import (
	"os"

	execBuilders "github.com/arthurbcp/kuma/v2/cmd/commands/exec/builders"
	"github.com/arthurbcp/kuma/v2/pkg/style"
)

func HandleWhen(module string, data map[string]interface{}, vars map[string]interface{}) {
	isTrue, err := execBuilders.BuildBoolValue("condition", data, vars, true)
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}
	run, err := execBuilders.BuildStringValue("run", data, vars, true)
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}
	vars = map[string]interface{}{
		"data": map[string]interface{}{},
	}
	if isTrue {
		HandleRun(run, module, vars)
	}
}
