// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// SetTodoStatusHandlerFunc turns a function with the right signature into a set todo status handler
type SetTodoStatusHandlerFunc func(SetTodoStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SetTodoStatusHandlerFunc) Handle(params SetTodoStatusParams) middleware.Responder {
	return fn(params)
}

// SetTodoStatusHandler interface for that can handle valid set todo status params
type SetTodoStatusHandler interface {
	Handle(SetTodoStatusParams) middleware.Responder
}

// NewSetTodoStatus creates a new http.Handler for the set todo status operation
func NewSetTodoStatus(ctx *middleware.Context, handler SetTodoStatusHandler) *SetTodoStatus {
	return &SetTodoStatus{Context: ctx, Handler: handler}
}

/*SetTodoStatus swagger:route POST /todo/{todoId}/status setTodoStatus

SetTodoStatus set todo status API

*/
type SetTodoStatus struct {
	Context *middleware.Context
	Handler SetTodoStatusHandler
}

func (o *SetTodoStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSetTodoStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
