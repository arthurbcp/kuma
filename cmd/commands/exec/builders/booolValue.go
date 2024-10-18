package execBuilders

import "strconv"

func BuildBoolValue(key string, input map[string]interface{}, vars map[string]interface{}, required bool) (bool, error) {
	val := false
	val, ok := input[key].(bool)
	if !ok {
		if valStr, err := BuildStringValue(key, input, vars, required); err == nil {
			if valStr == "" {
				return false, nil
			}
			val, err = strconv.ParseBool(valStr)
			if err != nil {
				return false, err
			}
		}
	}
	return val, nil
}
