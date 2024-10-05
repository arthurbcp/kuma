package helpers

import (
	"strings"
	"text/template"

	"github.com/go-sprout/sprout/sprigin"
	"gopkg.in/yaml.v3"
)

func ToYaml(data interface{}) []string {
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		panic("Error parsing YAML template: " + err.Error())
	}
	lines := strings.Split(string(yamlData), "\n")
	return lines
}

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

func (h *Helpers) GetFuncMap() template.FuncMap {
	fnMap := sprigin.TxtFuncMap()
	fnMap["toYaml"] = ToYaml
	fnMap["getRefFrom"] = GetRefFrom
	fnMap["getPathsByTag"] = GetPathsByTag
	fnMap["getParamsByType"] = GetParamsByType
	return fnMap
}

func (h *Helpers) ReplaceVars(text string, vars interface{}, funcs template.FuncMap) (string, error) {
	t, err := template.New("").Funcs(h.GetFuncMap()).Parse(text)
	if err != nil {
		return "", err
	}
	var buf strings.Builder
	err = t.Execute(&buf, vars)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
