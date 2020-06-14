package projecthandlers

import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
	"github.com/ju-zp/tasker/svc/restapi/operations"
)

type Context struct {
	DB *gorm.DB
}

func (ctx *Context)CreateProject(params operations.CreateProjectParams) middleware.Responder {
	fmt.Println("heellooo")

	operations.NewCreateProjectOK().WithPayload("success")
}