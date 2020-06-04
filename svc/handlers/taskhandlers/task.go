package taskhandlers

import (
	"fmt"

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