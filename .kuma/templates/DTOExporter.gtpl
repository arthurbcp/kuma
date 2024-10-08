{{range .data}}import { {{ toPascalCase . }} } from "./{{ toSnakeCase . }}" 
{{end}}

export {
{{range .data}}{{toPascalCase .}},
{{end -}}
}