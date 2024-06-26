// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/api/v1/convert": {
            "get": {
                "description": "get conversion",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "convert"
                ],
                "summary": "get conversion",
                "parameters": [
                    {
                        "maxLength": 50,
                        "minLength": 2,
                        "type": "string",
                        "description": "symbol base",
                        "name": "from",
                        "in": "query"
                    },
                    {
                        "maxLength": 50,
                        "minLength": 2,
                        "type": "string",
                        "description": "symbol quote",
                        "name": "to",
                        "in": "query"
                    },
                    {
                        "type": "number",
                        "description": "amount",
                        "name": "amount",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Conversion"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Conversion": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "base": {
                    "type": "string"
                },
                "quote": {
                    "type": "string"
                },
                "rate": {
                    "type": "number"
                },
                "value": {
                    "type": "number"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Calculator",
	Description:      "Calculator.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
