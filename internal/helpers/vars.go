package helpers

import (
	"html/template"
	"strings"
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
