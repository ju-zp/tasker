package projecthandlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
	"github.com/ju-zp/tasker/svc/project"
	"github.com/ju-zp/tasker/svc/services/tasker/restapi/operations"
)

type Context struct {
	DB *gorm.DB
	Repository *project.Repository
}

// GetProjects gets all the projects
func (ctx *Context)GetProjects(params operations.GetProjectsParams) middleware.Responder {
	projects, err := ctx.Repository.FindAll()
	if err != nil {
		return nil
	}

	return operations.NewGetProjectsOK().WithPayload(projects)
}

// CreateProject creates a new project
func (ctx *Context)CreateProject(params operations.CreateProjectParams) middleware.Responder {
	err := ctx.Repository.Create(params.Body.Title, &params.Body.Description)
	if err != nil {
		return nil
	}

	return operations.NewCreateProjectOK().WithPayload("success")
}

// GetProject gets a single project by id and its associated tasks and todos
func (ctx *Context)GetProject(params operations.GetProjectParams) middleware.Responder {
	project, taskTodos, err := ctx.Repository.FindByID(params.ProjectID)
	if err != nil {
		return nil
	}

	return operations.NewGetProjectOK().WithPayload(&operations.GetProjectOKBody{
		Project: project,
		Tasks:   taskTodos,
	})
}

// DeleteProject deletes the project and it's associated tasks and todos
func (ctx *Context)DeleteProject(params operations.DeleteProjectParams) middleware.Responder {
	err := ctx.Repository.DeleteById(params.ProjectID)
	if err != nil {
		return nil
	}

	return operations.NewDeleteProjectOK()
}