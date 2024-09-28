package helpers

import (
	"encoding/json"
	"html/template"
	"strings"

	"gopkg.in/yaml.v3"
)

func ReplaceVars(text string, vars interface{}, funcs template.FuncMap) (string, error) {
	t, err := template.New("").Funcs(funcs).Parse(text)
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

func ParseYamlTemplate(data interface{}) []string {
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		panic("Error parsing YAML template: " + err.Error())
	}
	lines := strings.Split(string(yamlData), "\n")
	return lines
}

func ParseJsonTemplate(data interface{}) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic("Error parsing JSON template: " + err.Error())
	}
	return string(jsonData)
}
