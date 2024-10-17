package execHandlers

import (
	"os"

	execBuilders "github.com/arthurbcp/kuma/cmd/commands/exec/builders"
	"github.com/arthurbcp/kuma/pkg/style"
)

func HandleWhen(module string, data map[string]interface{}, vars map[string]interface{}) {
	isTrue, err := execBuilders.BuildBoolValue("condition", data, vars)
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}
	run, err := execBuilders.BuildStringValue("run", data, vars)
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
