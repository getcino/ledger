// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/_info": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Server Info",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/config.ConfigInfo"
                        }
                    }
                }
            }
        },
        "/{ledger}/accounts": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "List All Accounts",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ledger",
                        "name": "ledger",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/controllers.BaseResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "cursor": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/query.Cursor"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "data": {
                                                            "type": "array",
                                                            "items": {
                                                                "$ref": "#/definitions/core.Account"
                                                            }
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/{ledger}/accounts/{accountId}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get account by address",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ledger",
                        "name": "ledger",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "accountId",
                        "name": "accountId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/controllers.BaseResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "account": {
                                            "$ref": "#/definitions/core.Account"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/{ledger}/accounts/{accountId}/metadata": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Add metadata to account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ledger",
                        "name": "ledger",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "accountId",
                        "name": "accountId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.BaseResponse"
                        }
                    }
                }
            }
        },
        "/{ledger}/script": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Execute a Numscript and commit transaction if any",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ledger",
                        "name": "ledger",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "script",
                        "name": "script",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/core.Script"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.BaseResponse"
                        }
                    }
                }
            }
        },
        "/{ledger}/stats": {
            "get": {
                "description": "The stats for account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get ledger stats (aggregate metrics on accounts and transactions)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ledger",
                        "name": "ledger",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ledger.Stats"
                        }
                    }
                }
            }
        },
        "/{ledger}/transactions": {
            "get": {
                "description": "List transactions",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get Transactions",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ledger",
                        "name": "ledger",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/controllers.BaseResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "cursor": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/query.Cursor"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "data": {
                                                            "type": "array",
                                                            "items": {
                                                                "$ref": "#/definitions/core.Transaction"
                                                            }
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "description": "Commit a new transaction to the ledger",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Commit a new transaction to the ledger",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ledger",
                        "name": "ledger",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "transaction",
                        "name": "transaction",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/core.Transaction"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.BaseResponse"
                        }
                    }
                }
            }
        },
        "/{ledger}/transactions/{reference}/metadata": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Set metadata on transaction",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ledger",
                        "name": "ledger",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "reference",
                        "name": "reference",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.BaseResponse"
                        }
                    }
                }
            }
        },
        "/{ledger}/transactions/{transactionId}/revert": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Revert transaction",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ledger",
                        "name": "ledger",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "transactionId",
                        "name": "transactionId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.BaseResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "config.Config": {
            "type": "object",
            "properties": {
                "storage": {
                    "$ref": "#/definitions/config.LedgerStorage"
                }
            }
        },
        "config.ConfigInfo": {
            "type": "object",
            "properties": {
                "config": {
                    "$ref": "#/definitions/config.Config"
                },
                "server": {
                    "type": "string"
                },
                "version": {}
            }
        },
        "config.LedgerStorage": {
            "type": "object",
            "properties": {
                "driver": {},
                "ledgers": {}
            }
        },
        "controllers.BaseResponse": {
            "type": "object",
            "properties": {
                "ok": {
                    "type": "boolean"
                }
            }
        },
        "core.Account": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string",
                    "example": "users:001"
                },
                "balances": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "integer"
                    },
                    "example": {
                        "COIN": 100
                    }
                },
                "contract": {
                    "type": "string",
                    "example": "default"
                },
                "metadata": {
                    "type": "object"
                },
                "type": {
                    "type": "string",
                    "example": "virtual"
                },
                "volumes": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "object",
                        "additionalProperties": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "core.Posting": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "asset": {
                    "type": "string"
                },
                "destination": {
                    "type": "string"
                },
                "source": {
                    "type": "string"
                }
            }
        },
        "core.Script": {
            "type": "object",
            "properties": {
                "plain": {
                    "type": "string"
                },
                "vars": {
                    "type": "object"
                }
            }
        },
        "core.Transaction": {
            "type": "object",
            "properties": {
                "metadata": {
                    "type": "object"
                },
                "postings": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/core.Posting"
                    }
                },
                "reference": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "string"
                },
                "txid": {
                    "type": "integer"
                }
            }
        },
        "ledger.Stats": {
            "type": "object",
            "properties": {
                "accounts": {
                    "type": "integer"
                },
                "transactions": {
                    "type": "integer"
                }
            }
        },
        "query.Cursor": {
            "type": "object",
            "properties": {
                "data": {},
                "has_more": {
                    "type": "boolean"
                },
                "next": {
                    "type": "string"
                },
                "page_size": {
                    "type": "integer"
                },
                "previous": {
                    "type": "string"
                },
                "remaining_results": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
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
	Host:        "localhost:3068",
	BasePath:    "",
	Schemes:     []string{"http", "https"},
	Title:       "Ledger API",
	Description: "",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
	swag.Register("swagger", &s{})
}
