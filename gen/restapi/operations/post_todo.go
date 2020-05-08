// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostTodoHandlerFunc turns a function with the right signature into a post todo handler
type PostTodoHandlerFunc func(PostTodoParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostTodoHandlerFunc) Handle(params PostTodoParams) middleware.Responder {
	return fn(params)
}

// PostTodoHandler interface for that can handle valid post todo params
type PostTodoHandler interface {
	Handle(PostTodoParams) middleware.Responder
}

// NewPostTodo creates a new http.Handler for the post todo operation
func NewPostTodo(ctx *middleware.Context, handler PostTodoHandler) *PostTodo {
	return &PostTodo{Context: ctx, Handler: handler}
}

/*PostTodo swagger:route POST /todo postTodo

PostTodo post todo API

*/
type PostTodo struct {
	Context *middleware.Context
	Handler PostTodoHandler
}

func (o *PostTodo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostTodoParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}