// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/ju-zp/tasker/integration/services/tasker/models"
)

// NewSetTodoStatusParams creates a new SetTodoStatusParams object
// with the default values initialized.
func NewSetTodoStatusParams() *SetTodoStatusParams {
	var ()
	return &SetTodoStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetTodoStatusParamsWithTimeout creates a new SetTodoStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetTodoStatusParamsWithTimeout(timeout time.Duration) *SetTodoStatusParams {
	var ()
	return &SetTodoStatusParams{

		timeout: timeout,
	}
}

// NewSetTodoStatusParamsWithContext creates a new SetTodoStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetTodoStatusParamsWithContext(ctx context.Context) *SetTodoStatusParams {
	var ()
	return &SetTodoStatusParams{

		Context: ctx,
	}
}

// NewSetTodoStatusParamsWithHTTPClient creates a new SetTodoStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetTodoStatusParamsWithHTTPClient(client *http.Client) *SetTodoStatusParams {
	var ()
	return &SetTodoStatusParams{
		HTTPClient: client,
	}
}

/*SetTodoStatusParams contains all the parameters to send to the API endpoint
for the set todo status operation typically these are written to a http.Request
*/
type SetTodoStatusParams struct {

	/*Body*/
	Body *models.TodoStatus
	/*TodoID
	  ID of a todo

	*/
	TodoID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set todo status params
func (o *SetTodoStatusParams) WithTimeout(timeout time.Duration) *SetTodoStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set todo status params
func (o *SetTodoStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set todo status params
func (o *SetTodoStatusParams) WithContext(ctx context.Context) *SetTodoStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set todo status params
func (o *SetTodoStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set todo status params
func (o *SetTodoStatusParams) WithHTTPClient(client *http.Client) *SetTodoStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set todo status params
func (o *SetTodoStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set todo status params
func (o *SetTodoStatusParams) WithBody(body *models.TodoStatus) *SetTodoStatusParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set todo status params
func (o *SetTodoStatusParams) SetBody(body *models.TodoStatus) {
	o.Body = body
}

// WithTodoID adds the todoID to the set todo status params
func (o *SetTodoStatusParams) WithTodoID(todoID string) *SetTodoStatusParams {
	o.SetTodoID(todoID)
	return o
}

// SetTodoID adds the todoId to the set todo status params
func (o *SetTodoStatusParams) SetTodoID(todoID string) {
	o.TodoID = todoID
}

// WriteToRequest writes these params to a swagger request
func (o *SetTodoStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param todoId
	if err := r.SetPathParam("todoId", o.TodoID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
