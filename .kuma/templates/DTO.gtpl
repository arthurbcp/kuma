{{- range .data.properties -}}
{{$ref :=  getRefFrom . }}
{{- if not $ref  }}{{if getRefFrom .items}}{{$ref = getRefFrom .items}}{{- end}}{{- end}}
{{- if $ref }}
import { {{ toPascalCase $ref }} } from './{{ toSnakeCase $ref }}'
{{- end -}}
{{- end -}}
{{"\n"}}
{{- if .data.description }}
// {{ .data.description }}
{{- end }}
export type {{ toPascalCase .data.name }} = { 
    {{- range $name, $prop := .data.properties}}
    {{- if $prop.description }}
    // {{ $prop.description }}
    {{- end }}
    {{ $name }}{{- if not $prop.required -}}?{{- end -}}: {{- block "TypeResolver" $prop }}{{end}}
    {{- end}}
}