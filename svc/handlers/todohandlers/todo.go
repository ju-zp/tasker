package todohandlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/ju-zp/tasker/svc/services/tasker/restapi/operations"
	"github.com/ju-zp/tasker/svc/todo"
)

type Context struct {
	Repository *todo.Repository
}

// SetTodoStatus updates the done field kfo a given todo
func (ctx *Context)SetTodoStatus(params operations.SetTodoStatusParams) middleware.Responder {
	err := ctx.Repository.UpdateStatus(params.TodoID, swag.BoolValue(params.Body.Status))

	if err != nil {
		return nil
	}

	return operations.NewSetTodoStatusOK().WithPayload("Success")
}

// DeleteTodo delete the todo
func (ctx *Context)DeleteTodo(params operations.DeleteTodoParams) middleware.Responder {
	err := ctx.Repository.Delete(params.TodoID)

	if err != nil {
		return nil
	}

	return operations.NewDeleteTodoOK().WithPayload("Success")
}
