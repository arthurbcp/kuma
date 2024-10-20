package execHandlers

import (
	execBuilders "github.com/arthurbcp/kuma/v2/cmd/commands/exec/builders"
	"github.com/arthurbcp/kuma/v2/cmd/constants"
)

func HandleWhen(module string, params map[string]interface{}, vars map[string]interface{}) error {
	isTrue, err := execBuilders.BuildBoolValue("condition", params, vars, true, constants.WhenHandler)
	if err != nil {
		return err
	}
	run, err := execBuilders.BuildStringValue("run", params, vars, true, constants.WhenHandler)
	if err != nil {
		return err
	}
	if isTrue {
		err := HandleRun(run, module, vars)
		if err != nil {
			return err
		}
	}

	return nil
}
