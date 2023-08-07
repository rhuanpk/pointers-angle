// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Rhuan Patriky",
            "url": "https://linktr.ee/rhuanpk",
            "email": "support@rhuanpk.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/clock/{hour}": {
            "get": {
                "description": "Calculate the angle between the clock pointers.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "clock"
                ],
                "summary": "ANGLE",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "The hour.",
                        "name": "hour",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "The minute.",
                        "name": "minute",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Angle JSON response.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.HTTP"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.HTTP"
                        }
                    }
                }
            }
        },
        "/clock/{hour}/{minute}": {
            "get": {
                "description": "Calculate the angle between the clock pointers.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "clock"
                ],
                "summary": "ANGLE",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "The hour.",
                        "name": "hour",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "The minute.",
                        "name": "minute",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Angle JSON response.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.HTTP"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.HTTP"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.HTTP": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "error": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.0",
	Host:             "localhost:8080",
	BasePath:         "/v0/rest",
	Schemes:          []string{},
	Title:            "Pointer Angles API",
	Description:      "API for calculate the less angle between hour and minute pointer.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
