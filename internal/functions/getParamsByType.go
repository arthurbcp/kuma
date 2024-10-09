package functions

func GetParamsByType(params []interface{}, paramType string) []interface{} {
	filteredParams := make([]interface{}, 0)
	for _, param := range params {
		if paramMap, ok := param.(map[string]interface{}); ok {
			if paramTypeStr, ok := paramMap["in"].(string); ok {
				if paramTypeStr == paramType {
					filteredParams = append(filteredParams, param)
				}
			}
		}
	}
	return filteredParams
}
