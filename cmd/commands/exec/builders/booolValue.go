package execBuilders

func BuildBoolValue(key string, input map[string]interface{}, vars map[string]interface{}) (bool, error) {
	val := false
	if o, ok := input[key]; ok {
		val = o.(bool)
	}
	return val, nil
}
