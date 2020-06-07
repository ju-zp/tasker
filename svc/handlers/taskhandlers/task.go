package taskhandlers

import (
	"fmt"
	"strconv"

	"github.com/ju-zp/tasker/svc/models"

	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
	"github.com/ju-zp/tasker/svc/restapi/operations"
)

// Context for task handlers
type Context struct {
	DB *gorm.DB
}

// GetTasks gets all the tasks
func (ctx Context) GetTasks(params operations.GetTasksParams) middleware.Responder {
	var tasks []*models.Task

	ctx.DB.Find(&tasks)

	return operations.NewGetTasksOK().WithPayload(tasks)
}

// CreateTask creates a task
func (ctx Context) CreateTask(params operations.CreateTaskParams) middleware.Responder {
	task := models.Task{
		Title: params.Body.Title,
		Done:  false,
	}

	err := ctx.DB.Create(&task).Error

	if err != nil {
		fmt.Println("unable to save task")
		fmt.Println("error: ", err.Error())
	}

	return operations.NewCreateTaskOK().WithPayload(&task)
}

// GetTaskTodos gets the task and the associated todos
func (ctx Context) GetTaskTodos(params operations.GetTaskTodosParams) middleware.Responder {
	id, _ := strconv.Atoi(params.TaskID)

	task := models.Task{
		ID: int64(id),
	}

	ctx.DB.Find(&task)

	var todos []*models.Todo

	ctx.DB.Where("task_id = ?", params.TaskID).Find(&todos)

	return operations.NewGetTaskTodosOK().WithPayload(&operations.GetTaskTodosOKBody{
		Task:  &task,
		Todos: todos,
	})
}

// CreateTaskTodo creates a todo for a given task
func (ctx Context) CreateTaskTodo(params operations.CreateTaskTodoParams) middleware.Responder {
	id, _ := strconv.Atoi(params.TaskID)

	task := models.Task{
		ID: int64(id),
	}

	err := ctx.DB.Find(&task).Error

	if err != nil {
		return nil
	}

	todo := models.Todo{
		Todo: params.Body.Todo,
		TaskID: task.ID,
	}

	err = ctx.DB.Create(&todo).Error

	if err != nil {
		fmt.Println("unable to save todo")
		fmt.Println("error: ", err.Error())
	}

	return operations.NewCreateTaskTodoOK().WithPayload("success")
}
