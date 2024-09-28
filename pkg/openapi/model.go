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
	BasePath  string                    `json:"BasePath,omitempty"`
	Endpoints []OpenApiTemplateEndpoint `json:"Endpoints,omitempty"`
}

// OpenApiTemplateEndpoint represents an API endpoint.
type OpenApiTemplateEndpoint struct {
	Name        string                     `json:"Name,omitempty"`
	Summary     string                     `json:"Summary,omitempty"`
	Description string                     `json:"Description,omitempty"`
	Route       string                     `json:"Route,omitempty"`
	HttpMethod  string                     `json:"HttpMethod,omitempty"`
	QueryParams []OpenApiTemplateParam     `json:"QueryParams,omitempty"`
	PathParams  []OpenApiTemplateParam     `json:"PathParams,omitempty"`
	Headers     []OpenApiTemplateParam     `json:"Headers,omitempty"`
	Cookies     []OpenApiTemplateParam     `json:"Cookies,omitempty"`
	RequestBody OpenApiTemplateRequestBody `json:"RequestBody,omitempty"`
	Responses   []OpenApiTemplateResponse  `json:"Responses,omitempty"`
}

// OpenApiTemplateParam represents a parameter in an endpoint.
type OpenApiTemplateParam struct {
	Name        string                 `json:"Name,omitempty"`
	Description string                 `json:"Description,omitempty"`
	Required    bool                   `json:"Required,omitempty"`
	Schema      map[string]interface{} `json:"Schema,omitempty"`
}

// OpenApiTemplateRequestBody represents the request body of an endpoint.
type OpenApiTemplateRequestBody struct {
	Required    bool                   `json:"Required,omitempty"`
	Description string                 `json:"Description,omitempty"`
	Content     OpenApiTemplateContent `json:"Content,omitempty"`
}

// OpenApiTemplateContent represents the content of request bodies and responses.
type OpenApiTemplateContent struct {
	MediaTypes []OpenAPITemplateMediaType `json:"MediaTypes,omitempty"`
}

// OpenAPITemplateMediaType represents a media type in content.
type OpenAPITemplateMediaType struct {
	Type   string                 `json:"Type,omitempty"`
	Schema map[string]interface{} `json:"Schema,omitempty"`
}

// OpenApiTemplateResponse represents a response from an endpoint.
type OpenApiTemplateResponse struct {
	Description string                 `json:"Description,omitempty"`
	StatusCode  int                    `json:"StatusCode,omitempty"`
	Content     OpenApiTemplateContent `json:"Content,omitempty"`
}

// OpenApiTemplateComponent represents a component schema in the OpenAPI specification.
type OpenApiTemplateComponent struct {
	Name        string                 `json:"Name,omitempty"`
	Description string                 `json:"Description,omitempty"`
	Schema      map[string]interface{} `json:"Schema,omitempty"`
}
