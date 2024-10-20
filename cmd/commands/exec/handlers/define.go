package execHandlers

import (
	execBuilders "github.com/arthurbcp/kuma/v2/cmd/commands/exec/builders"
	"github.com/arthurbcp/kuma/v2/cmd/constants"
)

func HandleDefine(params map[string]interface{}, vars map[string]interface{}) error {
	data := vars["data"].(map[string]interface{})
	variable, err := execBuilders.BuildStringValue("variable", params, vars, true, constants.DefineHandler)
	if err != nil {
		return err
	}
	var value any
	value, err = execBuilders.BuildBoolValue("value", params, vars, true, constants.DefineHandler)
	if err != nil {
		value, err = execBuilders.BuildIntValue("value", params, vars, true, constants.DefineHandler)
		if err != nil {
			value, err = execBuilders.BuildStringValue("value", params, vars, true, constants.DefineHandler)
			if err != nil {
				return err
			}
		}
	}
	data[variable] = value
	return nil
}
