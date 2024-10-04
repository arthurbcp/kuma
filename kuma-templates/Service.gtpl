import { AxiosRequestConfig, AxiosResponse } from "axios";
import { IHttpProvider } from "../../providers/http/http_provider_interface";{{"\n"}}
{{- $imports := list -}}
{{- range $path, $op :=  .Data.operations -}}
  {{- range  $method, $data := $op -}}
    {{- $body := getParamsByType $data.parameters "body" -}}
    {{- range $body -}}
      {{- if getRefFrom .schema -}}
        {{- $import := getRefFrom .schema -}}
        {{- $hasImport := $imports | has $import }}
        {{- if not $hasImport -}}
          {{- $imports = $imports | append $import -}}
import { {{toPascalCase $import}} } from "../../dto/{{ toSnakeCase $import }}"{{"\n"}}
        {{- end -}}
      {{- end -}}
    {{- end -}}
    {{- range $data.responses -}}
      {{- if getRefFrom .schema -}}
        {{- $import := getRefFrom .schema -}}
        {{- $hasImport := $imports | has $import }}
        {{- if not $hasImport -}}
           {{- $imports = $imports | append $import -}}
           import { {{toPascalCase $import}} } from "../../dto/{{ toSnakeCase $import }}"{{"\n"}}
        {{- end -}}
      {{- end -}}
    {{- end -}}
  {{- end -}}
{{- end -}}
{{"\n"}}

export class {{toPascalCase .Data.name}}Service {
  private http: IHttpProvider;

  constructor(http: IHttpProvider) {
    this.http = http;
  }
  {{"\n"}}
  {{- range $path, $op :=  .Data.operations -}}
  {{ range  $method, $data := $op }}
    {{- $params := getParamsByType $data.parameters "path" -}}
    {{- $query := getParamsByType $data.parameters "query" -}}
    {{- $body := getParamsByType $data.parameters "body" -}}
    {{ if $data.description }}//{{ $data.description }}{{end}}
    async {{ $data.operationId }}(data?: {
      {{- if $params }}params?: { {{ range $params }}
        {{ .name }}{{- if not .required -}}?{{- end -}}: {{- block "TypeResolver" . }}{{end}},
      {{- end -}} },{{- end -}}
      {{- if $query }}query?: { {{ range $query }}
        {{ .name }}{{- if not .required -}}?{{- end -}}: {{- block "TypeResolver" . }}{{end}},
      {{- end -}} },{{- end -}}
      {{- if $body }}body?: {{ range $body }}
        {{- block "TypeResolver" .schema  }}{{end}},
      {{- end -}}{{- end -}}
    },   config?: AxiosRequestConfig):
     
      {{- $responses := list -}}
      {{- range $index, $response := $data.responses -}}
      {{- if $response.schema -}}
        {{- $inResponse := $responses | has $response.schema -}}
        {{- if not $inResponse -}}
          {{- $responses = $responses | append $response.schema -}}
        {{- end -}}
      {{- end -}}
    {{- end -}}
    Promise<AxiosResponse<{{- block "ResponseTypeResolver" $responses -}}{{end}}>> {
      return await this.http.request<{{- block "ResponseTypeResolver" $responses -}}{{end}}>("{{ toLower $method }}","{{ $path }}", data, config);
    }

  {{end}}
  {{end}}
}