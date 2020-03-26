// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-03-26 19:21:03.072643322 +0700 WIB m=+0.033396382

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "richstain",
            "email": "richstain2u@gmail.com"
        },
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/client/login": {
            "post": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "login returns token and expire time in seconds",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Client"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "JSON request body",
                        "name": "RequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.LoginPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.LoginResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/helper.EchoResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helper.EchoResp"
                        }
                    }
                }
            }
        },
        "/user/user": {
            "get": {
                "security": [
                    {
                        "OAuth2Password": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "View current user datas",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.ViewCurrentUserResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/helper.EchoResp"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/helper.EchoResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helper.EchoResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.LoginPayload": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "example": "pass123"
                },
                "username": {
                    "type": "string",
                    "example": "user123"
                }
            }
        },
        "handlers.LoginResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "Error code explanations\n0000 = Success\n0001 = Error Validation\n0002 = Function Error\n0003 = Permission denied",
                    "type": "string",
                    "example": "0001"
                },
                "details": {
                    "type": "object",
                    "$ref": "#/definitions/handlers.LoginResponseDetails"
                },
                "message": {
                    "type": "string",
                    "example": "this is example message"
                }
            }
        },
        "handlers.LoginResponseDetails": {
            "type": "object",
            "properties": {
                "expires_in": {
                    "type": "integer",
                    "example": 3600
                },
                "token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJle..."
                }
            }
        },
        "handlers.ViewCurrentUserResponse": {
            "type": "object",
            "properties": {
                "username": {
                    "type": "string"
                }
            }
        },
        "helper.EchoResp": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "Error code explanations\n0000 = Success\n0001 = Error Validation\n0002 = Function Error\n0003 = Permission denied",
                    "type": "string",
                    "example": "0001"
                },
                "details": {
                    "type": "object"
                },
                "message": {
                    "type": "string",
                    "example": "this is example message"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        },
        "OAuth2Password": {
            "type": "oauth2",
            "flow": "password",
            "tokenUrl": "client/login"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Pando API",
	Description: "This is a sample server Petstore server.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}