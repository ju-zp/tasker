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