package todohandlers

import (
	"fmt"

	"github.com/go-openapi/runtime/middleware"
	"github.com/ju-zp/tasker/gen/restapi/operations"
)

// GetTodo gets all todos
func GetTodo(params *operations.GetTodoParams) middleware.Responder {
	fmt.Println("here")
	return operations.NewGetTodoOK()
}
