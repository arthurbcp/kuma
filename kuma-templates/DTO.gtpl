{{- range .Data.properties -}}
{{$ref :=  getRefFrom . }}
{{- if not $ref  }}{{if getRefFrom .items}}{{$ref = getRefFrom .items}}{{- end}}{{- end}}
{{- if $ref }}
import { {{ toPascalCase $ref }} } from './{{ toSnakeCase $ref }}'
{{- end -}}
{{- end -}}
{{"\n"}}
{{- if .Data.description }}
// {{ .Data.description }}
{{- end }}
export type {{ toPascalCase .Data.name }} = { 
    {{- range $name, $prop := .Data.properties}}
    {{- if $prop.description }}
    // {{ $prop.description }}
    {{- end }}
    {{ $name }}{{- if not $prop.required -}}?{{- end -}}: {{- block "TypeResolver" $prop }}{{end}}
    {{- end}}
}