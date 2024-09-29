This README provides a comprehensive overview of how the OpenAPI specification properties are mapped to the corresponding fields in the parsed JSON structure. This mapping facilitates the transformation of OpenAPI definitions into a structured format suitable for application use.

## Table of Contents

- [Introduction](#introduction)
- [Overall Structure](#overall-structure)
- [Detailed Mapping](#detailed-mapping)
  - [Version and Information](#version-and-information)
  - [Servers](#servers)
  - [Controllers](#controllers)
    - [Controller Properties](#controller-properties)
    - [Endpoints](#endpoints)
      - [Endpoint Properties](#endpoint-properties)
      - [Parameters](#parameters)
      - [Request Body](#request-body)
      - [Responses](#responses)
  - [Components](#components)
    - [Component Properties](#component-properties)
- [Parsed Data Example](#parsed-data-example)
- [Conclusion](#conclusion)

## Introduction

The OpenAPI specification provides a standardized format to describe RESTful APIs. To leverage this specification within applications, the OpenAPI definitions are parsed into a structured JSON format. This mapping ensures that all relevant information from the OpenAPI file is accurately represented and accessible within the application.

## Overall Structure

The parsed OpenAPI data is encapsulated within a structured JSON object that organizes the information into several key components:

- **Version and Information:** General metadata about the API.
- **Servers:** Details about the API servers.
- **Controllers:** Groupings of related API endpoints.
- **Components:** Reusable schemas and security schemes.

## Detailed Mapping

### Version and Information

**OpenAPI Property:**

```json
{
  "openapi": "3.0.0",
  "info": {
    "title": "GitHub Repositories API",
    "description": "API for managing repositories on GitHub",
    "version": "1.0.0"
  }
}
```

**Mapped to Parsed JSON:**

```json
{
  "Version": "3.0.0",
  "InfoTitle": "GitHub Repositories API",
  "InfoDescription": "API for managing repositories on GitHub",
  "InfoVersion": "1.0.0"
}
```

- **\`openapi\`** ➔ `Version`
- **\`info.title\`** ➔ `InfoTitle`
- **\`info.description\`** ➔ `InfoDescription`
- **\`info.version\`** ➔ `InfoVersion`

### Servers

**OpenAPI Property:**

```json
{
  "servers": [
    {
      "url": "https://api.github.com",
      "description": "Public GitHub API"
    }
  ]
}
```

**Mapped to Parsed JSON:**

```json
{
  "Servers": [
    {
      "Url": "https://api.github.com",
      "Description": "Public GitHub API"
    }
  ]
}
```

- **\`servers\`** ➔ `Servers`
  - **\`url\`** ➔ `Url`
  - **\`description\`** ➔ `Description`

### Controllers

Controllers group related endpoints based on tags in the OpenAPI specification.

**OpenAPI Property:**

```json
"tags": ["Repositories"]
```

**Mapped to Parsed JSON:**

```json
{
  "Controllers": [
    {
      "Name": "Repositories",
      "BasePath": "/user",
      "Endpoints": [
        /* Endpoints listed here */
      ]
    }
  ]
}
```

- **\`tags\`** ➔ `Controllers.Name`
- **`BasePath`** is derived from the common path prefix (e.g., `/user` for `/user/repos`)
- **Endpoints** within each controller are grouped accordingly.

#### Controller Properties

- **Name:** Derived from the tag name (e.g., "Repositories").
- **BasePath:** Common path segment shared by endpoints under the controller.

### Endpoints

Endpoints represent individual API operations (e.g., GET, POST).

**OpenAPI Property:**

Each path and its operations (GET, POST, etc.)

**Mapped to Parsed JSON:**

```json
{
"Endpoints": [
{
"Summary": "List repositories of the authenticated user",
"Route": "/user/repos",
"HttpMethod": "GET",
"QueryParams": [ /* Query Parameters */ ],
"Responses": [ /* Responses */ ]
},
{
"Summary": "Create a new repository",
"Route": "/user/repos",
"HttpMethod": "POST",
"RequestBody": { /_ Request Body _/ },
"Responses": [ /* Responses */ ]
}
// ... other endpoints
]
}
```

#### Endpoint Properties

- **Summary:** Mapped from `summary`.
- **Description:** Mapped from `description` (if available).
- **Route:** Mapped from the path (e.g., `/user/repos`).
- **HttpMethod:** Mapped from the HTTP method (e.g., `GET`, `POST`).

#### Parameters

Parameters can be categorized into Query, Path, Header, and Cookie parameters.

**OpenAPI Property:**

```json
"parameters": [
{
"in": "query",
"name": "visibility",
"schema": { "type": "string", "enum": ["all", "public", "private"] },
"description": "Filter repositories by visibility"
},
{
"in": "query",
"name": "affiliation",
"schema": { "type": "string", "enum": ["owner", "collaborator", "organization_member"] },
"description": "Filter repositories by affiliation"
}
]
```

**Mapped to Parsed JSON:**

```json
{
  "QueryParams": [
    {
      "Name": "visibility",
      "Description": "Filter repositories by visibility",
      "Schema": { "type": "string" }
    },
    {
      "Name": "affiliation",
      "Description": "Filter repositories by affiliation",
      "Schema": { "type": "string" }
    }
  ]
}
```

- **\`in\`** determines the parameter category (`QueryParams`, `PathParams`, etc.).
- **\`name\`**, **\`description\`**, and **\`schema\`** are directly mapped.
- **\`required\`** status is also captured if specified.

#### Request Body

**OpenAPI Property:**

```json
"requestBody": {
"required": true,
"content": {
"application/json": {
"schema": { "$ref": "#/components/schemas/NewRepository" }
}
}
}
```

**Mapped to Parsed JSON:**

```json
{
  "RequestBody": {
    "Required": true,
    "Content": {
      "MediaTypes": [
        {
          "Type": "application/json",
          "Schema": { "$ref": "NewRepository" }
        }
      ]
    }
  }
}
```

- **\`required\`** ➔ `RequestBody.Required`
- **\`content\`** ➔ `RequestBody.Content`
  - **Media Types** such as `application/json` are captured with their respective schemas.

#### Responses

**OpenAPI Property:**

```json
"responses": {
"200": {
"description": "List of repositories",
"content": {
"application/json": {
"schema": { "$ref": "#/components/schemas/Repository" }
      }
    }
  },
  "404": {
    "description": "Repository not found",
    "content": {
      "application/json": {
        "schema": { "$ref": "#/components/schemas/Error" }
}
}
}
}
```

**Mapped to Parsed JSON:**

```json
{
  "Responses": [
    {
      "Description": "List of repositories",
      "StatusCode": 200,
      "Content": {
        "MediaTypes": [
          {
            "Type": "application/json",
            "Schema": { "$ref": "Repository" }
          }
        ]
      }
    },
    {
      "Description": "Repository not found",
      "StatusCode": 404,
      "Content": {
        "MediaTypes": [
          {
            "Type": "application/json",
            "Schema": { "$ref": "Error" }
          }
        ]
      }
    }
  ]
}
```

- **Status Codes** (e.g., `200`, `404`) are mapped to `StatusCode`.
- **Description** and **Content** are directly mapped.

### Components

Components include reusable schemas, security schemes, and other definitions.

**OpenAPI Property:**

```json
"components": {
"schemas": {
"Repository": { /_ schema definition _/ },
"NewRepository": { /_ schema definition _/ },
"UpdateRepository": { /_ schema definition _/ },
"User": { /_ schema definition _/ },
"Error": { /_ schema definition _/ }
},
"securitySchemes": {
"bearerAuth": { /_ security scheme definition _/ }
}
}
```

**Mapped to Parsed JSON:**

```json
{
  "Components": [
    {
      "Name": "Repository",
      "Schema": {
        "type": "object",
        "properties": {
          "id": { "type": "integer" },
          "name": { "type": "string" },
          "full_name": { "type": "string" },
          "private": { "type": "boolean" },
          "owner": { "$ref": "User" },
          "html_url": { "type": "string", "format": "uri" },
          "description": { "type": "string" },
          "fork": { "type": "boolean" },
          "url": { "type": "string", "format": "uri" },
          "created_at": { "type": "string", "format": "date-time" },
          "updated_at": { "type": "string", "format": "date-time" },
          "pushed_at": { "type": "string", "format": "date-time" }
        }
      }
    },
    {
      "Name": "NewRepository",
      "Schema": {
        "type": "object",
        "required": ["name"],
        "properties": {
          "name": {
            "description": "Name of the new repository",
            "type": "string"
          },
          "description": {
            "description": "Repository description",
            "type": "string"
          },
          "private": {
            "description": "Specifies if the repository will be private",
            "type": "boolean"
          }
        }
      }
    }
    // ... other components
  ]
}
```

- **Schemas** and **Security Schemes** are treated as components.
- Each component includes its **Name**, **Description** (if available), and **Schema**.

#### Component Properties

- **Name:** Derived from the schema or security scheme name (e.g., "Repository", "User").
- **Description:** Mapped from the component's description (if available).
- **Schema:** The detailed schema definition, including properties, types, and references.

## Parsed Data Example

Here is an example of the parsed JSON data based on the provided OpenAPI specification:

```json
{
  "Version": "3.0.0",
  "InfoTitle": "GitHub Repositories API",
  "InfoDescription": "API for managing repositories on GitHub",
  "InfoVersion": "1.0.0",
  "Servers": [
    {
      "Url": "https://api.github.com",
      "Description": "Public GitHub API"
    }
  ],
  "Controllers": [
    {
      "Name": "Repositories",
      "BasePath": "/user",
      "Endpoints": [
        {
          "Summary": "List repositories of the authenticated user",
          "Route": "/user/repos",
          "HttpMethod": "GET",
          "QueryParams": [
            {
              "Name": "visibility",
              "Description": "Filter repositories by visibility",
              "Schema": { "type": "string" }
            },
            {
              "Name": "affiliation",
              "Description": "Filter repositories by affiliation",
              "Schema": { "type": "string" }
            }
          ],
          "Responses": [
            {
              "Description": "List of repositories",
              "StatusCode": 200,
              "Content": {
                "MediaTypes": [
                  {
                    "Type": "application/json",
                    "Schema": { "type": "array" }
                  }
                ]
              }
            }
          ]
        },
        {
          "Summary": "Create a new repository",
          "Route": "/user/repos",
          "HttpMethod": "POST",
          "RequestBody": {
            "Required": true,
            "Content": {
              "MediaTypes": [
                {
                  "Type": "application/json",
                  "Schema": { "$ref": "NewRepository" }
                }
              ]
            }
          },
          "Responses": [
            {
              "Description": "Repository created successfully",
              "StatusCode": 201,
              "Content": {
                "MediaTypes": [
                  {
                    "Type": "application/json",
                    "Schema": { "$ref": "Repository" }
                  }
                ]
              }
            },
            {
              "Description": "Validation error",
              "StatusCode": 422,
              "Content": {
                "MediaTypes": [
                  {
                    "Type": "application/json",
                    "Schema": { "$ref": "Error" }
                  }
                ]
              }
            }
          ]
        },
        {
          "Summary": "Get details of a specific repository",
          "Route": "/repos/{owner}/{repo}",
          "HttpMethod": "GET",
          "PathParams": [
            {
              "Name": "owner",
              "Description": "Username or organization name of the repository owner",
              "Required": true,
              "Schema": { "type": "string" }
            },
            {
              "Name": "repo",
              "Description": "Repository name",
              "Required": true,
              "Schema": { "type": "string" }
            }
          ],
          "Responses": [
            {
              "Description": "Repository details",
              "StatusCode": 200,
              "Content": {
                "MediaTypes": [
                  {
                    "Type": "application/json",
                    "Schema": { "$ref": "Repository" }
                  }
                ]
              }
            },
            {
              "Description": "Repository not found",
              "StatusCode": 404,
              "Content": {
                "MediaTypes": [
                  {
                    "Type": "application/json",
                    "Schema": { "$ref": "Error" }
                  }
                ]
              }
            }
          ]
        },
        {
          "Summary": "Update an existing repository",
          "Route": "/repos/{owner}/{repo}",
          "HttpMethod": "PATCH",
          "PathParams": [
            {
              "Name": "owner",
              "Description": "Username or organization name of the repository owner",
              "Required": true,
              "Schema": { "type": "string" }
            },
            {
              "Name": "repo",
              "Description": "Repository name",
              "Required": true,
              "Schema": { "type": "string" }
            }
          ],
          "RequestBody": {
            "Required": true,
            "Content": {
              "MediaTypes": [
                {
                  "Type": "application/json",
                  "Schema": { "$ref": "UpdateRepository" }
                }
              ]
            }
          },
          "Responses": [
            {
              "Description": "Repository updated successfully",
              "StatusCode": 200,
              "Content": {
                "MediaTypes": [
                  {
                    "Type": "application/json",
                    "Schema": { "$ref": "Repository" }
                  }
                ]
              }
            },
            {
              "Description": "Repository not found",
              "StatusCode": 404,
              "Content": {
                "MediaTypes": [
                  {
                    "Type": "application/json",
                    "Schema": { "$ref": "Error" }
                  }
                ]
              }
            }
          ]
        },
        {
          "Summary": "Delete a repository",
          "Route": "/repos/{owner}/{repo}",
          "HttpMethod": "DELETE",
          "PathParams": [
            {
              "Name": "owner",
              "Description": "Username or organization name of the repository owner",
              "Required": true,
              "Schema": { "type": "string" }
            },
            {
              "Name": "repo",
              "Description": "Repository name",
              "Required": true,
              "Schema": { "type": "string" }
            }
          ],
          "Responses": [
            {
              "Description": "Repository not found",
              "StatusCode": 404,
              "Content": {
                "MediaTypes": [
                  {
                    "Type": "application/json",
                    "Schema": { "$ref": "Error" }
                  }
                ]
              }
            },
            {
              "Description": "Repository deleted successfully",
              "StatusCode": 204,
              "Content": {}
            }
          ]
        }
      ]
    }
  ],
  "Components": [
    {
      "Name": "Repository",
      "Schema": {
        "type": "object",
        "properties": {
          "id": { "type": "integer" },
          "name": { "type": "string" },
          "full_name": { "type": "string" },
          "private": { "type": "boolean" },
          "owner": { "$ref": "User" },
          "html_url": { "type": "string", "format": "uri" },
          "description": { "type": "string" },
          "fork": { "type": "boolean" },
          "url": { "type": "string", "format": "uri" },
          "created_at": { "type": "string", "format": "date-time" },
          "updated_at": { "type": "string", "format": "date-time" },
          "pushed_at": { "type": "string", "format": "date-time" }
        }
      }
    },
    {
      "Name": "NewRepository",
      "Schema": {
        "type": "object",
        "required": ["name"],
        "properties": {
          "name": {
            "description": "Name of the new repository",
            "type": "string"
          },
          "description": {
            "description": "Repository description",
            "type": "string"
          },
          "private": {
            "description": "Specifies if the repository will be private",
            "type": "boolean"
          }
        }
      }
    },
    {
      "Name": "UpdateRepository",
      "Schema": {
        "type": "object",
        "properties": {
          "name": {
            "description": "New repository name",
            "type": "string"
          },
          "description": {
            "description": "New repository description",
            "type": "string"
          },
          "private": {
            "description": "Specifies if the repository will be private",
            "type": "boolean"
          }
        }
      }
    },
    {
      "Name": "User",
      "Schema": {
        "type": "object",
        "properties": {
          "login": { "type": "string" },
          "id": { "type": "integer" },
          "avatar_url": { "type": "string", "format": "uri" },
          "html_url": { "type": "string", "format": "uri" }
        }
      }
    },
    {
      "Name": "Error",
      "Schema": {
        "type": "object",
        "properties": {
          "message": { "type": "string" },
          "documentation_url": { "type": "string", "format": "uri" }
        }
      }
    }
  ]
}
```

This JSON structure aligns with the parsed data format, ensuring all OpenAPI properties are appropriately captured.

## Conclusion

This mapping ensures that all elements of the OpenAPI specification are accurately represented within the application, facilitating seamless interaction with the defined API endpoints. By adhering to this structure, developers can efficiently parse and utilize OpenAPI definitions, enabling robust API management and integration.
