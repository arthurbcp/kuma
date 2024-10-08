package execBuilders

import (
	"fmt"

	"github.com/arthurbcp/kuma-cli/internal/helpers"
)

func BuildStringValue(key string, input map[string]interface{}, vars map[string]interface{}) (string, error) {
	var err error
	val, ok := input[key].(string)
	if !ok {
		return "", fmt.Errorf("%s is required for input", key)
	}
	val, err = helpers.ReplaceVars(val, vars, helpers.GetFuncMap())
	if err != nil {
		return "", err
	}
	return val, nil
}
