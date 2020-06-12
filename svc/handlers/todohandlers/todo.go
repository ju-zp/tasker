package todohandlers

import (
	"fmt"
	"github.com/go-openapi/swag"
	"github.com/jinzhu/gorm"

	"github.com/go-openapi/runtime/middleware"
	"github.com/ju-zp/tasker/svc/restapi/operations"
)

type Context struct {
	DB *gorm.DB
}

// GetTodos gets all todos
func GetTodos(params operations.GetTodosParams) middleware.Responder {
	fmt.Println("here")
	return operations.NewGetTodosOK()
}

// CreateTodo makes a new todo entry in db
func CreateTodo(params operations.CreateTodoParams) middleware.Responder {
	fmt.Println(swag.StringValue(params.Body.Todo))
	return operations.NewCreateTodoOK()
}

// SetTodoStatus updates the done field kfo a given todo
func (c *Context)SetTodoStatus(params operations.SetTodoStatusParams) middleware.Responder {
	fmt.Println("here")
	return operations.NewSetTodoStatusOK()
}
