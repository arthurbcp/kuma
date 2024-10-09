package functions

import "strings"

func GetRefFrom(object map[string]interface{}) string {
	ref, ok := object["$ref"].(string)
	if !ok {
		return ""
	}
	const refPrefix = "#/definitions/"
	if strings.HasPrefix(ref, refPrefix) {
		return strings.TrimPrefix(ref, refPrefix)
	}
	return ""
}
