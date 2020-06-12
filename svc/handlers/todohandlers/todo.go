package todohandlers

import (
	"fmt"
	"github.com/go-openapi/swag"
	"github.com/jinzhu/gorm"
	"github.com/ju-zp/tasker/svc/models"
	"strconv"

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
func (ctx *Context)SetTodoStatus(params operations.SetTodoStatusParams) middleware.Responder {
	id, _ := strconv.Atoi(params.TodoID)

	todo := models.Todo{
		ID:     int64(id),
	}

	err := ctx.DB.Find(&todo).Error

	if err != nil {
		return nil
	}

	todo.Done = params.Body.Status

	err = ctx.DB.Save(&todo).Error

	if err != nil {
		return nil
	}

	return operations.NewSetTodoStatusOK().WithPayload("Success")
}

// DeleteTodo delete the todo
func (ctx *Context)DeleteTodo(params operations.DeleteTodoParams) middleware.Responder {
	id, _ := strconv.Atoi(params.TodoID)

	todo := models.Todo{
		ID:     int64(id),
	}

	err := ctx.DB.Delete(&todo).Error

	if err != nil {
		return nil
	}

	return operations.NewDeleteTodoOK().WithPayload("Success")
}
