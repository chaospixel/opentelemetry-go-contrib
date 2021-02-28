// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

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
    "description": "demonstrate otel middleware",
    "title": "example server",
    "version": "0.17.0"
  },
  "paths": {
    "/users/{userId}": {
      "get": {
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "userId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "user info",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "404": {
            "description": "user not found"
          }
        }
      }
    }
  },
  "definitions": {
    "user": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "name": {
          "type": "string",
          "minLength": 1
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
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
    "description": "demonstrate otel middleware",
    "title": "example server",
    "version": "0.17.0"
  },
  "paths": {
    "/users/{userId}": {
      "get": {
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "userId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "user info",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "404": {
            "description": "user not found"
          }
        }
      }
    }
  },
  "definitions": {
    "user": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "name": {
          "type": "string",
          "minLength": 1
        }
      }
    }
  }
}`))
}
