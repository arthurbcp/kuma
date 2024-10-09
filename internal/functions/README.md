# Kuma Functions

This document provides an overview of the functions available in this package. These functions are designed to extract and manipulate data from OpenAPI 2.0 files and convert data into YAML format for use in Go templates.

## Functions

### Parser functions

#### ToYaml

ToYaml converts a Go data structure into a YAML-formatted slice of strings, where each element represents a line of the resulting YAML. This function is useful for rendering data in a human-readable format for templating purposes.

**Signature:**

```go
func ToYaml(data interface{}) []string
```

**Parameters:**

- `data`: The data structure to be converted into YAML.

**Returns:**

A slice of strings containing the YAML representation of data.

**Example**

```yaml
# input:
#    {
#         "items":  {
#             "subItems": [ "item1",   "item2",  "item3"],
#         },
#   }
data:
   {{ range toYaml . }}{{ . }}{{end}}
#output:
# data:
#   items:
#     subItems:
#         - item1
#         - item2
#         - item3
```

---

### GetRefFrom

GetRefFrom extracts the reference identifier from an OpenAPI 2.0 object if it exists. The function expects the reference to be in the format of a JSON pointer within the OpenAPI specification.

**Signature:**

```go
func GetRefFrom(object map[string]interface{}) string
```

**Parameters:**

- `object`: A map representing an OpenAPI object which might contain a $ref field.

**Returns:**

A string containing the reference identifier or an empty string if no valid reference is found.

**Example**

```yaml
# input:
# {
#     "schema": {
#         "$ref": "#/definitions/ApiResponse"
#     }
# }
"{{getRefFrom  .schema}}"
# output:
# "ApiResponse"
```

---

### GetPathsByTag

GetPathsByTag filters OpenAPI paths by a specified tag. It returns a subset of paths that are associated with the provided tag, useful for generating documentation for specific sections of an API.

**Signature:**

```go
func GetPathsByTag(paths map[string]interface{},
tag string) map[string]interface{}
```

**Parameters:**

- `paths`: A map where keys are path names and values are path item objects.
- `tag`: The tag used to filter paths.

**Returns:**

A map containing the paths that include the specified tag.

**Example**

```yaml
# input:
# paths: {
#   "/pet/{petId}/uploadImage": {
#       "post": {
#         "tags": ["pet"],
#         "summary": "uploads an image",
#       }
#       ...
#   },
#   "/user": {
#       "post": {
#         "tags": ["user"],
#         "summary": "create a user",
#       },
#   },
#   "/user/{userId}": {
#       "put": {
#         "tags": ["user"],
#         "summary": "update user by id",
#       },
#       "get": {
#         "tags": ["user"],
#         "summary": "get user by id",
#       }
#       ...
#  }
#}
"{{ getPathsByTag .paths "user" }}"
# output:
# {
#   "/user": {
#       "post": {
#         "tags": ["user"],
#         "summary": "create a user",
#       },
#   },
#   "/user/{userId}": {
#       "put": {
#         "tags": ["user"],
#         "summary": "update user by id",
#       },
#       "get": {
#         "tags": ["user"],
#         "summary": "get user by id",
#       }
#       ...
#  }
#}
```

---

### GetParamsByType

GetParamsByType filters parameters based on their `in` field type (e.g., query, header, path, formData). This function helps in extracting parameters of a certain type from an OpenAPI operation.

**Signature:**

```go
func GetParamsByType(params []interface{},
paramType string) []interface{}
```

**Parameters:**

- `params`: A slice of parameter objects.
- `paramType`: The type of parameters to filter by (e.g., "query").

**Returns:**

A slice of parameters that match the specified type.

**Example**

```yaml
# input:
#     {
#       "parameters": [
#             {
#                 "name": "petId",
#                 "in": "path",
#             },
#             {
#                 "name": "body",
#                 "in": "body",
#             },
#             {
#                 "name": "userId",
#                 "in": "path,
#             }
#     ],
# }
"{{ getParamsByType $data.parameters "path" }}"
# output:
# [
#     {
#         "name": "userId"
#     },
#      {
#        "name": "petId"
#      }
#  ]
```

---
