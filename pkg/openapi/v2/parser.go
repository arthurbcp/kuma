package openapi2

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ParseToOpenAPITemplate parses the OpenAPI 2.0 file represented as a map[string]interface{}
// into an OpenAPITemplate struct, including definitions and handling $ref references.
func ParseToOpenAPITemplate(openAPIFile map[string]interface{}) OpenAPITemplate {
	template := OpenAPITemplate{}

	// Parse Info
	if info, ok := openAPIFile["info"].(map[string]interface{}); ok {
		if title, ok := info["title"].(string); ok {
			template.Title = title
		}
		if description, ok := info["description"].(string); ok {
			template.Description = description
		}
		if version, ok := info["version"].(string); ok {
			template.Version = version
		}
	}

	// Parse Host
	if host, ok := openAPIFile["host"].(string); ok {
		template.Host = host
	}

	// Parse BasePath
	if basePath, ok := openAPIFile["basePath"].(string); ok {
		template.BasePath = basePath
	}

	// Parse Schemes
	if schemes, ok := openAPIFile["schemes"].([]interface{}); ok {
		for _, scheme := range schemes {
			if schemeStr, ok := scheme.(string); ok {
				template.Schemes = append(template.Schemes, schemeStr)
			}
		}
	}

	// Parse Definitions
	if definitions, ok := openAPIFile["definitions"].(map[string]interface{}); ok {
		for defName, defValue := range definitions {
			if defMap, ok := defValue.(map[string]interface{}); ok {
				schema := parseSchema(defName, defMap)
				template.Definitions = append(template.Definitions, schema)
			}
		}
	}

	// Initialize a map to group operations by tags
	groupMap := make(map[string]*OpenApiTemplateOperationGroup)

	// Parse Paths
	if paths, ok := openAPIFile["paths"].(map[string]interface{}); ok {
		for _, pathItem := range paths {
			if pathMap, ok := pathItem.(map[string]interface{}); ok {
				// Iterate over HTTP methods
				for method, operation := range pathMap {
					// HTTP methods in OpenAPI 2.0 are lowercase: get, post, put, delete, etc.
					if isHTTPMethod(method) {
						if opMap, ok := operation.(map[string]interface{}); ok {
							operationStruct := OpenApiTemplateOperation{
								HTTPMethod: strings.ToUpper(method),
							}

							// OperationId, Summary, Description
							if opID, ok := opMap["operationId"].(string); ok {
								operationStruct.Name = opID
							}
							if summary, ok := opMap["summary"].(string); ok {
								operationStruct.Summary = summary
							}
							if description, ok := opMap["description"].(string); ok {
								operationStruct.Description = description
							}

							// Consumes and Produces (operation level overrides global)
							if consumes, ok := opMap["consumes"].([]interface{}); ok {
								for _, consume := range consumes {
									if consumeStr, ok := consume.(string); ok {
										operationStruct.Consumes = append(operationStruct.Consumes, consumeStr)
									}
								}
							} else if globalConsumes, ok := openAPIFile["consumes"].([]interface{}); ok {
								for _, consume := range globalConsumes {
									if consumeStr, ok := consume.(string); ok {
										operationStruct.Consumes = append(operationStruct.Consumes, consumeStr)
									}
								}
							}

							if produces, ok := opMap["produces"].([]interface{}); ok {
								for _, produce := range produces {
									if produceStr, ok := produce.(string); ok {
										operationStruct.Produces = append(operationStruct.Produces, produceStr)
									}
								}
							} else if globalProduces, ok := openAPIFile["produces"].([]interface{}); ok {
								for _, produce := range globalProduces {
									if produceStr, ok := produce.(string); ok {
										operationStruct.Produces = append(operationStruct.Produces, produceStr)
									}
								}
							}

							// Parameters
							if params, ok := opMap["parameters"].([]interface{}); ok {
								for _, param := range params {
									if paramMap, ok := param.(map[string]interface{}); ok {
										parsedParam := parseParameter(paramMap)
										switch parsedParam.In {
										case "query":
											operationStruct.QueryParams = append(operationStruct.QueryParams, parsedParam.Schema)
										case "path":
											operationStruct.PathParams = append(operationStruct.PathParams, parsedParam.Schema)
										case "header":
											operationStruct.Headers = append(operationStruct.Headers, parsedParam.Header)
										case "body":
											operationStruct.Body = &parsedParam.Schema
											// case "formData":
											// 	// Handle formData if necessary
										}
									}
								}
							}

							// Responses
							if responses, ok := opMap["responses"].(map[string]interface{}); ok {
								for statusCode, resp := range responses {
									if respMap, ok := resp.(map[string]interface{}); ok {
										parsedResp := parseResponse(respMap, statusCode)
										operationStruct.Responses = append(operationStruct.Responses, parsedResp)
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

							// Assign operation to groups based on tags
							for _, tag := range tags {
								// Check if group for this tag already exists
								group, exists := groupMap[tag]
								if !exists {
									// Create a new group for the tag
									group = &OpenApiTemplateOperationGroup{
										Name:       tag,
										Operations: []OpenApiTemplateOperation{},
									}
									groupMap[tag] = group
								}
								// Append the operation to the group's operations
								group.Operations = append(group.Operations, operationStruct)
							}
						}
					}
				}
			}
		}
	}

	// Convert groupMap to slice
	for _, group := range groupMap {
		template.Groups = append(template.Groups, *group)
	}

	return template
}

// parseSchema parses a schema map into an OpenApiTemplateSchema struct.
func parseSchema(name string, schemaMap map[string]interface{}) OpenApiTemplateSchema {
	schema := OpenApiTemplateSchema{
		Name: name,
	}

	// Description
	if desc, ok := schemaMap["description"].(string); ok {
		schema.Description = desc
	}

	// Type
	if typ, ok := schemaMap["type"].(string); ok {
		schema.Type = typ
	}

	// Required
	if required, ok := schemaMap["required"].([]interface{}); ok {
		for _, req := range required {
			if reqStr, ok := req.(string); ok {
				if reqStr == name {
					schema.Required = true
					break
				}
			}
		}
	}

	// Default
	if def, ok := schemaMap["default"]; ok {
		// Convert default value to string
		defBytes, err := json.Marshal(def)
		if err == nil {
			schema.Default = string(defBytes)
		}
	}

	// Enum
	if enums, ok := schemaMap["enum"].([]interface{}); ok {
		for _, enumVal := range enums {
			if enumStr, ok := enumVal.(string); ok {
				schema.Enum = append(schema.Enum, enumStr)
			} else {
				// Handle non-string enums if necessary
				enumBytes, err := json.Marshal(enumVal)
				if err == nil {
					schema.Enum = append(schema.Enum, string(enumBytes))
				}
			}
		}
	}

	// $ref
	if ref, ok := schemaMap["$ref"].(string); ok {
		schema.Ref = getDefinitionName(ref)
	}

	// Format
	if format, ok := schemaMap["format"].(string); ok {
		schema.Format = format
	}

	// Minimum
	if min, ok := schemaMap["minimum"].(float64); ok {
		schema.Minimum = min
	}

	// Maximum
	if max, ok := schemaMap["maximum"].(float64); ok {
		schema.Maximum = max
	}

	// CollectionFormat
	if colFmt, ok := schemaMap["collectionFormat"].(string); ok {
		schema.CollectionFormat = colFmt
	}

	// Items (for arrays)
	if items, ok := schemaMap["items"].(map[string]interface{}); ok {
		itemSchema := parseSchema(name+"Item", items)
		schema.Items = &itemSchema
	}

	// MaxItems
	if maxItems, ok := schemaMap["maxItems"].(float64); ok {
		schema.MaxItems = int(maxItems)
	}

	// MinItems
	if minItems, ok := schemaMap["minItems"].(float64); ok {
		schema.MinItems = int(minItems)
	}

	// UniqueItems
	if unique, ok := schemaMap["uniqueItems"].(bool); ok {
		schema.UniqueItems = unique
	}

	// Properties (for objects)
	if properties, ok := schemaMap["properties"].(map[string]interface{}); ok {
		for propName, propValue := range properties {
			if propMap, ok := propValue.(map[string]interface{}); ok {
				propSchema := parseSchema(propName, propMap)
				schema.Properties = append(schema.Properties, propSchema)
			}
		}
	}

	return schema
}

// parseParameter parses a parameter map into a schema and header if applicable.
func parseParameter(paramMap map[string]interface{}) struct {
	Schema OpenApiTemplateSchema
	Header OpenApiTemplateHeader
	In     string
} {
	result := struct {
		Schema OpenApiTemplateSchema
		Header OpenApiTemplateHeader
		In     string
	}{}

	// Check if the parameter uses a $ref
	if ref, ok := paramMap["$ref"].(string); ok {
		result.Schema.Ref = getDefinitionName(ref)
		return result
	}

	// Parameter Location
	if in, ok := paramMap["in"].(string); ok {
		result.In = in
	}

	// Common fields
	if name, ok := paramMap["name"].(string); ok {
		result.Schema.Name = name
		result.Header.Name = name
	}
	if description, ok := paramMap["description"].(string); ok {
		result.Schema.Description = description
		result.Header.Description = description
	}
	if required, ok := paramMap["required"].(bool); ok {
		result.Schema.Required = required
	}
	if typ, ok := paramMap["type"].(string); ok {
		result.Schema.Type = typ
		result.Header.Type = typ
	}
	if format, ok := paramMap["format"].(string); ok {
		result.Schema.Format = format
	}
	if def, ok := paramMap["default"]; ok {
		defBytes, err := json.Marshal(def)
		if err == nil {
			result.Schema.Default = string(defBytes)
		}
	}
	if enums, ok := paramMap["enum"].([]interface{}); ok {
		for _, enumVal := range enums {
			if enumStr, ok := enumVal.(string); ok {
				result.Schema.Enum = append(result.Schema.Enum, enumStr)
			} else {
				// Handle non-string enums if necessary
				enumBytes, err := json.Marshal(enumVal)
				if err == nil {
					result.Schema.Enum = append(result.Schema.Enum, string(enumBytes))
				}
			}
		}
	}

	// Specific handling based on 'in' field
	switch result.In {
	case "header":
		// For headers, type is required
		if typ, ok := paramMap["type"].(string); ok {
			result.Header.Type = typ
		}
	case "body":
		// Body parameters have a schema
		if schema, ok := paramMap["schema"].(map[string]interface{}); ok {
			result.Schema = parseSchema(result.Schema.Name, schema)
		}
	}

	return result
}

// parseResponse parses a response map into an OpenApiTemplateResponse struct.
func parseResponse(respMap map[string]interface{}, statusCode string) OpenApiTemplateResponse {
	response := OpenApiTemplateResponse{}

	// Description
	if desc, ok := respMap["description"].(string); ok {
		response.Description = desc
	}

	// Schema or $ref
	if schema, ok := respMap["schema"].(map[string]interface{}); ok {
		if ref, ok := schema["$ref"].(string); ok {
			response.Ref = getDefinitionName(ref)
		} else {
			// Parse inline schema
			inlineSchema := parseSchema("", schema)
			response.Schema = inlineSchema
		}
	}

	// Headers
	if headers, ok := respMap["headers"].(map[string]interface{}); ok {
		for headerName, headerValue := range headers {
			if headerMap, ok := headerValue.(map[string]interface{}); ok {
				header := OpenApiTemplateHeader{
					Name: headerName,
				}
				if desc, ok := headerMap["description"].(string); ok {
					header.Description = desc
				}
				if typ, ok := headerMap["type"].(string); ok {
					header.Type = typ
				}
				response.Headers = append(response.Headers, header)
			}
		}
	}

	// Convert statusCode to int
	if sc, err := parseStatusCode(statusCode); err == nil {
		response.StatusCode = sc
	} else {
		// Handle non-integer status codes like "default" by setting to 0 or a special value
		response.StatusCode = 0
	}

	return response
}

// isHTTPMethod checks if a given method string is a valid HTTP method.
func isHTTPMethod(method string) bool {
	httpMethods := []string{"get", "post", "put", "delete", "patch", "head", "options"}
	method = strings.ToLower(method)
	for _, m := range httpMethods {
		if method == m {
			return true
		}
	}
	return false
}

// getDefinitionName extracts the definition name from a $ref string.
// Example: "#/definitions/User" returns "User"
func getDefinitionName(ref string) string {
	parts := strings.Split(ref, "/")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return ""
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
