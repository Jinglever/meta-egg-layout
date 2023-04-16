// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/v1/users/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "获取用户详情",
                "operationId": "GetUserDetail",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer \u003cjwt-token\u003e",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "用户ID",
                        "name": "id",
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
                                    "$ref": "#/definitions/internal_handler_http.RspData"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/internal_handler_http.UserDetail"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/internal_handler_http.RspBase"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "internal_handler_http.RspBase": {
            "type": "object",
            "properties": {
                "code": {
                    "$ref": "#/definitions/meta-egg-layout_api_meta_egg_layout.ErrCode"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "internal_handler_http.RspData": {
            "type": "object",
            "properties": {
                "code": {
                    "$ref": "#/definitions/meta-egg-layout_api_meta_egg_layout.ErrCode"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "internal_handler_http.UserDetail": {
            "type": "object",
            "properties": {
                "created_at": {
                    "description": "创建时间",
                    "type": "integer"
                },
                "created_by": {
                    "description": "创建者",
                    "type": "integer"
                },
                "gender": {
                    "description": "性别",
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "description": "用户名",
                    "type": "string"
                },
                "updated_at": {
                    "description": "更新时间",
                    "type": "integer"
                },
                "updated_by": {
                    "description": "更新者",
                    "type": "integer"
                }
            }
        },
        "meta-egg-layout_api_meta_egg_layout.ErrCode": {
            "type": "integer",
            "enum": [
                0,
                2,
                3,
                5,
                6,
                7,
                8,
                9,
                10,
                11,
                13,
                14,
                15,
                16
            ],
            "x-enum-comments": {
                "ErrCode_Aborted": "操作被中止",
                "ErrCode_AlreadyExists": "创建实体时冲突",
                "ErrCode_DataLoss": "数据丢失或损坏",
                "ErrCode_FailedPrecondition": "前置条件失败",
                "ErrCode_Internal": "内部错误",
                "ErrCode_InvalidArgument": "参数错",
                "ErrCode_NotFound": "实体不存在",
                "ErrCode_Ok": "成功",
                "ErrCode_OutOfRange": "超出范围",
                "ErrCode_PermissionDenied": "权限不足",
                "ErrCode_ResourceExhausted": "资源不足",
                "ErrCode_Unauthenticated": "未认证，客户端未提供凭据或提供的凭据无效",
                "ErrCode_Unavailable": "服务不可用，请重试",
                "ErrCode_Unknown": "未知错误"
            },
            "x-enum-varnames": [
                "ErrCode_Ok",
                "ErrCode_Unknown",
                "ErrCode_InvalidArgument",
                "ErrCode_NotFound",
                "ErrCode_AlreadyExists",
                "ErrCode_PermissionDenied",
                "ErrCode_ResourceExhausted",
                "ErrCode_FailedPrecondition",
                "ErrCode_Aborted",
                "ErrCode_OutOfRange",
                "ErrCode_Internal",
                "ErrCode_Unavailable",
                "ErrCode_DataLoss",
                "ErrCode_Unauthenticated"
            ]
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "xxx",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "meta-egg-layout",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
