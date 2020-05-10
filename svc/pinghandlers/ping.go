package pinghandlers

import (
	"fmt"

	"github.com/go-openapi/runtime/middleware"
	"github.com/ju-zp/tasker/gen/restapi/operations"
)

// GetPing returns the status of the server
func GetPing(params operations.GetPingParams) middleware.Responder {
	fmt.Println("here")
	return operations.NewGetPingOK().WithPayload("Server is running")
}
