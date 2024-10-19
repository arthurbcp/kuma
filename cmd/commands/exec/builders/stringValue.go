package execBuilders

import (
	"fmt"

	"github.com/arthurbcp/kuma/v2/internal/helpers"
)

func BuildStringValue(key string, input map[string]interface{}, vars map[string]interface{}, required bool, component string) (string, error) {
	var err error
	val, ok := input[key].(string)
	if !ok {
		if required {
			return "", fmt.Errorf("%s is required for %s", key, component)
		}
		return "", nil
	}
	val, err = helpers.ReplaceVars(val, vars, helpers.GetFuncMap())
	if err != nil {
		return "", err
	}
	return val, nil
}
