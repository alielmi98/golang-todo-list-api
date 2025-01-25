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
        "/v1/users/login-by-username": {
            "post": {
                "description": "LoginByUsername",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "LoginByUsername",
                "parameters": [
                    {
                        "description": "LoginByUsernameRequest",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_golang-todo-list-api_api_dto.LoginByUsernameRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_golang-todo-list-api_api_helper.BaseHttpResponse"
                        }
                    },
                    "400": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_golang-todo-list-api_api_helper.BaseHttpResponse"
                        }
                    },
                    "409": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_golang-todo-list-api_api_helper.BaseHttpResponse"
                        }
                    }
                }
            }
        },
        "/v1/users/register-by-username": {
            "post": {
                "description": "RegisterByUsername",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "RegisterByUsername",
                "parameters": [
                    {
                        "description": "RegisterUserByUsernameRequest",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_golang-todo-list-api_api_dto.RegisterUserByUsernameRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_golang-todo-list-api_api_helper.BaseHttpResponse"
                        }
                    },
                    "400": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_golang-todo-list-api_api_helper.BaseHttpResponse"
                        }
                    },
                    "409": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_golang-todo-list-api_api_helper.BaseHttpResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_alielmi98_golang-todo-list-api_api_dto.LoginByUsernameRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "minLength": 6
                },
                "username": {
                    "type": "string",
                    "minLength": 5
                }
            }
        },
        "github_com_alielmi98_golang-todo-list-api_api_dto.RegisterUserByUsernameRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "minLength": 6
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                },
                "username": {
                    "type": "string",
                    "minLength": 5
                }
            }
        },
        "github_com_alielmi98_golang-todo-list-api_api_helper.BaseHttpResponse": {
            "type": "object",
            "properties": {
                "error": {},
                "result": {},
                "resultCode": {
                    "$ref": "#/definitions/github_com_alielmi98_golang-todo-list-api_api_helper.ResultCode"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "github_com_alielmi98_golang-todo-list-api_api_helper.ResultCode": {
            "type": "integer",
            "enum": [
                0,
                40001,
                40101,
                40301,
                40401,
                42901,
                42902,
                50001,
                50002
            ],
            "x-enum-varnames": [
                "Success",
                "ValidationError",
                "AuthError",
                "ForbiddenError",
                "NotFoundError",
                "LimiterError",
                "OtpLimiterError",
                "CustomRecovery",
                "InternalError"
            ]
        }
    },
    "securityDefinitions": {
        "AuthBearer": {
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
