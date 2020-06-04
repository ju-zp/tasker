// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateTodoHandlerFunc turns a function with the right signature into a create todo handler
type CreateTodoHandlerFunc func(CreateTodoParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateTodoHandlerFunc) Handle(params CreateTodoParams) middleware.Responder {
	return fn(params)
}

// CreateTodoHandler interface for that can handle valid create todo params
type CreateTodoHandler interface {
	Handle(CreateTodoParams) middleware.Responder
}

// NewCreateTodo creates a new http.Handler for the create todo operation
func NewCreateTodo(ctx *middleware.Context, handler CreateTodoHandler) *CreateTodo {
	return &CreateTodo{Context: ctx, Handler: handler}
}

/*CreateTodo swagger:route POST /todo createTodo

CreateTodo create todo API

*/
type CreateTodo struct {
	Context *middleware.Context
	Handler CreateTodoHandler
}

func (o *CreateTodo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateTodoParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}