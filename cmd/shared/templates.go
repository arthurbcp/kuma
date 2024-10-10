package shared

import "github.com/arthurbcp/kuma/internal/domain"

var Templates = map[string]domain.Template{
	"arthurbcp/typescript-rest-openapi-services": domain.NewTemplate(
		"TypeScript Rest Services (OpenAPI 2.0)",
		"Create a library TypeScript with services typed for all endpoints described in a file Open API 2.0",
		[]string{"typescript", "openapi", "rest", "library"},
	),
	"arthurbcp/kuma-hello-world": domain.NewTemplate(
		"Hello World",
		"A simple Hello World in Go!",
		[]string{"golang", "example"},
	),
	"arthurbcp/kuma-changelog-generator": domain.NewTemplate(
		"Changelog Generator",
		"Helper to write a good changelog to your project",
		[]string{"changelog", "helper", "markdown"},
	),
}
