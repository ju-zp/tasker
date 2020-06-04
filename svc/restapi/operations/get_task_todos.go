// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/ju-zp/tasker/svc/models"
)

// GetTaskTodosHandlerFunc turns a function with the right signature into a get task todos handler
type GetTaskTodosHandlerFunc func(GetTaskTodosParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTaskTodosHandlerFunc) Handle(params GetTaskTodosParams) middleware.Responder {
	return fn(params)
}

// GetTaskTodosHandler interface for that can handle valid get task todos params
type GetTaskTodosHandler interface {
	Handle(GetTaskTodosParams) middleware.Responder
}

// NewGetTaskTodos creates a new http.Handler for the get task todos operation
func NewGetTaskTodos(ctx *middleware.Context, handler GetTaskTodosHandler) *GetTaskTodos {
	return &GetTaskTodos{Context: ctx, Handler: handler}
}

/*GetTaskTodos swagger:route GET /task/{taskId} getTaskTodos

GetTaskTodos get task todos API

*/
type GetTaskTodos struct {
	Context *middleware.Context
	Handler GetTaskTodosHandler
}

func (o *GetTaskTodos) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetTaskTodosParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetTaskTodosOKBody the task and associated todos
//
// swagger:model GetTaskTodosOKBody
type GetTaskTodosOKBody struct {

	// task
	Task *models.Task `json:"task,omitempty"`

	// Array of todos
	Todos []*models.Todo `json:"todos"`
}

// Validate validates this get task todos o k body
func (o *GetTaskTodosOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTask(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTodos(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTaskTodosOKBody) validateTask(formats strfmt.Registry) error {

	if swag.IsZero(o.Task) { // not required
		return nil
	}

	if o.Task != nil {
		if err := o.Task.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getTaskTodosOK" + "." + "task")
			}
			return err
		}
	}

	return nil
}

func (o *GetTaskTodosOKBody) validateTodos(formats strfmt.Registry) error {

	if swag.IsZero(o.Todos) { // not required
		return nil
	}

	for i := 0; i < len(o.Todos); i++ {
		if swag.IsZero(o.Todos[i]) { // not required
			continue
		}

		if o.Todos[i] != nil {
			if err := o.Todos[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getTaskTodosOK" + "." + "todos" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetTaskTodosOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTaskTodosOKBody) UnmarshalBinary(b []byte) error {
	var res GetTaskTodosOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
