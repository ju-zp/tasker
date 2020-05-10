package todohandlers

import (
	"fmt"

	"github.com/go-openapi/runtime/middleware"
	"github.com/ju-zp/tasker/gen/restapi/operations"
)

// GetTodos gets all todos
func GetTodos(params *operations.GetTodoParams) middleware.Responder {
	fmt.Println("here")
	return operations.NewGetTodoOK()
}

// CreateTodo makes a new todo entry in db
func CreateTodo(params *operations.PostTodoParams) middleware.Responder {
	fmt.Println("or here")
	return operations.NewPostTodoOK()
}
