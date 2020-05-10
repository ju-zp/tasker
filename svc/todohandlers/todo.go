package todohandlers

import (
	"fmt"

	"github.com/go-openapi/runtime/middleware"
	"github.com/ju-zp/tasker/svc/restapi/operations"
)

// GetTodos gets all todos
func GetTodos(params operations.GetTodosParams) middleware.Responder {

	fmt.Println("here")
	return operations.NewGetTodosOK()
}

// CreateTodo makes a new todo entry in db
func CreateTodo(params operations.CreateTodoParams) middleware.Responder {
	fmt.Println("or here")
	return operations.NewCreateTodoOK()
}
