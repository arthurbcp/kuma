package helpers

import (
	"strings"
	"text/template"

	"github.com/arthurbcp/kuma-cli/internal/functions"
	"github.com/go-sprout/sprout/sprigin"
)

func GetFuncMap() template.FuncMap {
	fnMap := sprigin.TxtFuncMap()
	fnMap["toYaml"] = functions.ToYaml
	fnMap["getRefFrom"] = functions.GetRefFrom
	fnMap["getPathsByTag"] = functions.GetPathsByTag
	fnMap["getParamsByType"] = functions.GetParamsByType
	return fnMap
}

func ReplaceVars(text string, vars interface{}, funcs template.FuncMap) (string, error) {
	t, err := template.New("").Funcs(GetFuncMap()).Parse(text)
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
