package pinghandlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/ju-zp/tasker/svc/restapi/operations"
)

// GetPing returns the status of the server
func GetPing(params operations.GetPingParams) middleware.Responder {
	return operations.NewGetPingOK().WithPayload("Server is running")
}
