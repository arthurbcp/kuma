{{- define "TypeResolver" -}}
    {{- if getRefFrom . -}}
        {{ getRefFrom . }}
    {{- else if eq .type "integer" -}}
        number
    {{- else if eq .type "number" -}}
        number
    {{- else if eq .type "string" -}}
        {{- if .enum }}
            {{- range $index, $enumVal := .enum -}}
                {{- if $index }} | {{ end }}"{{ $enumVal }}"
            {{- end -}}
        {{- else -}}
            string
        {{- end -}}
    {{- else if eq .type "array" -}}
        {{- block "TypeResolver" .items }}{{- end -}}[]
    {{- else if eq .type "object" -}}
        { [key: string]: any }
    {{- else if eq .type "boolean" -}}
        boolean
    {{- else -}}
        undefined
    {{- end -}}
{{- end -}}