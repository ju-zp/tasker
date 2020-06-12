// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	"github.com/ju-zp/tasker/svc/models"
)

// NewSetTodoStatusParams creates a new SetTodoStatusParams object
// no default values defined in spec.
func NewSetTodoStatusParams() SetTodoStatusParams {

	return SetTodoStatusParams{}
}

// SetTodoStatusParams contains all the bound params for the set todo status operation
// typically these are obtained from a http.Request
//
// swagger:parameters setTodoStatus
type SetTodoStatusParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: body
	*/
	Body *models.TodoStatus
	/*ID of a todo
	  Required: true
	  In: path
	*/
	TodoID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewSetTodoStatusParams() beforehand.
func (o *SetTodoStatusParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.TodoStatus
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("body", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}
	}
	rTodoID, rhkTodoID, _ := route.Params.GetOK("todoId")
	if err := o.bindTodoID(rTodoID, rhkTodoID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindTodoID binds and validates parameter TodoID from path.
func (o *SetTodoStatusParams) bindTodoID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.TodoID = raw

	return nil
}