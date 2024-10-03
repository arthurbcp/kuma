package openapi2

// OpenAPITemplate represents the overall structure of the parsed OpenAPI 2.0 specification.
type OpenAPITemplate struct {
	// Info title
	Title string `json:"InfoTitle,omitempty"`
	// Info description
	Description string `json:"InfoDescription,omitempty"`
	// Info version
	Version  string   `json:"InfoVersion,omitempty"`
	Host     string   `json:"Host,omitempty"`
	BasePath string   `json:"BasePath,omitempty"`
	Schemes  []string `json:"Schemes,omitempty"`

	// paths grouped by tag
	Groups []OpenApiTemplateOperationGroup `json:"Groups,omitempty"`

	// definitions
	Definitions []OpenApiTemplateSchema `json:"Structs,omitempty"`
}

// OpenApiTemplateOperationGroup groups operations under a controller (based on tags).
type OpenApiTemplateOperationGroup struct {
	// tag
	Name       string                     `json:"Name,omitempty"`
	Operations []OpenApiTemplateOperation `json:"Operations,omitempty"`
}

// OpenApiTemplateOperation represents an API operation (e.g., GET, POST).
type OpenApiTemplateOperation struct {
	HTTPMethod  string   `json:"HttpMethod,omitempty"`
	Name        string   `json:"Name,omitempty"`
	Summary     string   `json:"Summary,omitempty"`
	Description string   `json:"Description,omitempty"`
	Route       string   `json:"Route,omitempty"`
	Consumes    []string `json:"Consumes,omitempty"`
	Produces    []string `json:"Produces,omitempty"`

	// Params are grouped by type (query, path, body, header)
	Headers     []OpenApiTemplateHeader `json:"Headers,omitempty"`
	QueryParams []OpenApiTemplateSchema `json:"QueryParams,omitempty"`
	PathParams  []OpenApiTemplateSchema `json:"PathParams,omitempty"`
	Body        *OpenApiTemplateSchema  `json:"BodyParams,omitempty"`

	Responses []OpenApiTemplateResponse `json:"Responses,omitempty"`

	// TODO: Add security support
}

// OpenApiTemplateSchema represents a parameter in an operation.
type OpenApiTemplateSchema struct {
	Name        string   `json:"Name,omitempty"`
	Description string   `json:"Description,omitempty"`
	Required    bool     `json:"Required,omitempty"`
	Default     string   `json:"Default,omitempty"`
	Type        string   `json:"Type,omitempty"`
	Format      string   `json:"Format,omitempty"`
	Enum        []string `json:"Enum,omitempty"`
	// Ex: $ref: #/definitions/User to User
	Ref              string                  `json:"Ref,omitempty"`
	Minimum          float64                 `json:"Minimum,omitempty"`
	Maximum          float64                 `json:"Maximum,omitempty"`
	CollectionFormat string                  `json:"CollectionFormat,omitempty"`
	Items            *OpenApiTemplateSchema  `json:"Items,omitempty"`
	MaxItems         int                     `json:"MaxItems,omitempty"`
	MinItems         int                     `json:"MinItems,omitempty"`
	UniqueItems      bool                    `json:"UniqueItems,omitempty"`
	Properties       []OpenApiTemplateSchema `json:"Properties,omitempty"`
}

// OpenApiTemplateResponse represents a response from an operation.
type OpenApiTemplateResponse struct {
	// Response status code
	StatusCode  int                     `json:"StatusCode,omitempty"`
	Description string                  `json:"Description,omitempty"`
	Schema      OpenApiTemplateSchema   `json:"Schema,omitempty"`
	Ref         string                  `json:"Ref,omitempty"`
	Headers     []OpenApiTemplateHeader `json:"Headers,omitempty"`
}

type OpenApiTemplateHeader struct {
	Name        string `json:"Name,omitempty"`
	Description string `json:"Description,omitempty"`
	// Header key
	Type string `json:"Type,omitempty"`
}
