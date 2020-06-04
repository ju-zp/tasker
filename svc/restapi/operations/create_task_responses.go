// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ju-zp/tasker/svc/models"
)

// CreateTaskOKCode is the HTTP code returned for type CreateTaskOK
const CreateTaskOKCode int = 200

/*CreateTaskOK Created task

swagger:response createTaskOK
*/
type CreateTaskOK struct {

	/*
	  In: Body
	*/
	Payload *models.Task `json:"body,omitempty"`
}

// NewCreateTaskOK creates CreateTaskOK with default headers values
func NewCreateTaskOK() *CreateTaskOK {

	return &CreateTaskOK{}
}

// WithPayload adds the payload to the create task o k response
func (o *CreateTaskOK) WithPayload(payload *models.Task) *CreateTaskOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create task o k response
func (o *CreateTaskOK) SetPayload(payload *models.Task) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTaskOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateTaskBadRequestCode is the HTTP code returned for type CreateTaskBadRequest
const CreateTaskBadRequestCode int = 400

/*CreateTaskBadRequest Incomplete data

swagger:response createTaskBadRequest
*/
type CreateTaskBadRequest struct {

	/*Body of the request was Incomplete
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewCreateTaskBadRequest creates CreateTaskBadRequest with default headers values
func NewCreateTaskBadRequest() *CreateTaskBadRequest {

	return &CreateTaskBadRequest{}
}

// WithPayload adds the payload to the create task bad request response
func (o *CreateTaskBadRequest) WithPayload(payload string) *CreateTaskBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create task bad request response
func (o *CreateTaskBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTaskBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}