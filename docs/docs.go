package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "https://github.com/verssache/chatgpt-creator"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/accounts": {
            "get": {
                "description": "Get all registered accounts",
                "produces": ["application/json"],
                "tags": ["accounts"],
                "summary": "List accounts",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "accounts": {
                                    "type": "array",
                                    "items": {
                                        "$ref": "#/definitions/Account"
                                    }
                                }
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new account",
                "produces": ["application/json"],
                "tags": ["accounts"],
                "summary": "Create account",
                "parameters": [
                    {
                        "description": "Account data",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Account"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "account": {
                                    "$ref": "#/definitions/Account"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/email-domains": {
            "get": {
                "description": "Get all configured email domains",
                "produces": ["application/json"],
                "tags": ["email-domains"],
                "summary": "List email domains",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "domains": {
                                    "type": "array",
                                    "items": {
                                        "$ref": "#/definitions/EmailDomain"
                                    }
                                }
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Add a new email domain for registration",
                "produces": ["application/json"],
                "tags": ["email-domains"],
                "summary": "Create email domain",
                "parameters": [
                    {
                        "description": "Domain data",
                        "name": "domain",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/EmailDomain"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "domain": {
                                    "$ref": "#/definitions/EmailDomain"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/batch-jobs": {
            "get": {
                "description": "Get all batch registration jobs",
                "produces": ["application/json"],
                "tags": ["batch-jobs"],
                "summary": "List batch jobs",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "jobs": {
                                    "type": "array",
                                    "items": {
                                        "$ref": "#/definitions/BatchJob"
                                    }
                                }
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new batch registration job",
                "produces": ["application/json"],
                "tags": ["batch-jobs"],
                "summary": "Create batch job",
                "parameters": [
                    {
                        "description": "Job configuration",
                        "name": "job",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/BatchJob"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "job": {
                                    "$ref": "#/definitions/BatchJob"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/batch-jobs/{id}/start": {
            "post": {
                "description": "Start a pending batch job",
                "produces": ["application/json"],
                "tags": ["batch-jobs"],
                "summary": "Start batch job",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Job ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "status": {
                                    "type": "string"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/configurations": {
            "get": {
                "description": "Get all application configurations",
                "produces": ["application/json"],
                "tags": ["configurations"],
                "summary": "List configurations",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "configurations": {
                                    "type": "array",
                                    "items": {
                                        "$ref": "#/definitions/Configuration"
                                    }
                                }
                            }
                        }
                    }
                }
            }
        },
        "/stats/dashboard": {
            "get": {
                "description": "Get dashboard statistics",
                "produces": ["application/json"],
                "tags": ["stats"],
                "summary": "Dashboard stats",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/DashboardStats"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "Account": {
            "type": "object",
            "properties": {
                "id": {"type": "string"},
                "email": {"type": "string"},
                "status": {"type": "string", "enum": ["active", "inactive", "suspended"]},
                "batch_job_id": {"type": "string"},
                "created_at": {"type": "string"},
                "updated_at": {"type": "string"}
            }
        },
        "EmailDomain": {
            "type": "object",
            "properties": {
                "id": {"type": "string"},
                "domain": {"type": "string"},
                "priority": {"type": "integer"},
                "is_active": {"type": "boolean"},
                "source": {"type": "string", "enum": ["generator", "custom"]},
                "health_status": {"type": "string", "enum": ["healthy", "unhealthy", "unknown"]},
                "created_at": {"type": "string"},
                "updated_at": {"type": "string"}
            }
        },
        "BatchJob": {
            "type": "object",
            "properties": {
                "id": {"type": "string"},
                "target_count": {"type": "integer"},
                "success_count": {"type": "integer"},
                "failure_count": {"type": "integer"},
                "status": {"type": "string", "enum": ["pending", "running", "completed", "cancelled", "failed"]},
                "max_workers": {"type": "integer"},
                "created_at": {"type": "string"}
            }
        },
        "Configuration": {
            "type": "object",
            "properties": {
                "id": {"type": "string"},
                "key": {"type": "string"},
                "value": {"type": "string"},
                "created_at": {"type": "string"},
                "updated_at": {"type": "string"}
            }
        },
        "DashboardStats": {
            "type": "object",
            "properties": {
                "total_accounts": {"type": "integer"},
                "active_accounts": {"type": "integer"},
                "total_batch_jobs": {"type": "integer"},
                "running_batch_jobs": {"type": "integer"},
                "active_email_domains": {"type": "integer"}
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
	Schemes     []string
}

var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "/api",
	Title:       "ChatGPT Creator API",
	Description: "API for ChatGPT account registration bot with batch processing, email domain management, and real-time progress tracking.",
	Schemes:     []string{"http", "https"},
}

type S struct{}

func (s *S) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = "API for ChatGPT account registration bot with batch processing, email domain management, and real-time progress tracking."
	return docTemplate
}

func init() {
	swag.Register(swag.Name, &S{})
}
