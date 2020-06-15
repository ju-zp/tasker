package taskhandlers

import (
	"fmt"
	"github.com/ju-zp/tasker/svc/models"
	"github.com/ju-zp/tasker/svc/task"
	"github.com/ju-zp/tasker/svc/todo"

	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
	"github.com/ju-zp/tasker/svc/restapi/operations"
)

// Context for task handlers
type Context struct {
	DB *gorm.DB
	Repository *task.Repository
}

// GetTasks gets all the tasks
func (ctx Context) GetTasks(params operations.GetTasksParams) middleware.Responder {
	var tasks []*models.Task

	ctx.DB.Find(&tasks)

	return operations.NewGetTasksOK().WithPayload(tasks)
}

// CreateTask creates a task
func (ctx Context) CreateTask(params operations.CreateTaskParams) middleware.Responder {
	task, err := ctx.Repository.Create(params.Body.Title, params.Body.ProjectID)

	if err != nil {
		fmt.Println("Handle this error")
		return nil
	}

	return operations.NewCreateTaskOK().WithPayload(task)
}

// GetTaskTodos gets the task and the associated todos
func (ctx Context) GetTaskTodos(params operations.GetTaskTodosParams) middleware.Responder {
	task, err := ctx.Repository.Find(params.TaskID)

	todos, err := todo.CreateRepository(ctx.DB).FindByTaskId(task.ID)

	if err != nil {
		fmt.Println("Handle this error")
		return nil
	}

	return operations.NewGetTaskTodosOK().WithPayload(&operations.GetTaskTodosOKBody{
		Task:  task,
		Todos: todos,
	})
}

// CreateTaskTodo creates a todo for a given task
func (ctx Context) CreateTaskTodo(params operations.CreateTaskTodoParams) middleware.Responder {
	task, err := ctx.Repository.Find(params.TaskID)

	if err != nil {
		fmt.Println("Handle this error")
		return nil
	}

	err = todo.CreateRepository(ctx.DB).Create(&params.Body.Todo, task.ID)

	if err != nil {
		fmt.Println("Handle this error")
		return nil
	}

	return operations.NewCreateTaskTodoOK().WithPayload("success")
}

// DeleteTask deletes the given task and associated todos
func (ctx Context) DeleteTask(params operations.DeleteTaskParams) middleware.Responder {
	err := ctx.Repository.Delete(params.TaskID)

	if err != nil {
		fmt.Println("Handle this error")
		return nil
	}

	return operations.NewDeleteTodoOK().WithPayload("Success")
}
