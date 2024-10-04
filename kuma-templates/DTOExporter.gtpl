{{range .Data}}import { {{ toPascalCase . }} } from "./{{ toSnakeCase . }}" 
{{end}}

export {
{{range .Data}}{{toPascalCase .}},
{{end -}}
}