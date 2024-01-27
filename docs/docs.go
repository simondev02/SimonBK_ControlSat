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
        "/ControlSat/Results": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Obtine ControlSat",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Controlsat"
                ],
                "summary": "Obtiene ControlSat",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/swagger.GormModelStub"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/swagger.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/Finandina": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Obtine ControlSat",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Controlsat"
                ],
                "summary": "Obtiene ControlSat",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/swagger.GormModelStub"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/swagger.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/getAll": {
            "get": {
                "responses": {}
            }
        }
    },
    "definitions": {
        "swagger.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 500
                },
                "message": {
                    "type": "string",
                    "example": "Error message"
                }
            }
        },
        "swagger.GormModelStub": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "2023-10-05T15:04:05Z"
                },
                "deleted_at": {
                    "type": "string",
                    "example": "2023-10-05T15:04:05Z"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "updated_at": {
                    "type": "string",
                    "example": "2023-10-05T15:04:05Z"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
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
	BasePath:         "",
	Schemes:          []string{},
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
