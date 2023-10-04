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
        "/v1/images": {
            "get": {
                "parameters": [
                    {
                        "type": "string",
                        "description": "query",
                        "name": "q",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Image"
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
        },
        "/v1/lgtms": {
            "get": {
                "parameters": [
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "after",
                        "name": "after",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "random",
                        "name": "random",
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
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.createLGTMInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.LGTM"
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
        },
        "/v1/news": {
            "get": {
                "parameters": [
                    {
                        "type": "string",
                        "description": "locale",
                        "name": "locale",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.News"
                            }
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
        },
        "/v1/reports": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.createReportInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Report"
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
        "models.Image": {
            "type": "object",
            "required": [
                "title",
                "url"
            ],
            "properties": {
                "title": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "models.LGTM": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "models.News": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.Report": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "lgtm_id": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                },
                "type": {
                    "$ref": "#/definitions/models.ReportType"
                }
            }
        },
        "models.ReportType": {
            "type": "string",
            "enum": [
                "illegal",
                "inappropriate",
                "other"
            ],
            "x-enum-varnames": [
                "ReportTypeIllegal",
                "ReportTypeInappropriate",
                "ReportTypeOther"
            ]
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
            "required": [
                "code"
            ],
            "properties": {
                "code": {
                    "$ref": "#/definitions/service.ErrCode"
                }
            }
        },
        "service.createLGTMInput": {
            "type": "object",
            "properties": {
                "base64": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "service.createReportInput": {
            "type": "object",
            "required": [
                "lgtm_id",
                "text",
                "type"
            ],
            "properties": {
                "lgtm_id": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                },
                "type": {
                    "$ref": "#/definitions/models.ReportType"
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
