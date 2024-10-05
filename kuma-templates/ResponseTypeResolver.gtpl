{{- define "ResponseTypeResolver" -}}
{{- if gt (len .) 0 -}}
    {{range $index, $response := .}}{{- if $index -}} | {{end}} {{- block "TypeResolver" $response  }}{{end}}{{end}}
{{- else -}}
    undefined
{{- end -}}
{{- end -}}