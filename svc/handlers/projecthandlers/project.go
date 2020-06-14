package projecthandlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
	"github.com/ju-zp/tasker/svc/models"
	"github.com/ju-zp/tasker/svc/restapi/operations"
)

type Context struct {
	DB *gorm.DB
}

// GetProjects gets all the projects
func (ctx *Context)GetProjects(params operations.GetProjectsParams) middleware.Responder {
	var projects []*models.Project

	ctx.DB.Find(&projects)

	return operations.NewGetProjectsOK().WithPayload(projects)
}

// CreateProject creates a new project
func (ctx *Context)CreateProject(params operations.CreateProjectParams) middleware.Responder {
	project := models.Project{
		Description: params.Body.Description,
		Title:       params.Body.Title,
	}

	err := ctx.DB.Create(&project).Error

	if err != nil {
		return nil
	}

	return operations.NewCreateProjectOK().WithPayload("success")
}