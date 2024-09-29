package openapi

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/arthurbcp/kuma-cli/internal/helpers"
)

// ParseToOpenAPITemplate parses the OpenAPI file represented as a map[string]interface{}
// into an OpenAPITemplate struct, including components and handling $ref references.
func ParseToOpenAPITemplate(helpers helpers.HelpersInterface, openAPIFile map[string]interface{}) OpenAPITemplate {
	template := OpenAPITemplate{}

	// Parse OpenAPI version
	if version, ok := openAPIFile["openapi"].(string); ok {
		template.Version = version
	}

	// Parse Info
	if info, ok := openAPIFile["info"].(map[string]interface{}); ok {
		if title, ok := info["title"].(string); ok {
			template.InfoTitle = title
		}
		if description, ok := info["description"].(string); ok {
			template.InfoDescription = description
		}
		if infoVersion, ok := info["version"].(string); ok {
			template.InfoVersion = infoVersion
		}
	}

	// Parse Servers
	if servers, ok := openAPIFile["servers"].([]interface{}); ok {
		for _, server := range servers {
			if serverMap, ok := server.(map[string]interface{}); ok {
				openAPIServer := OpenApiTemplateServer{}
				if url, ok := serverMap["url"].(string); ok {
					openAPIServer.Url = url
				}
				if description, ok := serverMap["description"].(string); ok {
					openAPIServer.Description = description
				}
				if variables, ok := serverMap["variables"].(map[string]interface{}); ok {
					for varName, varValue := range variables {
						if varMap, ok := varValue.(map[string]interface{}); ok {
							openApiVar := OpenApiTemplateVariable{
								Name: varName,
							}
							if desc, ok := varMap["description"].(string); ok {
								openApiVar.Description = desc
							}
							if def, ok := varMap["default"].(string); ok {
								openApiVar.Default = def
							}
							if enum, ok := varMap["enum"].([]interface{}); ok {
								for _, enumVal := range enum {
									if enumStr, ok := enumVal.(string); ok {
										openApiVar.Enum = append(openApiVar.Enum, enumStr)
									}
								}
							}
							openAPIServer.Variables = append(openAPIServer.Variables, openApiVar)
						}
					}
				}
				template.Servers = append(template.Servers, openAPIServer)
			}
		}
	}

	// Parse Components
	if components, ok := openAPIFile["components"].(map[string]interface{}); ok {
		// Parse Schemas
		if schemas, ok := components["schemas"].(map[string]interface{}); ok {
			for name, schema := range schemas {
				if schemaMap, ok := schema.(map[string]interface{}); ok {
					component := parseSchema(helpers, name, schemaMap)
					template.Components = append(template.Components, component)
				}
			}
		}
	}

	// Initialize a map to group controllers by tag names
	controllerMap := make(map[string]*OpenApiTemplateController)

	// Parse Paths
	if paths, ok := openAPIFile["paths"].(map[string]interface{}); ok {
		for path, pathItem := range paths {
			if pathMap, ok := pathItem.(map[string]interface{}); ok {
				// Iterate over HTTP methods
				for method, operation := range pathMap {
					// HTTP methods in OpenAPI are lowercase: get, post, put, delete, etc.
					if isHTTPMethod(method) {
						if opMap, ok := operation.(map[string]interface{}); ok {
							endpoint := OpenApiTemplateEndpoint{
								// Assign only the URL path to Route
								Route: path,
								// Assign the HTTP method to HttpMethod
								HttpMethod: strings.ToUpper(method),
							}

							// Name, Summary and Description
							if name, ok := opMap["operationId"].(string); ok {
								endpoint.Name = name
							}
							if summary, ok := opMap["summary"].(string); ok {
								endpoint.Summary = summary
							}
							if description, ok := opMap["description"].(string); ok {
								endpoint.Description = description
							}

							// Parameters
							if params, ok := opMap["parameters"].([]interface{}); ok {
								for _, param := range params {
									if paramMap, ok := param.(map[string]interface{}); ok {
										parsedParam := parseParameter(paramMap)
										switch parsedParam.In {
										case "query":
											endpoint.QueryParams = append(endpoint.QueryParams, parsedParam.Param)
										case "path":
											endpoint.PathParams = append(endpoint.PathParams, parsedParam.Param)
										case "header":
											endpoint.Headers = append(endpoint.Headers, parsedParam.Param)
										case "cookie":
											endpoint.Cookies = append(endpoint.Cookies, parsedParam.Param)
										}
									}
								}
							}

							// Request Body
							if reqBody, ok := opMap["requestBody"].(map[string]interface{}); ok {
								endpoint.RequestBody = parseRequestBody(reqBody)
							}

							// Responses
							if responses, ok := opMap["responses"].(map[string]interface{}); ok {
								for statusCode, resp := range responses {
									if respMap, ok := resp.(map[string]interface{}); ok {
										parsedResp := parseResponse(respMap, statusCode)
										endpoint.Responses = append(endpoint.Responses, parsedResp)
									}
								}
							}

							// Tags
							tags := []string{"default"} // Default tag if none provided
							if opTags, ok := opMap["tags"].([]interface{}); ok && len(opTags) > 0 {
								tags = []string{}
								for _, tag := range opTags {
									if tagStr, ok := tag.(string); ok {
										tags = append(tags, tagStr)
									}
								}
							}

							// Assign endpoint to controllers based on tags
							for _, tag := range tags {
								// Check if controller for this tag already exists
								controller, exists := controllerMap[tag]
								if !exists {
									// Create a new controller for the tag
									controller = &OpenApiTemplateController{
										Name:      tag,
										Endpoints: []OpenApiTemplateEndpoint{},
									}
									controllerMap[tag] = controller
								}
								// Append the endpoint to the controller's endpoints
								controller.Endpoints = append(controller.Endpoints, endpoint)
							}
						}
					}
				}
			}
		}
	}

	// Optionally, implement getBasePath logic if you intend to use BasePath
	for _, controller := range controllerMap {
		if len(controller.Endpoints) > 0 {
			path := controller.Endpoints[0].Route
			// Assign BasePath based on the first endpoint's path
			controller.BasePath = getBasePath(path)
		}
		template.Controllers = append(template.Controllers, *controller)
	}

	return template
}

// parseSchema parses a schema map into an OpenApiTemplateComponent struct.
func parseSchema(helpers helpers.HelpersInterface, name string, schemaMap map[string]interface{}) OpenApiTemplateComponent {
	component := OpenApiTemplateComponent{
		Name: name,
	}

	// Description
	if desc, ok := schemaMap["description"].(string); ok {
		component.Description = desc
	}

	// Type
	if typ, ok := schemaMap["type"].(string); ok {
		component.Type = typ
	}

	// Description
	var required []interface{}
	if req, ok := schemaMap["required"].([]interface{}); ok {
		required = req
	}

	// Parse Properties
	if properties, ok := schemaMap["properties"].(map[string]interface{}); ok {
		for propName, propValue := range properties {
			if propMap, ok := propValue.(map[string]interface{}); ok {
				property := parseProperty(helpers, propName, propMap, required)
				component.Properties = append(component.Properties, property)
			}
		}
	}

	return component
}

// parseProperty parses a property map into an OpenAPITemplateComponentProperty struct.
func parseProperty(helpers helpers.HelpersInterface, name string, propMap map[string]interface{}, required []interface{}) OpenAPITemplateComponentProperty {
	property := OpenAPITemplateComponentProperty{
		Name: name,
	}

	// Required
	if helpers.InterfaceContains(required, name) {
		property.Required = helpers.InterfaceContains(required, name)
	}

	// Description
	if desc, ok := propMap["description"].(string); ok {
		property.Description = desc
	}

	// Type
	if typ, ok := propMap["type"].(string); ok {
		property.Type = typ
	}

	if ref, ok := propMap["$ref"].(string); ok {
		property.Ref = getComponentName(ref)
	}

	// Format
	if format, ok := propMap["format"].(string); ok {
		property.Format = format
	}

	// Nullable
	if nullable, ok := propMap["nullable"].(bool); ok {
		property.Nullable = nullable
	}

	// ReadOnly
	if readOnly, ok := propMap["readOnly"].(bool); ok {
		property.ReadOnly = readOnly
	}

	// WriteOnly
	if writeOnly, ok := propMap["writeOnly"].(bool); ok {
		property.WriteOnly = writeOnly
	}

	// Deprecated
	if deprecated, ok := propMap["deprecated"].(bool); ok {
		property.Deprecated = deprecated
	}

	// Default
	if def, ok := propMap["default"]; ok {
		// Convert default value to string
		defBytes, err := json.Marshal(def)
		if err == nil {
			property.Default = string(defBytes)
		}
	}

	// Enum
	if enums, ok := propMap["enum"].([]interface{}); ok {
		for _, enumVal := range enums {
			if enumStr, ok := enumVal.(string); ok {
				property.Enum = append(property.Enum, enumStr)
			} else {
				// Handle non-string enums if necessary
				enumBytes, err := json.Marshal(enumVal)
				if err == nil {
					property.Enum = append(property.Enum, string(enumBytes))
				}
			}
		}
	}

	// Items (for arrays)
	if items, ok := propMap["items"].(map[string]interface{}); ok {
		itemProperty := OpenAPITemplateComponentProperty{}
		if itemType, ok := items["type"].(string); ok {
			itemProperty.Type = itemType
		}
		if itemFormat, ok := items["format"].(string); ok {
			itemProperty.Format = itemFormat
		}
		// Recursive parsing for nested items
		if nestedItems, ok := items["items"].(map[string]interface{}); ok {
			nestedProperty := parseProperty(helpers, name, nestedItems, required)
			itemProperty.Items = append(itemProperty.Items, nestedProperty)
		}
		property.Items = append(property.Items, itemProperty)
	}

	// Pattern
	if pattern, ok := propMap["pattern"].(string); ok {
		property.Pattern = pattern
	}

	// Minimum and Maximum
	if min, ok := propMap["minimum"].(float64); ok {
		property.Minimum = min
	}
	if max, ok := propMap["maximum"].(float64); ok {
		property.Maximum = max
	}
	if exMin, ok := propMap["exclusiveMinimum"].(float64); ok {
		property.ExclusiveMinimum = exMin
	}
	if exMax, ok := propMap["exclusiveMaximum"].(float64); ok {
		property.ExclusiveMaximum = exMax
	}

	// UniqueItems
	if unique, ok := propMap["uniqueItems"].(bool); ok {
		property.UniqueItems = unique
	}

	// AllOf, OneOf, AnyOf, Not
	if allOf, ok := propMap["allOf"].([]interface{}); ok {
		for _, item := range allOf {
			if ref, ok := item.(map[string]interface{})["$ref"].(string); ok {
				componentName := getComponentName(ref)
				if componentName != "" {
					property.AllOf = append(property.AllOf, componentName)
				}
			}
		}
	}
	if oneOf, ok := propMap["oneOf"].([]interface{}); ok {
		for _, item := range oneOf {
			if ref, ok := item.(map[string]interface{})["$ref"].(string); ok {
				componentName := getComponentName(ref)
				if componentName != "" {
					property.OneOf = append(property.OneOf, componentName)
				}
			}
		}
	}
	if anyOf, ok := propMap["anyOf"].([]interface{}); ok {
		for _, item := range anyOf {
			if ref, ok := item.(map[string]interface{})["$ref"].(string); ok {
				componentName := getComponentName(ref)
				if componentName != "" {
					property.AnyOf = append(property.AnyOf, componentName)
				}
			}
		}
	}
	if not, ok := propMap["not"].([]interface{}); ok {
		for _, item := range not {
			if ref, ok := item.(map[string]interface{})["$ref"].(string); ok {
				componentName := getComponentName(ref)
				if componentName != "" {
					property.Not = append(property.Not, componentName)
				}
			}
		}
	}

	// Discriminator
	if discriminator, ok := propMap["discriminator"].(map[string]interface{}); ok {
		discriminatorStruct := OpenAPITemplateComponentDiscriminator{}
		if propertyName, ok := discriminator["propertyName"].(string); ok {
			discriminatorStruct.PropertyName = propertyName
		}
		if mapping, ok := discriminator["mapping"].(map[string]interface{}); ok {
			discriminatorStruct.Mapping = make(map[string]string)
			for key, value := range mapping {
				if valStr, ok := value.(string); ok {
					discriminatorStruct.Mapping[key] = valStr
				}
			}
		}
		property.Discriminator = &discriminatorStruct
	}

	return property
}

// parseParameter parses a parameter map into a parameterWithIn struct.
// It handles $ref references to component schemas.
func parseParameter(paramMap map[string]interface{}) parameterWithIn {
	param := OpenApiTemplateParam{}

	// Check if the parameter uses a $ref
	if ref, ok := paramMap["$ref"].(string); ok {
		param.Ref = getComponentName(ref)
	} else {
		if name, ok := paramMap["name"].(string); ok {
			param.Name = name
		}
		if description, ok := paramMap["description"].(string); ok {
			param.Description = description
		}
		if required, ok := paramMap["required"].(bool); ok {
			param.Required = required
		}
	}

	inField := ""
	if inVal, ok := paramMap["in"].(string); ok {
		inField = inVal
	}

	return parameterWithIn{
		Param: param,
		In:    inField,
	}
}

// parseRequestBody parses a request body map into an OpenApiTemplateRequestBody struct.
// It handles $ref references to component schemas.
func parseRequestBody(reqBody map[string]interface{}) *OpenApiTemplateRequestBody {
	requestBody := OpenApiTemplateRequestBody{}

	if required, ok := reqBody["required"].(bool); ok {
		requestBody.Required = required
	}
	if description, ok := reqBody["description"].(string); ok {
		requestBody.Description = description
	}
	if content, ok := reqBody["content"].(map[string]interface{}); ok {
		parsedContent := OpenApiTemplateContent{}
		for mediaType, mediaObj := range content {
			if mediaMap, ok := mediaObj.(map[string]interface{}); ok {
				mediaTypeStruct := OpenAPITemplateMediaType{
					Type: mediaType,
				}
				if schema, ok := mediaMap["schema"].(map[string]interface{}); ok {
					if ref, ok := schema["$ref"].(string); ok {
						mediaTypeStruct.Ref = getComponentName(ref)
					}
				}
				parsedContent.MediaTypes = append(parsedContent.MediaTypes, mediaTypeStruct)
			}
		}
		requestBody.Content = &parsedContent
	}

	return &requestBody
}

// parseResponse parses a response map into an OpenApiTemplateResponse struct.
// It handles $ref references to component schemas.
func parseResponse(respMap map[string]interface{}, statusCode string) OpenApiTemplateResponse {
	response := OpenApiTemplateResponse{}

	if description, ok := respMap["description"].(string); ok {
		response.Description = description
	}
	// Convert statusCode to int
	if sc, err := parseStatusCode(statusCode); err == nil {
		response.StatusCode = sc
	} else {
		// Handle non-integer status codes like "default" by setting to 0 or a special value
		response.StatusCode = 0
	}
	if content, ok := respMap["content"].(map[string]interface{}); ok {
		parsedContent := OpenApiTemplateContent{}
		for mediaType, mediaObj := range content {
			if mediaMap, ok := mediaObj.(map[string]interface{}); ok {
				mediaTypeStruct := OpenAPITemplateMediaType{
					Type: mediaType,
				}
				if schema, ok := mediaMap["schema"].(map[string]interface{}); ok {
					if ref, ok := schema["$ref"].(string); ok {
						mediaTypeStruct.Ref = getComponentName(ref)
					}
				}
				parsedContent.MediaTypes = append(parsedContent.MediaTypes, mediaTypeStruct)
			}
		}
		response.Content = &parsedContent
	}

	return response
}

// isHTTPMethod checks if a given method string is a valid HTTP method.
func isHTTPMethod(method string) bool {
	httpMethods := []string{"get", "post", "put", "delete", "patch", "head", "options", "trace"}
	method = strings.ToLower(method)
	for _, m := range httpMethods {
		if method == m {
			return true
		}
	}
	return false
}

// parameterWithIn is a helper struct to hold a parameter and its "in" field.
type parameterWithIn struct {
	Param OpenApiTemplateParam
	In    string
}

// parseStatusCode converts a status code string to an integer.
// It returns an error if the status code is not a valid integer (e.g., "default").
func parseStatusCode(statusCode string) (int, error) {
	// Handle "default" and other non-integer status codes if necessary
	if statusCode == "default" {
		return 0, fmt.Errorf("default status code is not handled")
	}
	var sc int
	_, err := fmt.Sscanf(statusCode, "%d", &sc)
	return sc, err
}

// getComponentName extracts the component name from a $ref string.
// Example: "#/components/schemas/User" returns "User"
func getComponentName(ref string) string {
	return ref[strings.LastIndex(ref, "/")+1:]
}

// getBasePath extracts the base path from a full API path.
// For example, "/users/{id}/details" returns "/users"
func getBasePath(path string) string {
	segments := strings.Split(strings.Trim(path, "/"), "/")
	if len(segments) > 0 {
		return "/" + segments[0]
	}
	return "/"
}
