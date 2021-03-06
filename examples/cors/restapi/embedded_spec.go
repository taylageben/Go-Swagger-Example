// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "## Welcome\n\nThis is a place to put general notes and extra information, for internal use.\n\nTo get started designing/documenting this API, select a version on the left. ## Welcome\n\nThis is a place to put general notes and extra information, for internal use.\n\nTo get started designing/documenting this API, select a version on the left.",
    "title": "To-do Demo",
    "version": "1.0"
  },
  "host": "todos.stoplight.io",
  "paths": {
    "/mocked": {
      "get": {
        "description": "This is an example of a mocked endpoint.\n\nThis endpoint does not actually exist in the api - try visiting [http://todos.stoplight.io/mocked](http://todos.stoplight.io). You will see the default response (same as you get when you visit the root \"/\" url).\n\nWe have defined a 200 response below, with an example, and then explicitly turned on mocking for this endpoint via the dropdown in the right sidebar.\n\nWith this enabled, if you visit {your prism instance host}/mocked, you'll see the mocked example. You can find the mock server host for this API by clicking the \"Design Dashboard\" button at the top of this screen. You can also try sending a test request by clicking \"send test request\" button in the right sidebar.",
        "tags": [
          "Misc"
        ],
        "summary": "Mock Example",
        "operationId": "get_mocked",
        "responses": {
          "200": {
            "schema": {
              "type": "object",
              "properties": {
                "mocked": {
                  "type": "boolean"
                },
                "this": {
                  "type": "string"
                }
              }
            }
          }
        }
      }
    },
    "/todos": {
      "get": {
        "tags": [
          "Todos"
        ],
        "summary": "List Todos",
        "operationId": "get_todos",
        "security": [
          {
            "apikey": []
          }
        ],
        "parameters": [
          {
            "maximum": 100,
            "type": "integer",
            "description": "This is how it works.",
            "name": "limit",
            "in": "query"
          },
          {
            "type": "string",
            "name": "skip",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/todo-full"
              }
            }
          },
          "401": {
            "$ref": "#/responses/trait:standardErrors:401"
          },
          "403": {
            "$ref": "#/responses/trait:standardErrors:403"
          },
          "404": {
            "$ref": "#/responses/trait:standardErrors:404"
          },
          "500": {
            "$ref": "#/responses/trait:standardErrors:500"
          }
        }
      },
      "post": {
        "tags": [
          "Todos"
        ],
        "summary": "Create Todo",
        "operationId": "post_todos",
        "security": [
          {
            "apikey": []
          }
        ],
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/todo-partial",
              "example": {
                "completed": false,
                "name": "my todo's name"
              }
            }
          }
        ],
        "responses": {
          "201": {
            "schema": {
              "$ref": "#/definitions/todo-full"
            }
          },
          "401": {
            "$ref": "#/responses/trait:standardErrors:401"
          },
          "403": {
            "schema": {
              "type": "object"
            }
          },
          "404": {
            "$ref": "#/responses/trait:standardErrors:404"
          },
          "500": {
            "$ref": "#/responses/trait:standardErrors:500"
          }
        }
      }
    },
    "/todos/{todoId}": {
      "get": {
        "tags": [
          "Todos"
        ],
        "summary": "Get Todo",
        "operationId": "get_todo",
        "responses": {
          "200": {
            "schema": {
              "$ref": "#/definitions/todo-full"
            }
          },
          "401": {
            "$ref": "#/responses/trait:standardErrors:401"
          },
          "403": {
            "$ref": "#/responses/trait:standardErrors:403"
          },
          "404": {
            "$ref": "#/responses/trait:standardErrors:404"
          },
          "500": {
            "$ref": "#/responses/trait:standardErrors:500"
          }
        }
      },
      "put": {
        "tags": [
          "Todos"
        ],
        "summary": "Update Todo",
        "operationId": "put_todos",
        "security": [
          {
            "apikey": []
          }
        ],
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/todo-partial",
              "example": {
                "completed": false,
                "name": "my todo's new name"
              }
            }
          }
        ],
        "responses": {
          "200": {
            "schema": {
              "$ref": "#/definitions/todo-full"
            }
          },
          "401": {
            "$ref": "#/responses/trait:standardErrors:401"
          },
          "403": {
            "schema": {
              "type": "object"
            }
          },
          "404": {
            "$ref": "#/responses/trait:standardErrors:404"
          },
          "500": {
            "$ref": "#/responses/trait:standardErrors:500"
          }
        }
      },
      "delete": {
        "tags": [
          "Todos"
        ],
        "summary": "Delete Todo",
        "operationId": "delete_todo",
        "security": [
          {
            "apikey": []
          }
        ],
        "responses": {
          "204": {
            "schema": {
              "type": "object"
            }
          },
          "401": {
            "$ref": "#/responses/trait:standardErrors:401"
          },
          "403": {
            "$ref": "#/responses/trait:standardErrors:403"
          },
          "404": {
            "$ref": "#/responses/trait:standardErrors:404"
          },
          "500": {
            "$ref": "#/responses/trait:standardErrors:500"
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "todoId",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "todo-full": {
      "title": "Todo Full",
      "allOf": [
        {
          "$ref": "#/definitions/todo-partial"
        },
        {
          "type": "object",
          "required": [
            "id"
          ],
          "properties": {
            "completed_at": {
              "type": [
                "string",
                "null"
              ],
              "format": "date-time"
            },
            "created_at": {
              "type": "string",
              "format": "date-time"
            },
            "id": {
              "type": "integer"
            },
            "updated_at": {
              "type": "string",
              "format": "date-time"
            }
          }
        }
      ]
    },
    "todo-partial": {
      "type": "object",
      "title": "Todo Partial",
      "required": [
        "name"
      ],
      "properties": {
        "completed": {
          "type": [
            "boolean",
            "null"
          ]
        },
        "name": {
          "type": "string"
        }
      }
    }
  },
  "parameters": {
    "trait:paged:limit": {
      "maximum": 100,
      "type": "integer",
      "description": "This is how it works.",
      "name": "limit",
      "in": "query"
    },
    "trait:paged:skip": {
      "type": "string",
      "name": "skip",
      "in": "query"
    }
  },
  "responses": {
    "trait:standardErrors:401": {
      "schema": {
        "type": "object"
      }
    },
    "trait:standardErrors:403": {
      "schema": {
        "type": "object",
        "required": [
          "message"
        ],
        "properties": {
          "message": {
            "type": "string"
          }
        }
      }
    },
    "trait:standardErrors:404": {
      "schema": {
        "type": "object",
        "required": [
          "status",
          "error"
        ],
        "properties": {
          "error": {
            "type": "string"
          },
          "status": {
            "type": "string"
          }
        }
      }
    },
    "trait:standardErrors:500": {
      "schema": {
        "type": "object"
      }
    }
  },
  "securityDefinitions": {
    "apikey": {
      "type": "apiKey",
      "name": "apikey",
      "in": "query"
    }
  }
}`))
}
