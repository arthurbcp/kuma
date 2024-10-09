package functions

func GetPathsByTag(paths map[string]interface{}, tag string) map[string]interface{} {
	filteredPaths := make(map[string]interface{})
	for path, pathItem := range paths {
		if pathMap, ok := pathItem.(map[string]interface{}); ok {
			for _, operation := range pathMap {
				if operationMap, ok := operation.(map[string]interface{}); ok {
					if pathTags, ok := operationMap["tags"].([]interface{}); ok {
						for _, tagItem := range pathTags {
							if tagStr, ok := tagItem.(string); ok {
								if tagStr == tag {
									filteredPaths[path] = pathItem
								}
							}
						}
					}
				}
			}
		}
	}
	return filteredPaths
}
