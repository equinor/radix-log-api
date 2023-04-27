// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/applications/{appName}/environments/{envName}/components/{componentName}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Inventory"
                ],
                "summary": "Get inventory (pods and their containers) for a component",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Application Name",
                        "name": "appName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Environment Name",
                        "name": "envName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Component Name",
                        "name": "componentName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "403": {
                        "description": "Forbidden"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/applications/{appName}/environments/{envName}/components/{componentName}/logs": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Logs"
                ],
                "summary": "Get log for a component",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Application Name",
                        "name": "appName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Environment Name",
                        "name": "envName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Component Name",
                        "name": "componentName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "403": {
                        "description": "Forbidden"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "description": "Bearer is currently not supported by go-swag. Use \"Bearer \u003cJWT\u003e\" in value.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{"http", "https"},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
