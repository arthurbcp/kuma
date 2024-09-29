package helpers

import (
	"encoding/json"
	"html/template"
	"strings"
	"unicode"

	"gopkg.in/yaml.v3"
)

var FuncMap = template.FuncMap{
	"startsWith": strings.HasPrefix,
	"toUpper":    strings.ToUpper,
	"toLower":    strings.ToLower,
	"toSnake":    ToSnakeCase,
	"toPascal":   ToPascalCase,
	"yaml":       ParseYamlTemplate,
	"json":       ParseJsonTemplate,
}

func ToSnakeCase(str string) string {
	var result []rune
	for i, r := range str {
		if i == 0 {
			result = append(result, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			result = append(result, '_')
		}
		result = append(result, unicode.ToLower(r))
	}
	return string(result)
}

func ToPascalCase(str string) string {
	if len(str) == 0 {
		return str
	}

	return strings.ToUpper(string(str[0])) + str[1:]
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

func (h *Helpers) ReplaceVars(text string, vars interface{}, funcs template.FuncMap) (string, error) {
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
