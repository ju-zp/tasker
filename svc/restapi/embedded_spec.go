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
    "description": "as in title",
    "title": "Just a Todo",
    "version": "1.0.0"
  },
  "paths": {
    "/ping": {
      "get": {
        "operationId": "getPing",
        "responses": {
          "200": {
            "description": "server is operational",
            "schema": {
              "description": "Server is up",
              "type": "string"
            }
          },
          "500": {
            "description": "server is not responding",
            "schema": {
              "description": "Server is down",
              "type": "string"
            }
          }
        }
      }
    },
    "/task": {
      "get": {
        "operationId": "getTasks",
        "responses": {
          "200": {
            "description": "Get all the tasks",
            "schema": {
              "description": "array of all the tasks",
              "type": "array",
              "items": {
                "$ref": "#/definitions/task"
              }
            }
          },
          "400": {
            "description": "Incomplete data",
            "schema": {
              "description": "Body of the request was incomplete",
              "type": "string"
            }
          }
        }
      },
      "post": {
        "operationId": "createTask",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/task"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Created task",
            "schema": {
              "$ref": "#/definitions/task"
            }
          },
          "400": {
            "description": "Incomplete data",
            "schema": {
              "description": "Body of the request was Incomplete",
              "type": "string"
            }
          }
        }
      }
    },
    "/task/{taskId}": {
      "get": {
        "operationId": "getTaskTodos",
        "parameters": [
          {
            "$ref": "#/parameters/path_taskId"
          }
        ],
        "responses": {
          "200": {
            "description": "Get the todos associated to a task",
            "schema": {
              "description": "the task and associated todos",
              "type": "object",
              "properties": {
                "task": {
                  "$ref": "#/definitions/task"
                },
                "todos": {
                  "description": "Array of todos",
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/todo"
                  }
                }
              }
            }
          }
        }
      }
    },
    "/task/{taskId}/todo": {
      "post": {
        "operationId": "createTaskTodo",
        "parameters": [
          {
            "$ref": "#/parameters/path_taskId"
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "description": "Body request",
              "type": "object",
              "properties": {
                "todo": {
                  "description": "content of the todo",
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "success message",
            "schema": {
              "description": "operation was successful",
              "type": "string"
            }
          }
        }
      }
    },
    "/todo": {
      "get": {
        "operationId": "getTodos",
        "responses": {
          "200": {
            "description": "Get all todos",
            "schema": {
              "description": "array of all the todos",
              "type": "array",
              "items": {
                "$ref": "#/definitions/todo"
              }
            }
          },
          "400": {
            "description": "Incomplete data",
            "schema": {
              "description": "Body of the request was Incomplete",
              "type": "string"
            }
          }
        }
      },
      "post": {
        "operationId": "createTodo",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/todo"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Created todo",
            "schema": {
              "$ref": "#/definitions/todo"
            }
          },
          "400": {
            "description": "Incomplete data",
            "schema": {
              "description": "Body of the request was Incomplete",
              "type": "string"
            }
          }
        }
      }
    },
    "/todo/{todoId}": {
      "delete": {
        "operationId": "deleteTodo",
        "parameters": [
          {
            "$ref": "#/parameters/path_todoId"
          }
        ],
        "responses": {
          "200": {
            "description": "Delete a todo",
            "schema": {
              "description": "Successfully deleted a todo",
              "type": "string"
            }
          }
        }
      }
    },
    "/todo/{todoId}/status": {
      "post": {
        "operationId": "setTodoStatus",
        "parameters": [
          {
            "$ref": "#/parameters/path_todoId"
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/todoStatus"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Updated the status of todo",
            "schema": {
              "description": "Successfully updated todo",
              "type": "string"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "task": {
      "type": "object",
      "required": [
        "title"
      ],
      "properties": {
        "done": {
          "type": "boolean"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "title": {
          "description": "title of a task",
          "type": "string"
        }
      }
    },
    "todo": {
      "type": "object",
      "required": [
        "todo",
        "done",
        "id"
      ],
      "properties": {
        "done": {
          "type": "boolean"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "taskId": {
          "description": "task it is associated with",
          "type": "integer",
          "format": "int64"
        },
        "todo": {
          "description": "the content of the todo",
          "type": "string"
        }
      }
    },
    "todoStatus": {
      "type": "object",
      "required": [
        "status"
      ],
      "properties": {
        "status": {
          "type": "boolean"
        }
      }
    }
  },
  "parameters": {
    "path_taskId": {
      "type": "string",
      "description": "ID of a task",
      "name": "taskId",
      "in": "path",
      "required": true
    },
    "path_todoId": {
      "type": "string",
      "description": "ID of a todo",
      "name": "todoId",
      "in": "path",
      "required": true
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
    "description": "as in title",
    "title": "Just a Todo",
    "version": "1.0.0"
  },
  "paths": {
    "/ping": {
      "get": {
        "operationId": "getPing",
        "responses": {
          "200": {
            "description": "server is operational",
            "schema": {
              "description": "Server is up",
              "type": "string"
            }
          },
          "500": {
            "description": "server is not responding",
            "schema": {
              "description": "Server is down",
              "type": "string"
            }
          }
        }
      }
    },
    "/task": {
      "get": {
        "operationId": "getTasks",
        "responses": {
          "200": {
            "description": "Get all the tasks",
            "schema": {
              "description": "array of all the tasks",
              "type": "array",
              "items": {
                "$ref": "#/definitions/task"
              }
            }
          },
          "400": {
            "description": "Incomplete data",
            "schema": {
              "description": "Body of the request was incomplete",
              "type": "string"
            }
          }
        }
      },
      "post": {
        "operationId": "createTask",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/task"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Created task",
            "schema": {
              "$ref": "#/definitions/task"
            }
          },
          "400": {
            "description": "Incomplete data",
            "schema": {
              "description": "Body of the request was Incomplete",
              "type": "string"
            }
          }
        }
      }
    },
    "/task/{taskId}": {
      "get": {
        "operationId": "getTaskTodos",
        "parameters": [
          {
            "type": "string",
            "description": "ID of a task",
            "name": "taskId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Get the todos associated to a task",
            "schema": {
              "description": "the task and associated todos",
              "type": "object",
              "properties": {
                "task": {
                  "$ref": "#/definitions/task"
                },
                "todos": {
                  "description": "Array of todos",
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/todo"
                  }
                }
              }
            }
          }
        }
      }
    },
    "/task/{taskId}/todo": {
      "post": {
        "operationId": "createTaskTodo",
        "parameters": [
          {
            "type": "string",
            "description": "ID of a task",
            "name": "taskId",
            "in": "path",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "description": "Body request",
              "type": "object",
              "properties": {
                "todo": {
                  "description": "content of the todo",
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "success message",
            "schema": {
              "description": "operation was successful",
              "type": "string"
            }
          }
        }
      }
    },
    "/todo": {
      "get": {
        "operationId": "getTodos",
        "responses": {
          "200": {
            "description": "Get all todos",
            "schema": {
              "description": "array of all the todos",
              "type": "array",
              "items": {
                "$ref": "#/definitions/todo"
              }
            }
          },
          "400": {
            "description": "Incomplete data",
            "schema": {
              "description": "Body of the request was Incomplete",
              "type": "string"
            }
          }
        }
      },
      "post": {
        "operationId": "createTodo",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/todo"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Created todo",
            "schema": {
              "$ref": "#/definitions/todo"
            }
          },
          "400": {
            "description": "Incomplete data",
            "schema": {
              "description": "Body of the request was Incomplete",
              "type": "string"
            }
          }
        }
      }
    },
    "/todo/{todoId}": {
      "delete": {
        "operationId": "deleteTodo",
        "parameters": [
          {
            "type": "string",
            "description": "ID of a todo",
            "name": "todoId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Delete a todo",
            "schema": {
              "description": "Successfully deleted a todo",
              "type": "string"
            }
          }
        }
      }
    },
    "/todo/{todoId}/status": {
      "post": {
        "operationId": "setTodoStatus",
        "parameters": [
          {
            "type": "string",
            "description": "ID of a todo",
            "name": "todoId",
            "in": "path",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/todoStatus"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Updated the status of todo",
            "schema": {
              "description": "Successfully updated todo",
              "type": "string"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "task": {
      "type": "object",
      "required": [
        "title"
      ],
      "properties": {
        "done": {
          "type": "boolean"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "title": {
          "description": "title of a task",
          "type": "string"
        }
      }
    },
    "todo": {
      "type": "object",
      "required": [
        "todo",
        "done",
        "id"
      ],
      "properties": {
        "done": {
          "type": "boolean"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "taskId": {
          "description": "task it is associated with",
          "type": "integer",
          "format": "int64"
        },
        "todo": {
          "description": "the content of the todo",
          "type": "string"
        }
      }
    },
    "todoStatus": {
      "type": "object",
      "required": [
        "status"
      ],
      "properties": {
        "status": {
          "type": "boolean"
        }
      }
    }
  },
  "parameters": {
    "path_taskId": {
      "type": "string",
      "description": "ID of a task",
      "name": "taskId",
      "in": "path",
      "required": true
    },
    "path_todoId": {
      "type": "string",
      "description": "ID of a todo",
      "name": "todoId",
      "in": "path",
      "required": true
    }
  }
}`))
}
