import { IHttpProvider } from "../providers/http";
import { {{- range $index, $service := .Data}}{{if $index}},{{end}}{{toPascalCase .}}Service{{end}} } from "./";

export const buildServices = (provider: IHttpProvider) => {
  return {
  {{- range .Data}}
    {{toCamelCase .}}: new {{toPascalCase .}}Service(provider),
  {{- end -}}
  };
};
