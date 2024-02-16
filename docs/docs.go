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
        "/api/asset": {
            "post": {
                "description": "Creates a new asset based on the title, privacy setting, asset URL, and asset type provided in the request.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "asset"
                ],
                "summary": "Create new asset",
                "parameters": [
                    {
                        "description": "Create Asset Request",
                        "name": "CreateAssetRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.CreateAssetRequest"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Asset creation successful, returns created asset details along with a success message.",
                        "schema": {
                            "$ref": "#/definitions/api.CreateAssetsResponse"
                        }
                    }
                }
            }
        },
        "/auth/google": {
            "post": {
                "description": "Authenticate user with Google OAuth token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Google Auth",
                "parameters": [
                    {
                        "description": "Google OAuth token",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.googleRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User info from Google",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/auth/signin": {
            "post": {
                "description": "Login user with the provided credentials",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "User login details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.loginUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User login successful",
                        "schema": {
                            "$ref": "#/definitions/api.loginUserResponse"
                        }
                    }
                }
            }
        },
        "/auth/signup": {
            "post": {
                "description": "Register a new user with the provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "User registration details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.registerUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User registration successful",
                        "schema": {
                            "$ref": "#/definitions/api.registerUserResponse"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieve user information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get user data",
                "responses": {
                    "200": {
                        "description": "User information retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/api.UserResponse"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update user information based on the provided user ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Update user information",
                "parameters": [
                    {
                        "description": "User update details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.updateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User information updated successfully",
                        "schema": {
                            "$ref": "#/definitions/api.updateUserResponse"
                        }
                    }
                }
            }
        },
        "/user/password": {
            "patch": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Change user password based on the provided user ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Change user password",
                "parameters": [
                    {
                        "description": "User password update details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.changeUserPasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User password updated successfully",
                        "schema": {
                            "$ref": "#/definitions/api.changeUserPasswordResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.AssetResponse": {
            "type": "object",
            "properties": {
                "assetType": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "gaussianUrl": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "isPrivate": {
                    "type": "boolean"
                },
                "likes": {
                    "type": "integer"
                },
                "pointCloudUrl": {
                    "type": "string"
                },
                "slug": {
                    "type": "string"
                },
                "thumbnailUrl": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/api.UserResponse"
                }
            }
        },
        "api.CreateAssetRequest": {
            "type": "object",
            "required": [
                "assetType",
                "assetUrl",
                "isPrivate",
                "title"
            ],
            "properties": {
                "assetType": {
                    "type": "string"
                },
                "assetUrl": {
                    "type": "string"
                },
                "isPrivate": {
                    "type": "boolean"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "api.CreateAssetsResponse": {
            "type": "object",
            "properties": {
                "asset": {
                    "$ref": "#/definitions/api.AssetResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "api.UserResponse": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "passwordChangedAt": {
                    "type": "string"
                },
                "provider": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "api.changeUserPasswordRequest": {
            "type": "object",
            "properties": {
                "newPassword": {
                    "type": "string"
                },
                "oldPassword": {
                    "type": "string"
                }
            }
        },
        "api.changeUserPasswordResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/api.UserResponse"
                }
            }
        },
        "api.googleRequest": {
            "type": "object",
            "required": [
                "token"
            ],
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "api.loginUserRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 8
                }
            }
        },
        "api.loginUserResponse": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/api.UserResponse"
                }
            }
        },
        "api.registerUserRequest": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 8
                }
            }
        },
        "api.registerUserResponse": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/api.UserResponse"
                }
            }
        },
        "api.updateUserRequest": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "api.updateUserResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/api.UserResponse"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Segment3d App API Documentation",
	Description:      "This is a documentation for Segment3d App API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
