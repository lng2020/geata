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
        "/api/v1/audit_log": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AuditLog"
                ],
                "summary": "Get all audit logs",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/service.AuditLog"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/audit_log/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AuditLog"
                ],
                "summary": "Get audit log by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Audit log ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.AuditLog"
                        }
                    }
                }
            }
        },
        "/api/v1/iec61850/data_object/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "IEC61850"
                ],
                "summary": "Get DataObject by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "DataObject ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.DataObject"
                        }
                    }
                }
            }
        },
        "/api/v1/iec61850/data_object/{object_id}/node": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "IEC61850"
                ],
                "summary": "Get Nodes by DataObject ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "DataObject ID",
                        "name": "object_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/service.DataAttribute"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/iec61850/model/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "IEC61850"
                ],
                "summary": "Get IEC61850 model by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "IEC61850 model ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.IEC61850Model"
                        }
                    }
                }
            }
        },
        "/api/v1/iec61850/node/ref/{ref}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "IEC61850"
                ],
                "summary": "Get Node by IEC61850 reference",
                "parameters": [
                    {
                        "type": "string",
                        "description": "IEC61850 reference",
                        "name": "ref",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.DataAttribute"
                        }
                    }
                }
            }
        },
        "/api/v1/iec61850/node/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "IEC61850"
                ],
                "summary": "Get Node by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Node ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.DataAttribute"
                        }
                    }
                }
            }
        },
        "/api/v1/iec61850/node/{id}/data_source": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "IEC61850"
                ],
                "summary": "Update DataSource of a Node",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Node ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "DataSource",
                        "name": "dataSource",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/mapping_rule/{id}": {
            "get": {
                "tags": [
                    "MappingRule"
                ],
                "summary": "Get a MappingRule by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "MappingRule ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.MappingRule"
                        }
                    }
                }
            },
            "put": {
                "tags": [
                    "MappingRule"
                ],
                "summary": "Update a MappingRule",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "MappingRule ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "MappingRule",
                        "name": "rule",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.MappingRule"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.MappingRule"
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "MappingRule"
                ],
                "summary": "Delete a MappingRule",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "MappingRule ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/api/v1/modbus_detail/{ruleID}": {
            "get": {
                "tags": [
                    "ModbusDetail"
                ],
                "summary": "Get ModbusDetail by RuleID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Rule ID",
                        "name": "ruleID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.ModbusDetail"
                        }
                    }
                }
            }
        },
        "/api/v1/mqtt_detail/{ruleID}": {
            "get": {
                "tags": [
                    "MqttDetail"
                ],
                "summary": "Get MqttDetail by RuleID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Rule ID",
                        "name": "ruleID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.MqttDetail"
                        }
                    }
                }
            }
        },
        "/api/v1/station": {
            "get": {
                "description": "List all stations",
                "produces": [
                    "application/json"
                ],
                "summary": "List all stations",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/service.Station"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create station",
                "produces": [
                    "application/json"
                ],
                "summary": "Create station",
                "parameters": [
                    {
                        "description": "Station",
                        "name": "station",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.Station"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.Station"
                        }
                    }
                }
            }
        },
        "/api/v1/station/{stationID}/mapping_rule": {
            "get": {
                "tags": [
                    "MappingRule"
                ],
                "summary": "List all MappingRule for a station",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Station ID",
                        "name": "stationID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/service.MappingRule"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/station/{station_id}": {
            "get": {
                "description": "Get station by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Get station by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Station ID",
                        "name": "station_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.Station"
                        }
                    }
                }
            },
            "put": {
                "description": "Update station by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Update station by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Station ID",
                        "name": "station_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Station",
                        "name": "station",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.Station"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.Station"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete station by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Delete station by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Station ID",
                        "name": "station_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.Station"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "service.AuditLog": {
            "type": "object",
            "properties": {
                "create_time": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "service.DataAttribute": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "data_object_id": {
                    "type": "integer"
                },
                "data_source": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "iec61850_ref": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "service.DataObject": {
            "type": "object",
            "properties": {
                "dataAttribute": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/service.DataAttribute"
                    }
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "logicalNodeID": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "service.IEC61850Model": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "logicalDevice": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/service.LogicalDevice"
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "service.LogicalDevice": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "logicalNode": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/service.LogicalNode"
                    }
                },
                "modelID": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "service.LogicalNode": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "logicalDeviceID": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "service.MappingRule": {
            "type": "object",
            "properties": {
                "create_time": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "iec61850_ref": {
                    "type": "string"
                },
                "station_id": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                },
                "update_time": {
                    "type": "string"
                }
            }
        },
        "service.ModbusDetail": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "rule_id": {
                    "type": "integer"
                }
            }
        },
        "service.MqttDetail": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "rule_id": {
                    "type": "integer"
                }
            }
        },
        "service.Station": {
            "type": "object",
            "properties": {
                "configs": {
                    "type": "object",
                    "additionalProperties": {}
                },
                "handlers": {
                    "type": "object",
                    "additionalProperties": {}
                },
                "host": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "port": {
                    "type": "integer"
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
