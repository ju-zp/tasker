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
)

// NewCreateTaskTodoParams creates a new CreateTaskTodoParams object
// with the default values initialized.
func NewCreateTaskTodoParams() *CreateTaskTodoParams {
	var ()
	return &CreateTaskTodoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTaskTodoParamsWithTimeout creates a new CreateTaskTodoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateTaskTodoParamsWithTimeout(timeout time.Duration) *CreateTaskTodoParams {
	var ()
	return &CreateTaskTodoParams{

		timeout: timeout,
	}
}

// NewCreateTaskTodoParamsWithContext creates a new CreateTaskTodoParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateTaskTodoParamsWithContext(ctx context.Context) *CreateTaskTodoParams {
	var ()
	return &CreateTaskTodoParams{

		Context: ctx,
	}
}

// NewCreateTaskTodoParamsWithHTTPClient creates a new CreateTaskTodoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateTaskTodoParamsWithHTTPClient(client *http.Client) *CreateTaskTodoParams {
	var ()
	return &CreateTaskTodoParams{
		HTTPClient: client,
	}
}

/*CreateTaskTodoParams contains all the parameters to send to the API endpoint
for the create task todo operation typically these are written to a http.Request
*/
type CreateTaskTodoParams struct {

	/*Body*/
	Body CreateTaskTodoBody
	/*TaskID
	  ID of a task

	*/
	TaskID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create task todo params
func (o *CreateTaskTodoParams) WithTimeout(timeout time.Duration) *CreateTaskTodoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create task todo params
func (o *CreateTaskTodoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create task todo params
func (o *CreateTaskTodoParams) WithContext(ctx context.Context) *CreateTaskTodoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create task todo params
func (o *CreateTaskTodoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create task todo params
func (o *CreateTaskTodoParams) WithHTTPClient(client *http.Client) *CreateTaskTodoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create task todo params
func (o *CreateTaskTodoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create task todo params
func (o *CreateTaskTodoParams) WithBody(body CreateTaskTodoBody) *CreateTaskTodoParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create task todo params
func (o *CreateTaskTodoParams) SetBody(body CreateTaskTodoBody) {
	o.Body = body
}

// WithTaskID adds the taskID to the create task todo params
func (o *CreateTaskTodoParams) WithTaskID(taskID string) *CreateTaskTodoParams {
	o.SetTaskID(taskID)
	return o
}

// SetTaskID adds the taskId to the create task todo params
func (o *CreateTaskTodoParams) SetTaskID(taskID string) {
	o.TaskID = taskID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTaskTodoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param taskId
	if err := r.SetPathParam("taskId", o.TaskID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
