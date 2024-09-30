package openapi

// OpenAPITemplate represents the overall structure of the parsed OpenAPI specification.
type OpenAPITemplate struct {
	Version         string                      `json:"Version,omitempty"`
	InfoTitle       string                      `json:"InfoTitle,omitempty"`
	InfoDescription string                      `json:"InfoDescription,omitempty"`
	InfoVersion     string                      `json:"InfoVersion,omitempty"`
	Servers         []OpenApiTemplateServer     `json:"Servers,omitempty"`
	Controllers     []OpenApiTemplateController `json:"Controllers,omitempty"`
	Components      []OpenApiTemplateComponent  `json:"Components,omitempty"`
}

// OpenApiTemplateServer represents a server in the OpenAPI specification.
type OpenApiTemplateServer struct {
	Url         string                    `json:"Url,omitempty"`
	Description string                    `json:"Description,omitempty"`
	Variables   []OpenApiTemplateVariable `json:"Variables,omitempty"`
}

// OpenApiTemplateVariable represents a server variable.
type OpenApiTemplateVariable struct {
	Name        string   `json:"Name,omitempty"`
	Description string   `json:"Description,omitempty"`
	Default     string   `json:"Default,omitempty"`
	Enum        []string `json:"Enum,omitempty"`
}

// OpenApiTemplateController groups endpoints under a controller (based on tags).
type OpenApiTemplateController struct {
	Name      string                    `json:"Name,omitempty"`
	Endpoints []OpenApiTemplateEndpoint `json:"Endpoints,omitempty"`
}

// OpenApiTemplateEndpoint represents an API endpoint.
type OpenApiTemplateEndpoint struct {
	Name        string                      `json:"Name,omitempty"`
	Summary     string                      `json:"Summary,omitempty"`
	Description string                      `json:"Description,omitempty"`
	Route       string                      `json:"Route,omitempty"`
	HttpMethod  string                      `json:"HttpMethod,omitempty"`
	QueryParams []OpenApiTemplateParam      `json:"QueryParams,omitempty"`
	PathParams  []OpenApiTemplateParam      `json:"PathParams,omitempty"`
	Headers     []OpenApiTemplateParam      `json:"Headers,omitempty"`
	Cookies     []OpenApiTemplateParam      `json:"Cookies,omitempty"`
	RequestBody *OpenApiTemplateRequestBody `json:"RequestBody,omitempty"`
	Responses   []OpenApiTemplateResponse   `json:"Responses,omitempty"`
}

// OpenApiTemplateParam represents a parameter in an endpoint.
type OpenApiTemplateParam struct {
	Name        string `json:"Name,omitempty"`
	Ref         string `json:"Ref,omitempty"`
	Description string `json:"Description,omitempty"`
	Required    bool   `json:"Required,omitempty"`
}

// OpenApiTemplateRequestBody represents the request body of an endpoint.
type OpenApiTemplateRequestBody struct {
	Required    bool                    `json:"Required,omitempty"`
	Description string                  `json:"Description,omitempty"`
	Content     *OpenApiTemplateContent `json:"Content,omitempty"`
	Ref         string                  `json:"Ref,omitempty"`
}

// OpenApiTemplateContent represents the content of request bodies and responses.
type OpenApiTemplateContent struct {
	MediaTypes []OpenAPITemplateMediaType `json:"MediaTypes,omitempty"`
}

// OpenAPITemplateMediaType represents a media type in content.
type OpenAPITemplateMediaType struct {
	Type string `json:"Type,omitempty"`
	Ref  string `json:"Ref,omitempty"`
}

// OpenApiTemplateResponse represents a response from an endpoint.
type OpenApiTemplateResponse struct {
	Description string                  `json:"Description,omitempty"`
	StatusCode  int                     `json:"StatusCode,omitempty"`
	Content     *OpenApiTemplateContent `json:"Content,omitempty"`
	Ref         string                  `json:"Ref,omitempty"`
}
type OpenApiTemplateComponent struct {
	Name            string                             `json:"Name,omitempty"`
	Description     string                             `json:"Description,omitempty"`
	Type            string                             `json:"Type,omitempty"`
	Properties      []OpenAPITemplateComponentProperty `json:"Properties,omitempty"`
	Responses       map[string]interface{}             `json:"Responses,omitempty"`
	Examples        map[string]interface{}             `json:"Examples,omitempty"`
	Parameters      map[string]interface{}             `json:"Parameters,omitempty"`
	RequestBodies   map[string]interface{}             `json:"RequestBodies,omitempty"`
	Headers         map[string]interface{}             `json:"Headers,omitempty"`
	SecuritySchemes map[string]interface{}             `json:"SecuritySchemes,omitempty"`
	Links           map[string]interface{}             `json:"Links,omitempty"`
	Callbacks       map[string]interface{}             `json:"Callbacks,omitempty"`
}

type OpenAPITemplateComponentProperty struct {
	Name             string                                 `json:"Name,omitempty"`
	Type             string                                 `json:"Type,omitempty"`
	Ref              string                                 `json:"Ref,omitempty"`
	Description      string                                 `json:"Description,omitempty"`
	Required         bool                                   `json:"Required,omitempty"`
	Nullable         bool                                   `json:"Nullable,omitempty"`
	Format           string                                 `json:"Format,omitempty"`
	Default          string                                 `json:"Default,omitempty"`
	Enum             []string                               `json:"Enum,omitempty"`
	ReadOnly         bool                                   `json:"ReadOnly,omitempty"`
	WriteOnly        bool                                   `json:"WriteOnly,omitempty"`
	Deprecated       bool                                   `json:"Deprecated,omitempty"`
	Items            *OpenAPITemplateComponentProperty      `json:"Items,omitempty"`
	MinLength        int                                    `json:"MinLength,omitempty"`
	MaxLength        int                                    `json:"MaxLength,omitempty"`
	Pattern          string                                 `json:"Pattern,omitempty"`
	Minimum          float64                                `json:"Minimum,omitempty"`
	Maximum          float64                                `json:"Maximum,omitempty"`
	ExclusiveMinimum float64                                `json:"ExclusiveMinimum,omitempty"`
	ExclusiveMaximum float64                                `json:"ExclusiveMaximum,omitempty"`
	UniqueItems      bool                                   `json:"UniqueItems,omitempty"`
	AllOf            []string                               `json:"AllOf,omitempty"`
	OneOf            []string                               `json:"OneOf,omitempty"`
	AnyOf            []string                               `json:"AnyOf,omitempty"`
	Not              []string                               `json:"Not,omitempty"`
	Discriminator    *OpenAPITemplateComponentDiscriminator `json:"Discriminator,omitempty"`
}

type OpenAPITemplateComponentDiscriminator struct {
	PropertyName string            `json:"PropertyName,omitempty"`
	Mapping      map[string]string `json:"Mapping,omitempty"`
}
