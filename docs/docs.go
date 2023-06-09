package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Anmol Singh",
            "email": "anmolsinghturka@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api": {
            "get": {
                "description": "Will respond I'm Alive! if the API is up.",
                "produces": [
                    "application/json"
                ],
                "summary": "Health Check",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/api/api-docs": {
            "get": {
                "description": "Swagger docs.",
                "produces": [
                    "application/json"
                ],
                "summary": "Swagger docs.",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/api/products/{id}": {
            "get": {
                "description": "Will respond I'm Alive! if the API is up.",
                "produces": [
                    "application/json"
                ],
                "summary": "Delete product",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            },
            "put": {
                "description": "Will respond I'm Alive! if the API is up.",
                "produces": [
                    "application/json"
                ],
                "summary": "Delete product",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            },
            "post": {
                "description": "Will respond I'm Alive! if the API is up.",
                "produces": [
                    "application/json"
                ],
                "summary": "Delete product",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            },
            "delete": {
                "description": "Will respond I'm Alive! if the API is up.",
                "produces": [
                    "application/json"
                ],
                "summary": "Delete product",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Golang Backend CRUD API Completed By Anmol Singh",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
