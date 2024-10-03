package helpers

import (
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
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

func (h *Helpers) GetFuncMap() template.FuncMap {
	fnMap := sprig.TxtFuncMap()
	fnMap["toYaml"] = ToYaml
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
