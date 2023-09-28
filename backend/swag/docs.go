// Package swag Code generated by swaggo/swag. DO NOT EDIT
package swag

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
        "/v1/lgtms": {
            "get": {
                "parameters": [
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.LGTM"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/service.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.LGTM": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                }
            }
        },
        "service.ErrCode": {
            "type": "string",
            "enum": [
                "BAD_REQUEST",
                "UNSUPPORTED_IMAGE_FORMAT",
                "FAILED_TO_GET_IMAGE",
                "NOT_FOUND",
                "RATE_LIMIT_REACHED",
                "INTERNAL_SERVER_ERROR"
            ],
            "x-enum-varnames": [
                "ErrCodeBadRequest",
                "ErrCodeUnsupportedImageFormat",
                "ErrCodeFailedToGetImage",
                "ErrCodeNotFound",
                "ErrCodeRateLimitReached",
                "ErrCodeInternalServerError"
            ]
        },
        "service.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "$ref": "#/definitions/service.ErrCode"
                }
            }
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
